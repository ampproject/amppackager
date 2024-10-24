package clients

import (
	"fmt"

	clientservices "github.com/selectel/go-selvpcclient/v3/selvpcclient/clients/services"
)

const (
	ResellServiceType = "resell"
	ResellAPIVersion  = "v2"
)

// ResellClient resell client with X-Auth-Token authorization.
type ResellClient struct {
	Requests *clientservices.RequestService
	catalog  *clientservices.CatalogService
	region   string
}

func NewResellClient(
	requestService *clientservices.RequestService,
	catalogService *clientservices.CatalogService,
	region string,
) *ResellClient {
	return &ResellClient{
		Requests: requestService,
		catalog:  catalogService,
		region:   region,
	}
}

// GetEndpoint - returns service url.
func (c *ResellClient) GetEndpoint() (string, error) {
	endpoint, err := c.catalog.GetEndpoint(ResellServiceType, c.region)
	if err != nil {
		return "", fmt.Errorf("failed to resolve endpoint for %s, err: %w", ResellServiceType, err)
	}

	url := fmt.Sprintf("%s/%s", endpoint.URL, ResellAPIVersion)

	return url, nil
}
