package clients

import (
	"fmt"

	clientservices "github.com/selectel/go-selvpcclient/v3/selvpcclient/clients/services"
)

const (
	QuotaManagerServiceType = "quota-manager"
)

// QuotaManagerClient quota-manager client with X-Auth-Token authorization.
type QuotaManagerClient struct {
	Requests *clientservices.RequestService
	catalog  *clientservices.CatalogService
}

func NewQuotaManagerClient(
	requestService *clientservices.RequestService,
	catalogService *clientservices.CatalogService,
) *QuotaManagerClient {
	return &QuotaManagerClient{
		Requests: requestService,
		catalog:  catalogService,
	}
}

// GetEndpoint - returns service url in specific region.
func (c *QuotaManagerClient) GetEndpoint(regionName string) (string, error) {
	endpoint, err := c.catalog.GetEndpoint(QuotaManagerServiceType, regionName)
	if err != nil {
		return "", fmt.Errorf(
			"failed to resolve endpoint for %s in %s, err: %w", QuotaManagerServiceType, regionName, err,
		)
	}

	return endpoint.URL, nil
}
