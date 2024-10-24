package client

import (
	"context"
	"os"
	"strings"

	"github.com/ultradns/ultradns-go-sdk/internal/token"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
	"golang.org/x/oauth2"
)

func NewClient(config Config) (client *Client, err error) {
	client, err = validateClientConfig(&config)

	if err != nil {
		return nil, err
	}

	tokenSource := token.TokenSource{
		BaseURL:  client.baseURL,
		Username: config.Username,
		Password: config.Password,
	}

	client.httpClient = oauth2.NewClient(context.Background(), oauth2.ReuseTokenSource(nil, &tokenSource))

	return
}

func validateClientConfig(config *Config) (*Client, error) {
	config.checkEnvConfig()
	errStr := ""

	if ok := validateParameter(config.Username); !ok {
		errStr += " username,"
	}

	if ok := validateParameter(config.Password); !ok {
		errStr += " password,"
	}

	if ok := validateParameter(config.HostURL); !ok {
		errStr += " host url"
	}

	if errStr != "" {
		return nil, errors.ValidationError(strings.TrimSuffix(errStr, ","))
	}

	hostURL := strings.TrimSuffix(config.HostURL, "/")
	client := &Client{
		baseURL:   hostURL,
		userAgent: config.UserAgent,
	}

	return client, nil
}

func validateParameter(value string) bool {
	return value != ""
}

func (c *Config) checkEnvConfig() {
	if c.Username == "" {
		c.Username = os.Getenv("ULTRADNS_USERNAME")
	}

	if c.Password == "" {
		c.Password = os.Getenv("ULTRADNS_PASSWORD")
	}

	if c.HostURL == "" {
		c.HostURL = os.Getenv("ULTRADNS_HOST_URL")
	}
}
