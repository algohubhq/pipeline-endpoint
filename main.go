package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"deployment-endpoint/pkg/config"
	"deployment-endpoint/pkg/kafka"
	"deployment-endpoint/pkg/logger"
	"deployment-endpoint/pkg/servers"
	"deployment-endpoint/pkg/uploader"

	ckg "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/prometheus/client_golang/prometheus"
)

// it's needed to pass wal config.
// We need to refactor whole config chain to be parsed the same way - using yaml.Unmarshal
type Config struct {
	Producer kafka.Config `yaml:"producer"`
}

var (
	defaults = map[string]interface{}{
		"global.log.level":            "info",
		"global.log.encoding":         "json",
		"global.log.outputPaths":      ("stdout"),
		"global.log.errorOutputPaths": ("stderr"),
		"global.log.encoderConfig":    logger.NewEncoderConfig(),
		"server.http.listen":          ":18080",
		"server.grpc.listen":          ":18282",
		"server.monitoring.listen":    ":28080",
		"producer.cb.interval":        0,
		"producer.cb.timeout":         "20s",
		"producer.cb.fails":           5,
		"producer.cb.requests":        3,
		"producer.resend.period":      "33s",
		"producer.resend.rate_limit":  10000,
		"producer.wal.mode":           "fallback",
		"producer.wal.path":           "/data/wal",
		// Please take a look at:
		// https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
		// for more configuration parameters
		"kafka.compression.codec":                     "gzip",
		"kafka.batch.num.messages":                    100000,
		"kafka.socket.timeout.ms":                     10000, // mark connection as stalled
		"kafka.message.timeout.ms":                    60000, // try to deliver message with retries
		"kafka.max.in.flight.requests.per.connection": 20,
		"server.grpc.max.request.size":                4 * 1024 * 1024,
		"server.grpc.monitoring.histogram.enable":     true,
		"server.grpc.monitoring.enable":               true,
	}
)

func main() {

	healthyChan := make(chan bool)
	var err error
	var configPathName string
	flag.StringVar(&configPathName, "config", "", "Configuration file to load")
	flag.Parse()

	// We need to shut down gracefully when the user hits Ctrl-C.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGHUP)

	s := new(servers.T)
	// Health channel
	s.HealthyChan = healthyChan
	c := &config.T{
		Filename:  configPathName,
		EnvPrefix: "EP",
	}
	s.Config, err = c.ReadConfig(defaults)
	if err != nil {
		panic(fmt.Sprintf("Could not read config file: %v", err))
	}
	// logs
	cfg := logger.NewLogConfig(s.Config.Sub("global.log"))
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	s.Logger = logger.Sugar()
	// metrics
	s.Prometheus = prometheus.NewRegistry()

	// endpoint outputs config
	var endpointPaths []config.EndpointPath
	err = s.Config.UnmarshalKey("paths", &endpointPaths)
	if err != nil {
		s.Logger.Errorf("Unable to deserialize endpoint paths [%v]", err)
	}

	outputMap := make(map[string]config.EndpointPath)
	for _, output := range endpointPaths {
		outputMap[output.Name] = output
	}
	s.EndpointPaths = outputMap

	// servers
	kafkaParams, err := kafka.Viper2Config(s.Config)
	if err != nil {
		os.Exit(1)
	}

	// separate config read for wal. This is to be refactored
	// Get full config from the envar
	var anotherConfig Config
	configData := os.Getenv("ENDPOINT_CONFIG")
	if configData != "" {
		if err := json.Unmarshal([]byte(configData), &anotherConfig); err != nil {
			s.Logger.Fatalf("error: %v", err)
			os.Exit(1)
		}
	}

	// Create the uploader
	uploaderConfig := uploader.UploaderConfig(s.Config, s.Logger)
	if err != nil {
		os.Exit(1)
	}

	uploader, err := uploader.New(uploaderConfig, s.Prometheus, s.Logger, healthyChan)
	if err != nil {
		os.Exit(1)
	}
	uploader.HealthyChan = healthyChan

	s.Uploader = uploader

	producer := &kafka.T{}
	producer.HealthyChan = healthyChan
	producer.Logger = s.Logger
	producer.Config = kafka.ProducerConfig(s.Config)
	producer.Config.Wal = anotherConfig.Producer.Wal
	err = producer.Init(&kafkaParams, s.Prometheus)
	defer producer.CloseWalDB()
	if err != nil {
		s.Logger.Fatal("Could not initialize producer")
		os.Exit(1)
	}

	s.Producer = producer
	s.Start()

	s.Logger.Info("All Servers Started")

	healthyChan <- true

	s.Logger.Info("Healthy")

	for {
		signal := <-sig
		switch signal {
		case syscall.SIGHUP:
			s.Logger.Info("Got SIGHUP: setting up new Kafka producer")
			kp, err := ckg.NewProducer(&kafkaParams)
			if err != nil {
				s.Logger.Errorf("ERROR. Could not create producer on SIGHUP due to: %v", err)
			} else {
				s.Producer.AddActiveProducer(kp, &kafkaParams)
			}
		case syscall.SIGTERM, syscall.SIGINT:
			s.Stop()
			s.Producer.Shutdown()
			for {
				if !s.Producer.QueueIsEmpty() {
					s.Logger.Info("We still have messages in queue, waiting")
					time.Sleep(5 * time.Second)
				} else {
					s.Logger.Info("Queue is empty, shut down properly")
					s.Producer.GetProducer().Producer.Close()
					return
				}
			}
		}
	}
}
