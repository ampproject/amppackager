package v2

import (
	"context"
	"errors"
	"fmt"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	papi "github.com/exoscale/egoscale/v2/internal/public-api"
)

// SKSNodepool represents a SKS Nodepool.
type SKSNodepool struct {
	AntiAffinityGroupIDs []string `reset:"anti-affinity-groups"`
	CreatedAt            time.Time
	Description          string `reset:"description"`
	DiskSize             int64
	ID                   string
	InstancePoolID       string
	InstanceTypeID       string
	Name                 string
	SecurityGroupIDs     []string `reset:"security-groups"`
	Size                 int64
	State                string
	TemplateID           string
	Version              string
}

func sksNodepoolFromAPI(n *papi.SksNodepool) *SKSNodepool {
	return &SKSNodepool{
		AntiAffinityGroupIDs: func() []string {
			ids := make([]string, 0)
			if n.AntiAffinityGroups != nil {
				for _, aag := range *n.AntiAffinityGroups {
					aag := aag
					ids = append(ids, *aag.Id)
				}
			}
			return ids
		}(),
		CreatedAt:      *n.CreatedAt,
		Description:    papi.OptionalString(n.Description),
		DiskSize:       papi.OptionalInt64(n.DiskSize),
		ID:             papi.OptionalString(n.Id),
		InstancePoolID: papi.OptionalString(n.InstancePool.Id),
		InstanceTypeID: papi.OptionalString(n.InstanceType.Id),
		Name:           papi.OptionalString(n.Name),
		SecurityGroupIDs: func() []string {
			ids := make([]string, 0)
			if n.SecurityGroups != nil {
				for _, sg := range *n.SecurityGroups {
					sg := sg
					ids = append(ids, *sg.Id)
				}
			}
			return ids
		}(),
		Size:       papi.OptionalInt64(n.Size),
		State:      papi.OptionalString(n.State),
		TemplateID: papi.OptionalString(n.Template.Id),
		Version:    papi.OptionalString(n.Version),
	}
}

// SKSCluster represents a SKS cluster.
type SKSCluster struct {
	AddOns       []string
	CNI          string
	CreatedAt    time.Time
	Description  string `reset:"description"`
	Endpoint     string
	ID           string
	Name         string
	Nodepools    []*SKSNodepool
	ServiceLevel string
	State        string
	Version      string

	c    *Client
	zone string
}

func sksClusterFromAPI(c *papi.SksCluster) *SKSCluster {
	return &SKSCluster{
		AddOns: func() []string {
			addOns := make([]string, 0)
			if c.Addons != nil {
				addOns = append(addOns, *c.Addons...)
			}
			return addOns
		}(),
		CNI:         papi.OptionalString(c.Cni),
		CreatedAt:   *c.CreatedAt,
		Description: papi.OptionalString(c.Description),
		Endpoint:    papi.OptionalString(c.Endpoint),
		ID:          papi.OptionalString(c.Id),
		Name:        papi.OptionalString(c.Name),
		Nodepools: func() []*SKSNodepool {
			nodepools := make([]*SKSNodepool, 0)
			if c.Nodepools != nil {
				for _, n := range *c.Nodepools {
					n := n
					nodepools = append(nodepools, sksNodepoolFromAPI(&n))
				}
			}
			return nodepools
		}(),
		ServiceLevel: papi.OptionalString(c.Level),
		State:        papi.OptionalString(c.State),
		Version:      papi.OptionalString(c.Version),
	}
}

// RotateCCMCredentials rotates the Exoscale IAM credentials managed by the SKS control plane for the
// Kubernetes Exoscale Cloud Controller Manager.
func (c *SKSCluster) RotateCCMCredentials(ctx context.Context) error {
	_, err := c.c.RotateSksCcmCredentialsWithResponse(apiv2.WithZone(ctx, c.zone), c.ID)
	return err
}

// AuthorityCert returns the SKS cluster base64-encoded certificate content for the specified authority.
func (c *SKSCluster) AuthorityCert(ctx context.Context, authority string) (string, error) {
	if authority == "" {
		return "", errors.New("authority not specified")
	}

	resp, err := c.c.GetSksClusterAuthorityCertWithResponse(apiv2.WithZone(ctx, c.zone), c.ID, authority)
	if err != nil {
		return "", err
	}

	return papi.OptionalString(resp.JSON200.Cacert), nil
}

// RequestKubeconfig returns a base64-encoded kubeconfig content for the specified user name,
// optionally associated to specified groups for a duration d (default API-set TTL applies if not specified).
// Fore more information: https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/
func (c *SKSCluster) RequestKubeconfig(ctx context.Context, user string, groups []string,
	d time.Duration) (string, error) {
	if user == "" {
		return "", errors.New("user not specified")
	}

	resp, err := c.c.GenerateSksClusterKubeconfigWithResponse(
		apiv2.WithZone(ctx, c.zone),
		c.ID,
		papi.GenerateSksClusterKubeconfigJSONRequestBody{
			User:   &user,
			Groups: &groups,
			Ttl: func() *int64 {
				ttl := int64(d.Seconds())
				if ttl > 0 {
					return &ttl
				}
				return nil
			}(),
		})
	if err != nil {
		return "", err
	}

	return papi.OptionalString(resp.JSON200.Kubeconfig), nil
}

// AddNodepool adds a Nodepool to the SKS cluster.
func (c *SKSCluster) AddNodepool(ctx context.Context, np *SKSNodepool) (*SKSNodepool, error) {
	resp, err := c.c.CreateSksNodepoolWithResponse(
		apiv2.WithZone(ctx, c.zone),
		c.ID,
		papi.CreateSksNodepoolJSONRequestBody{
			AntiAffinityGroups: func() *[]papi.AntiAffinityGroup {
				if l := len(np.AntiAffinityGroupIDs); l > 0 {
					list := make([]papi.AntiAffinityGroup, l)
					for i, v := range np.AntiAffinityGroupIDs {
						v := v
						list[i] = papi.AntiAffinityGroup{Id: &v}
					}
					return &list
				}
				return nil
			}(),
			Description: func() *string {
				if np.Description != "" {
					return &np.Description
				}
				return nil
			}(),
			DiskSize:     np.DiskSize,
			InstanceType: papi.InstanceType{Id: &np.InstanceTypeID},
			Name:         np.Name,
			SecurityGroups: func() *[]papi.SecurityGroup {
				if l := len(np.SecurityGroupIDs); l > 0 {
					list := make([]papi.SecurityGroup, l)
					for i, v := range np.SecurityGroupIDs {
						v := v
						list[i] = papi.SecurityGroup{Id: &v}
					}
					return &list
				}
				return nil
			}(),
			Size: np.Size,
		})
	if err != nil {
		return nil, err
	}

	res, err := papi.NewPoller().
		WithTimeout(c.c.timeout).
		Poll(ctx, c.c.OperationPoller(c.zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	nodepoolRes, err := c.c.GetSksNodepoolWithResponse(ctx, c.ID, *res.(*papi.Reference).Id)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Nodepool: %s", err)
	}

	return sksNodepoolFromAPI(nodepoolRes.JSON200), nil
}

// UpdateNodepool updates the specified SKS cluster Nodepool.
func (c *SKSCluster) UpdateNodepool(ctx context.Context, np *SKSNodepool) error {
	resp, err := c.c.UpdateSksNodepoolWithResponse(
		apiv2.WithZone(ctx, c.zone),
		c.ID,
		np.ID,
		papi.UpdateSksNodepoolJSONRequestBody{
			AntiAffinityGroups: func() *[]papi.AntiAffinityGroup {
				if l := len(np.AntiAffinityGroupIDs); l > 0 {
					list := make([]papi.AntiAffinityGroup, l)
					for i, v := range np.AntiAffinityGroupIDs {
						v := v
						list[i] = papi.AntiAffinityGroup{Id: &v}
					}
					return &list
				}
				return nil
			}(),
			Description: func() *string {
				if np.Description != "" {
					return &np.Description
				}
				return nil
			}(),
			DiskSize: func() *int64 {
				if np.DiskSize > 0 {
					return &np.DiskSize
				}
				return nil
			}(),
			InstanceType: func() *papi.InstanceType {
				if np.InstanceTypeID != "" {
					return &papi.InstanceType{Id: &np.InstanceTypeID}
				}
				return nil
			}(),
			Name: func() *string {
				if np.Name != "" {
					return &np.Name
				}
				return nil
			}(),
			SecurityGroups: func() *[]papi.SecurityGroup {
				if l := len(np.SecurityGroupIDs); l > 0 {
					list := make([]papi.SecurityGroup, l)
					for i, v := range np.SecurityGroupIDs {
						v := v
						list[i] = papi.SecurityGroup{Id: &v}
					}
					return &list
				}
				return nil
			}(),
		})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(c.c.timeout).
		Poll(ctx, c.c.OperationPoller(c.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// ScaleNodepool scales the SKS cluster Nodepool to the specified number of Kubernetes Nodes.
func (c *SKSCluster) ScaleNodepool(ctx context.Context, np *SKSNodepool, nodes int64) error {
	resp, err := c.c.ScaleSksNodepoolWithResponse(
		apiv2.WithZone(ctx, c.zone),
		c.ID,
		np.ID,
		papi.ScaleSksNodepoolJSONRequestBody{Size: nodes},
	)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(c.c.timeout).
		Poll(ctx, c.c.OperationPoller(c.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// EvictNodepoolMembers evicts the specified members (identified by their Compute instance ID) from the
// SKS cluster Nodepool.
func (c *SKSCluster) EvictNodepoolMembers(ctx context.Context, np *SKSNodepool, members []string) error {
	resp, err := c.c.EvictSksNodepoolMembersWithResponse(
		apiv2.WithZone(ctx, c.zone),
		c.ID,
		np.ID,
		papi.EvictSksNodepoolMembersJSONRequestBody{Instances: &members},
	)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(c.c.timeout).
		Poll(ctx, c.c.OperationPoller(c.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// DeleteNodepool deletes the specified Nodepool from the SKS cluster.
func (c *SKSCluster) DeleteNodepool(ctx context.Context, np *SKSNodepool) error {
	resp, err := c.c.DeleteSksNodepoolWithResponse(
		apiv2.WithZone(ctx, c.zone),
		c.ID,
		np.ID,
	)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(c.c.timeout).
		Poll(ctx, c.c.OperationPoller(c.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// ResetField resets the specified SKS cluster field to its default value.
// The value expected for the field parameter is a pointer to the SKSCluster field to reset.
func (c *SKSCluster) ResetField(ctx context.Context, field interface{}) error {
	resetField, err := resetFieldName(c, field)
	if err != nil {
		return err
	}

	resp, err := c.c.ResetSksClusterFieldWithResponse(apiv2.WithZone(ctx, c.zone), c.ID, resetField)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(c.c.timeout).
		Poll(ctx, c.c.OperationPoller(c.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// ResetNodepoolField resets the specified SKS Nodepool field to its default value.
// The value expected for the field parameter is a pointer to the SKSNodepool field to reset.
func (c *SKSCluster) ResetNodepoolField(ctx context.Context, np *SKSNodepool, field interface{}) error {
	resetField, err := resetFieldName(np, field)
	if err != nil {
		return err
	}

	resp, err := c.c.ResetSksNodepoolFieldWithResponse(apiv2.WithZone(ctx, c.zone), c.ID, np.ID, resetField)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(c.c.timeout).
		Poll(ctx, c.c.OperationPoller(c.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// CreateSKSCluster creates a SKS cluster in the specified zone.
func (c *Client) CreateSKSCluster(ctx context.Context, zone string, cluster *SKSCluster) (*SKSCluster, error) {
	resp, err := c.CreateSksClusterWithResponse(
		apiv2.WithZone(ctx, zone),
		papi.CreateSksClusterJSONRequestBody{
			Addons: func() *[]string {
				var addOns *[]string
				if len(cluster.AddOns) > 0 {
					addOns = &cluster.AddOns
				}
				return addOns
			}(),
			Cni: func() *string {
				var cni *string
				if cluster.CNI != "" {
					cni = &cluster.CNI
				}
				return cni
			}(),
			Description: &cluster.Description,
			Level:       cluster.ServiceLevel,
			Name:        cluster.Name,
			Version:     cluster.Version,
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

	return c.GetSKSCluster(ctx, zone, *res.(*papi.Reference).Id)
}

// ListSKSClusters returns the list of existing SKS clusters in the specified zone.
func (c *Client) ListSKSClusters(ctx context.Context, zone string) ([]*SKSCluster, error) {
	list := make([]*SKSCluster, 0)

	resp, err := c.ListSksClustersWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.SksClusters != nil {
		for i := range *resp.JSON200.SksClusters {
			cluster := sksClusterFromAPI(&(*resp.JSON200.SksClusters)[i])
			cluster.c = c
			cluster.zone = zone

			list = append(list, cluster)
		}
	}

	return list, nil
}

// ListSKSClusterVersions returns the list of Kubernetes versions supported during SKS cluster creation.
func (c *Client) ListSKSClusterVersions(ctx context.Context) ([]string, error) {
	list := make([]string, 0)

	resp, err := c.ListSksClusterVersionsWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.JSON200.SksClusterVersions != nil {
		for i := range *resp.JSON200.SksClusterVersions {
			version := &(*resp.JSON200.SksClusterVersions)[i]
			list = append(list, *version)
		}
	}

	return list, nil
}

// GetSKSCluster returns the SKS cluster corresponding to the specified ID in the specified zone.
func (c *Client) GetSKSCluster(ctx context.Context, zone, id string) (*SKSCluster, error) {
	resp, err := c.GetSksClusterWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	cluster := sksClusterFromAPI(resp.JSON200)
	cluster.c = c
	cluster.zone = zone

	return cluster, nil
}

// UpdateSKSCluster updates the specified SKS cluster in the specified zone.
func (c *Client) UpdateSKSCluster(ctx context.Context, zone string, cluster *SKSCluster) error {
	resp, err := c.UpdateSksClusterWithResponse(
		apiv2.WithZone(ctx, zone),
		cluster.ID,
		papi.UpdateSksClusterJSONRequestBody{
			Description: func() *string {
				if cluster.Description != "" {
					return &cluster.Description
				}
				return nil
			}(),
			Name: func() *string {
				if cluster.Name != "" {
					return &cluster.Name
				}
				return nil
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

// UpgradeSKSCluster upgrades the SKS cluster corresponding to the specified ID in the specified zone to the
// requested Kubernetes version.
func (c *Client) UpgradeSKSCluster(ctx context.Context, zone, id, version string) error {
	resp, err := c.UpgradeSksClusterWithResponse(
		apiv2.WithZone(ctx, zone),
		id,
		papi.UpgradeSksClusterJSONRequestBody{Version: version})
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

// DeleteSKSCluster deletes the specified SKS cluster in the specified zone.
func (c *Client) DeleteSKSCluster(ctx context.Context, zone, id string) error {
	resp, err := c.DeleteSksClusterWithResponse(apiv2.WithZone(ctx, zone), id)
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
