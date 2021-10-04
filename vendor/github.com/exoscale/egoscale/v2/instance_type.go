package v2

import (
	"context"
	"strings"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	papi "github.com/exoscale/egoscale/v2/internal/public-api"
)

// InstanceType represents a Compute instance type.
type InstanceType struct {
	Authorized *bool
	CPUs       *int64
	Family     *string
	GPUs       *int64
	ID         *string
	Memory     *int64
	Size       *string
}

// ToAPIMock returns the low-level representation of the resource. This is intended for testing purposes.
func (t InstanceType) ToAPIMock() interface{} {
	return papi.InstanceType{
		Authorized: t.Authorized,
		Cpus:       t.CPUs,
		Family:     (*papi.InstanceTypeFamily)(t.Family),
		Gpus:       t.GPUs,
		Id:         t.ID,
		Memory:     t.Memory,
		Size:       (*papi.InstanceTypeSize)(t.Size),
	}
}

func instanceTypeFromAPI(t *papi.InstanceType) *InstanceType {
	return &InstanceType{
		Authorized: t.Authorized,
		CPUs:       t.Cpus,
		Family:     (*string)(t.Family),
		GPUs:       t.Gpus,
		ID:         t.Id,
		Memory:     t.Memory,
		Size:       (*string)(t.Size),
	}
}

// ListInstanceTypes returns the list of existing Instance types in the specified zone.
func (c *Client) ListInstanceTypes(ctx context.Context, zone string) ([]*InstanceType, error) {
	list := make([]*InstanceType, 0)

	resp, err := c.ListInstanceTypesWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.InstanceTypes != nil {
		for i := range *resp.JSON200.InstanceTypes {
			list = append(list, instanceTypeFromAPI(&(*resp.JSON200.InstanceTypes)[i]))
		}
	}

	return list, nil
}

// GetInstanceType returns the Instance type corresponding to the specified ID in the specified zone.
func (c *Client) GetInstanceType(ctx context.Context, zone, id string) (*InstanceType, error) {
	resp, err := c.GetInstanceTypeWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return instanceTypeFromAPI(resp.JSON200), nil
}

// FindInstanceType attempts to find an Instance type by family+size or ID in the specified zone.
// To search by family+size, the expected format for v is "[FAMILY.]SIZE" (e.g. "large", "gpu.medium"),
// with family defaulting to "standard" if not specified.
func (c *Client) FindInstanceType(ctx context.Context, zone, v string) (*InstanceType, error) {
	var typeFamily, typeSize string

	parts := strings.SplitN(v, ".", 2)
	if l := len(parts); l > 0 {
		if l == 1 {
			typeFamily, typeSize = "standard", strings.ToLower(parts[0])
		} else {
			typeFamily, typeSize = strings.ToLower(parts[0]), strings.ToLower(parts[1])
		}
	}

	res, err := c.ListInstanceTypes(ctx, zone)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		if *r.ID == v || (*r.Family == typeFamily && *r.Size == typeSize) {
			return c.GetInstanceType(ctx, zone, *r.ID)
		}
	}

	return nil, apiv2.ErrNotFound
}
