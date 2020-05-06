package servers

import (
	"sync"

	"pipeline-endpoint/pkg/server"
	"pipeline-endpoint/pkg/servers/grpcserver"
	"pipeline-endpoint/pkg/servers/httpserver"
	"pipeline-endpoint/pkg/servers/monitoring"
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

	if s.Config.IsSet(monitoringPath + ".listen") {
		monitSrv := &monitoring.Server{
			HealthyChan:   s.HealthyChan,
			Producer:      s.Producer,
			Uploader:      s.Uploader,
			EndpointPaths: s.EndpointPaths,
			Config:        s.Config,
			Prometheus:    s.Prometheus,
			Logger:        s.Logger,
			Wg:            new(sync.WaitGroup),
			Done:          make(chan bool),
		}
		go monitSrv.Start(monitoringPath)
	}
	if s.Config.IsSet(httpPath + ".listen") {
		httpSrv := &httpserver.Server{
			HealthyChan:   s.HealthyChan,
			Producer:      s.Producer,
			Uploader:      s.Uploader,
			EndpointPaths: s.EndpointPaths,
			Config:        s.Config,
			Prometheus:    s.Prometheus,
			Logger:        s.Logger,
			Wg:            new(sync.WaitGroup),
			Done:          make(chan bool),
		}
		go httpSrv.Start(httpPath)
		s.Servers = append(s.Servers, httpSrv)
	}
	if s.Config.IsSet(grpcPath + ".listen") {
		grpcSrv := &grpcserver.Server{
			HealthyChan:   s.HealthyChan,
			Producer:      s.Producer,
			Uploader:      s.Uploader,
			Config:        s.Config,
			EndpointPaths: s.EndpointPaths,
			Prometheus:    s.Prometheus,
			Logger:        s.Logger,
			Wg:            new(sync.WaitGroup),
			Done:          make(chan bool),
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
