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
	"net"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// Find is fake implementation
func (o *VPCRouterOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.VPCRouterFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.VPCRouter
	for _, res := range results {
		dest := &iaas.VPCRouter{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.VPCRouterFindResult{
		Total:      len(results),
		Count:      len(results),
		From:       0,
		VPCRouters: values,
	}, nil
}

// Create is fake implementation
func (o *VPCRouterOp) Create(ctx context.Context, zone string, param *iaas.VPCRouterCreateRequest) (*iaas.VPCRouter, error) {
	result := &iaas.VPCRouter{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Class = "vpcrouter"
	result.Availability = types.Availabilities.Migrating
	result.ZoneID = zoneIDs[zone]
	result.SettingsHash = ""
	if result.Version == 0 {
		result.Version = 2
	}

	ifOp := NewInterfaceOp()
	swOp := NewSwitchOp()

	ifCreateParam := &iaas.InterfaceCreateRequest{}
	if param.Switch.Scope == types.Scopes.Shared {
		ifCreateParam.ServerID = result.ID
	} else {
		_, err := swOp.Read(ctx, zone, param.Switch.ID)
		if err != nil {
			return nil, newErrorConflict(o.key, types.ID(0), err.Error())
		}
	}

	iface, err := ifOp.Create(ctx, zone, ifCreateParam)
	if err != nil {
		return nil, newErrorConflict(o.key, types.ID(0), err.Error())
	}

	if param.Switch.Scope == types.Scopes.Shared {
		if err := ifOp.ConnectToSharedSegment(ctx, zone, iface.ID); err != nil {
			return nil, newErrorConflict(o.key, types.ID(0), err.Error())
		}
	} else {
		if err := ifOp.ConnectToSwitch(ctx, zone, iface.ID, param.Switch.ID); err != nil {
			return nil, newErrorConflict(o.key, types.ID(0), err.Error())
		}
	}

	iface, err = ifOp.Read(ctx, zone, iface.ID)
	if err != nil {
		return nil, newErrorConflict(o.key, types.ID(0), err.Error())
	}

	vpcRouterInterface := &iaas.VPCRouterInterface{}
	copySameNameField(iface, vpcRouterInterface)
	if param.Switch.Scope == types.Scopes.Shared {
		sharedIP := pool().nextSharedIP()
		vpcRouterInterface.IPAddress = sharedIP.String()
		vpcRouterInterface.SubnetNetworkMaskLen = sharedSegmentSwitch.NetworkMaskLen

		ipv4Mask := net.CIDRMask(pool().SharedNetMaskLen, 32)
		vpcRouterInterface.SubnetNetworkAddress = sharedIP.Mask(ipv4Mask).String()
		vpcRouterInterface.SubnetDefaultRoute = pool().SharedDefaultGateway.String()
	}
	result.Interfaces = append(result.Interfaces, vpcRouterInterface)

	putVPCRouter(zone, result)

	id := result.ID
	startMigration(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})
	return result, nil
}

// Read is fake implementation
func (o *VPCRouterOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.VPCRouter, error) {
	value := getVPCRouterByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.VPCRouter{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *VPCRouterOp) Update(ctx context.Context, zone string, id types.ID, param *iaas.VPCRouterUpdateRequest) (*iaas.VPCRouter, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putVPCRouter(zone, value)
	return value, nil
}

// UpdateSettings is fake implementation
func (o *VPCRouterOp) UpdateSettings(ctx context.Context, zone string, id types.ID, param *iaas.VPCRouterUpdateSettingsRequest) (*iaas.VPCRouter, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putVPCRouter(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *VPCRouterOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, zone, id)
	return nil
}

// Config is fake implementation
func (o *VPCRouterOp) Config(ctx context.Context, zone string, id types.ID) error {
	return nil
}

// Boot is fake implementation
func (o *VPCRouterOp) Boot(ctx context.Context, zone string, id types.ID) error {
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

// Shutdown is fake implementation
func (o *VPCRouterOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *iaas.ShutdownOption) error {
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
func (o *VPCRouterOp) Reset(ctx context.Context, zone string, id types.ID) error {
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

// ConnectToSwitch is fake implementation
func (o *VPCRouterOp) ConnectToSwitch(ctx context.Context, zone string, id types.ID, nicIndex int, switchID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	for _, nic := range value.Interfaces {
		if nic.Index == nicIndex {
			return newErrorBadRequest(o.key, id, fmt.Sprintf("nic[%d] already connected to switch", nicIndex))
		}
	}

	// find switch
	swOp := NewSwitchOp()
	_, err = swOp.Read(ctx, zone, switchID)
	if err != nil {
		return fmt.Errorf("ConnectToSwitch is failed: %s", err)
	}

	// create interface
	ifOp := NewInterfaceOp()
	iface, err := ifOp.Create(ctx, zone, &iaas.InterfaceCreateRequest{ServerID: id})
	if err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}

	if err := ifOp.ConnectToSwitch(ctx, zone, iface.ID, switchID); err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}

	iface, err = ifOp.Read(ctx, zone, iface.ID)
	if err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}

	vpcRouterInterface := &iaas.VPCRouterInterface{}
	copySameNameField(iface, vpcRouterInterface)
	vpcRouterInterface.Index = nicIndex
	value.Interfaces = append(value.Interfaces, vpcRouterInterface)

	putVPCRouter(zone, value)
	return nil
}

// DisconnectFromSwitch is fake implementation
func (o *VPCRouterOp) DisconnectFromSwitch(ctx context.Context, zone string, id types.ID, nicIndex int) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	var exists bool
	var nicID types.ID
	var interfaces []*iaas.VPCRouterInterface

	for _, nic := range value.Interfaces {
		if nic.Index == nicIndex {
			exists = true
			nicID = nic.ID
		} else {
			interfaces = append(interfaces, nic)
		}
	}
	if !exists {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("nic[%d] is not exists", nicIndex))
	}

	ifOp := NewInterfaceOp()
	if err := ifOp.DisconnectFromSwitch(ctx, zone, nicID); err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}

	value.Interfaces = interfaces
	putVPCRouter(zone, value)
	return nil
}

// MonitorInterface is fake implementation
func (o *VPCRouterOp) MonitorInterface(ctx context.Context, zone string, id types.ID, index int, condition *iaas.MonitorCondition) (*iaas.InterfaceActivity, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &iaas.InterfaceActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &iaas.MonitorInterfaceValue{
			Time:    now.Add(time.Duration(i*-5) * time.Minute),
			Send:    float64(random(1000)),
			Receive: float64(random(1000)),
		})
	}

	return res, nil
}

// Status is fake implementation
func (o *VPCRouterOp) Status(ctx context.Context, zone string, id types.ID) (*iaas.VPCRouterStatus, error) {
	v, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	if v.InstanceStatus.IsUp() && v.Settings.WireGuardEnabled.Bool() {
		return &iaas.VPCRouterStatus{
			WireGuard: &iaas.WireGuardStatus{
				PublicKey: "fake-public-key",
			},
			SessionAnalysis: &iaas.VPCRouterSessionAnalysis{
				SourceAndDestination: []*iaas.VPCRouterStatisticsValue{
					{Name: "UDP src:127.0.0.1 dst:127.0.0.1:53", Count: 1},
				},
				DestinationAddress: []*iaas.VPCRouterStatisticsValue{
					{Name: "127.0.0.1", Count: 1},
				},
				DestinationPort: []*iaas.VPCRouterStatisticsValue{
					{Name: "UDP:53", Count: 1},
				},
				SourceAddress: []*iaas.VPCRouterStatisticsValue{
					{Name: "127.0.0.1", Count: 1},
				},
			},
		}, nil
	}
	return &iaas.VPCRouterStatus{}, nil
}

// MonitorCPU is fake implementation
func (o *VPCRouterOp) MonitorCPU(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.CPUTimeActivity, error) {
	_, err := o.Read(ctx, zone, id)
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
			CPUTime: float64(random(1000)),
		})
	}

	return res, nil
}

// Logs is fake implementation
func (o *VPCRouterOp) Logs(ctx context.Context, zone string, id types.ID) (*iaas.VPCRouterLog, error) {
	return &iaas.VPCRouterLog{Log: "fake"}, nil
}

func (o *VPCRouterOp) Ping(ctx context.Context, zone string, id types.ID, destination string) (*iaas.VPCRouterPingResults, error) {
	return &iaas.VPCRouterPingResults{Result: []string{"fake"}}, nil
}
