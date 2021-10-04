package v2

import (
	"context"
	"net"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	papi "github.com/exoscale/egoscale/v2/internal/public-api"
)

// InstanceManager represents a Compute instance manager.
type InstanceManager struct {
	ID   string
	Type string
}

// Instance represents a Compute instance.
type Instance struct {
	AntiAffinityGroupIDs *[]string
	CreatedAt            *time.Time
	DeployTargetID       *string
	DiskSize             *int64 `req-if:"create"`
	ElasticIPIDs         *[]string
	ID                   *string `req-if:"update"`
	IPv6Address          *net.IP
	IPv6Enabled          *bool
	InstanceTypeID       *string `req-if:"create"`
	Labels               *map[string]string
	Manager              *InstanceManager
	Name                 *string `req-if:"create"`
	PrivateNetworkIDs    *[]string
	PublicIPAddress      *net.IP
	SSHKey               *string
	SecurityGroupIDs     *[]string
	SnapshotIDs          *[]string
	State                *string
	TemplateID           *string `req-if:"create"`
	UserData             *string

	c    *Client
	zone string
}

func instanceFromAPI(client *Client, zone string, i *papi.Instance) *Instance {
	return &Instance{
		AntiAffinityGroupIDs: func() (v *[]string) {
			if i.AntiAffinityGroups != nil && len(*i.AntiAffinityGroups) > 0 {
				ids := make([]string, len(*i.AntiAffinityGroups))
				for i, item := range *i.AntiAffinityGroups {
					ids[i] = *item.Id
				}
				v = &ids
			}
			return
		}(),
		CreatedAt: i.CreatedAt,
		DeployTargetID: func() (v *string) {
			if i.DeployTarget != nil {
				v = i.DeployTarget.Id
			}
			return
		}(),
		DiskSize: i.DiskSize,
		ElasticIPIDs: func() (v *[]string) {
			if i.ElasticIps != nil && len(*i.ElasticIps) > 0 {
				ids := make([]string, len(*i.ElasticIps))
				for i, item := range *i.ElasticIps {
					ids[i] = *item.Id
				}
				v = &ids
			}
			return
		}(),
		ID: i.Id,
		IPv6Address: func() (v *net.IP) {
			if i.Ipv6Address != nil {
				ip := net.ParseIP(*i.Ipv6Address)
				v = &ip
			}
			return
		}(),
		IPv6Enabled: func() (v *bool) {
			if i.Ipv6Address != nil {
				ipv6Enabled := i.Ipv6Address != nil
				v = &ipv6Enabled
			}
			return
		}(),
		InstanceTypeID: i.InstanceType.Id,
		Labels: func() (v *map[string]string) {
			if i.Labels != nil && len(i.Labels.AdditionalProperties) > 0 {
				v = &i.Labels.AdditionalProperties
			}
			return
		}(),
		Manager: func() *InstanceManager {
			if i.Manager != nil {
				return &InstanceManager{
					ID:   *i.Manager.Id,
					Type: string(*i.Manager.Type),
				}
			}
			return nil
		}(),
		Name: i.Name,
		PrivateNetworkIDs: func() (v *[]string) {
			if i.PrivateNetworks != nil && len(*i.PrivateNetworks) > 0 {
				ids := make([]string, len(*i.PrivateNetworks))
				for i, item := range *i.PrivateNetworks {
					ids[i] = *item.Id
				}
				v = &ids
			}
			return
		}(),
		PublicIPAddress: func() (v *net.IP) {
			if i.PublicIp != nil {
				ip := net.ParseIP(*i.PublicIp)
				v = &ip
			}
			return
		}(),
		SSHKey: func() (v *string) {
			if i.SshKey != nil {
				v = i.SshKey.Name
			}
			return
		}(),
		SecurityGroupIDs: func() (v *[]string) {
			if i.SecurityGroups != nil && len(*i.SecurityGroups) > 0 {
				ids := make([]string, len(*i.SecurityGroups))
				for i, item := range *i.SecurityGroups {
					ids[i] = *item.Id
				}
				v = &ids
			}
			return
		}(),
		SnapshotIDs: func() (v *[]string) {
			if i.Snapshots != nil && len(*i.Snapshots) > 0 {
				ids := make([]string, len(*i.Snapshots))
				for i, item := range *i.Snapshots {
					ids[i] = *item.Id
				}
				v = &ids
			}
			return
		}(),
		State:      (*string)(i.State),
		TemplateID: i.Template.Id,
		UserData:   i.UserData,

		c:    client,
		zone: zone,
	}
}

func (i Instance) get(ctx context.Context, client *Client, zone, id string) (interface{}, error) {
	return client.GetInstance(ctx, zone, id)
}

// AntiAffinityGroups returns the list of Anti-Affinity Groups applied to the Compute instance.
func (i *Instance) AntiAffinityGroups(ctx context.Context) ([]*AntiAffinityGroup, error) {
	if i.AntiAffinityGroupIDs != nil {
		res, err := i.c.fetchFromIDs(ctx, i.zone, *i.AntiAffinityGroupIDs, new(AntiAffinityGroup))
		return res.([]*AntiAffinityGroup), err
	}
	return nil, nil
}

// AttachElasticIP attaches the Compute instance to the specified Elastic IP.
func (i *Instance) AttachElasticIP(ctx context.Context, elasticIP *ElasticIP) error {
	resp, err := i.c.AttachInstanceToElasticIpWithResponse(
		apiv2.WithZone(ctx, i.zone), *elasticIP.ID, papi.AttachInstanceToElasticIpJSONRequestBody{
			Instance: papi.Instance{Id: i.ID},
		})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// AttachPrivateNetwork attaches the Compute instance to the specified Private Network.
func (i *Instance) AttachPrivateNetwork(ctx context.Context, privateNetwork *PrivateNetwork, address net.IP) error {
	resp, err := i.c.AttachInstanceToPrivateNetworkWithResponse(
		apiv2.WithZone(ctx, i.zone), *privateNetwork.ID, papi.AttachInstanceToPrivateNetworkJSONRequestBody{
			Instance: papi.Instance{Id: i.ID},
			Ip: func() *string {
				if len(address) > 0 {
					ip := address.String()
					return &ip
				}
				return nil
			}(),
		})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// AttachSecurityGroup attaches the Compute instance to the specified Security Group.
func (i *Instance) AttachSecurityGroup(ctx context.Context, securityGroup *SecurityGroup) error {
	resp, err := i.c.AttachInstanceToSecurityGroupWithResponse(
		apiv2.WithZone(ctx, i.zone), *securityGroup.ID, papi.AttachInstanceToSecurityGroupJSONRequestBody{
			Instance: papi.Instance{Id: i.ID},
		})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// CreateSnapshot creates a Snapshot of the Compute instance storage volume.
func (i *Instance) CreateSnapshot(ctx context.Context) (*Snapshot, error) {
	resp, err := i.c.CreateSnapshotWithResponse(apiv2.WithZone(ctx, i.zone), *i.ID)
	if err != nil {
		return nil, err
	}

	res, err := papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	return i.c.GetSnapshot(ctx, i.zone, *res.(*papi.Reference).Id)
}

// DetachElasticIP detaches the Compute instance from the specified Elastic IP.
func (i *Instance) DetachElasticIP(ctx context.Context, elasticIP *ElasticIP) error {
	resp, err := i.c.DetachInstanceFromElasticIpWithResponse(
		apiv2.WithZone(ctx, i.zone), *elasticIP.ID, papi.DetachInstanceFromElasticIpJSONRequestBody{
			Instance: papi.Instance{Id: i.ID},
		})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// DetachPrivateNetwork detaches the Compute instance from the specified Private Network.
func (i *Instance) DetachPrivateNetwork(ctx context.Context, privateNetwork *PrivateNetwork) error {
	resp, err := i.c.DetachInstanceFromPrivateNetworkWithResponse(
		apiv2.WithZone(ctx, i.zone), *privateNetwork.ID, papi.DetachInstanceFromPrivateNetworkJSONRequestBody{
			Instance: papi.Instance{Id: i.ID},
		})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// DetachSecurityGroup detaches the Compute instance from the specified Security Group.
func (i *Instance) DetachSecurityGroup(ctx context.Context, securityGroup *SecurityGroup) error {
	resp, err := i.c.DetachInstanceFromSecurityGroupWithResponse(
		apiv2.WithZone(ctx, i.zone), *securityGroup.ID, papi.DetachInstanceFromSecurityGroupJSONRequestBody{
			Instance: papi.Instance{Id: i.ID},
		})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// ElasticIPs returns the list of Elastic IPs attached to the Compute instance.
func (i *Instance) ElasticIPs(ctx context.Context) ([]*ElasticIP, error) {
	if i.ElasticIPIDs != nil {
		res, err := i.c.fetchFromIDs(ctx, i.zone, *i.ElasticIPIDs, new(ElasticIP))
		return res.([]*ElasticIP), err
	}
	return nil, nil
}

// PrivateNetworks returns the list of Private Networks attached to the Compute instance.
func (i *Instance) PrivateNetworks(ctx context.Context) ([]*PrivateNetwork, error) {
	if i.PrivateNetworkIDs != nil {
		res, err := i.c.fetchFromIDs(ctx, i.zone, *i.PrivateNetworkIDs, new(PrivateNetwork))
		return res.([]*PrivateNetwork), err
	}
	return nil, nil
}

// Reboot reboots the Compute instance.
func (i *Instance) Reboot(ctx context.Context) error {
	resp, err := i.c.RebootInstanceWithResponse(apiv2.WithZone(ctx, i.zone), *i.ID)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// Reset resets the Compute instance to a base template state (the current instance template if not specified),
// and optionally resizes the disk size.
func (i *Instance) Reset(ctx context.Context, template *Template, diskSize int64) error {
	resp, err := i.c.ResetInstanceWithResponse(
		apiv2.WithZone(ctx, i.zone),
		*i.ID,
		papi.ResetInstanceJSONRequestBody{
			DiskSize: func() (v *int64) {
				if diskSize > 0 {
					v = &diskSize
				}
				return
			}(),
			Template: func() (v *papi.Template) {
				if template != nil {
					v = &papi.Template{Id: template.ID}
				}
				return
			}(),
		})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// ResizeDisk resizes the Compute instance's disk to a larger size.
func (i *Instance) ResizeDisk(ctx context.Context, size int64) error {
	resp, err := i.c.ResizeInstanceDiskWithResponse(
		apiv2.WithZone(ctx, i.zone),
		*i.ID,
		papi.ResizeInstanceDiskJSONRequestBody{DiskSize: size})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// RevertToSnapshot reverts the Compute instance storage volume to the specified Snapshot.
func (i *Instance) RevertToSnapshot(ctx context.Context, snapshot *Snapshot) error {
	resp, err := i.c.RevertInstanceToSnapshotWithResponse(
		apiv2.WithZone(ctx, i.zone),
		*i.ID,
		papi.RevertInstanceToSnapshotJSONRequestBody{Id: *snapshot.ID})
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// Scale scales the Compute instance type.
func (i *Instance) Scale(ctx context.Context, instanceType *InstanceType) error {
	resp, err := i.c.ScaleInstanceWithResponse(
		apiv2.WithZone(ctx, i.zone),
		*i.ID,
		papi.ScaleInstanceJSONRequestBody{InstanceType: papi.InstanceType{Id: instanceType.ID}},
	)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// SecurityGroups returns the list of Security Groups attached to the Compute instance.
func (i *Instance) SecurityGroups(ctx context.Context) ([]*SecurityGroup, error) {
	if i.SecurityGroupIDs != nil {
		res, err := i.c.fetchFromIDs(ctx, i.zone, *i.SecurityGroupIDs, new(SecurityGroup))
		return res.([]*SecurityGroup), err
	}
	return nil, nil
}

// Start starts the Compute instance.
func (i *Instance) Start(ctx context.Context) error {
	resp, err := i.c.StartInstanceWithResponse(apiv2.WithZone(ctx, i.zone), *i.ID)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// Stop stops the Compute instance.
func (i *Instance) Stop(ctx context.Context) error {
	resp, err := i.c.StopInstanceWithResponse(apiv2.WithZone(ctx, i.zone), *i.ID)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(i.c.timeout).
		WithInterval(i.c.pollInterval).
		Poll(ctx, i.c.OperationPoller(i.zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// ToAPIMock returns the low-level representation of the resource. This is intended for testing purposes.
func (i Instance) ToAPIMock() interface{} {
	return papi.Instance{
		AntiAffinityGroups: func() *[]papi.AntiAffinityGroup {
			if i.AntiAffinityGroupIDs != nil {
				list := make([]papi.AntiAffinityGroup, len(*i.AntiAffinityGroupIDs))
				for j, id := range *i.AntiAffinityGroupIDs {
					id := id
					list[j] = papi.AntiAffinityGroup{Id: &id}
				}
				return &list
			}
			return nil
		}(),
		CreatedAt:    i.CreatedAt,
		DeployTarget: &papi.DeployTarget{Id: i.DeployTargetID},
		DiskSize:     i.DiskSize,
		ElasticIps: func() *[]papi.ElasticIp {
			if i.ElasticIPIDs != nil {
				list := make([]papi.ElasticIp, len(*i.ElasticIPIDs))
				for j, id := range *i.ElasticIPIDs {
					id := id
					list[j] = papi.ElasticIp{Id: &id}
				}
				return &list
			}
			return nil
		}(),
		Id:           i.ID,
		InstanceType: &papi.InstanceType{Id: i.InstanceTypeID},
		Ipv6Address: func() *string {
			if i.IPv6Address != nil {
				v := i.IPv6Address.String()
				return &v
			}
			return nil
		}(),
		Labels: func() *papi.Labels {
			if i.Labels != nil {
				return &papi.Labels{AdditionalProperties: *i.Labels}
			}
			return nil
		}(),
		Manager: func() *papi.Manager {
			if i.Manager != nil {
				return &papi.Manager{
					Id:   &i.Manager.ID,
					Type: (*papi.ManagerType)(&i.Manager.Type),
				}
			}
			return nil
		}(),
		Name: i.Name,
		PrivateNetworks: func() *[]papi.PrivateNetwork {
			if i.PrivateNetworkIDs != nil {
				list := make([]papi.PrivateNetwork, len(*i.PrivateNetworkIDs))
				for j, id := range *i.PrivateNetworkIDs {
					id := id
					list[j] = papi.PrivateNetwork{Id: &id}
				}
				return &list
			}
			return nil
		}(),
		PublicIp: func() *string {
			if i.PublicIPAddress != nil {
				v := i.PublicIPAddress.String()
				return &v
			}
			return nil
		}(),
		SecurityGroups: func() *[]papi.SecurityGroup {
			if i.SecurityGroupIDs != nil {
				list := make([]papi.SecurityGroup, len(*i.SecurityGroupIDs))
				for j, id := range *i.SecurityGroupIDs {
					id := id
					list[j] = papi.SecurityGroup{Id: &id}
				}
				return &list
			}
			return nil
		}(),
		Snapshots: func() *[]papi.Snapshot {
			if i.SnapshotIDs != nil {
				list := make([]papi.Snapshot, len(*i.SnapshotIDs))
				for j, id := range *i.SnapshotIDs {
					id := id
					list[j] = papi.Snapshot{Id: &id}
				}
				return &list
			}
			return nil
		}(),
		SshKey: func() *papi.SshKey {
			if i.SSHKey != nil {
				return &papi.SshKey{Name: i.SSHKey}
			}
			return nil
		}(),
		State:    (*papi.InstanceState)(i.State),
		Template: &papi.Template{Id: i.TemplateID},
		UserData: i.UserData,
	}
}

// CreateInstance creates a Compute instance in the specified zone.
func (c *Client) CreateInstance(ctx context.Context, zone string, instance *Instance) (*Instance, error) {
	if err := validateOperationParams(instance, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreateInstanceWithResponse(
		apiv2.WithZone(ctx, zone),
		papi.CreateInstanceJSONRequestBody{
			AntiAffinityGroups: func() (v *[]papi.AntiAffinityGroup) {
				if instance.AntiAffinityGroupIDs != nil {
					ids := make([]papi.AntiAffinityGroup, len(*instance.AntiAffinityGroupIDs))
					for i, item := range *instance.AntiAffinityGroupIDs {
						item := item
						ids[i] = papi.AntiAffinityGroup{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			DeployTarget: func() (v *papi.DeployTarget) {
				if instance.DeployTargetID != nil {
					v = &papi.DeployTarget{Id: instance.DeployTargetID}
				}
				return
			}(),
			DiskSize:     *instance.DiskSize,
			InstanceType: papi.InstanceType{Id: instance.InstanceTypeID},
			Ipv6Enabled:  instance.IPv6Enabled,
			Labels: func() (v *papi.Labels) {
				if instance.Labels != nil {
					v = &papi.Labels{AdditionalProperties: *instance.Labels}
				}
				return
			}(),
			Name: instance.Name,
			SecurityGroups: func() (v *[]papi.SecurityGroup) {
				if instance.SecurityGroupIDs != nil {
					ids := make([]papi.SecurityGroup, len(*instance.SecurityGroupIDs))
					for i, item := range *instance.SecurityGroupIDs {
						item := item
						ids[i] = papi.SecurityGroup{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			SshKey: func() (v *papi.SshKey) {
				if instance.SSHKey != nil {
					v = &papi.SshKey{Name: instance.SSHKey}
				}
				return
			}(),
			Template: papi.Template{Id: instance.TemplateID},
			UserData: instance.UserData,
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

	return c.GetInstance(ctx, zone, *res.(*papi.Reference).Id)
}

// ListInstances returns the list of existing Compute instances in the specified zone.
func (c *Client) ListInstances(ctx context.Context, zone string) ([]*Instance, error) {
	list := make([]*Instance, 0)

	resp, err := c.ListInstancesWithResponse(apiv2.WithZone(ctx, zone), &papi.ListInstancesParams{})
	if err != nil {
		return nil, err
	}

	if resp.JSON200.Instances != nil {
		for i := range *resp.JSON200.Instances {
			list = append(list, instanceFromAPI(c, zone, &(*resp.JSON200.Instances)[i]))
		}
	}

	return list, nil
}

// GetInstance returns the Instance  corresponding to the specified ID in the specified zone.
func (c *Client) GetInstance(ctx context.Context, zone, id string) (*Instance, error) {
	resp, err := c.GetInstanceWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return instanceFromAPI(c, zone, resp.JSON200), nil
}

// FindInstance attempts to find a Compute instance by name or ID in the specified zone.
// In case the identifier is a name and multiple resources match, an ErrTooManyFound error is returned.
func (c *Client) FindInstance(ctx context.Context, zone, v string) (*Instance, error) {
	res, err := c.ListInstances(ctx, zone)
	if err != nil {
		return nil, err
	}

	var found *Instance
	for _, r := range res {
		if *r.ID == v {
			return c.GetInstance(ctx, zone, *r.ID)
		}

		// Historically, the Exoscale API allowed users to create multiple Compute instances sharing a common name.
		// This function being expected to return one resource at most, in case the specified identifier is a name
		// we have to check that there aren't more than one matching result before returning it.
		if *r.Name == v {
			if found != nil {
				return nil, apiv2.ErrTooManyFound
			}
			found = r
		}
	}

	if found != nil {
		return c.GetInstance(ctx, zone, *found.ID)
	}

	return nil, apiv2.ErrNotFound
}

// UpdateInstance updates the specified Compute instance in the specified zone.
func (c *Client) UpdateInstance(ctx context.Context, zone string, instance *Instance) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}

	resp, err := c.UpdateInstanceWithResponse(
		apiv2.WithZone(ctx, zone),
		*instance.ID,
		papi.UpdateInstanceJSONRequestBody{
			Labels: func() (v *papi.Labels) {
				if instance.Labels != nil {
					v = &papi.Labels{AdditionalProperties: *instance.Labels}
				}
				return
			}(),
			Name:     instance.Name,
			UserData: instance.UserData,
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

// DeleteInstance deletes the specified Compute instance in the specified zone.
func (c *Client) DeleteInstance(ctx context.Context, zone, id string) error {
	resp, err := c.DeleteInstanceWithResponse(apiv2.WithZone(ctx, zone), id)
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
