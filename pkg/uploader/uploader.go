package uploader

import (
	"bytes"
	"deployment-endpoint/pkg/logger"
	"fmt"
	"strings"

	"github.com/minio/minio-go/v6"
	"github.com/prometheus/client_golang/prometheus"
)

type I interface {
	Upload([]byte) error
}

type Uploader struct {
	Config *Config
	logger logger.Logger
	client *minio.Client
}

func New(conf *Config, prom *prometheus.Registry, logger logger.Logger) (*Uploader, error) {

	// Initialize minio client object.
	minioClient, err := minio.New(conf.host, conf.accessKeyID, conf.secretAccessKey, conf.useSSL)
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
		client: minioClient,
		Config: conf,
		logger: logger,
	}

	return uploader, nil

}

func (u *Uploader) Upload(byteData []byte) error {

	objectName := "golden-oldies.zip"
	destBucket := strings.ToLower(fmt.Sprintf("algorun/%s/%s",
		u.Config.deploymentOwnerUserName,
		u.Config.deploymentName))

	dataReader := bytes.NewReader(byteData)

	// Upload file with PutObject
	n, err := u.client.PutObject(destBucket, objectName, dataReader, int64(len(byteData)), minio.PutObjectOptions{})
	if err != nil {
		u.logger.Errorf("Error uploading file [%s] [%v]", objectName, err)
	}

	u.logger.Infof("Successfully uploaded %s of size %d\n", objectName, n)

	return err

}
