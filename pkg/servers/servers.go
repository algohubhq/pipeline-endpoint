package servers

import (
	"fmt"
	"sync"

	"deployment-endpoint/pkg/config"
	"deployment-endpoint/pkg/server"
	"deployment-endpoint/pkg/servers/grpcserver"
	"deployment-endpoint/pkg/servers/httpserver"
	"deployment-endpoint/pkg/servers/monitoring"
)

const (
	grpcPath       = "server.grpc"
	httpPath       = "server.http"
	monitoringPath = "server.monitoring"
)

type T struct {
	server.T
	Servers []server.I
}

func (s *T) Start() {

	var endpointOutputs []config.EndpointOutput
	err := s.Config.UnmarshalKey("outputs", &endpointOutputs)
	if err != nil {
		s.Logger.Errorf("Unable to deserialize endpoint outputs [%v]", err)
	}

	outputMap := make(map[string]*config.EndpointOutput)
	for _, output := range endpointOutputs {
		outputMap[output.Name] = &output
	}

	test := s.Config.GetStringMap("outputs")
	fmt.Printf("%v", test)

	if s.Config.IsSet(monitoringPath + ".listen") {
		monitSrv := &monitoring.Server{
			Producer:        s.Producer,
			Uploader:        s.Uploader,
			EndpointOutputs: outputMap,
			Config:          s.Config,
			Prometheus:      s.Prometheus,
			Logger:          s.Logger,
			Wg:              new(sync.WaitGroup),
			Done:            make(chan bool),
		}
		go monitSrv.Start(monitoringPath)
	}
	if s.Config.IsSet(httpPath + ".listen") {
		httpSrv := &httpserver.Server{
			Producer:        s.Producer,
			Uploader:        s.Uploader,
			EndpointOutputs: outputMap,
			Config:          s.Config,
			Prometheus:      s.Prometheus,
			Logger:          s.Logger,
			Wg:              new(sync.WaitGroup),
			Done:            make(chan bool),
		}
		go httpSrv.Start(httpPath)
		s.Servers = append(s.Servers, httpSrv)
	}
	if s.Config.IsSet(grpcPath + ".listen") {
		grpcSrv := &grpcserver.Server{
			Producer:        s.Producer,
			Uploader:        s.Uploader,
			Config:          s.Config,
			EndpointOutputs: outputMap,
			Prometheus:      s.Prometheus,
			Logger:          s.Logger,
			Wg:              new(sync.WaitGroup),
			Done:            make(chan bool),
		}
		go grpcSrv.Start(grpcPath)
		s.Servers = append(s.Servers, grpcSrv)
	}
}

func (s *T) Stop() {
	for i, _ := range s.Servers {
		s.Servers[i].Stop()
	}
}
