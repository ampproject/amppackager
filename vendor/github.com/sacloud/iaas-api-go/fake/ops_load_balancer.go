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
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// Find is fake implementation
func (o *LoadBalancerOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.LoadBalancerFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.LoadBalancer
	for _, res := range results {
		dest := &iaas.LoadBalancer{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.LoadBalancerFindResult{
		Total:         len(results),
		Count:         len(results),
		From:          0,
		LoadBalancers: values,
	}, nil
}

// Create is fake implementation
func (o *LoadBalancerOp) Create(ctx context.Context, zone string, param *iaas.LoadBalancerCreateRequest) (*iaas.LoadBalancer, error) {
	result := &iaas.LoadBalancer{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Class = "loadbalancer"
	result.Availability = types.Availabilities.Migrating
	result.ZoneID = zoneIDs[zone]
	result.SettingsHash = ""
	for _, vip := range result.VirtualIPAddresses {
		if vip.DelayLoop == 0 {
			vip.DelayLoop = 10 // default value
		}
	}

	putLoadBalancer(zone, result)

	id := result.ID
	startPowerOn(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})
	return result, nil
}

// Read is fake implementation
func (o *LoadBalancerOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.LoadBalancer, error) {
	value := getLoadBalancerByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &iaas.LoadBalancer{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *LoadBalancerOp) Update(ctx context.Context, zone string, id types.ID, param *iaas.LoadBalancerUpdateRequest) (*iaas.LoadBalancer, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	for _, vip := range value.VirtualIPAddresses {
		if vip.DelayLoop == 0 {
			vip.DelayLoop = 10 // default value
		}
	}
	putLoadBalancer(zone, value)
	return value, nil
}

// UpdateSettings is fake implementation
func (o *LoadBalancerOp) UpdateSettings(ctx context.Context, zone string, id types.ID, param *iaas.LoadBalancerUpdateSettingsRequest) (*iaas.LoadBalancer, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	for _, vip := range value.VirtualIPAddresses {
		if vip.DelayLoop == 0 {
			vip.DelayLoop = 10 // default value
		}
	}
	putLoadBalancer(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *LoadBalancerOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, zone, id)
	return nil
}

// Config is fake implementation
func (o *LoadBalancerOp) Config(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	return err
}

// Boot is fake implementation
func (o *LoadBalancerOp) Boot(ctx context.Context, zone string, id types.ID) error {
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
func (o *LoadBalancerOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *iaas.ShutdownOption) error {
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
func (o *LoadBalancerOp) Reset(ctx context.Context, zone string, id types.ID) error {
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

// MonitorInterface is fake implementation
func (o *LoadBalancerOp) MonitorInterface(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.InterfaceActivity, error) {
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
func (o *LoadBalancerOp) Status(ctx context.Context, zone string, id types.ID) (*iaas.LoadBalancerStatusResult, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	var results []*iaas.LoadBalancerStatus
	for _, vip := range value.VirtualIPAddresses {
		status := &iaas.LoadBalancerStatus{
			VirtualIPAddress: vip.VirtualIPAddress,
			Port:             vip.Port,
			CPS:              types.StringNumber(random(100)),
		}
		var servers []*iaas.LoadBalancerServerStatus
		for _, server := range vip.Servers {
			servers = append(servers, &iaas.LoadBalancerServerStatus{
				ActiveConn: types.StringNumber(random(10)),
				Status:     types.ServerInstanceStatuses.Up,
				IPAddress:  server.IPAddress,
				Port:       server.Port,
				CPS:        types.StringNumber(random(100)),
			})
		}
		status.Servers = servers

		results = append(results, status)
	}

	return &iaas.LoadBalancerStatusResult{
		Status: results,
	}, nil
}

// MonitorCPU is fake implementation
func (o *LoadBalancerOp) MonitorCPU(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.CPUTimeActivity, error) {
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
