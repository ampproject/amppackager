package clientservices

import (
	"errors"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
)

const publicInterface = string(gophercloud.AvailabilityPublic)

var (
	ErrEndpointNotFound    = errors.New("endpoint not found")
	ErrEndpointsNotFound   = errors.New("endpoint not found")
	ErrServiceTypeNotFound = errors.New("service type not found")
)

type CatalogService struct {
	serviceClient *gophercloud.ServiceClient
	catalog       *tokens.ServiceCatalog
}

func NewCatalogService(serviceClient *gophercloud.ServiceClient) (*CatalogService, error) {
	service := &CatalogService{serviceClient: serviceClient}

	// Cache warming.
	_, err := service.GetCatalog()
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoints catalog from identity service, err: %w", err)
	}

	return service, nil
}

// GetEndpoint - returns endpoint url from Keystone catalog for ServiceType in RegionName.
// ServiceType is type of service (for example: compute, storage, quota-manager), not their names (nova, cinder).
func (cs *CatalogService) GetEndpoint(serviceType, regionName string) (tokens.Endpoint, error) {
	errMsg := "failed to get endpoint for %s in %s, err: %w"

	endpoints, err := cs.findServiceTypeEndpoints(serviceType)
	if err != nil {
		return tokens.Endpoint{}, fmt.Errorf(errMsg, serviceType, regionName, err)
	}

	endpoint, err := cs.findRegionalEndpoint(endpoints, regionName)
	if err != nil {
		return tokens.Endpoint{}, fmt.Errorf(errMsg, serviceType, regionName, err)
	}

	return endpoint, nil
}

// GetEndpoints - returns endpoints from Keystone catalog for ServiceType from all regions.
func (cs *CatalogService) GetEndpoints(serviceType string) ([]tokens.Endpoint, error) {
	endpoints, err := cs.findServiceTypeEndpoints(serviceType)
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoints for %s, err: %w", serviceType, err)
	}

	if len(endpoints) == 0 {
		return nil, ErrEndpointsNotFound
	}

	return endpoints, nil
}

// GetCatalog - returns endpoints catalog from Keystone or cache.
func (cs *CatalogService) GetCatalog() (*tokens.ServiceCatalog, error) {
	if cs.catalog != nil {
		return cs.catalog, nil
	}

	catalog, err := tokens.Get(cs.serviceClient, cs.serviceClient.Token()).ExtractServiceCatalog()
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoints catalog, err: %w", err)
	}

	cs.catalog = catalog

	return catalog, nil
}

// findServiceTypeEndpoints - returns all public endpoints for ServiceType from endpoints catalog.
func (cs *CatalogService) findServiceTypeEndpoints(serviceType string) ([]tokens.Endpoint, error) {
	for _, service := range cs.catalog.Entries {
		if service.Type == serviceType {
			return cs.findPublicEndpoints(service.Endpoints), nil
		}
	}

	return nil, ErrServiceTypeNotFound
}

// findPublicEndpoints - returns all public endpoints from input endpoints slice.
func (cs *CatalogService) findPublicEndpoints(endpoints []tokens.Endpoint) []tokens.Endpoint {
	publicEndpoints := make([]tokens.Endpoint, 0)

	for _, endpoint := range endpoints {
		if endpoint.Interface == publicInterface {
			publicEndpoints = append(publicEndpoints, endpoint)
		}
	}

	return publicEndpoints
}

// findRegionalEndpoint - returns single public endpoint for service in the specified region.
func (cs *CatalogService) findRegionalEndpoint(endpoints []tokens.Endpoint, regionName string) (tokens.Endpoint, error) {
	for _, endpoint := range endpoints {
		if endpoint.Interface == publicInterface && endpoint.RegionID == regionName {
			return endpoint, nil
		}
	}

	return tokens.Endpoint{}, ErrEndpointNotFound
}
