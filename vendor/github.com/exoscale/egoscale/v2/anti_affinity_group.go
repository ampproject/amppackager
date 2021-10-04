package v2

import (
	"context"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	papi "github.com/exoscale/egoscale/v2/internal/public-api"
)

// AntiAffinityGroup represents an Anti-Affinity Group.
type AntiAffinityGroup struct {
	Description *string
	ID          *string
	Name        *string `req-for:"create"`
}

func antiAffinityGroupFromAPI(a *papi.AntiAffinityGroup) *AntiAffinityGroup {
	return &AntiAffinityGroup{
		Description: a.Description,
		ID:          a.Id,
		Name:        a.Name,
	}
}

func (a AntiAffinityGroup) get(ctx context.Context, client *Client, zone, id string) (interface{}, error) {
	return client.GetAntiAffinityGroup(ctx, zone, id)
}

// CreateAntiAffinityGroup creates an Anti-Affinity Group in the specified zone.
func (c *Client) CreateAntiAffinityGroup(
	ctx context.Context,
	zone string,
	antiAffinityGroup *AntiAffinityGroup,
) (*AntiAffinityGroup, error) {
	if err := validateOperationParams(antiAffinityGroup, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreateAntiAffinityGroupWithResponse(
		apiv2.WithZone(ctx, zone),
		papi.CreateAntiAffinityGroupJSONRequestBody{
			Description: antiAffinityGroup.Description,
			Name:        *antiAffinityGroup.Name,
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

	return c.GetAntiAffinityGroup(ctx, zone, *res.(*papi.Reference).Id)
}

// ListAntiAffinityGroups returns the list of existing Anti-Affinity Groups in the specified zone.
func (c *Client) ListAntiAffinityGroups(ctx context.Context, zone string) ([]*AntiAffinityGroup, error) {
	list := make([]*AntiAffinityGroup, 0)

	resp, err := c.ListAntiAffinityGroupsWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.AntiAffinityGroups != nil {
		for i := range *resp.JSON200.AntiAffinityGroups {
			list = append(list, antiAffinityGroupFromAPI(&(*resp.JSON200.AntiAffinityGroups)[i]))
		}
	}

	return list, nil
}

// GetAntiAffinityGroup returns the Anti-Affinity Group corresponding to the specified ID in the specified zone.
func (c *Client) GetAntiAffinityGroup(ctx context.Context, zone, id string) (*AntiAffinityGroup, error) {
	resp, err := c.GetAntiAffinityGroupWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return antiAffinityGroupFromAPI(resp.JSON200), nil
}

// FindAntiAffinityGroup attempts to find an Anti-Affinity Group by name or ID in the specified zone.
func (c *Client) FindAntiAffinityGroup(ctx context.Context, zone, v string) (*AntiAffinityGroup, error) {
	res, err := c.ListAntiAffinityGroups(ctx, zone)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		if *r.ID == v || *r.Name == v {
			return c.GetAntiAffinityGroup(ctx, zone, *r.ID)
		}
	}

	return nil, apiv2.ErrNotFound
}

// DeleteAntiAffinityGroup deletes the specified Anti-Affinity Group in the specified zone.
func (c *Client) DeleteAntiAffinityGroup(ctx context.Context, zone, id string) error {
	resp, err := c.DeleteAntiAffinityGroupWithResponse(apiv2.WithZone(ctx, zone), id)
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
