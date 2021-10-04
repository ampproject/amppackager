package v2

import (
	"context"
	"errors"
	"net"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	papi "github.com/exoscale/egoscale/v2/internal/public-api"
)

// NetworkLoadBalancerServerStatus represents a Network Load Balancer service target server status.
type NetworkLoadBalancerServerStatus struct {
	InstanceIP *net.IP
	Status     *string
}

func nlbServerStatusFromAPI(st *papi.LoadBalancerServerStatus) *NetworkLoadBalancerServerStatus {
	return &NetworkLoadBalancerServerStatus{
		InstanceIP: func() (v *net.IP) {
			if st.PublicIp != nil {
				ip := net.ParseIP(*st.PublicIp)
				v = &ip
			}
			return
		}(),
		Status: (*string)(st.Status),
	}
}

// NetworkLoadBalancerServiceHealthcheck represents a Network Load Balancer service healthcheck.
type NetworkLoadBalancerServiceHealthcheck struct {
	Interval *time.Duration `req-for:"create,update"`
	Mode     *string        `req-for:"create,update"`
	Port     *uint16        `req-for:"create,update"`
	Retries  *int64
	TLSSNI   *string
	Timeout  *time.Duration
	URI      *string
}

// NetworkLoadBalancerService represents a Network Load Balancer service.
type NetworkLoadBalancerService struct {
	Description       *string
	Healthcheck       *NetworkLoadBalancerServiceHealthcheck `req-for:"create"`
	HealthcheckStatus []*NetworkLoadBalancerServerStatus
	ID                *string `req-for:"update,delete"`
	InstancePoolID    *string `req-for:"create"`
	Name              *string `req-for:"create"`
	Port              *uint16 `req-for:"create"`
	Protocol          *string `req-for:"create"`
	State             *string
	Strategy          *string `req-for:"create"`
	TargetPort        *uint16 `req-for:"create"`
}

func nlbServiceFromAPI(svc *papi.LoadBalancerService) *NetworkLoadBalancerService {
	var (
		port       = uint16(*svc.Port)
		targetPort = uint16(*svc.TargetPort)
		hcPort     = uint16(*svc.Healthcheck.Port)
		hcInterval = time.Duration(*svc.Healthcheck.Interval) * time.Second
		hcTimeout  = time.Duration(*svc.Healthcheck.Timeout) * time.Second
	)

	return &NetworkLoadBalancerService{
		Description: svc.Description,
		Healthcheck: &NetworkLoadBalancerServiceHealthcheck{
			Interval: &hcInterval,
			Mode:     (*string)(svc.Healthcheck.Mode),
			Port:     &hcPort,
			Retries:  svc.Healthcheck.Retries,
			TLSSNI:   svc.Healthcheck.TlsSni,
			Timeout:  &hcTimeout,
			URI:      svc.Healthcheck.Uri,
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
		ID:             svc.Id,
		InstancePoolID: svc.InstancePool.Id,
		Name:           svc.Name,
		Port:           &port,
		Protocol:       (*string)(svc.Protocol),
		Strategy:       (*string)(svc.Strategy),
		TargetPort:     &targetPort,
		State:          (*string)(svc.State),
	}
}

// NetworkLoadBalancer represents a Network Load Balancer instance.
type NetworkLoadBalancer struct {
	CreatedAt   *time.Time
	Description *string
	ID          *string `req-for:"update"`
	IPAddress   *net.IP
	Labels      *map[string]string
	Name        *string `req-for:"create"`
	Services    []*NetworkLoadBalancerService
	State       *string

	c    *Client
	zone string
}

func nlbFromAPI(client *Client, zone string, nlb *papi.LoadBalancer) *NetworkLoadBalancer {
	return &NetworkLoadBalancer{
		CreatedAt:   nlb.CreatedAt,
		Description: nlb.Description,
		ID:          nlb.Id,
		IPAddress: func() (v *net.IP) {
			if nlb.Ip != nil {
				ip := net.ParseIP(*nlb.Ip)
				v = &ip
			}
			return
		}(),
		Labels: func() (v *map[string]string) {
			if nlb.Labels != nil && len(nlb.Labels.AdditionalProperties) > 0 {
				v = &nlb.Labels.AdditionalProperties
			}
			return
		}(),
		Name: nlb.Name,
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
		State: (*string)(nlb.State),

		c:    client,
		zone: zone,
	}
}

// AddService adds a service to the Network Load Balancer instance.
func (nlb *NetworkLoadBalancer) AddService(
	ctx context.Context,
	svc *NetworkLoadBalancerService,
) (*NetworkLoadBalancerService, error) {
	if err := validateOperationParams(svc, "create"); err != nil {
		return nil, err
	}
	if err := validateOperationParams(svc.Healthcheck, "create"); err != nil {
		return nil, err
	}

	var (
		port                = int64(*svc.Port)
		targetPort          = int64(*svc.TargetPort)
		healthcheckPort     = int64(*svc.Healthcheck.Port)
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
		services[*svc.ID] = struct{}{}
	}

	resp, err := nlb.c.AddServiceToLoadBalancerWithResponse(
		apiv2.WithZone(ctx, nlb.zone),
		*nlb.ID,
		papi.AddServiceToLoadBalancerJSONRequestBody{
			Description: svc.Description,
			Healthcheck: papi.LoadBalancerServiceHealthcheck{
				Interval: &healthcheckInterval,
				Mode:     (*papi.LoadBalancerServiceHealthcheckMode)(svc.Healthcheck.Mode),
				Port:     &healthcheckPort,
				Retries:  svc.Healthcheck.Retries,
				Timeout:  &healthcheckTimeout,
				TlsSni:   svc.Healthcheck.TLSSNI,
				Uri:      svc.Healthcheck.URI,
			},
			InstancePool: papi.InstancePool{Id: svc.InstancePoolID},
			Name:         *svc.Name,
			Port:         port,
			Protocol:     papi.AddServiceToLoadBalancerJSONBodyProtocol(*svc.Protocol),
			Strategy:     papi.AddServiceToLoadBalancerJSONBodyStrategy(*svc.Strategy),
			TargetPort:   targetPort,
		})
	if err != nil {
		return nil, err
	}

	res, err := papi.NewPoller().
		WithTimeout(nlb.c.timeout).
		WithInterval(nlb.c.pollInterval).
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
		if _, ok := services[*s.ID]; !ok && *s.Name == *svc.Name {
			return s, nil
		}
	}

	return nil, errors.New("unable to identify the service created")
}

// UpdateService updates the specified Network Load Balancer service.
func (nlb *NetworkLoadBalancer) UpdateService(ctx context.Context, svc *NetworkLoadBalancerService) error {
	if err := validateOperationParams(svc, "update"); err != nil {
		return err
	}
	if svc.Healthcheck != nil {
		if err := validateOperationParams(svc.Healthcheck, "update"); err != nil {
			return err
		}
	}

	resp, err := nlb.c.UpdateLoadBalancerServiceWithResponse(
		apiv2.WithZone(ctx, nlb.zone),
		*nlb.ID,
		*svc.ID,
		papi.UpdateLoadBalancerServiceJSONRequestBody{
			Description: svc.Description,
			Healthcheck: &papi.LoadBalancerServiceHealthcheck{
				Interval: func() (v *int64) {
					if svc.Healthcheck.Interval != nil {
						interval := int64(svc.Healthcheck.Interval.Seconds())
						v = &interval
					}
					return
				}(),
				Mode: (*papi.LoadBalancerServiceHealthcheckMode)(svc.Healthcheck.Mode),
				Port: func() (v *int64) {
					if svc.Healthcheck.Port != nil {
						port := int64(*svc.Healthcheck.Port)
						v = &port
					}
					return
				}(),
				Retries: svc.Healthcheck.Retries,
				Timeout: func() (v *int64) {
					if svc.Healthcheck.Timeout != nil {
						interval := int64(svc.Healthcheck.Timeout.Seconds())
						v = &interval
					}
					return
				}(),
				TlsSni: svc.Healthcheck.TLSSNI,
				Uri:    svc.Healthcheck.URI,
			},
			Name: svc.Name,
			Port: func() (v *int64) {
				if svc.Port != nil {
					port := int64(*svc.Port)
					v = &port
				}
				return
			}(),
			Protocol: (*papi.UpdateLoadBalancerServiceJSONBodyProtocol)(svc.Protocol),
			Strategy: (*papi.UpdateLoadBalancerServiceJSONBodyStrategy)(svc.Strategy),
			TargetPort: func() (v *int64) {
				if svc.TargetPort != nil {
					port := int64(*svc.TargetPort)
					v = &port
				}
				return
			}(),
		})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(nlb.c.timeout).
		WithInterval(nlb.c.pollInterval).
		Poll(ctx, nlb.c.OperationPoller(nlb.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// DeleteService deletes the specified service from the Network Load Balancer instance.
func (nlb *NetworkLoadBalancer) DeleteService(ctx context.Context, svc *NetworkLoadBalancerService) error {
	if err := validateOperationParams(svc, "delete"); err != nil {
		return err
	}

	resp, err := nlb.c.DeleteLoadBalancerServiceWithResponse(
		apiv2.WithZone(ctx, nlb.zone),
		*nlb.ID,
		*svc.ID,
	)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(nlb.c.timeout).
		WithInterval(nlb.c.pollInterval).
		Poll(ctx, nlb.c.OperationPoller(nlb.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// CreateNetworkLoadBalancer creates a Network Load Balancer instance in the specified zone.
func (c *Client) CreateNetworkLoadBalancer(
	ctx context.Context,
	zone string,
	nlb *NetworkLoadBalancer,
) (*NetworkLoadBalancer, error) {
	if err := validateOperationParams(nlb, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreateLoadBalancerWithResponse(
		apiv2.WithZone(ctx, zone),
		papi.CreateLoadBalancerJSONRequestBody{
			Description: nlb.Description,
			Labels: func() (v *papi.Labels) {
				if nlb.Labels != nil {
					v = &papi.Labels{AdditionalProperties: *nlb.Labels}
				}
				return
			}(),
			Name: *nlb.Name,
		})
	if err != nil {
		return nil, err
	}

	res, err := papi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
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
			list = append(list, nlbFromAPI(c, zone, &(*resp.JSON200.LoadBalancers)[i]))
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

	return nlbFromAPI(c, zone, resp.JSON200), nil
}

// FindNetworkLoadBalancer attempts to find a Network Load Balancer by name or ID in the specified zone.
func (c *Client) FindNetworkLoadBalancer(ctx context.Context, zone, v string) (*NetworkLoadBalancer, error) {
	res, err := c.ListNetworkLoadBalancers(ctx, zone)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		if *r.ID == v || *r.Name == v {
			return c.GetNetworkLoadBalancer(ctx, zone, *r.ID)
		}
	}

	return nil, apiv2.ErrNotFound
}

// UpdateNetworkLoadBalancer updates the specified Network Load Balancer instance in the specified zone.
func (c *Client) UpdateNetworkLoadBalancer(ctx context.Context, zone string, nlb *NetworkLoadBalancer) error {
	if err := validateOperationParams(nlb, "update"); err != nil {
		return err
	}

	resp, err := c.UpdateLoadBalancerWithResponse(
		apiv2.WithZone(ctx, zone),
		*nlb.ID,
		papi.UpdateLoadBalancerJSONRequestBody{
			Description: nlb.Description,
			Labels: func() (v *papi.Labels) {
				if nlb.Labels != nil {
					v = &papi.Labels{AdditionalProperties: *nlb.Labels}
				}
				return
			}(),
			Name: nlb.Name,
		})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// DeleteNetworkLoadBalancer deletes the specified Network Load Balancer instance in the specified zone.
func (c *Client) DeleteNetworkLoadBalancer(ctx context.Context, zone, id string) error {
	resp, err := c.DeleteLoadBalancerWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}
