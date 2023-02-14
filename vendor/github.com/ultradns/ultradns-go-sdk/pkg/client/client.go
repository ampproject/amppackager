package client

import (
	"context"
	"strings"
	"time"

	"github.com/ultradns/ultradns-go-sdk/internal/token"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
	"golang.org/x/oauth2"
)

const ctxTimeout = 1

func NewClient(config Config) (client *Client, err error) {
	client, err = validateClientConfig(config)

	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(ctxTimeout)*time.Minute)
	_ = cancelFunc
	tokenSource := token.TokenSource{
		Ctx:      ctx,
		BaseURL:  client.baseURL,
		Username: config.Username,
		Password: config.Password,
	}
	client.httpClient = oauth2.NewClient(ctx, oauth2.ReuseTokenSource(nil, &tokenSource))

	return
}

func validateClientConfig(config Config) (*Client, error) {
	if ok := validateParameter(config.Username); !ok {
		return nil, errors.ValidationError("username")
	}

	if ok := validateParameter(config.Password); !ok {
		return nil, errors.ValidationError("password")
	}

	if ok := validateParameter(config.HostURL); !ok {
		return nil, errors.ValidationError("host url")
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
