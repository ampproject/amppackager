package v2

import (
	"context"
	"net"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// ElasticIPHealthcheck represents an Elastic IP healthcheck.
type ElasticIPHealthcheck struct {
	Interval      *time.Duration
	Mode          *string `req-for:"create,update"`
	Port          *uint16 `req-for:"create,update"`
	StrikesFail   *int64
	StrikesOK     *int64
	TLSSNI        *string
	TLSSkipVerify *bool
	Timeout       *time.Duration
	URI           *string
}

// ElasticIP represents an Elastic IP.
type ElasticIP struct {
	Description   *string
	Healthcheck   *ElasticIPHealthcheck
	ID            *string `req-for:"update,delete"`
	IPAddress     *net.IP
	Labels        *map[string]string
	Zone          *string
	CIDR          *string
	AddressFamily *string
}

func elasticIPFromAPI(e *oapi.ElasticIp, zone string) *ElasticIP {
	ipAddress := net.ParseIP(*e.Ip)

	return &ElasticIP{
		Description: e.Description,
		Healthcheck: func() *ElasticIPHealthcheck {
			if hc := e.Healthcheck; hc != nil {
				port := uint16(hc.Port)
				interval := time.Duration(oapi.OptionalInt64(hc.Interval)) * time.Second
				timeout := time.Duration(oapi.OptionalInt64(hc.Timeout)) * time.Second

				return &ElasticIPHealthcheck{
					Interval:      &interval,
					Mode:          (*string)(&hc.Mode),
					Port:          &port,
					StrikesFail:   hc.StrikesFail,
					StrikesOK:     hc.StrikesOk,
					TLSSNI:        hc.TlsSni,
					TLSSkipVerify: hc.TlsSkipVerify,
					Timeout:       &timeout,
					URI:           hc.Uri,
				}
			}
			return nil
		}(),
		ID:        e.Id,
		IPAddress: &ipAddress,
		Labels: func() (v *map[string]string) {
			if e.Labels != nil && len(e.Labels.AdditionalProperties) > 0 {
				v = &e.Labels.AdditionalProperties
			}
			return
		}(),
		Zone:          &zone,
		CIDR:          e.Cidr,
		AddressFamily: (*string)(e.Addressfamily),
	}
}

// CreateElasticIP creates an Elastic IP.
func (c *Client) CreateElasticIP(ctx context.Context, zone string, elasticIP *ElasticIP) (*ElasticIP, error) {
	if err := validateOperationParams(elasticIP, "create"); err != nil {
		return nil, err
	}
	if elasticIP.Healthcheck != nil {
		if err := validateOperationParams(elasticIP.Healthcheck, "create"); err != nil {
			return nil, err
		}
	}

	var addressFamily *oapi.CreateElasticIpJSONBodyAddressfamily
	if elasticIP.AddressFamily != nil {
		addressFamily = (*oapi.CreateElasticIpJSONBodyAddressfamily)(elasticIP.AddressFamily)
	}

	resp, err := c.CreateElasticIpWithResponse(
		apiv2.WithZone(ctx, zone),
		oapi.CreateElasticIpJSONRequestBody{
			Description:   elasticIP.Description,
			Addressfamily: addressFamily,
			Healthcheck: func() *oapi.ElasticIpHealthcheck {
				if hc := elasticIP.Healthcheck; hc != nil {
					var (
						port     = int64(*hc.Port)
						interval = int64(hc.Interval.Seconds())
						timeout  = int64(hc.Timeout.Seconds())
					)

					return &oapi.ElasticIpHealthcheck{
						Interval:      &interval,
						Mode:          oapi.ElasticIpHealthcheckMode(*hc.Mode),
						Port:          port,
						StrikesFail:   hc.StrikesFail,
						StrikesOk:     hc.StrikesOK,
						Timeout:       &timeout,
						TlsSkipVerify: hc.TLSSkipVerify,
						TlsSni:        hc.TLSSNI,
						Uri:           hc.URI,
					}
				}
				return nil
			}(),
			Labels: func() (v *oapi.Labels) {
				if elasticIP.Labels != nil {
					v = &oapi.Labels{AdditionalProperties: *elasticIP.Labels}
				}
				return
			}(),
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

	return c.GetElasticIP(ctx, zone, *res.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
}

// DeleteElasticIP deletes an Elastic IP.
func (c *Client) DeleteElasticIP(ctx context.Context, zone string, elasticIP *ElasticIP) error {
	if err := validateOperationParams(elasticIP, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteElasticIpWithResponse(apiv2.WithZone(ctx, zone), *elasticIP.ID)
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

// FindElasticIP attempts to find an Elastic IP by IP address or ID.
func (c *Client) FindElasticIP(ctx context.Context, zone, x string) (*ElasticIP, error) {
	res, err := c.ListElasticIPs(ctx, zone)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		if *r.ID == x || r.IPAddress.String() == x {
			return c.GetElasticIP(ctx, zone, *r.ID)
		}
	}

	return nil, apiv2.ErrNotFound
}

// GetElasticIP returns the Elastic IP corresponding to the specified ID.
func (c *Client) GetElasticIP(ctx context.Context, zone, id string) (*ElasticIP, error) {
	resp, err := c.GetElasticIpWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return elasticIPFromAPI(resp.JSON200, zone), nil
}

// ListElasticIPs returns the list of existing Elastic IPs.
func (c *Client) ListElasticIPs(ctx context.Context, zone string) ([]*ElasticIP, error) {
	list := make([]*ElasticIP, 0)

	resp, err := c.ListElasticIpsWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.ElasticIps != nil {
		for i := range *resp.JSON200.ElasticIps {
			list = append(list, elasticIPFromAPI(&(*resp.JSON200.ElasticIps)[i], zone))
		}
	}

	return list, nil
}

// UpdateElasticIP updates an Elastic IP.
func (c *Client) UpdateElasticIP(ctx context.Context, zone string, elasticIP *ElasticIP) error {
	if err := validateOperationParams(elasticIP, "update"); err != nil {
		return err
	}
	if elasticIP.Healthcheck != nil {
		if err := validateOperationParams(elasticIP.Healthcheck, "update"); err != nil {
			return err
		}
	}

	resp, err := c.UpdateElasticIpWithResponse(
		apiv2.WithZone(ctx, zone),
		*elasticIP.ID,
		oapi.UpdateElasticIpJSONRequestBody{
			Description: oapi.NilableString(elasticIP.Description),
			Healthcheck: func() *oapi.ElasticIpHealthcheck {
				if hc := elasticIP.Healthcheck; hc != nil {
					port := int64(*hc.Port)

					return &oapi.ElasticIpHealthcheck{
						Interval: func() (v *int64) {
							if hc.Interval != nil {
								interval := int64(hc.Interval.Seconds())
								v = &interval
							}
							return
						}(),
						Mode:        oapi.ElasticIpHealthcheckMode(*hc.Mode),
						Port:        port,
						StrikesFail: hc.StrikesFail,
						StrikesOk:   hc.StrikesOK,
						Timeout: func() (v *int64) {
							if hc.Timeout != nil {
								timeout := int64(hc.Timeout.Seconds())
								v = &timeout
							}
							return
						}(),
						TlsSkipVerify: hc.TLSSkipVerify,
						TlsSni:        hc.TLSSNI,
						Uri:           hc.URI,
					}
				}
				return nil
			}(),
			Labels: func() (v *oapi.Labels) {
				if elasticIP.Labels != nil {
					v = &oapi.Labels{AdditionalProperties: *elasticIP.Labels}
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

// GetElasticIPReverseDNS returns the Reverse DNS record corresponding to the specified Elastic IP ID.
func (c *Client) GetElasticIPReverseDNS(ctx context.Context, zone, id string) (string, error) {
	resp, err := c.GetReverseDnsElasticIpWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return "", err
	}

	if resp.JSON200 == nil || resp.JSON200.DomainName == nil {
		return "", apiv2.ErrNotFound
	}

	return string(*resp.JSON200.DomainName), nil
}

// DeleteElasticIPReverseDNS deletes a Reverse DNS record of a Elastic IP.
func (c *Client) DeleteElasticIPReverseDNS(ctx context.Context, zone string, id string) error {
	resp, err := c.DeleteReverseDnsElasticIpWithResponse(apiv2.WithZone(ctx, zone), id)
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

// UpdateElasticIPReverseDNS updates a Reverse DNS record for a ElasticIP.
func (c *Client) UpdateElasticIPReverseDNS(ctx context.Context, zone, id, domain string) error {
	resp, err := c.UpdateReverseDnsElasticIpWithResponse(
		apiv2.WithZone(ctx, zone),
		id,
		oapi.UpdateReverseDnsElasticIpJSONRequestBody{
			DomainName: &domain,
		},
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
