package uploader

import (
	"deployment-endpoint/pkg/logger"
	"fmt"
	"net/url"
	"regexp"

	"github.com/spf13/viper"
)

type Config struct {
	Host                    string `yaml:"host"`
	deploymentOwnerUserName string `yaml:"deploymentOwnerUserName"`
	deploymentName          string `yaml:"deploymentName"`
	accessKeyID             string `yaml:"accessKeyID"`
	secretAccessKey         string `yaml:"secretAccessKey"`
	useSSL                  bool   `yaml:"useSSL"`
}

func UploaderConfig(c *viper.Viper, logger logger.Logger) *Config {

	u := new(Config)

	deploymentOwnerUserName := c.GetString("deploymentOwnerUserName")
	deploymentName := c.GetString("deploymentName")

	if deploymentOwnerUserName == "" || deploymentName == "" {
		logger.Panic("Deployment owner and name configuration must not be empty")
	}

	s3ConnectionString := c.GetString("uploader.host")

	// Parse the host connection string
	host, accessKey, secret, err := parseEnvURLStr(s3ConnectionString)
	if err != nil {
		logger.Panic("Could not parse s3 connection string: %s [%v]", s3ConnectionString, err)
	}

	u.deploymentOwnerUserName = deploymentOwnerUserName
	u.deploymentName = deploymentName
	u.Host = host.Host
	u.accessKeyID = accessKey
	u.secretAccessKey = secret
	u.useSSL = host.Scheme == "https"

	return u
}

// parse url usually obtained from env.
func parseEnvURL(envURL string) (*url.URL, string, string, error) {
	u, e := url.Parse(envURL)
	if e != nil {
		return nil, "", "", fmt.Errorf("S3 Endpoint url invalid [%s]", envURL)
	}

	var accessKey, secretKey string
	// Check if username:password is provided in URL, with no
	// access keys or secret we proceed and perform anonymous
	// requests.
	if u.User != nil {
		accessKey = u.User.Username()
		secretKey, _ = u.User.Password()
	}

	// Look for if URL has invalid values and return error.
	if !((u.Scheme == "http" || u.Scheme == "https") &&
		(u.Path == "/" || u.Path == "") && u.Opaque == "" &&
		!u.ForceQuery && u.RawQuery == "" && u.Fragment == "") {
		return nil, "", "", fmt.Errorf("S3 Endpoint url invalid [%s]", u.String())
	}

	// Now that we have validated the URL to be in expected style.
	u.User = nil

	return u, accessKey, secretKey, nil
}

// parse url usually obtained from env.
func parseEnvURLStr(envURL string) (*url.URL, string, string, error) {
	var envURLStr string
	u, accessKey, secretKey, err := parseEnvURL(envURL)
	if err != nil {
		// url parsing can fail when accessKey/secretKey contains non url encoded values
		// such as #. Strip accessKey/secretKey from envURL and parse again.
		re := regexp.MustCompile("^(https?://)(.*?):(.*?)@(.*?)$")
		res := re.FindAllStringSubmatch(envURL, -1)
		// regex will return full match, scheme, accessKey, secretKey and endpoint:port as
		// captured groups.
		if res == nil || len(res[0]) != 5 {
			return nil, "", "", err
		}
		for k, v := range res[0] {
			if k == 2 {
				accessKey = v
			}
			if k == 3 {
				secretKey = v
			}
			if k == 1 || k == 4 {
				envURLStr = fmt.Sprintf("%s%s", envURLStr, v)
			}
		}
		u, _, _, err = parseEnvURL(envURLStr)
		if err != nil {
			return nil, "", "", err
		}
	}
	// Check if username:password is provided in URL, with no
	// access keys or secret we proceed and perform anonymous
	// requests.
	if u.User != nil {
		accessKey = u.User.Username()
		secretKey, _ = u.User.Password()
	}
	return u, accessKey, secretKey, nil
}
