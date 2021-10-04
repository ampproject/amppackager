package v2

import (
	"context"
	"net"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	papi "github.com/exoscale/egoscale/v2/internal/public-api"
)

// PrivateNetworkLease represents a managed Private Network lease.
type PrivateNetworkLease struct {
	InstanceID *string
	IPAddress  *net.IP
}

// PrivateNetwork represents a Private Network.
type PrivateNetwork struct {
	Description *string
	EndIP       *net.IP
	ID          *string `req-for:"update"`
	Name        *string `req-for:"create"`
	Netmask     *net.IP
	StartIP     *net.IP
	Leases      []*PrivateNetworkLease

	c    *Client
	zone string
}

func privateNetworkFromAPI(client *Client, zone string, p *papi.PrivateNetwork) *PrivateNetwork {
	return &PrivateNetwork{
		Description: p.Description,
		EndIP: func() (v *net.IP) {
			if p.EndIp != nil {
				ip := net.ParseIP(*p.EndIp)
				v = &ip
			}
			return
		}(),
		ID:   p.Id,
		Name: p.Name,
		Netmask: func() (v *net.IP) {
			if p.Netmask != nil {
				ip := net.ParseIP(*p.Netmask)
				v = &ip
			}
			return
		}(),
		StartIP: func() (v *net.IP) {
			if p.StartIp != nil {
				ip := net.ParseIP(*p.StartIp)
				v = &ip
			}
			return
		}(),
		Leases: func() (v []*PrivateNetworkLease) {
			if p.Leases != nil {
				v = make([]*PrivateNetworkLease, len(*p.Leases))
				for i, lease := range *p.Leases {
					v[i] = &PrivateNetworkLease{
						InstanceID: lease.InstanceId,
						IPAddress:  func() *net.IP { ip := net.ParseIP(*lease.Ip); return &ip }(),
					}
				}
			}
			return
		}(),

		c:    client,
		zone: zone,
	}
}

func (p PrivateNetwork) get(ctx context.Context, client *Client, zone, id string) (interface{}, error) {
	return client.GetPrivateNetwork(ctx, zone, id)
}

// UpdateInstanceIPAddress updates the IP address of a Compute instance attached to the managed Private Network.
func (p *PrivateNetwork) UpdateInstanceIPAddress(ctx context.Context, instance *Instance, ip net.IP) error {
	resp, err := p.c.UpdatePrivateNetworkInstanceIpWithResponse(
		apiv2.WithZone(ctx, p.zone),
		*p.ID,
		papi.UpdatePrivateNetworkInstanceIpJSONRequestBody{
			Instance: papi.Instance{Id: instance.ID},
			Ip: func() *string {
				s := ip.String()
				return &s
			}(),
		})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(p.c.timeout).
		WithInterval(p.c.pollInterval).
		Poll(ctx, p.c.OperationPoller(p.zone, *resp.JSON200.Id))

	return err
}

// CreatePrivateNetwork creates a Private Network in the specified zone.
func (c *Client) CreatePrivateNetwork(
	ctx context.Context,
	zone string,
	privateNetwork *PrivateNetwork,
) (*PrivateNetwork, error) {
	if err := validateOperationParams(privateNetwork, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreatePrivateNetworkWithResponse(
		apiv2.WithZone(ctx, zone),
		papi.CreatePrivateNetworkJSONRequestBody{
			Description: privateNetwork.Description,
			EndIp: func() (ip *string) {
				if privateNetwork.EndIP != nil {
					v := privateNetwork.EndIP.String()
					return &v
				}
				return
			}(),
			Name: *privateNetwork.Name,
			Netmask: func() (ip *string) {
				if privateNetwork.Netmask != nil {
					v := privateNetwork.Netmask.String()
					return &v
				}
				return
			}(),
			StartIp: func() (ip *string) {
				if privateNetwork.StartIP != nil {
					v := privateNetwork.StartIP.String()
					return &v
				}
				return
			}(),
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

	return c.GetPrivateNetwork(ctx, zone, *res.(*papi.Reference).Id)
}

// ListPrivateNetworks returns the list of existing Private Networks in the specified zone.
func (c *Client) ListPrivateNetworks(ctx context.Context, zone string) ([]*PrivateNetwork, error) {
	list := make([]*PrivateNetwork, 0)

	resp, err := c.ListPrivateNetworksWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.PrivateNetworks != nil {
		for i := range *resp.JSON200.PrivateNetworks {
			list = append(list, privateNetworkFromAPI(c, zone, &(*resp.JSON200.PrivateNetworks)[i]))
		}
	}

	return list, nil
}

// GetPrivateNetwork returns the Private Network corresponding to the specified ID in the specified zone.
func (c *Client) GetPrivateNetwork(ctx context.Context, zone, id string) (*PrivateNetwork, error) {
	resp, err := c.GetPrivateNetworkWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return privateNetworkFromAPI(c, zone, resp.JSON200), nil
}

// FindPrivateNetwork attempts to find a Private Network by name or ID in the specified zone.
// In case the identifier is a name and multiple resources match, an ErrTooManyFound error is returned.
func (c *Client) FindPrivateNetwork(ctx context.Context, zone, v string) (*PrivateNetwork, error) {
	res, err := c.ListPrivateNetworks(ctx, zone)
	if err != nil {
		return nil, err
	}

	var found *PrivateNetwork
	for _, r := range res {
		if *r.ID == v {
			return c.GetPrivateNetwork(ctx, zone, *r.ID)
		}

		// Historically, the Exoscale API allowed users to create multiple Private Networks sharing a common name.
		// This function being expected to return one resource at most, in case the specified identifier is a name
		// we have to check that there aren't more that one matching result before returning it.
		if *r.Name == v {
			if found != nil {
				return nil, apiv2.ErrTooManyFound
			}
			found = r
		}
	}

	if found != nil {
		return c.GetPrivateNetwork(ctx, zone, *found.ID)
	}

	return nil, apiv2.ErrNotFound
}

// UpdatePrivateNetwork updates the specified Private Network in the specified zone.
func (c *Client) UpdatePrivateNetwork(ctx context.Context, zone string, privateNetwork *PrivateNetwork) error {
	if err := validateOperationParams(privateNetwork, "update"); err != nil {
		return err
	}

	resp, err := c.UpdatePrivateNetworkWithResponse(
		apiv2.WithZone(ctx, zone),
		*privateNetwork.ID,
		papi.UpdatePrivateNetworkJSONRequestBody{
			Description: privateNetwork.Description,
			EndIp: func() (ip *string) {
				if privateNetwork.EndIP != nil {
					v := privateNetwork.EndIP.String()
					return &v
				}
				return
			}(),
			Name: privateNetwork.Name,
			Netmask: func() (ip *string) {
				if privateNetwork.Netmask != nil {
					v := privateNetwork.Netmask.String()
					return &v
				}
				return
			}(),
			StartIp: func() (ip *string) {
				if privateNetwork.StartIP != nil {
					v := privateNetwork.StartIP.String()
					return &v
				}
				return
			}(),
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

// DeletePrivateNetwork deletes the specified Private Network in the specified zone.
func (c *Client) DeletePrivateNetwork(ctx context.Context, zone, id string) error {
	resp, err := c.DeletePrivateNetworkWithResponse(apiv2.WithZone(ctx, zone), id)
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
