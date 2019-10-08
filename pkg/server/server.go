package server

import (
	"sync"

	"deployment-endpoint/pkg/config"
	"deployment-endpoint/pkg/kafka"
	"deployment-endpoint/pkg/logger"
	"deployment-endpoint/pkg/uploader"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
)

// We need this package to prevent cyclic dependencies

type I interface {
	Start(string)
	Stop()
}

type T struct {
	Producer        kafka.I
	Uploader        uploader.I
	EndpointOutputs map[string]*config.EndpointOutput
	Logger          logger.Logger
	Prometheus      *prometheus.Registry
	Config          *viper.Viper
	Wg              *sync.WaitGroup
	Done            chan bool
}

func (s *T) Stop() {

}
