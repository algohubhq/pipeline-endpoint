package uploader

import (
	"bytes"
	"deployment-endpoint/openapi"
	"deployment-endpoint/pkg/logger"
	"fmt"
	"strings"

	"github.com/minio/minio-go/v6"
	"github.com/prometheus/client_golang/prometheus"
)

type Uploader struct {
	HealthyChan chan<- bool
	Config      *Config
	logger      logger.Logger
	client      *minio.Client
}

func New(conf *Config, prom *prometheus.Registry, logger logger.Logger, healthyChan chan<- bool) (*Uploader, error) {

	// Initialize minio client object.
	minioClient, err := minio.New(conf.Host, conf.accessKeyID, conf.secretAccessKey, conf.useSSL)
	if err != nil {
		logger.Errorf("Error initializing minio client [%v]", err)
	}

	destBucket := strings.ToLower(fmt.Sprintf("algorun.%s.%s",
		conf.deploymentOwnerUserName,
		conf.deploymentName))
	// location := "us-east-1"

	// err = minioClient.MakeBucket(destBucket, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(destBucket)
		if errBucketExists != nil && !exists {
			logger.Errorf("Error making bucket [%v]", err)
		}
	} else {
		logger.Infof("Successfully created bucket %s\n", destBucket)
	}

	uploader := &Uploader{
		HealthyChan: healthyChan,
		client:      minioClient,
		Config:      conf,
		logger:      logger,
	}

	return uploader, nil

}

func (u *Uploader) Upload(fileReference openapi.FileReference, byteData []byte) error {

	dataReader := bytes.NewReader(byteData)

	// Upload file with PutObject
	n, err := u.client.PutObject(fileReference.Bucket, fileReference.File, dataReader, int64(len(byteData)), minio.PutObjectOptions{})
	if err != nil {
		u.logger.Errorf("Error uploading file [%s] [%v]", fileReference.File, err)
		u.HealthyChan <- false
		return err
	}

	u.logger.Infof("Successfully uploaded %s of size %d\n", fileReference.File, n)

	return err

}
