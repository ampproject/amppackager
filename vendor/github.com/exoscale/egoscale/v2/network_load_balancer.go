package v2

import (
	"context"
	"errors"
	"net"
	"strings"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	papi "github.com/exoscale/egoscale/v2/internal/public-api"
)

// NetworkLoadBalancerServerStatus represents a Network Load Balancer service target server status.
type NetworkLoadBalancerServerStatus struct {
	InstanceIP net.IP
	Status     string
}

func nlbServerStatusFromAPI(st *papi.LoadBalancerServerStatus) *NetworkLoadBalancerServerStatus {
	return &NetworkLoadBalancerServerStatus{
		InstanceIP: net.ParseIP(papi.OptionalString(st.PublicIp)),
		Status:     papi.OptionalString(st.Status),
	}
}

// NetworkLoadBalancerServiceHealthcheck represents a Network Load Balancer service healthcheck.
type NetworkLoadBalancerServiceHealthcheck struct {
	Interval time.Duration
	Mode     string
	Port     uint16
	Retries  int64
	TLSSNI   string
	Timeout  time.Duration
	URI      string
}

// NetworkLoadBalancerService represents a Network Load Balancer service.
type NetworkLoadBalancerService struct {
	Description       string
	Healthcheck       NetworkLoadBalancerServiceHealthcheck
	HealthcheckStatus []*NetworkLoadBalancerServerStatus
	ID                string
	InstancePoolID    string
	Name              string
	Port              uint16
	Protocol          string
	State             string
	Strategy          string
	TargetPort        uint16
}

func nlbServiceFromAPI(svc *papi.LoadBalancerService) *NetworkLoadBalancerService {
	return &NetworkLoadBalancerService{
		Description: papi.OptionalString(svc.Description),
		Healthcheck: NetworkLoadBalancerServiceHealthcheck{
			Interval: time.Duration(svc.Healthcheck.Interval) * time.Second,
			Mode:     svc.Healthcheck.Mode,
			Port:     uint16(svc.Healthcheck.Port),
			Retries:  svc.Healthcheck.Retries,
			TLSSNI:   papi.OptionalString(svc.Healthcheck.TlsSni),
			Timeout:  time.Duration(svc.Healthcheck.Timeout) * time.Second,
			URI:      papi.OptionalString(svc.Healthcheck.Uri),
		},
		HealthcheckStatus: func() []*NetworkLoadBalancerServerStatus {
			statuses := make([]*NetworkLoadBalancerServerStatus, 0)
			if svc.HealthcheckStatus != nil {
				for _, st := range *svc.HealthcheckStatus {
					st := st
					statuses = append(statuses, nlbServerStatusFromAPI(&st))
				}
			}
			return statuses
		}(),
		ID:             papi.OptionalString(svc.Id),
		InstancePoolID: papi.OptionalString(svc.InstancePool.Id),
		Name:           papi.OptionalString(svc.Name),
		Port:           uint16(papi.OptionalInt64(svc.Port)),
		Protocol:       papi.OptionalString(svc.Protocol),
		Strategy:       papi.OptionalString(svc.Strategy),
		TargetPort:     uint16(papi.OptionalInt64(svc.TargetPort)),
		State:          papi.OptionalString(svc.State),
	}
}

// NetworkLoadBalancer represents a Network Load Balancer instance.
type NetworkLoadBalancer struct {
	CreatedAt   time.Time
	Description string
	ID          string
	IPAddress   net.IP
	Name        string
	Services    []*NetworkLoadBalancerService
	State       string

	c    *Client
	zone string
}

func nlbFromAPI(nlb *papi.LoadBalancer) *NetworkLoadBalancer {
	return &NetworkLoadBalancer{
		CreatedAt:   *nlb.CreatedAt,
		Description: papi.OptionalString(nlb.Description),
		ID:          papi.OptionalString(nlb.Id),
		IPAddress:   net.ParseIP(papi.OptionalString(nlb.Ip)),
		Name:        papi.OptionalString(nlb.Name),
		Services: func() []*NetworkLoadBalancerService {
			services := make([]*NetworkLoadBalancerService, 0)
			if nlb.Services != nil {
				for _, svc := range *nlb.Services {
					svc := svc
					services = append(services, nlbServiceFromAPI(&svc))
				}
			}
			return services
		}(),
		State: papi.OptionalString(nlb.State),
	}
}

// AddService adds a service to the Network Load Balancer instance.
func (nlb *NetworkLoadBalancer) AddService(ctx context.Context,
	svc *NetworkLoadBalancerService) (*NetworkLoadBalancerService, error) {
	var (
		port                = int64(svc.Port)
		targetPort          = int64(svc.TargetPort)
		healthcheckPort     = int64(svc.Healthcheck.Port)
		healthcheckInterval = int64(svc.Healthcheck.Interval.Seconds())
		healthcheckTimeout  = int64(svc.Healthcheck.Timeout.Seconds())
	)

	// The API doesn't return the NLB service created directly, so in order to return a
	// *NetworkLoadBalancerService corresponding to the new service we have to manually
	// compare the list of services on the NLB instance before and after the service
	// creation, and identify the service that wasn't there before.
	// Note: in case of multiple services creation in parallel this technique is subject
	// to race condition as we could return an unrelated service. To prevent this, we
	// also compare the name of the new service to the name specified in the svc
	// parameter.
	services := make(map[string]struct{})
	for _, svc := range nlb.Services {
		services[svc.ID] = struct{}{}
	}

	resp, err := nlb.c.AddServiceToLoadBalancerWithResponse(
		apiv2.WithZone(ctx, nlb.zone),
		nlb.ID,
		papi.AddServiceToLoadBalancerJSONRequestBody{
			Description: &svc.Description,
			Healthcheck: papi.LoadBalancerServiceHealthcheck{
				Interval: healthcheckInterval,
				Mode:     svc.Healthcheck.Mode,
				Port:     healthcheckPort,
				Retries:  svc.Healthcheck.Retries,
				Timeout:  healthcheckTimeout,
				TlsSni: func() *string {
					if svc.Healthcheck.Mode == "https" && svc.Healthcheck.TLSSNI != "" {
						return &svc.Healthcheck.TLSSNI
					}
					return nil
				}(),
				Uri: func() *string {
					if strings.HasPrefix(svc.Healthcheck.Mode, "http") {
						return &svc.Healthcheck.URI
					}
					return nil
				}(),
			},
			InstancePool: papi.InstancePool{Id: &svc.InstancePoolID},
			Name:         svc.Name,
			Port:         port,
			Protocol:     svc.Protocol,
			Strategy:     svc.Strategy,
			TargetPort:   targetPort,
		})
	if err != nil {
		return nil, err
	}

	res, err := papi.NewPoller().
		WithTimeout(nlb.c.timeout).
		Poll(ctx, nlb.c.OperationPoller(nlb.zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	nlbUpdated, err := nlb.c.GetNetworkLoadBalancer(ctx, nlb.zone, *res.(*papi.Reference).Id)
	if err != nil {
		return nil, err
	}

	// Look for an unknown service: if we find one we hope it's the one we've just created.
	for _, s := range nlbUpdated.Services {
		if _, ok := services[svc.ID]; !ok && s.Name == svc.Name {
			return s, nil
		}
	}

	return nil, errors.New("unable to identify the service created")
}

// UpdateService updates the specified Network Load Balancer service.
func (nlb *NetworkLoadBalancer) UpdateService(ctx context.Context, svc *NetworkLoadBalancerService) error {
	var (
		healthcheckPort     = int64(svc.Healthcheck.Port)
		healthcheckInterval = int64(svc.Healthcheck.Interval.Seconds())
		healthcheckTimeout  = int64(svc.Healthcheck.Timeout.Seconds())
	)

	resp, err := nlb.c.UpdateLoadBalancerServiceWithResponse(
		apiv2.WithZone(ctx, nlb.zone),
		nlb.ID,
		svc.ID,
		papi.UpdateLoadBalancerServiceJSONRequestBody{
			Description: func() *string {
				if svc.Description != "" {
					return &svc.Description
				}
				return nil
			}(),
			Healthcheck: &papi.LoadBalancerServiceHealthcheck{
				Interval: healthcheckInterval,
				Mode:     svc.Healthcheck.Mode,
				Port:     healthcheckPort,
				Retries:  svc.Healthcheck.Retries,
				Timeout:  healthcheckTimeout,
				TlsSni: func() *string {
					if svc.Healthcheck.Mode == "https" && svc.Healthcheck.TLSSNI != "" {
						return &svc.Healthcheck.TLSSNI
					}
					return nil
				}(),
				Uri: func() *string {
					if strings.HasPrefix(svc.Healthcheck.Mode, "http") {
						return &svc.Healthcheck.URI
					}
					return nil
				}(),
			},
			Name: func() *string {
				if svc.Name != "" {
					return &svc.Name
				}
				return nil
			}(),
			Port: func() *int64 {
				if v := svc.Port; v > 0 {
					port := int64(v)
					return &port
				}
				return nil
			}(),
			Protocol: func() *string {
				if svc.Protocol != "" {
					return &svc.Protocol
				}
				return nil
			}(),
			Strategy: func() *string {
				if svc.Strategy != "" {
					return &svc.Strategy
				}
				return nil
			}(),
			TargetPort: func() *int64 {
				if v := svc.TargetPort; v > 0 {
					targetPort := int64(v)
					return &targetPort
				}
				return nil
			}(),
		})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(nlb.c.timeout).
		Poll(ctx, nlb.c.OperationPoller(nlb.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// DeleteService deletes the specified service from the Network Load Balancer instance.
func (nlb *NetworkLoadBalancer) DeleteService(ctx context.Context, svc *NetworkLoadBalancerService) error {
	resp, err := nlb.c.DeleteLoadBalancerServiceWithResponse(
		apiv2.WithZone(ctx, nlb.zone),
		nlb.ID,
		svc.ID,
	)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(nlb.c.timeout).
		Poll(ctx, nlb.c.OperationPoller(nlb.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// CreateNetworkLoadBalancer creates a Network Load Balancer instance in the specified zone.
func (c *Client) CreateNetworkLoadBalancer(ctx context.Context, zone string,
	nlb *NetworkLoadBalancer) (*NetworkLoadBalancer, error) {
	resp, err := c.CreateLoadBalancerWithResponse(
		apiv2.WithZone(ctx, zone),
		papi.CreateLoadBalancerJSONRequestBody{
			Description: &nlb.Description,
			Name:        nlb.Name,
		})
	if err != nil {
		return nil, err
	}

	res, err := papi.NewPoller().
		WithTimeout(c.timeout).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	return c.GetNetworkLoadBalancer(ctx, zone, *res.(*papi.Reference).Id)
}

// ListNetworkLoadBalancers returns the list of existing Network Load Balancers in the
// specified zone.
func (c *Client) ListNetworkLoadBalancers(ctx context.Context, zone string) ([]*NetworkLoadBalancer, error) {
	list := make([]*NetworkLoadBalancer, 0)

	resp, err := c.ListLoadBalancersWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.LoadBalancers != nil {
		for i := range *resp.JSON200.LoadBalancers {
			nlb := nlbFromAPI(&(*resp.JSON200.LoadBalancers)[i])
			nlb.c = c
			nlb.zone = zone

			list = append(list, nlb)
		}
	}

	return list, nil
}

// GetNetworkLoadBalancer returns the Network Load Balancer instance corresponding to the
// specified ID in the specified zone.
func (c *Client) GetNetworkLoadBalancer(ctx context.Context, zone, id string) (*NetworkLoadBalancer, error) {
	resp, err := c.GetLoadBalancerWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	nlb := nlbFromAPI(resp.JSON200)
	nlb.c = c
	nlb.zone = zone

	return nlb, nil
}

// UpdateNetworkLoadBalancer updates the specified Network Load Balancer instance in the specified zone.
func (c *Client) UpdateNetworkLoadBalancer(ctx context.Context, zone string, // nolint:dupl
	nlb *NetworkLoadBalancer) (*NetworkLoadBalancer, error) {
	resp, err := c.UpdateLoadBalancerWithResponse(
		apiv2.WithZone(ctx, zone),
		nlb.ID,
		papi.UpdateLoadBalancerJSONRequestBody{
			Description: func() *string {
				if nlb.Description != "" {
					return &nlb.Description
				}
				return nil
			}(),
			Name: func() *string {
				if nlb.Name != "" {
					return &nlb.Name
				}
				return nil
			}(),
		})
	if err != nil {
		return nil, err
	}

	res, err := papi.NewPoller().
		WithTimeout(c.timeout).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	return c.GetNetworkLoadBalancer(ctx, zone, *res.(*papi.Reference).Id)
}

// DeleteNetworkLoadBalancer deletes the specified Network Load Balancer instance in the specified zone.
func (c *Client) DeleteNetworkLoadBalancer(ctx context.Context, zone, id string) error {
	resp, err := c.DeleteLoadBalancerWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(c.timeout).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}
