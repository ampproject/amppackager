package v2

import (
	"context"
	"errors"
	"net"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// NetworkLoadBalancerServerStatus represents a Network Load Balancer service target server status.
type NetworkLoadBalancerServerStatus struct {
	InstanceIP *net.IP
	Status     *string
}

func nlbServerStatusFromAPI(st *oapi.LoadBalancerServerStatus) *NetworkLoadBalancerServerStatus {
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

func nlbServiceFromAPI(svc *oapi.LoadBalancerService) *NetworkLoadBalancerService {
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

// CreateNetworkLoadBalancerService creates a Network Load Balancer service.
func (c *Client) CreateNetworkLoadBalancerService(
	ctx context.Context,
	zone string,
	nlb *NetworkLoadBalancer,
	service *NetworkLoadBalancerService,
) (*NetworkLoadBalancerService, error) {
	if err := validateOperationParams(service, "create"); err != nil {
		return nil, err
	}
	if err := validateOperationParams(service.Healthcheck, "create"); err != nil {
		return nil, err
	}

	var (
		port                = int64(*service.Port)
		targetPort          = int64(*service.TargetPort)
		healthcheckPort     = int64(*service.Healthcheck.Port)
		healthcheckInterval = int64(service.Healthcheck.Interval.Seconds())
		healthcheckTimeout  = int64(service.Healthcheck.Timeout.Seconds())
	)

	// The API doesn't return the NLB service created directly, so in order to return a
	// *NetworkLoadBalancerService corresponding to the new service we have to manually
	// compare the list of services on the NLB before and after the service creation,
	// and identify the service that wasn't there before.
	// Note: in case of multiple services creation in parallel this technique is subject
	// to race condition as we could return an unrelated service. To prevent this, we
	// also compare the name of the new service to the name specified in the service
	// parameter.
	services := make(map[string]struct{})
	for _, svc := range nlb.Services {
		services[*svc.ID] = struct{}{}
	}

	resp, err := c.AddServiceToLoadBalancerWithResponse(
		apiv2.WithZone(ctx, zone),
		*nlb.ID,
		oapi.AddServiceToLoadBalancerJSONRequestBody{
			Description: service.Description,
			Healthcheck: oapi.LoadBalancerServiceHealthcheck{
				Interval: &healthcheckInterval,
				Mode:     (*oapi.LoadBalancerServiceHealthcheckMode)(service.Healthcheck.Mode),
				Port:     &healthcheckPort,
				Retries:  service.Healthcheck.Retries,
				Timeout:  &healthcheckTimeout,
				TlsSni:   service.Healthcheck.TLSSNI,
				Uri:      service.Healthcheck.URI,
			},
			InstancePool: oapi.InstancePool{Id: service.InstancePoolID},
			Name:         *service.Name,
			Port:         port,
			Protocol:     oapi.AddServiceToLoadBalancerJSONBodyProtocol(*service.Protocol),
			Strategy:     oapi.AddServiceToLoadBalancerJSONBodyStrategy(*service.Strategy),
			TargetPort:   targetPort,
		})
	if err != nil {
		return nil, err
	}

	res, err := oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	nlbUpdated, err := c.GetNetworkLoadBalancer(ctx, zone, *res.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
	if err != nil {
		return nil, err
	}

	// Look for an unknown service: if we find one we hope it's the one we've just created.
	for _, s := range nlbUpdated.Services {
		if _, ok := services[*s.ID]; !ok && *s.Name == *service.Name {
			return s, nil
		}
	}

	return nil, errors.New("unable to identify the service created")
}

// DeleteNetworkLoadBalancerService deletes a Network Load Balancer service.
func (c *Client) DeleteNetworkLoadBalancerService(
	ctx context.Context,
	zone string,
	nlb *NetworkLoadBalancer,
	service *NetworkLoadBalancerService,
) error {
	if err := validateOperationParams(nlb, "delete"); err != nil {
		return err
	}
	if err := validateOperationParams(service, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteLoadBalancerServiceWithResponse(
		apiv2.WithZone(ctx, zone),
		*nlb.ID,
		*service.ID,
	)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// UpdateNetworkLoadBalancerService updates a Network Load Balancer service.
func (c *Client) UpdateNetworkLoadBalancerService(
	ctx context.Context,
	zone string,
	nlb *NetworkLoadBalancer,
	service *NetworkLoadBalancerService,
) error {
	if err := validateOperationParams(service, "update"); err != nil {
		return err
	}
	if service.Healthcheck != nil {
		if err := validateOperationParams(service.Healthcheck, "update"); err != nil {
			return err
		}
	}

	resp, err := c.UpdateLoadBalancerServiceWithResponse(
		apiv2.WithZone(ctx, zone),
		*nlb.ID,
		*service.ID,
		oapi.UpdateLoadBalancerServiceJSONRequestBody{
			Description: oapi.NilableString(service.Description),
			Healthcheck: &oapi.LoadBalancerServiceHealthcheck{
				Interval: func() (v *int64) {
					if service.Healthcheck.Interval != nil {
						interval := int64(service.Healthcheck.Interval.Seconds())
						v = &interval
					}
					return
				}(),
				Mode: (*oapi.LoadBalancerServiceHealthcheckMode)(service.Healthcheck.Mode),
				Port: func() (v *int64) {
					if service.Healthcheck.Port != nil {
						port := int64(*service.Healthcheck.Port)
						v = &port
					}
					return
				}(),
				Retries: service.Healthcheck.Retries,
				Timeout: func() (v *int64) {
					if service.Healthcheck.Timeout != nil {
						interval := int64(service.Healthcheck.Timeout.Seconds())
						v = &interval
					}
					return
				}(),
				TlsSni: service.Healthcheck.TLSSNI,
				Uri:    service.Healthcheck.URI,
			},
			Name: service.Name,
			Port: func() (v *int64) {
				if service.Port != nil {
					port := int64(*service.Port)
					v = &port
				}
				return
			}(),
			Protocol: (*oapi.UpdateLoadBalancerServiceJSONBodyProtocol)(service.Protocol),
			Strategy: (*oapi.UpdateLoadBalancerServiceJSONBodyStrategy)(service.Strategy),
			TargetPort: func() (v *int64) {
				if service.TargetPort != nil {
					port := int64(*service.TargetPort)
					v = &port
				}
				return
			}(),
		})
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}
