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

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// Find is fake implementation
func (o *SwitchOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.SwitchFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.Switch
	for _, res := range results {
		dest := &iaas.Switch{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.SwitchFindResult{
		Total:    len(results),
		Count:    len(results),
		From:     0,
		Switches: values,
	}, nil
}

// Create is fake implementation
func (o *SwitchOp) Create(ctx context.Context, zone string, param *iaas.SwitchCreateRequest) (*iaas.Switch, error) {
	result := &iaas.Switch{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillAvailability, fillScope)
	result.Scope = types.Scopes.User
	putSwitch(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *SwitchOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.Switch, error) {
	value := getSwitchByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.Switch{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *SwitchOp) Update(ctx context.Context, zone string, id types.ID, param *iaas.SwitchUpdateRequest) (*iaas.Switch, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putSwitch(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *SwitchOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, zone, id)
	return nil
}

// ConnectToBridge is fake implementation
func (o *SwitchOp) ConnectToBridge(ctx context.Context, zone string, id types.ID, bridgeID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	bridgeOp := NewBridgeOp()
	bridge, err := bridgeOp.Read(ctx, zone, bridgeID)
	if err != nil {
		return fmt.Errorf("ConnectToBridge is failed: %s", err)
	}

	if bridge.SwitchInZone != nil {
		return newErrorConflict(o.key, id, fmt.Sprintf("Bridge[%d] already connected to switch", bridgeID))
	}

	value.BridgeID = bridgeID

	switchInZone := &iaas.BridgeSwitchInfo{}
	copySameNameField(value, switchInZone)
	bridge.SwitchInZone = switchInZone

	// bridge.BridgeInfo = append(bridge.BridgeInfo, &iaas.BridgeInfo{
	//	ID:     value.ID,
	//	Name:   value.Name,
	//	ZoneID: zoneIDs[zone],
	// })

	putBridge(zone, bridge)
	putSwitch(zone, value)
	return nil
}

// DisconnectFromBridge is fake implementation
func (o *SwitchOp) DisconnectFromBridge(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	if value.BridgeID.IsEmpty() {
		return newErrorConflict(o.key, id, fmt.Sprintf("Switch[%d] already disconnected from switch", id))
	}

	bridgeOp := NewBridgeOp()
	bridge, err := bridgeOp.Read(ctx, zone, value.BridgeID)
	if err != nil {
		return fmt.Errorf("DisconnectFromBridge is failed: %s", err)
	}

	// var bridgeInfo []*iaas.BridgeInfo
	// for _, i := range bridge.BridgeInfo {
	//	if i.ID != value.ID {
	//		bridgeInfo = append(bridgeInfo, i)
	//	}
	// }

	value.BridgeID = types.ID(0)
	bridge.SwitchInZone = nil
	// fakeドライバーではBridgeInfoに非対応
	// bridge.BridgeInfo = bridgeInfo

	putBridge(zone, bridge)
	putSwitch(zone, value)
	return nil
}

// GetServers is fake implementation
func (o *SwitchOp) GetServers(ctx context.Context, zone string, id types.ID) (*iaas.SwitchGetServersResult, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	res := &iaas.SwitchGetServersResult{}
	if value.ServerCount == 0 {
		return res, nil
	}

	searched, err := NewServerOp().Find(ctx, zone, nil)
	if err != nil {
		return nil, err
	}
	for _, server := range searched.Servers {
		for _, nic := range server.Interfaces {
			if nic.SwitchID == id {
				res.Servers = append(res.Servers, server)
				res.Count++
				res.Total++
				break
			}
		}
	}
	return res, nil
}
