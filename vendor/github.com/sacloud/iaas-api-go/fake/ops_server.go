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
func (o *ServerOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.ServerFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.Server
	for _, res := range results {
		dest := &iaas.Server{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.ServerFindResult{
		Total:   len(results),
		Count:   len(results),
		From:    0,
		Servers: values,
	}, nil
}

// Create is fake implementation
func (o *ServerOp) Create(ctx context.Context, zone string, param *iaas.ServerCreateRequest) (*iaas.Server, error) {
	result := &iaas.Server{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Availability = types.Availabilities.Migrating
	if param.ServerPlanGeneration == types.PlanGenerations.Default {
		switch zone {
		case "is1a":
			result.ServerPlanGeneration = types.PlanGenerations.G200
		default:
			result.ServerPlanGeneration = types.PlanGenerations.G100
		}
	}
	// TODO プランAPIを実装したら修正する
	result.ServerPlanID = types.StringID(fmt.Sprintf("%03d%03d%03d", result.ServerPlanGeneration, result.GetMemoryGB(), result.CPU))
	result.ServerPlanName = fmt.Sprintf("世代:%03d メモリ:%03d CPU:%03d", result.ServerPlanGeneration, result.GetMemoryGB(), result.CPU)

	// NIC操作のためにあらかじめ登録しておく
	putServer(zone, result)

	for _, cs := range param.ConnectedSwitches {
		ifOp := NewInterfaceOp()
		swOp := NewSwitchOp()

		ifCreateParam := &iaas.InterfaceCreateRequest{}
		if cs != nil {
			if cs.Scope != types.Scopes.Shared {
				_, err := swOp.Read(ctx, zone, cs.ID)
				if err != nil {
					return nil, newErrorConflict(o.key, types.ID(0), err.Error())
				}
			}
			ifCreateParam.ServerID = result.ID
		}

		iface, err := ifOp.Create(ctx, zone, ifCreateParam)
		if err != nil {
			return nil, newErrorConflict(o.key, types.ID(0), err.Error())
		}

		if cs != nil {
			if cs.Scope == types.Scopes.Shared {
				if err := ifOp.ConnectToSharedSegment(ctx, zone, iface.ID); err != nil {
					return nil, newErrorConflict(o.key, types.ID(0), err.Error())
				}
			} else {
				if err := ifOp.ConnectToSwitch(ctx, zone, iface.ID, cs.ID); err != nil {
					return nil, newErrorConflict(o.key, types.ID(0), err.Error())
				}
			}
		}

		iface, err = ifOp.Read(ctx, zone, iface.ID)
		if err != nil {
			return nil, newErrorConflict(o.key, types.ID(0), err.Error())
		}
		ifaceView := &iaas.InterfaceView{}
		copySameNameField(iface, ifaceView)

		// note: UserIPAddressとIPAddressはディスクの修正にて設定されるためここでは空となる。
		if cs != nil {
			if cs.Scope == types.Scopes.Shared {
				ifaceView.SwitchScope = sharedSegmentSwitch.Scope
				ifaceView.SwitchID = sharedSegmentSwitch.ID
				ifaceView.SwitchName = sharedSegmentSwitch.Name

				if len(sharedSegmentSwitch.Subnets) > 0 {
					ifaceView.UserSubnetDefaultRoute = sharedSegmentSwitch.Subnets[0].DefaultRoute
					ifaceView.UserSubnetNetworkMaskLen = sharedSegmentSwitch.Subnets[0].NetworkMaskLen
					ifaceView.SubnetDefaultRoute = sharedSegmentSwitch.Subnets[0].DefaultRoute
					ifaceView.SubnetNetworkAddress = sharedSegmentSwitch.Subnets[0].NetworkAddress
				}
			} else {
				ifaceView.SwitchScope = types.Scopes.User
				ifaceView.SwitchID = cs.ID

				sw, err := swOp.Read(ctx, zone, cs.ID)
				if err != nil {
					return nil, err
				}
				if len(sw.Subnets) > 0 {
					ifaceView.UserSubnetDefaultRoute = sw.Subnets[0].DefaultRoute
					ifaceView.UserSubnetNetworkMaskLen = sw.Subnets[0].NetworkMaskLen
					ifaceView.SubnetDefaultRoute = sw.Subnets[0].DefaultRoute
					ifaceView.SubnetNetworkAddress = sw.Subnets[0].NetworkAddress
				}
			}
		}

		result.Interfaces = append(result.Interfaces, ifaceView)
	}
	zoneOp := NewZoneOp()
	zones, _ := zoneOp.Find(ctx, nil)
	for _, z := range zones.Zones {
		if zone == z.Name {
			zoneInfo := &iaas.ZoneInfo{}
			copySameNameField(z, zoneInfo)
			result.Zone = zoneInfo
		}
	}

	result.Availability = types.Availabilities.Available
	putServer(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *ServerOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.Server, error) {
	value := getServerByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &iaas.Server{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *ServerOp) Update(ctx context.Context, zone string, id types.ID, param *iaas.ServerUpdateRequest) (*iaas.Server, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putServer(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *ServerOp) Delete(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	if value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, fmt.Sprintf("Server[%s] is still running", id))
	}

	ifOp := NewInterfaceOp()
	for _, iface := range value.Interfaces {
		if err := ifOp.Delete(ctx, zone, iface.ID); err != nil {
			return err
		}
	}

	diskOp := NewDiskOp()
	for _, disk := range value.Disks {
		if err := diskOp.DisconnectFromServer(ctx, zone, disk.ID); err != nil {
			return err
		}
	}

	ds().Delete(o.key, zone, id)
	return nil
}

// DeleteWithDisks is fake implementation
func (o *ServerOp) DeleteWithDisks(ctx context.Context, zone string, id types.ID, disks *iaas.ServerDeleteWithDisksRequest) error {
	if err := o.Delete(ctx, zone, id); err != nil {
		return err
	}
	diskOp := NewDiskOp()
	for _, diskID := range disks.IDs {
		if err := diskOp.Delete(ctx, zone, diskID); err != nil {
			return err
		}
	}
	return nil
}

// ChangePlan is fake implementation
func (o *ServerOp) ChangePlan(ctx context.Context, zone string, id types.ID, plan *iaas.ServerChangePlanRequest) (*iaas.Server, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	if value.InstanceStatus.IsUp() {
		return nil, newErrorConflict(o.key, id, fmt.Sprintf("Server[%d] is running", value.ID))
	}

	value.CPU = plan.CPU
	value.MemoryMB = plan.MemoryMB
	value.ServerPlanCommitment = plan.ServerPlanCommitment
	value.ServerPlanGeneration = plan.ServerPlanGeneration
	value.ServerPlanID = types.StringID(fmt.Sprintf("%03d%03d%03d", value.ServerPlanGeneration, value.GetMemoryGB(), value.CPU))
	value.ServerPlanName = fmt.Sprintf("世代:%03d メモリ:%03d CPU:%03d", value.ServerPlanGeneration, value.GetMemoryGB(), value.CPU)

	// ID変更
	ds().Delete(o.key, zone, value.ID)
	newServer := &iaas.Server{}
	copySameNameField(value, newServer)
	newServer.ID = pool().generateID()
	putServer(zone, newServer)

	// DiskのServerIDも変更
	searched, _ := NewDiskOp().Find(ctx, zone, nil)
	for _, disk := range searched.Disks {
		if disk.ServerID == value.ID {
			disk.ServerID = newServer.ID
			putDisk(zone, disk)
		}
	}
	for _, nic := range newServer.Interfaces {
		iface, err := NewInterfaceOp().Read(ctx, zone, nic.ID)
		if err == nil {
			iface.ServerID = newServer.ID
			putInterface(zone, iface)
		}
	}

	return newServer, nil
}

// InsertCDROM is fake implementation
func (o *ServerOp) InsertCDROM(ctx context.Context, zone string, id types.ID, insertParam *iaas.InsertCDROMRequest) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	cdromOp := NewCDROMOp()
	if _, err = cdromOp.Read(ctx, zone, insertParam.ID); err != nil {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("CDROM[%d] is not exists", insertParam.ID))
	}

	value.CDROMID = insertParam.ID
	putServer(zone, value)
	return nil
}

// EjectCDROM is fake implementation
func (o *ServerOp) EjectCDROM(ctx context.Context, zone string, id types.ID, insertParam *iaas.EjectCDROMRequest) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	cdromOp := NewCDROMOp()
	if _, err = cdromOp.Read(ctx, zone, insertParam.ID); err != nil {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("CDROM[%d] is not exists", insertParam.ID))
	}

	value.CDROMID = types.ID(0)
	putServer(zone, value)
	return nil
}

// Boot is fake implementation
func (o *ServerOp) Boot(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Boot is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})

	return err
}

// BootWithVariables is fake implementation
func (o *ServerOp) BootWithVariables(ctx context.Context, zone string, id types.ID, param *iaas.ServerBootVariables) error {
	return o.Boot(ctx, zone, id) // paramは非対応
}

// Shutdown is fake implementation
func (o *ServerOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *iaas.ShutdownOption) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Shutdown is failed")
	}

	startPowerOff(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})

	return err
}

// Reset is fake implementation
func (o *ServerOp) Reset(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Reset is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})

	return nil
}

// SendKey is fake implementation
func (o *ServerOp) SendKey(ctx context.Context, zone string, id types.ID, keyboardParam *iaas.SendKeyRequest) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	return nil
}

// SendNMI is fake implementation
func (o *ServerOp) SendNMI(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	return nil
}

// GetVNCProxy is fake implementation
func (o *ServerOp) GetVNCProxy(ctx context.Context, zone string, id types.ID) (*iaas.VNCProxyInfo, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	vncFileTemplate := `[connection]
host=sac-%s-vnc.cloud.sakura.ad.jp
port=51234
password=aaabababababa`

	return &iaas.VNCProxyInfo{
		Status:       "OK",
		Host:         "localhost",
		IOServerHost: fmt.Sprintf("sac-%s-vnc.cloud.sakura.ad.jp", zone),
		Port:         51234,
		Password:     "dummy",
		VNCFile:      fmt.Sprintf(vncFileTemplate, zone),
	}, nil
}

// Monitor is fake implementation
func (o *ServerOp) Monitor(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.CPUTimeActivity, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &iaas.CPUTimeActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &iaas.MonitorCPUTimeValue{
			Time:    now.Add(time.Duration(i*-5) * time.Minute),
			CPUTime: float64(random(value.CPU * 1000)),
		})
	}

	return res, nil
}

// MonitorCPU is fake implementation
func (o *ServerOp) MonitorCPU(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.CPUTimeActivity, error) {
	return o.Monitor(ctx, zone, id, condition)
}
