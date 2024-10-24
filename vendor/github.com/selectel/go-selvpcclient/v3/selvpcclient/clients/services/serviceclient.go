package clientservices

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

type ServiceClientOptions struct {
	// The name of the domain in which the token will be issued.
	DomainName string

	// Credentials to auth with.
	Username string
	Password string

	// Optional field for setting project scope.
	ProjectID string

	// Optional field. The name of the domain where the user resides (Identity v3).
	UserDomainName string

	// Field for setting Identity endpoint.
	AuthURL string

	// Field for setting location for endpoints like ResellAPI or Keystone.
	AuthRegion string

	// Optional field.
	HTTPClient *http.Client

	// Optional field.
	UserAgent string
}

func NewServiceClient(options *ServiceClientOptions) (*gophercloud.ServiceClient, error) {
	// UserDomainName field to specify the domain name where the user is located.
	// If this field is not specified, then we will think that the token will be
	// issued in the same domain where the user is located.
	if options.UserDomainName == "" {
		options.UserDomainName = options.DomainName
	}

	authOptions := gophercloud.AuthOptions{
		AllowReauth:      true,
		IdentityEndpoint: options.AuthURL,
		Username:         options.Username,
		Password:         options.Password,
		DomainName:       options.UserDomainName,
		Scope: &gophercloud.AuthScope{
			ProjectID: options.ProjectID,
		},
	}

	// If project scope is not set, we use domain scope.
	if authOptions.Scope.ProjectID == "" {
		authOptions.Scope.DomainName = options.DomainName
	}

	authProvider, err := openstack.AuthenticatedClient(authOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth provider, err: %w", err)
	}

	serviceClient, err := openstack.NewIdentityV3(authProvider, gophercloud.EndpointOpts{
		Availability: gophercloud.AvailabilityPublic,
		Region:       options.AuthRegion,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create service client, err: %w", err)
	}

	httpClient := options.HTTPClient
	if httpClient == nil {
		httpClient = NewHTTPClient()
	}
	serviceClient.HTTPClient = *httpClient

	if options.UserAgent != "" {
		userAgent := gophercloud.UserAgent{}
		userAgent.Prepend(options.UserAgent)
		serviceClient.UserAgent = userAgent
	}

	return serviceClient, nil
}

// ---------------------------------------------------------------------------------------------------------------------

const (
	// httpTimeout represents the default timeout (in seconds) for HTTP
	// requests.
	httpTimeout = 120

	// dialTimeout represents the default timeout (in seconds) for HTTP
	// connection establishments.
	dialTimeout = 60

	// keepaliveTimeout represents the default keep-alive period for an active
	// network connection.
	keepaliveTimeout = 60

	// maxIdleConns represents the maximum number of idle (keep-alive)
	// connections.
	maxIdleConns = 100

	// idleConnTimeout represents the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing itself.
	idleConnTimeout = 100

	// tlsHandshakeTimeout represents the default timeout (in seconds)
	// for TLS handshake.
	tlsHandshakeTimeout = 60

	// expectContinueTimeout represents the default amount of time to
	// wait for a server's first response headers.
	expectContinueTimeout = 1
)

func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * httpTimeout,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   dialTimeout * time.Second,
				KeepAlive: keepaliveTimeout * time.Second,
			}).DialContext,
			MaxIdleConns:          maxIdleConns,
			IdleConnTimeout:       idleConnTimeout * time.Second,
			TLSHandshakeTimeout:   tlsHandshakeTimeout * time.Second,
			ExpectContinueTimeout: expectContinueTimeout * time.Second,
		},
	}
}
