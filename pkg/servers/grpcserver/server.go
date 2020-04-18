package grpcserver

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/url"
	"strings"

	"deployment-endpoint/openapi"
	"deployment-endpoint/pkg/server"

	pb "deployment-endpoint/pkg/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Server server.T

func (s *Server) Start(configPath string) {
	var grpcSrv *grpc.Server
	c := s.Config.Sub(configPath)
	addr := c.GetString("listen")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		s.Logger.Fatal(err.Error())
	}

	if s.Config.GetBool(configPath + ".monitoring.enable") {
		s.Logger.Info("Monitoring is enabled, applying GRPC interceptors")
		// Create a gRPC Server with gRPC interceptor.
		grpcSrv = grpc.NewServer(
			grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
			grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
			grpc.MaxMsgSize(s.Config.GetInt(configPath+".max.request.size")),
		)
		// Initialize all metrics.
		s.RegisterMetrics()
		grpcMetrics.InitializeMetrics(grpcSrv)
		// Histograms can be expensive on Prometheus servers.
		if s.Config.GetBool(configPath + ".monitoring.histogram.enable") {
			grpcMetrics.EnableHandlingTimeHistogram()
		}
	} else {
		s.Logger.Info("Monitoring is NOT enabled, enable it if you would like to see prometheus metrics")
		grpcSrv = grpc.NewServer(
			grpc.MaxMsgSize(s.Config.GetInt(configPath + ".max.request.size")),
		)

	}
	pb.RegisterKafkaAmbassadorServer(grpcSrv, s)

	s.Logger.Infof("Listening for GRPC requests on %s", addr)

	go func() {
		if err = grpcSrv.Serve(lis); err != nil {
			s.Logger.Fatalf("failed to serve: %v", err)
			s.HealthyChan <- false
		}
	}()
	s.Wg.Add(1)
	s.Logger.Info("Registered GRPC in servers pool")
	s.Wg.Wait()
	// GracefulStop stops the gRPC server gracefully.
	// It stops the server from accepting new connections and RPCs
	// and blocks until all the pending RPCs are finished.
	s.Logger.Warn("Initiating graceful stop of GRPC server")
	grpcSrv.GracefulStop()
	close(s.Done)
}

func (s *Server) Produce(stream pb.KafkaAmbassador_ProduceServer) error {
	var res *pb.ProdRs
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			s.Logger.Errorf("Could not receive stream from client: %v", err)
			return err
		}

		deploymentOwner := s.Config.GetString("deploymentOwner")
		deploymentName := s.Config.GetString("deploymentName")
		endpointPath := req.EndpointPath
		runID := req.RunID
		contentType := req.ContentType

		headers := make(map[string][]byte)

		if runID == "" {
			runIDUuid := uuid.New()
			runID = runIDUuid.String()
		}

		// encode the parameters
		urlValues := url.Values{}
		for k, v := range req.Parameters {
			urlValues.Set(k, v)
		}

		headers["contentType"] = []byte(contentType)
		headers["endpointParams"] = []byte(urlValues.Encode())

		if deploymentOwner != req.DeploymentOwner ||
			deploymentName != req.DeploymentName {
			err = fmt.Errorf("Received message intended for deployment [%s/%s] but this endpoint handles [%s/%s]. Message dropped",
				req.DeploymentOwner, req.DeploymentName, deploymentOwner, deploymentName)
			s.Logger.Errorf("%v", err)
			return err
		}

		topic := strings.ToLower(fmt.Sprintf("algorun.%s.%s.endpoint.%s",
			deploymentOwner,
			deploymentName,
			req.EndpointPath))

		// Get the message type for this output
		if s.EndpointPaths[endpointPath].MessageDataType == "FileReference" {
			// Upload the file to storage and generate file reference
			// Create file uuid
			fileName := uuid.New()
			bucketName := fmt.Sprintf("%s.%s",
				strings.ToLower(s.Config.GetString("deploymentOwner")),
				strings.ToLower(s.Config.GetString("deploymentName")))
			fileReference := openapi.FileReference{
				Host:   s.Uploader.Config.Host,
				Bucket: bucketName,
				File:   fileName.String(),
			}
			err := s.Uploader.Upload(fileReference, req.Message)
			if err != nil {
				s.Logger.Errorf("Could not upload message to storage: %v. Error: %s", fileReference, err)
				return err
			}

			jsonBytes, jsonErr := json.Marshal(fileReference)
			if jsonErr != nil {
				s.Logger.Errorf("Error serializing the file reference: %v. Error: %s", fileReference, err)
				return err
			}

			s.Producer.Send(topic, headers, runID, jsonBytes)

		} else {
			s.Producer.Send(topic, headers, runID, req.Message)
		}

		res = &pb.ProdRs{StreamOffset: req.StreamOffset}
		err = stream.Send(res)
		if err != nil {
			s.Logger.Errorf("Could not stream (GRPC) to the client: %s", err)
			return err
		}

	}
}

func (s *Server) ListTopics(ctx context.Context, nothing *pb.Empty) (*pb.ListTopicsResponse, error) {
	ret := &pb.ListTopicsResponse{}
	topics, err := s.Producer.ListTopics()
	ret.Topics = topics
	return ret, err
}

func (s *Server) Stop() {
	s.Logger.Info("Stopping GRPC server")
	s.Wg.Done()
	<-s.Done
	s.Logger.Info("Stopped GRPC server")
}
