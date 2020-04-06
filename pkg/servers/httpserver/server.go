package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"deployment-endpoint/openapi"
	"deployment-endpoint/pkg/server"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"k8s.io/utils/pointer"
)

const (
	// HTTP headers used by the API.
	hdrContentLength = "Content-Length"
	hdrContentType   = "Content-Type"
	pathEndpoint     = "endpointPath"
)

var (
	healthy bool
)

type Server server.T

func (s *Server) getRouter(deploymentOwnerUserName string, deploymentName string) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc(fmt.Sprintf("/%s/%s/{%s}", deploymentOwnerUserName, deploymentName, pathEndpoint), s.messageHandler)
	r.HandleFunc(fmt.Sprintf("/%s/%s/{%s}/{runId}", deploymentOwnerUserName, deploymentName, pathEndpoint), s.messageHandler)

	r.HandleFunc("/topics", s.topicsHandler).Methods("GET")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if healthy {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(500)
		}

	})

	return r
}

func (s *Server) Start(configPath string) {

	go func() {
		for h := range s.HealthyChan {
			healthy = h
		}
	}()

	deploymentOwnerUserName := s.Config.GetString("deploymentOwnerUserName")
	deploymentName := s.Config.GetString("deploymentName")

	r := s.getRouter(deploymentOwnerUserName, deploymentName)
	c := s.Config.Sub(configPath)
	addr := c.GetString("listen")
	httpServer := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		s.Logger.Infof("Listening for HTTP requests on %s", addr)
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.Logger.Errorf("Unable to start http server: %v", err)
			s.HealthyChan <- false
		}
	}()

	s.Wg.Add(1)
	s.Logger.Info("Registered HTTP server in servers pool")
	s.Wg.Wait()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)
	close(s.Done)

}

func (s *Server) messageHandler(w http.ResponseWriter, r *http.Request) {

	var msg []byte
	var err error

	endpointPath := mux.Vars(r)[pathEndpoint]
	params := r.URL.Query()
	contentType := r.Header.Get(hdrContentType)
	runID := mux.Vars(r)["runId"]

	headers := make(map[string][]byte)

	if runID == "" {
		runIDUuid := uuid.New()
		runID = runIDUuid.String()
	}

	headers["endpointParams"] = []byte(params.Encode())
	headers["contentType"] = []byte(contentType)

	if _, ok := s.EndpointPaths[endpointPath]; !ok {
		// Create error response
		errMsg := openapi.ApiBadRequestResponse{
			StatusCode: 400,
			Message:    "Failed to run endpoint",
			Errors: []openapi.ErrorModel{
				openapi.ErrorModel{
					ErrorCode: 50002,
					Message:   pointer.StringPtr(fmt.Sprintf("Endpoint Output [%s] was not found", endpointPath)),
				},
			},
		}
		errBytes, _ := json.Marshal(errMsg)
		http.Error(w, string(errBytes), http.StatusBadRequest)
		return
	}

	topic := strings.ToLower(fmt.Sprintf("algorun.%s.%s.endpoint.%s",
		s.Config.GetString("deploymentOwnerUserName"),
		s.Config.GetString("deploymentName"),
		endpointPath))

	if msg, err = readMsg(r); err != nil {
		// Create error response
		errMsg := openapi.ApiBadRequestResponse{
			StatusCode: 400,
			Message:    "Failed to run endpoint",
			Errors: []openapi.ErrorModel{
				openapi.ErrorModel{
					ErrorCode: 50001,
					Message:   pointer.StringPtr(fmt.Sprintf("Error reading request body. %s", err.Error())),
				},
			},
		}
		errBytes, _ := json.Marshal(errMsg)
		http.Error(w, string(errBytes), http.StatusBadRequest)
		return
	}

	// Get the message type for this output
	if s.EndpointPaths[endpointPath].MessageDataType == "FileReference" {

		// Upload the file to storage and generate file reference
		// Create file uuid
		fileName := uuid.New()
		bucketName := fmt.Sprintf("%s.%s",
			strings.ToLower(s.Config.GetString("deploymentOwnerUserName")),
			strings.ToLower(s.Config.GetString("deploymentName")))
		fileReference := openapi.FileReference{
			Host:   s.Uploader.Config.Host,
			Bucket: bucketName,
			File:   fileName.String(),
		}
		err := s.Uploader.Upload(fileReference, msg)
		if err != nil {
			// Create error message
			errMsg := openapi.ApiBadRequestResponse{
				StatusCode: 400,
				Message:    "Failed to run endpoint",
				Errors: []openapi.ErrorModel{
					openapi.ErrorModel{
						ErrorCode: 50003,
						Message:   pointer.StringPtr(fmt.Sprintf("Error uploading to storage for file reference [%s]", fileReference.File)),
					},
					openapi.ErrorModel{
						ErrorCode: 50004,
						Message:   pointer.StringPtr(fmt.Sprintf("Storage error [%s]", err.Error())),
					},
				},
			}
			errBytes, _ := json.Marshal(errMsg)
			http.Error(w, string(errBytes), http.StatusBadRequest)
			return
		}

		jsonBytes, _ := json.Marshal(fileReference)

		s.Producer.Send(topic, headers, runID, jsonBytes)

	} else {

		s.Producer.Send(topic, headers, runID, msg)

	}

}

func readMsg(r *http.Request) ([]byte, error) {

	// TODO: Validate content type matches the endpoint path config

	if _, ok := r.Header[hdrContentLength]; !ok {
		return nil, errors.Errorf("missing %s header", hdrContentLength)
	}
	messageSizeStr := r.Header.Get(hdrContentLength)
	msgSize, err := strconv.Atoi(messageSizeStr)
	if err != nil {
		return nil, errors.Errorf("invalid %s header: %s", hdrContentLength, messageSizeStr)
	}
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read message")
	}
	if len(msg) != msgSize {
		return nil, errors.Errorf("message size does not match %s: expected=%v, actual=%v",
			hdrContentLength, msgSize, len(msg))
	}
	return msg, nil

}

func (s *Server) topicsHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Infof("Pulling TOPICS")
	topics, err := s.Producer.ListTopics()
	if err != nil {
		// Create error message
		errMsg := openapi.ApiBadRequestResponse{
			StatusCode: 400,
			Message:    "Failed to get topics",
			Errors: []openapi.ErrorModel{
				openapi.ErrorModel{
					ErrorCode: 50005,
					Message:   pointer.StringPtr(fmt.Sprintf("Error getting topics [%s]", err)),
				},
			},
		}
		errBytes, _ := json.Marshal(errMsg)
		http.Error(w, string(errBytes), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, strings.Join(topics, "\n"))
}

func (s *Server) Stop() {
	s.Logger.Info("Stopping HTTP server")
	s.Wg.Done()
	<-s.Done
	s.Logger.Info("Stopped HTTP server")
}
