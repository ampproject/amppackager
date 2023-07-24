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
func (o *MobileGatewayOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.MobileGatewayFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.MobileGateway
	for _, res := range results {
		dest := &iaas.MobileGateway{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.MobileGatewayFindResult{
		Total:          len(results),
		Count:          len(results),
		From:           0,
		MobileGateways: values,
	}, nil
}

// Create is fake implementation
func (o *MobileGatewayOp) Create(ctx context.Context, zone string, param *iaas.MobileGatewayCreateRequest) (*iaas.MobileGateway, error) {
	result := &iaas.MobileGateway{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Availability = types.Availabilities.Available
	result.Class = "mobilegateway"
	result.ZoneID = zoneIDs[zone]
	result.SettingsHash = ""

	// set interface
	ifOp := NewInterfaceOp()
	iface, err := ifOp.Create(ctx, zone, &iaas.InterfaceCreateRequest{ServerID: result.ID})
	if err != nil {
		return nil, newErrorConflict(o.key, types.ID(0), err.Error())
	}
	if err := ifOp.ConnectToSharedSegment(ctx, zone, iface.ID); err != nil {
		return nil, newErrorConflict(o.key, types.ID(0), err.Error())
	}
	iface, err = ifOp.Read(ctx, zone, iface.ID)
	if err != nil {
		return nil, newErrorConflict(o.key, types.ID(0), err.Error())
	}
	ifaceView := &iaas.MobileGatewayInterface{}
	copySameNameField(iface, ifaceView)
	result.Interfaces = append(result.Interfaces, ifaceView)

	putMobileGateway(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *MobileGatewayOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.MobileGateway, error) {
	value := getMobileGatewayByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.MobileGateway{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *MobileGatewayOp) Update(ctx context.Context, zone string, id types.ID, param *iaas.MobileGatewayUpdateRequest) (*iaas.MobileGateway, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putMobileGateway(zone, value)
	return value, nil
}

// UpdateSettings is fake implementation
func (o *MobileGatewayOp) UpdateSettings(ctx context.Context, zone string, id types.ID, param *iaas.MobileGatewayUpdateSettingsRequest) (*iaas.MobileGateway, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putMobileGateway(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *MobileGatewayOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, zone, id)
	return nil
}

// Config is fake implementation
func (o *MobileGatewayOp) Config(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	return err
}

// Boot is fake implementation
func (o *MobileGatewayOp) Boot(ctx context.Context, zone string, id types.ID) error {
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
func (o *MobileGatewayOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *iaas.ShutdownOption) error {
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
func (o *MobileGatewayOp) Reset(ctx context.Context, zone string, id types.ID) error {
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
func (o *MobileGatewayOp) ConnectToSwitch(ctx context.Context, zone string, id types.ID, switchID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	for _, nic := range value.Interfaces {
		if nic.Index == 1 {
			return newErrorBadRequest(o.key, id, fmt.Sprintf("nic[%d] already connected to switch", 1))
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

	mobileGatewayInterface := &iaas.MobileGatewayInterface{}
	copySameNameField(iface, mobileGatewayInterface)
	mobileGatewayInterface.Index = 1 // 1固定
	value.Interfaces = append(value.Interfaces, mobileGatewayInterface)

	putMobileGateway(zone, value)
	return nil
}

// DisconnectFromSwitch is fake implementation
func (o *MobileGatewayOp) DisconnectFromSwitch(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	var exists bool
	var nicID types.ID
	var interfaces []*iaas.MobileGatewayInterface

	for _, nic := range value.Interfaces {
		if nic != nil && nic.Index == 1 {
			exists = true
			nicID = nic.ID
		} else {
			interfaces = append(interfaces, nic)
		}
	}
	if !exists {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("nic[%d] is not exists", 1))
	}

	ifOp := NewInterfaceOp()
	if err := ifOp.DisconnectFromSwitch(ctx, zone, nicID); err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}

	value.Interfaces = interfaces
	putMobileGateway(zone, value)
	return nil
}

// GetDNS is fake implementation
func (o *MobileGatewayOp) GetDNS(ctx context.Context, zone string, id types.ID) (*iaas.MobileGatewayDNSSetting, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	dns := ds().Get(o.dnsStoreKey(), zone, id)
	if dns == nil {
		return &iaas.MobileGatewayDNSSetting{
			DNS1: "133.242.0.1",
			DNS2: "133.242.0.2",
		}, nil
	}
	return dns.(*iaas.MobileGatewayDNSSetting), nil
}

// SetDNS is fake implementation
func (o *MobileGatewayOp) SetDNS(ctx context.Context, zone string, id types.ID, param *iaas.MobileGatewayDNSSetting) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	ds().Put(o.dnsStoreKey(), zone, id, param)
	return nil
}

// GetSIMRoutes is fake implementation
func (o *MobileGatewayOp) GetSIMRoutes(ctx context.Context, zone string, id types.ID) (iaas.MobileGatewaySIMRoutes, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	routes := ds().Get(o.simRoutesStoreKey(), zone, id)
	if routes == nil {
		return nil, nil
	}

	rs := routes.(*[]*iaas.MobileGatewaySIMRoute)
	var res []*iaas.MobileGatewaySIMRoute
	res = append(res, *rs...)
	return res, nil
}

// SetSIMRoutes is fake implementation
func (o *MobileGatewayOp) SetSIMRoutes(ctx context.Context, zone string, id types.ID, param []*iaas.MobileGatewaySIMRouteParam) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	simOp := NewSIMOp()
	var values []*iaas.MobileGatewaySIMRoute
	for _, p := range param {
		sim, err := simOp.Read(ctx, types.StringID(p.ResourceID))
		if err != nil {
			return err
		}
		values = append(values, &iaas.MobileGatewaySIMRoute{
			ResourceID: p.ResourceID,
			Prefix:     p.Prefix,
			ICCID:      sim.ICCID,
		})
	}

	ds().Put(o.simRoutesStoreKey(), zone, id, &values)
	return nil
}

// ListSIM is fake implementation
func (o *MobileGatewayOp) ListSIM(ctx context.Context, zone string, id types.ID) (iaas.MobileGatewaySIMs, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	sims := ds().Get(o.simsStoreKey(), zone, id)
	if sims == nil {
		return nil, nil
	}

	ss := sims.(*[]*iaas.MobileGatewaySIMInfo)
	var res []*iaas.MobileGatewaySIMInfo
	res = append(res, *ss...)
	return res, nil
}

// AddSIM is fake implementation
func (o *MobileGatewayOp) AddSIM(ctx context.Context, zone string, id types.ID, param *iaas.MobileGatewayAddSIMRequest) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	var sims []*iaas.MobileGatewaySIMInfo
	rawSIMs := ds().Get(o.simsStoreKey(), zone, id)
	if rawSIMs != nil {
		sims = rawSIMs.([]*iaas.MobileGatewaySIMInfo)
		for _, sim := range sims {
			if sim.ResourceID == param.SIMID {
				return newErrorBadRequest(o.key, id, fmt.Sprintf("SIM %s already exists", param.SIMID))
			}
		}
	}

	simOp := NewSIMOp()
	simInfo, err := simOp.Status(context.Background(), types.StringID(param.SIMID))
	if err != nil {
		return err
	}
	sim := &iaas.MobileGatewaySIMInfo{}
	copySameNameField(simInfo, sim)

	sims = append(sims, sim)

	ds().Put(o.simsStoreKey(), zone, id, &sims)
	return nil
}

// DeleteSIM is fake implementation
func (o *MobileGatewayOp) DeleteSIM(ctx context.Context, zone string, id types.ID, simID types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	var updSIMs []*iaas.MobileGatewaySIMInfo
	rawSIMs := ds().Get(o.simsStoreKey(), zone, id)
	if rawSIMs != nil {
		ss := rawSIMs.(*[]*iaas.MobileGatewaySIMInfo)
		for _, sim := range *ss {
			if sim.ResourceID != simID.String() {
				updSIMs = append(updSIMs, sim)
			}
		}
		if len(*ss) != len(updSIMs) {
			ds().Put(o.simsStoreKey(), zone, id, &updSIMs)
			return nil
		}
	}
	return newErrorBadRequest(o.key, id, fmt.Sprintf("SIM %d is not exists", simID))
}

// Logs is fake implementation
func (o *MobileGatewayOp) Logs(ctx context.Context, zone string, id types.ID) ([]*iaas.MobileGatewaySIMLogs, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	return []*iaas.MobileGatewaySIMLogs{
		{
			Date:          time.Now(),
			SessionStatus: "UP",
			ResourceID:    types.ID(1).String(),
		},
	}, nil
}

// GetTrafficConfig is fake implementation
func (o *MobileGatewayOp) GetTrafficConfig(ctx context.Context, zone string, id types.ID) (*iaas.MobileGatewayTrafficControl, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	config := ds().Get(o.trafficConfigStoreKey(), zone, id)
	if config == nil {
		return nil, nil
	}
	return config.(*iaas.MobileGatewayTrafficControl), nil
}

// SetTrafficConfig is fake implementation
func (o *MobileGatewayOp) SetTrafficConfig(ctx context.Context, zone string, id types.ID, param *iaas.MobileGatewayTrafficControl) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	ds().Put(o.trafficConfigStoreKey(), zone, id, param)
	return nil
}

// DeleteTrafficConfig is fake implementation
func (o *MobileGatewayOp) DeleteTrafficConfig(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	ds().Delete(o.trafficConfigStoreKey(), zone, id)
	return nil
}

// TrafficStatus is fake implementation
func (o *MobileGatewayOp) TrafficStatus(ctx context.Context, zone string, id types.ID) (*iaas.MobileGatewayTrafficStatus, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	return &iaas.MobileGatewayTrafficStatus{
		UplinkBytes:    0,
		DownlinkBytes:  0,
		TrafficShaping: true,
	}, nil
}

// MonitorInterface is fake implementation
func (o *MobileGatewayOp) MonitorInterface(ctx context.Context, zone string, id types.ID, index int, condition *iaas.MonitorCondition) (*iaas.InterfaceActivity, error) {
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

func (o *MobileGatewayOp) dnsStoreKey() string {
	return o.key + "DNS"
}

func (o *MobileGatewayOp) simRoutesStoreKey() string {
	return o.key + "SIMRoutes"
}

func (o *MobileGatewayOp) simsStoreKey() string {
	return o.key + "SIMs"
}

func (o *MobileGatewayOp) trafficConfigStoreKey() string {
	return o.key + "TrafficConfig"
}
