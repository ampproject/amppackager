package v2

import (
	"context"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// SecurityGroup represents a Security Group.
type SecurityGroup struct {
	Description     *string
	ID              *string `req-for:"update,delete"`
	Name            *string `req-for:"create"`
	ExternalSources *[]string
	Rules           []*SecurityGroupRule
}

func securityGroupFromAPI(s *oapi.SecurityGroup) *SecurityGroup {
	return &SecurityGroup{
		Description:     s.Description,
		ID:              s.Id,
		Name:            s.Name,
		ExternalSources: s.ExternalSources,
		Rules: func() (rules []*SecurityGroupRule) {
			if s.Rules != nil {
				rules = make([]*SecurityGroupRule, 0)
				for _, rule := range *s.Rules {
					rule := rule
					rules = append(rules, securityGroupRuleFromAPI(&rule))
				}
			}
			return rules
		}(),
	}
}

// AddExternalSourceToSecurityGroup adds a new external source to a
// Security Group. This operation is idempotent.
func (c *Client) AddExternalSourceToSecurityGroup(
	ctx context.Context,
	zone string,
	securityGroup *SecurityGroup,
	cidr string,
) error {
	if err := validateOperationParams(securityGroup, "update"); err != nil {
		return err
	}

	resp, err := c.AddExternalSourceToSecurityGroupWithResponse(
		apiv2.WithZone(ctx, zone),
		*securityGroup.ID,
		oapi.AddExternalSourceToSecurityGroupJSONRequestBody{
			Cidr: cidr,
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

// CreateSecurityGroup creates a Security Group.
func (c *Client) CreateSecurityGroup(
	ctx context.Context,
	zone string,
	securityGroup *SecurityGroup,
) (*SecurityGroup, error) {
	if err := validateOperationParams(securityGroup, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreateSecurityGroupWithResponse(ctx, oapi.CreateSecurityGroupJSONRequestBody{
		Description: securityGroup.Description,
		Name:        *securityGroup.Name,
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

	return c.GetSecurityGroup(ctx, zone, *res.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
}

// DeleteSecurityGroup deletes a Security Group.
func (c *Client) DeleteSecurityGroup(ctx context.Context, zone string, securityGroup *SecurityGroup) error {
	if err := validateOperationParams(securityGroup, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteSecurityGroupWithResponse(apiv2.WithZone(ctx, zone), *securityGroup.ID)
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

// FindSecurityGroup attempts to find a Security Group by name or ID.
func (c *Client) FindSecurityGroup(ctx context.Context, zone, x string) (*SecurityGroup, error) {
	res, err := c.ListSecurityGroups(ctx, zone)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		if *r.ID == x || *r.Name == x {
			return c.GetSecurityGroup(ctx, zone, *r.ID)
		}
	}

	return nil, apiv2.ErrNotFound
}

// GetSecurityGroup returns the Security Group corresponding to the specified ID.
func (c *Client) GetSecurityGroup(ctx context.Context, zone, id string) (*SecurityGroup, error) {
	resp, err := c.GetSecurityGroupWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return securityGroupFromAPI(resp.JSON200), nil
}

// ListSecurityGroups returns the list of existing Security Groups.
func (c *Client) ListSecurityGroups(ctx context.Context, zone string) ([]*SecurityGroup, error) {
	list := make([]*SecurityGroup, 0)

	resp, err := c.ListSecurityGroupsWithResponse(apiv2.WithZone(ctx, zone), &oapi.ListSecurityGroupsParams{})
	if err != nil {
		return nil, err
	}

	if resp.JSON200.SecurityGroups != nil {
		for i := range *resp.JSON200.SecurityGroups {
			list = append(list, securityGroupFromAPI(&(*resp.JSON200.SecurityGroups)[i]))
		}
	}

	return list, nil
}

// FindSecurityGroups returns the list of existing Security Groups.
// The `params` allows specifying standard filters.
func (c *Client) FindSecurityGroups(ctx context.Context, zone string, params *oapi.ListSecurityGroupsParams) ([]*SecurityGroup, error) {
	list := make([]*SecurityGroup, 0)

	resp, err := c.ListSecurityGroupsWithResponse(apiv2.WithZone(ctx, zone), params)
	if err != nil {
		return nil, err
	}

	if resp.JSON200.SecurityGroups != nil {
		for i := range *resp.JSON200.SecurityGroups {
			list = append(list, securityGroupFromAPI(&(*resp.JSON200.SecurityGroups)[i]))
		}
	}

	return list, nil
}

// RemoveExternalSourceFromSecurityGroup removes an external source from
// a Security Group. This operation is idempotent.
func (c *Client) RemoveExternalSourceFromSecurityGroup(
	ctx context.Context,
	zone string,
	securityGroup *SecurityGroup,
	cidr string,
) error {
	if err := validateOperationParams(securityGroup, "update"); err != nil {
		return err
	}

	resp, err := c.RemoveExternalSourceFromSecurityGroupWithResponse(
		apiv2.WithZone(ctx, zone),
		*securityGroup.ID,
		oapi.RemoveExternalSourceFromSecurityGroupJSONRequestBody{
			Cidr: cidr,
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
