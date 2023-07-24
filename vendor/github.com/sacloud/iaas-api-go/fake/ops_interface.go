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
func (o *InterfaceOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.InterfaceFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.Interface
	for _, res := range results {
		dest := &iaas.Interface{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.InterfaceFindResult{
		Total:      len(results),
		Count:      len(results),
		From:       0,
		Interfaces: values,
	}, nil
}

// Create is fake implementation
func (o *InterfaceOp) Create(ctx context.Context, zone string, param *iaas.InterfaceCreateRequest) (*iaas.Interface, error) {
	result := &iaas.Interface{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.MACAddress = pool().nextMACAddress().String()

	// connect to server
	if param != nil && !param.ServerID.IsEmpty() {
		serverOp := NewServerOp()
		server, err := serverOp.Read(ctx, zone, param.ServerID)
		if err == nil {
			ifaceView := &iaas.InterfaceView{}
			copySameNameField(result, ifaceView)
			server.Interfaces = append(server.Interfaces, ifaceView)
			putServer(zone, server)
		}
	}

	putInterface(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *InterfaceOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.Interface, error) {
	value := getInterfaceByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &iaas.Interface{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *InterfaceOp) Update(ctx context.Context, zone string, id types.ID, param *iaas.InterfaceUpdateRequest) (*iaas.Interface, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	serverOp := NewServerOp()
	searched, err := serverOp.Find(ctx, zone, nil)
	if err == nil {
		for _, server := range searched.Servers {
			for _, iface := range server.Interfaces {
				if iface.ID == id {
					iface.UserIPAddress = param.UserIPAddress
					putServer(zone, server)
				}
			}
		}
	}

	putInterface(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *InterfaceOp) Delete(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, zone, id)

	if !value.ServerID.IsEmpty() {
		server, err := NewServerOp().Read(ctx, zone, value.ServerID)
		if err == nil {
			var deleted []*iaas.InterfaceView
			for _, iface := range server.Interfaces {
				if iface.ID != id {
					deleted = append(deleted, iface)
				}
			}
			server.Interfaces = deleted
			putServer(zone, server)
		}
	}

	return nil
}

// Monitor is fake implementation
func (o *InterfaceOp) Monitor(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.InterfaceActivity, error) {
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

// ConnectToSharedSegment is fake implementation
func (o *InterfaceOp) ConnectToSharedSegment(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	if !value.SwitchID.IsEmpty() {
		return newErrorConflict(o.key, id,
			fmt.Sprintf("Interface[%d] is already connected to switch[%d]", value.ID, value.SwitchID))
	}

	value.SwitchID = sharedSegmentSwitch.ID
	putInterface(zone, value)

	if !value.ServerID.IsEmpty() {
		server, err := NewServerOp().Read(ctx, zone, value.ServerID)
		if err == nil {
			for _, iface := range server.Interfaces {
				if iface.ID == id {
					iface.SwitchScope = types.Scopes.Shared
					iface.SwitchID = sharedSegmentSwitch.ID
					iface.SwitchName = sharedSegmentSwitch.Name
				}
			}
			putServer(zone, server)
		}
	}

	return nil
}

// ConnectToSwitch is fake implementation
func (o *InterfaceOp) ConnectToSwitch(ctx context.Context, zone string, id types.ID, switchID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.SwitchID.IsEmpty() {
		return newErrorConflict(o.key, id,
			fmt.Sprintf("Interface[%d] is already connected to switch[%d]", value.ID, switchID))
	}

	sw, err := NewSwitchOp().Read(ctx, zone, switchID)
	if err != nil {
		return err
	}
	sw.ServerCount++
	putSwitch(zone, sw)

	value.SwitchID = switchID
	putInterface(zone, value)

	if !value.ServerID.IsEmpty() {
		server, err := NewServerOp().Read(ctx, zone, value.ServerID)
		if err == nil {
			for _, iface := range server.Interfaces {
				if iface.ID == id {
					iface.SwitchScope = types.Scopes.User
					iface.SwitchID = sw.ID
					iface.SwitchName = sw.Name
				}
			}
			putServer(zone, server)
		}
	}

	return nil
}

// DisconnectFromSwitch is fake implementation
func (o *InterfaceOp) DisconnectFromSwitch(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if value.SwitchID.IsEmpty() {
		return newErrorConflict(o.key, id,
			fmt.Sprintf("Interface[%d] is already disconnected", value.ID))
	}

	value.SwitchID = types.ID(0)
	putInterface(zone, value)

	if !value.ServerID.IsEmpty() {
		server, err := NewServerOp().Read(ctx, zone, value.ServerID)
		if err == nil {
			for _, iface := range server.Interfaces {
				if iface.ID == id {
					iface.SwitchScope = types.EScope("")
					iface.SwitchID = types.ID(0)
					iface.SwitchName = ""
				}
			}
			putServer(zone, server)
		}
	}
	return nil
}

// ConnectToPacketFilter is fake implementation
func (o *InterfaceOp) ConnectToPacketFilter(ctx context.Context, zone string, id types.ID, packetFilterID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.PacketFilterID.IsEmpty() {
		return newErrorConflict(o.key, id,
			fmt.Sprintf("Interface[%d] is already connected to packetfilter[%s]", value.ID, value.PacketFilterID))
	}

	value.PacketFilterID = packetFilterID
	putInterface(zone, value)

	// server配下のInterfaceの修正
	searched, err := NewServerOp().Find(ctx, zone, nil)
	if err != nil {
		return err
	}
	for _, server := range searched.Servers {
		upd := false
		for _, nic := range server.Interfaces {
			if nic.ID == id {
				nic.PacketFilterID = packetFilterID
				upd = true
			}
		}
		if upd {
			putServer(zone, server)
		}
	}

	return nil
}

// DisconnectFromPacketFilter is fake implementation
func (o *InterfaceOp) DisconnectFromPacketFilter(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if value.PacketFilterID.IsEmpty() {
		return newErrorConflict(o.key, id,
			fmt.Sprintf("Interface[%d] is already disconnected", value.ID))
	}

	value.PacketFilterID = types.ID(0)
	putInterface(zone, value)

	// server配下のInterfaceの修正
	searched, err := NewServerOp().Find(ctx, zone, nil)
	if err != nil {
		return err
	}
	for _, server := range searched.Servers {
		upd := false
		for _, nic := range server.Interfaces {
			if nic.ID == id {
				nic.PacketFilterID = types.ID(0)
				upd = true
			}
		}
		if upd {
			putServer(zone, server)
		}
	}
	return nil
}
