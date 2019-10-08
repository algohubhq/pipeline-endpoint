package httpserver

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"deployment-endpoint/pkg/server"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const (
	// HTTP headers used by the API.
	hdrContentLength = "Content-Length"
	hdrContentType   = "Content-Type"
)

type Server server.T

func (s *Server) getRouter(deploymentOwnerUserName string, deploymentName string) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc(fmt.Sprintf("/%s/%s/{endpointOutput}", deploymentOwnerUserName, deploymentName), s.messageHandler)

	r.HandleFunc("/topics", s.topicsHandler).Methods("GET")

	return r
}

func (s *Server) Start(configPath string) {

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
		} else {
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

	endpointOutput := mux.Vars(r)["endpointOutput"]

	if s.EndpointOutputs[endpointOutput] == nil {
		http.Error(w, fmt.Sprintf("Endpoint Output [%s] was not found", endpointOutput), http.StatusNotFound)
	}

	topic := strings.ToLower(fmt.Sprintf("algorun.%s.%s.endpoint.%s",
		s.Config.GetString("deploymentOwnerUserName"),
		s.Config.GetString("deploymentName"),
		endpointOutput))

	if msg, err = readMsg(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the message type for this output
	if s.EndpointOutputs[endpointOutput].MessageDataType == "FileReference" {
		// TODO: Upload the file to the s3 bucket and generate file reference

	} else {
		s.Producer.Send(topic, msg)
	}

}

func readMsg(r *http.Request) ([]byte, error) {

	contentType := r.Header.Get(hdrContentType)
	if contentType == "text/plain" || contentType == "application/json" {
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

	return nil, errors.Errorf("unsupported content type %s", contentType)

}

func (s *Server) topicsHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Infof("Pulling TOPICS")
	topics, err := s.Producer.ListTopics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
