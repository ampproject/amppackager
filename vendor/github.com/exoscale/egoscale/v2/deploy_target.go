package v2

import (
	"context"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	papi "github.com/exoscale/egoscale/v2/internal/public-api"
)

// DeployTarget represents a Deploy Target.
type DeployTarget struct {
	Description *string
	ID          *string
	Name        *string
	Type        *string
}

// ToAPIMock returns the low-level representation of the resource. This is intended for testing purposes.
func (d DeployTarget) ToAPIMock() interface{} {
	return papi.DeployTarget{
		Description: d.Description,
		Id:          d.ID,
		Name:        d.Name,
		Type:        (*papi.DeployTargetType)(d.Type),
	}
}

func deployTargetFromAPI(d *papi.DeployTarget) *DeployTarget {
	return &DeployTarget{
		Description: d.Description,
		ID:          d.Id,
		Name:        d.Name,
		Type:        (*string)(d.Type),
	}
}

// ListDeployTargets returns the list of existing Deploy Targets in the specified zone.
func (c *Client) ListDeployTargets(ctx context.Context, zone string) ([]*DeployTarget, error) {
	list := make([]*DeployTarget, 0)

	resp, err := c.ListDeployTargetsWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.DeployTargets != nil {
		for i := range *resp.JSON200.DeployTargets {
			list = append(list, deployTargetFromAPI(&(*resp.JSON200.DeployTargets)[i]))
		}
	}

	return list, nil
}

// GetDeployTarget returns the Deploy Target corresponding to the specified ID in the specified zone.
func (c *Client) GetDeployTarget(ctx context.Context, zone, id string) (*DeployTarget, error) {
	resp, err := c.GetDeployTargetWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return deployTargetFromAPI(resp.JSON200), nil
}

// FindDeployTarget attempts to find a Deploy Target by name or ID in the specified zone.
func (c *Client) FindDeployTarget(ctx context.Context, zone, v string) (*DeployTarget, error) {
	res, err := c.ListDeployTargets(ctx, zone)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		if *r.ID == v || *r.Name == v {
			return c.GetDeployTarget(ctx, zone, *r.ID)
		}
	}

	return nil, apiv2.ErrNotFound
}
