// Copyright 2022-2023 The sacloud/iaas-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fake

import (
	"context"
	"fmt"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// Find is fake implementation
func (o *DiskOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.DiskFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.Disk
	for _, res := range results {
		dest := &iaas.Disk{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.DiskFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		Disks: values,
	}, nil
}

// Create is fake implementation
func (o *DiskOp) Create(ctx context.Context, zone string, param *iaas.DiskCreateRequest, distantFrom []types.ID) (*iaas.Disk, error) {
	result := &iaas.Disk{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillDiskPlan)
	result.Availability = types.Availabilities.Migrating

	result.Storage = &iaas.Storage{
		ID:   types.ID(123456789012),
		Name: "dummy",
	}

	if result.Connection == types.EDiskConnection("") {
		result.Connection = types.DiskConnections.VirtIO
	}
	if !param.SourceArchiveID.IsEmpty() {
		archiveOp := NewArchiveOp()
		source, err := archiveOp.Read(ctx, zone, param.SourceArchiveID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), "SourceArchive is not found")
		}
		result.SourceArchiveAvailability = source.Availability
	}
	if !param.SourceDiskID.IsEmpty() {
		source, err := o.Read(ctx, zone, param.SourceDiskID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), "SourceDisk is not found")
		}
		result.SourceDiskAvailability = source.Availability
	}
	if !param.ServerID.IsEmpty() {
		server, err := NewServerOp().Read(ctx, zone, param.ServerID)
		if err != nil {
			return nil, newErrorConflict(o.key, types.ID(0), "Server is not found")
		}
		server.Disks = append(server.Disks, &iaas.ServerConnectedDisk{
			ID:              result.ID,
			Name:            result.Name,
			Availability:    result.Availability,
			Connection:      result.Connection,
			ConnectionOrder: result.ConnectionOrder,
			ReinstallCount:  result.ReinstallCount,
			SizeMB:          result.SizeMB,
			DiskPlanID:      result.DiskPlanID,
			Storage:         result.Storage,
		})
		putServer(zone, server)
	}

	putDisk(zone, result)

	id := result.ID
	startDiskCopy(o.key, zone, func() (interface{}, error) {
		disk, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return disk, nil
	})

	return result, nil
}

// Config is fake implementation
func (o *DiskOp) Config(ctx context.Context, zone string, id types.ID, edit *iaas.DiskEditRequest) error {
	disk, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if disk.ServerID.IsEmpty() {
		return nil
	}

	serverOp := NewServerOp()
	server, err := serverOp.Read(ctx, zone, disk.ServerID)
	if err != nil {
		return err
	}

	if edit.HostName != "" {
		server.HostName = edit.HostName
		putServer(zone, server)
	}

	if len(server.Interfaces) > 0 {
		nic := server.Interfaces[0]
		if nic.SwitchScope == types.Scopes.Shared {
			nic.IPAddress = pool().nextSharedIP().String()
		} else {
			nic.UserIPAddress = edit.UserIPAddress
		}

		swOp := NewSwitchOp()
		sw, err := swOp.Read(ctx, zone, nic.SwitchID)
		if err != nil {
			return err
		}

		if len(sw.Subnets) == 0 {
			nic.UserSubnetDefaultRoute = edit.UserSubnet.DefaultRoute
			nic.UserSubnetNetworkMaskLen = edit.UserSubnet.NetworkMaskLen
		} else {
			nic.UserSubnetDefaultRoute = sw.Subnets[0].DefaultRoute
			nic.UserSubnetNetworkMaskLen = sw.Subnets[0].NetworkMaskLen
			nic.SubnetDefaultRoute = sw.Subnets[0].DefaultRoute
			nic.SubnetNetworkAddress = sw.Subnets[0].NetworkAddress
		}

		putServer(zone, server)
	}

	return nil
}

// CreateWithConfig is fake implementation
func (o *DiskOp) CreateWithConfig(ctx context.Context, zone string, createParam *iaas.DiskCreateRequest, editParam *iaas.DiskEditRequest, bootAtAvailable bool, distantFrom []types.ID) (*iaas.Disk, error) {
	// check
	if !createParam.ServerID.IsEmpty() {
		serverOp := NewServerOp()
		_, err := serverOp.Read(ctx, zone, createParam.ServerID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), fmt.Sprintf("Server %s is not found", createParam.ServerID))
		}
	}

	result, err := o.Create(ctx, zone, createParam, distantFrom)
	if err != nil {
		return nil, err
	}

	if err := o.Config(ctx, zone, result.ID, editParam); err != nil {
		return nil, err
	}

	if !createParam.ServerID.IsEmpty() && bootAtAvailable {
		waiter := iaas.WaiterForReady(func() (interface{}, error) {
			disk, err := o.Read(ctx, zone, result.ID)
			if err != nil {
				return nil, err
			}
			return disk, nil
		})
		res, err := waiter.WaitForState(ctx)
		if err != nil {
			return nil, err
		}
		result = res.(*iaas.Disk)

		// boot server
		serverOp := NewServerOp()
		if err := serverOp.Boot(ctx, zone, createParam.ServerID); err != nil {
			return nil, err
		}
	}
	return result, nil
}

// ResizePartition is fake implementation
func (o *DiskOp) ResizePartition(ctx context.Context, zone string, id types.ID, param *iaas.DiskResizePartitionRequest) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	return nil
}

// ConnectToServer is fake implementation
func (o *DiskOp) ConnectToServer(ctx context.Context, zone string, id types.ID, serverID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	serverOp := NewServerOp()
	server, err := serverOp.Read(ctx, zone, serverID)
	if err != nil {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("Server[%d] is not exists", serverID))
	}

	for _, connected := range server.Disks {
		if connected.ID == value.ID {
			return newErrorBadRequest(o.key, id, fmt.Sprintf("Disk[%d] is already connected to Server[%d]", id, serverID))
		}
	}

	// TODO とりあえず同時実行制御は考慮しない。更新対象リソースが増えるようであれば実装方法を考える

	connectedDisk := &iaas.ServerConnectedDisk{}
	copySameNameField(value, connectedDisk)
	server.Disks = append(server.Disks, connectedDisk)
	putServer(zone, server)
	value.ServerID = serverID
	value.ServerName = server.Name
	putDisk(zone, value)

	return nil
}

// DisconnectFromServer is fake implementation
func (o *DiskOp) DisconnectFromServer(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	if value.ServerID.IsEmpty() {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("Disk[%d] is not connected to Server", id))
	}

	serverOp := NewServerOp()
	server, err := serverOp.Read(ctx, zone, value.ServerID)
	if err != nil {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("Server[%d] is not exists", value.ServerID))
	}

	var disks []*iaas.ServerConnectedDisk
	for _, connected := range server.Disks {
		if connected.ID != value.ID {
			connectedDisk := &iaas.ServerConnectedDisk{}
			copySameNameField(value, connectedDisk)
			server.Disks = append(server.Disks, connectedDisk)
			disks = append(disks, connected)
		}
	}
	if len(disks) == len(server.Disks) {
		return newInternalServerError(o.key, id, fmt.Sprintf("Disk[%d] is not found on server's connected disks", id))
	}

	server.Disks = disks
	putServer(zone, server)
	value.ServerID = types.ID(0)
	value.ServerName = ""
	putDisk(zone, value)

	return nil
}

// Read is fake implementation
func (o *DiskOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.Disk, error) {
	value := getDiskByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &iaas.Disk{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *DiskOp) Update(ctx context.Context, zone string, id types.ID, param *iaas.DiskUpdateRequest) (*iaas.Disk, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putDisk(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *DiskOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, zone, id)
	return nil
}

// Monitor is fake implementation
func (o *DiskOp) Monitor(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.DiskActivity, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &iaas.DiskActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &iaas.MonitorDiskValue{
			Time:  now.Add(time.Duration(i*-5) * time.Minute),
			Read:  float64(random(1000)),
			Write: float64(random(1000)),
		})
	}

	return res, nil
}

// MonitorDisk is fake implementation
func (o *DiskOp) MonitorDisk(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.DiskActivity, error) {
	return o.Monitor(ctx, zone, id, condition)
}
