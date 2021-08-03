package v2

import (
	"context"
	"net"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	papi "github.com/exoscale/egoscale/v2/internal/public-api"
)

// PrivateNetwork represents a Private Network.
type PrivateNetwork struct {
	Description string
	EndIP       net.IP
	ID          string
	Name        string
	Netmask     net.IP
	StartIP     net.IP
}

func privateNetworkFromAPI(p *papi.PrivateNetwork) *PrivateNetwork {
	return &PrivateNetwork{
		Description: papi.OptionalString(p.Description),
		EndIP:       net.ParseIP(papi.OptionalString(p.EndIp)),
		ID:          papi.OptionalString(p.Id),
		Name:        papi.OptionalString(p.Name),
		Netmask:     net.ParseIP(papi.OptionalString(p.Netmask)),
		StartIP:     net.ParseIP(papi.OptionalString(p.StartIp)),
	}
}

// CreatePrivateNetwork creates a Private Network in the specified zone.
func (c *Client) CreatePrivateNetwork(ctx context.Context, zone string,
	privateNetwork *PrivateNetwork) (*PrivateNetwork, error) {
	resp, err := c.CreatePrivateNetworkWithResponse(
		apiv2.WithZone(ctx, zone),
		papi.CreatePrivateNetworkJSONRequestBody{
			Description: &privateNetwork.Description,
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
		return nil, err
	}

	res, err := papi.NewPoller().
		WithTimeout(c.timeout).
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
			list = append(list, privateNetworkFromAPI(&(*resp.JSON200.PrivateNetworks)[i]))
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

	return privateNetworkFromAPI(resp.JSON200), nil
}

// UpdatePrivateNetwork updates the specified Private Network in the specified zone.
func (c *Client) UpdatePrivateNetwork(ctx context.Context, zone string, privateNetwork *PrivateNetwork) error {
	resp, err := c.UpdatePrivateNetworkWithResponse(
		apiv2.WithZone(ctx, zone),
		privateNetwork.ID,
		papi.UpdatePrivateNetworkJSONRequestBody{
			Description: &privateNetwork.Description,
			EndIp: func() (ip *string) {
				if privateNetwork.EndIP != nil {
					v := privateNetwork.EndIP.String()
					return &v
				}
				return
			}(),
			Name: &privateNetwork.Name,
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
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}
