package v2

import (
	"context"
	"net"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
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
	DiskSize             *int64 `req-for:"create"`
	ElasticIPIDs         *[]string
	ID                   *string `req-for:"update,delete"`
	IPv6Address          *net.IP
	IPv6Enabled          *bool
	InstanceTypeID       *string `req-for:"create"`
	Labels               *map[string]string
	Manager              *InstanceManager
	Name                 *string `req-for:"create"`
	PrivateNetworkIDs    *[]string
	PublicIPAddress      *net.IP
	PublicIPAssignment   *string
	SSHKey               *string
	SecurityGroupIDs     *[]string
	SnapshotIDs          *[]string
	State                *string
	TemplateID           *string `req-for:"create"`
	UserData             *string
	Zone                 *string
}

// AttachInstanceToPrivateNetworkOpt represents an AttachInstanceToPrivateNetwork operation option.
type AttachInstanceToPrivateNetworkOpt func(*oapi.AttachInstanceToPrivateNetworkJSONRequestBody)

// AttachInstanceToPrivateNetworkWithIPAddress sets the Private Network lease IP address to
// request for the Compute instance.
func AttachInstanceToPrivateNetworkWithIPAddress(v net.IP) AttachInstanceToPrivateNetworkOpt {
	return func(b *oapi.AttachInstanceToPrivateNetworkJSONRequestBody) {
		if v != nil {
			ip := v.String()
			b.Ip = &ip
		}
	}
}

// ListInstancesOpt represents a ListInstances operation option.
type ListInstancesOpt func(params *oapi.ListInstancesParams)

// ListInstancesByManagerID sets a Compute instances listing filter based on a manager ID.
func ListInstancesByManagerID(v string) ListInstancesOpt {
	return func(p *oapi.ListInstancesParams) {
		if v != "" {
			p.ManagerId = &v
		}
	}
}

// ListInstancesByManagerType sets a Compute instances listing filter based on a manager type.
func ListInstancesByManagerType(v string) ListInstancesOpt {
	return func(p *oapi.ListInstancesParams) {
		if v != "" {
			p.ManagerType = (*oapi.ListInstancesParamsManagerType)(&v)
		}
	}
}

// ListInstancesByIpAddress sets a Compute instances listing filter based on an IP Address.
func ListInstancesByIpAddress(v string) ListInstancesOpt { //nolint:revive
	return func(p *oapi.ListInstancesParams) {
		if v != "" {
			p.IpAddress = &v
		}
	}
}

// ResetInstanceOpt represents a ResetInstance operation option.
type ResetInstanceOpt func(*oapi.ResetInstanceJSONRequestBody)

// ResetInstanceWithDiskSize sets a new disk size to set the Compute instance.
func ResetInstanceWithDiskSize(v int64) ResetInstanceOpt {
	return func(b *oapi.ResetInstanceJSONRequestBody) {
		if v > 0 {
			b.DiskSize = &v
		}
	}
}

// ResetInstanceWithTemplate sets a template to reset the Compute instance to.
func ResetInstanceWithTemplate(v *Template) ResetInstanceOpt {
	return func(b *oapi.ResetInstanceJSONRequestBody) {
		if v != nil {
			b.Template = &oapi.Template{Id: v.ID}
		}
	}
}

// StartInstanceOpt represents a StartInstance operation option.
type StartInstanceOpt func(*oapi.StartInstanceJSONRequestBody)

// StartInstanceWithRescueProfile sets the rescue profile to start a Compute instance with.
func StartInstanceWithRescueProfile(v string) StartInstanceOpt {
	return func(b *oapi.StartInstanceJSONRequestBody) {
		if v != "" {
			b.RescueProfile = (*oapi.StartInstanceJSONBodyRescueProfile)(&v)
		}
	}
}

func instanceFromAPI(i *oapi.Instance, zone string) *Instance {
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
				v = &i.DeployTarget.Id
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
		PublicIPAssignment: (*string)(i.PublicIpAssignment),
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
		Zone:       &zone,
	}
}

// AttachInstanceToElasticIP attaches a Compute instance to the specified Elastic IP.
func (c *Client) AttachInstanceToElasticIP(
	ctx context.Context,
	zone string,
	instance *Instance,
	elasticIP *ElasticIP,
) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(elasticIP, "update"); err != nil {
		return err
	}

	resp, err := c.AttachInstanceToElasticIpWithResponse(
		apiv2.WithZone(ctx, zone), *elasticIP.ID, oapi.AttachInstanceToElasticIpJSONRequestBody{
			Instance: oapi.Instance{Id: instance.ID},
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

// AttachInstanceToPrivateNetwork attaches a Compute instance to the specified Private Network.
func (c *Client) AttachInstanceToPrivateNetwork(
	ctx context.Context,
	zone string,
	instance *Instance,
	privateNetwork *PrivateNetwork,
	opts ...AttachInstanceToPrivateNetworkOpt,
) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(privateNetwork, "update"); err != nil {
		return err
	}

	var body oapi.AttachInstanceToPrivateNetworkJSONRequestBody
	for _, opt := range opts {
		opt(&body)
	}

	a := oapi.AttachInstanceToPrivateNetworkJSONBody{}
	a.Instance.Id = instance.ID
	body.Instance = a.Instance

	resp, err := c.AttachInstanceToPrivateNetworkWithResponse(apiv2.WithZone(ctx, zone), *privateNetwork.ID, body)
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

// AttachInstanceToSecurityGroup attaches a Compute instance to the specified Security Group.
func (c *Client) AttachInstanceToSecurityGroup(
	ctx context.Context,
	zone string,
	instance *Instance,
	securityGroup *SecurityGroup,
) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(securityGroup, "update"); err != nil {
		return err
	}

	resp, err := c.AttachInstanceToSecurityGroupWithResponse(
		apiv2.WithZone(ctx, zone), *securityGroup.ID, oapi.AttachInstanceToSecurityGroupJSONRequestBody{
			Instance: oapi.Instance{Id: instance.ID},
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

// CreateInstance creates a Compute instance.
func (c *Client) CreateInstance(ctx context.Context, zone string, instance *Instance) (*Instance, error) {
	if err := validateOperationParams(instance, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreateInstanceWithResponse(
		apiv2.WithZone(ctx, zone),
		oapi.CreateInstanceJSONRequestBody{
			AntiAffinityGroups: func() (v *[]oapi.AntiAffinityGroup) {
				if instance.AntiAffinityGroupIDs != nil {
					ids := make([]oapi.AntiAffinityGroup, len(*instance.AntiAffinityGroupIDs))
					for i, item := range *instance.AntiAffinityGroupIDs {
						item := item
						ids[i] = oapi.AntiAffinityGroup{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			DeployTarget: func() (v *oapi.DeployTarget) {
				if instance.DeployTargetID != nil {
					v = &oapi.DeployTarget{Id: *instance.DeployTargetID}
				}
				return
			}(),
			DiskSize:     *instance.DiskSize,
			InstanceType: oapi.InstanceType{Id: instance.InstanceTypeID},
			Ipv6Enabled:  instance.IPv6Enabled,
			Labels: func() (v *oapi.Labels) {
				if instance.Labels != nil {
					v = &oapi.Labels{AdditionalProperties: *instance.Labels}
				}
				return
			}(),
			Name:               instance.Name,
			PublicIpAssignment: (*oapi.PublicIpAssignment)(instance.PublicIPAssignment),
			SecurityGroups: func() (v *[]oapi.SecurityGroup) {
				if instance.SecurityGroupIDs != nil {
					ids := make([]oapi.SecurityGroup, len(*instance.SecurityGroupIDs))
					for i, item := range *instance.SecurityGroupIDs {
						item := item
						ids[i] = oapi.SecurityGroup{Id: &item}
					}
					v = &ids
				}
				return
			}(),
			SshKey: func() (v *oapi.SshKey) {
				if instance.SSHKey != nil {
					v = &oapi.SshKey{Name: instance.SSHKey}
				}
				return
			}(),
			Template: oapi.Template{Id: instance.TemplateID},
			UserData: instance.UserData,
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

	return c.GetInstance(ctx, zone, *res.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
}

// CreateInstanceSnapshot creates a Snapshot of a Compute instance storage volume.
func (c *Client) CreateInstanceSnapshot(ctx context.Context, zone string, instance *Instance) (*Snapshot, error) {
	if err := validateOperationParams(instance, "update"); err != nil {
		return nil, err
	}

	resp, err := c.CreateSnapshotWithResponse(apiv2.WithZone(ctx, zone), *instance.ID)
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

	return c.GetSnapshot(ctx, zone, *res.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
}

// DeleteInstance deletes a Compute instance.
func (c *Client) DeleteInstance(ctx context.Context, zone string, instance *Instance) error {
	if err := validateOperationParams(instance, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteInstanceWithResponse(apiv2.WithZone(ctx, zone), *instance.ID)
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

// DetachInstanceFromElasticIP detaches a Compute instance from the specified Elastic IP.
func (c *Client) DetachInstanceFromElasticIP(
	ctx context.Context,
	zone string,
	instance *Instance,
	elasticIP *ElasticIP,
) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(elasticIP, "update"); err != nil {
		return err
	}

	resp, err := c.DetachInstanceFromElasticIpWithResponse(
		apiv2.WithZone(ctx, zone), *elasticIP.ID, oapi.DetachInstanceFromElasticIpJSONRequestBody{
			Instance: oapi.Instance{Id: instance.ID},
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

// DetachInstanceFromPrivateNetwork detaches a Compute instance from the specified Private Network.
func (c *Client) DetachInstanceFromPrivateNetwork(
	ctx context.Context,
	zone string,
	instance *Instance,
	privateNetwork *PrivateNetwork,
) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(privateNetwork, "update"); err != nil {
		return err
	}

	resp, err := c.DetachInstanceFromPrivateNetworkWithResponse(
		apiv2.WithZone(ctx, zone), *privateNetwork.ID, oapi.DetachInstanceFromPrivateNetworkJSONRequestBody{
			Instance: oapi.Instance{Id: instance.ID},
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

// DetachInstanceFromSecurityGroup detaches a Compute instance from the specified Security Group.
func (c *Client) DetachInstanceFromSecurityGroup(
	ctx context.Context,
	zone string,
	instance *Instance,
	securityGroup *SecurityGroup,
) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(securityGroup, "update"); err != nil {
		return err
	}

	resp, err := c.DetachInstanceFromSecurityGroupWithResponse(
		apiv2.WithZone(ctx, zone), *securityGroup.ID, oapi.DetachInstanceFromSecurityGroupJSONRequestBody{
			Instance: oapi.Instance{Id: instance.ID},
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

// FindInstance attempts to find a Compute instance by name or ID.
// In case the identifier is a name and multiple resources match, an ErrTooManyFound error is returned.
func (c *Client) FindInstance(ctx context.Context, zone, x string) (*Instance, error) {
	res, err := c.ListInstances(ctx, zone)
	if err != nil {
		return nil, err
	}

	var found *Instance
	for _, r := range res {
		if *r.ID == x {
			return c.GetInstance(ctx, zone, *r.ID)
		}

		// Historically, the Exoscale API allowed users to create multiple Compute instances sharing a common name.
		// This function being expected to return one resource at most, in case the specified identifier is a name
		// we have to check that there aren't more than one matching result before returning it.
		if *r.Name == x {
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

// GetInstance returns the Instance corresponding to the specified ID.
func (c *Client) GetInstance(ctx context.Context, zone, id string) (*Instance, error) {
	resp, err := c.GetInstanceWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return instanceFromAPI(resp.JSON200, zone), nil
}

// ListInstances returns the list of existing Compute instances.
func (c *Client) ListInstances(ctx context.Context, zone string, opts ...ListInstancesOpt) ([]*Instance, error) {
	list := make([]*Instance, 0)

	var params oapi.ListInstancesParams
	for _, opt := range opts {
		opt(&params)
	}

	resp, err := c.ListInstancesWithResponse(apiv2.WithZone(ctx, zone), &params)
	if err != nil {
		return nil, err
	}

	if resp.JSON200.Instances != nil {
		for i := range *resp.JSON200.Instances {
			list = append(list, instanceFromAPI(&(*resp.JSON200.Instances)[i], zone))
		}
	}

	return list, nil
}

// RebootInstance reboots a Compute instance.
func (c *Client) RebootInstance(ctx context.Context, zone string, instance *Instance) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}

	resp, err := c.RebootInstanceWithResponse(apiv2.WithZone(ctx, zone), *instance.ID)
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

// ResetInstance resets a Compute instance to a base template state (the instance's current
// template if not specified).
func (c *Client) ResetInstance(ctx context.Context, zone string, instance *Instance, opts ...ResetInstanceOpt) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}

	var body oapi.ResetInstanceJSONRequestBody
	for _, opt := range opts {
		opt(&body)
	}

	resp, err := c.ResetInstanceWithResponse(apiv2.WithZone(ctx, zone), *instance.ID, body)
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

// ResizeInstanceDisk resizes a Compute instance's disk to a larger size.
func (c *Client) ResizeInstanceDisk(ctx context.Context, zone string, instance *Instance, size int64) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}

	resp, err := c.ResizeInstanceDiskWithResponse(
		apiv2.WithZone(ctx, zone),
		*instance.ID,
		oapi.ResizeInstanceDiskJSONRequestBody{DiskSize: size})
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

// RevertInstanceToSnapshot reverts a Compute instance storage volume to the specified Snapshot.
func (c *Client) RevertInstanceToSnapshot(
	ctx context.Context,
	zone string,
	instance *Instance,
	snapshot *Snapshot,
) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(snapshot, "delete"); err != nil {
		return err
	}

	resp, err := c.RevertInstanceToSnapshotWithResponse(
		apiv2.WithZone(ctx, zone),
		*instance.ID,
		oapi.RevertInstanceToSnapshotJSONRequestBody{Id: *snapshot.ID})
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

// ScaleInstance scales a Compute instance type.
func (c *Client) ScaleInstance(ctx context.Context, zone string, instance *Instance, instanceType *InstanceType) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}

	resp, err := c.ScaleInstanceWithResponse(
		apiv2.WithZone(ctx, zone),
		*instance.ID,
		oapi.ScaleInstanceJSONRequestBody{InstanceType: oapi.InstanceType{Id: instanceType.ID}},
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

// StartInstance starts a Compute instance.
func (c *Client) StartInstance(ctx context.Context, zone string, instance *Instance, opts ...StartInstanceOpt) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}

	var body oapi.StartInstanceJSONRequestBody
	for _, opt := range opts {
		opt(&body)
	}

	resp, err := c.StartInstanceWithResponse(apiv2.WithZone(ctx, zone), *instance.ID, body)
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

// StopInstance stops a Compute instance.
func (c *Client) StopInstance(ctx context.Context, zone string, instance *Instance) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}

	resp, err := c.StopInstanceWithResponse(apiv2.WithZone(ctx, zone), *instance.ID)
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

// RevealInstancePassword retrieves a recently started instance's root password if possible.
func (c *Client) RevealInstancePassword(ctx context.Context, zone string, instance *Instance) (string, error) {
	resp, err := c.RevealInstancePasswordWithResponse(apiv2.WithZone(ctx, zone), *instance.ID)
	if err != nil {
		return "", err
	}
	// If the password is unavailable, return an empty string
	if resp.JSON200.Password == nil {
		return "", nil
	}
	return *resp.JSON200.Password, nil
}

// UpdateInstance updates a Compute instance.
func (c *Client) UpdateInstance(ctx context.Context, zone string, instance *Instance) error {
	if err := validateOperationParams(instance, "update"); err != nil {
		return err
	}

	resp, err := c.UpdateInstanceWithResponse(
		apiv2.WithZone(ctx, zone),
		*instance.ID,
		oapi.UpdateInstanceJSONRequestBody{
			Labels: func() (v *oapi.Labels) {
				if instance.Labels != nil {
					v = &oapi.Labels{AdditionalProperties: *instance.Labels}
				}
				return
			}(),
			Name:     instance.Name,
			UserData: instance.UserData,
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

// GetInstanceReverseDNS returns the Reverse DNS record corresponding to the specified Instance ID.
func (c *Client) GetInstanceReverseDNS(ctx context.Context, zone, id string) (string, error) {
	resp, err := c.GetReverseDnsInstanceWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return "", err
	}

	if resp.JSON200 == nil || resp.JSON200.DomainName == nil {
		return "", apiv2.ErrNotFound
	}

	return string(*resp.JSON200.DomainName), nil
}

// DeleteInstanceReverseDNS deletes a Reverse DNS record of a Compute Instance.
func (c *Client) DeleteInstanceReverseDNS(ctx context.Context, zone string, id string) error {
	resp, err := c.DeleteReverseDnsInstanceWithResponse(apiv2.WithZone(ctx, zone), id)
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

// UpdateInstanceReverseDNS updates a Reverse DNS record for a Compute Instance.
func (c *Client) UpdateInstanceReverseDNS(ctx context.Context, zone, id, domain string) error {
	resp, err := c.UpdateReverseDnsInstanceWithResponse(
		apiv2.WithZone(ctx, zone),
		id,
		oapi.UpdateReverseDnsInstanceJSONRequestBody{
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
