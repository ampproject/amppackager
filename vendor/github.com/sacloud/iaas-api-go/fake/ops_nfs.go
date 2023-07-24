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
func (o *NFSOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.NFSFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.NFS
	for _, res := range results {
		dest := &iaas.NFS{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.NFSFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		NFS:   values,
	}, nil
}

// Create is fake implementation
func (o *NFSOp) Create(ctx context.Context, zone string, param *iaas.NFSCreateRequest) (*iaas.NFS, error) {
	result := &iaas.NFS{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Class = "nfs"
	result.Availability = types.Availabilities.Migrating
	result.ZoneID = zoneIDs[zone]

	putNFS(zone, result)

	id := result.ID
	startPowerOn(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})
	return result, nil
}

// Read is fake implementation
func (o *NFSOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.NFS, error) {
	value := getNFSByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.NFS{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *NFSOp) Update(ctx context.Context, zone string, id types.ID, param *iaas.NFSUpdateRequest) (*iaas.NFS, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putNFS(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *NFSOp) Delete(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, fmt.Sprintf("NFS[%s] is still running", id))
	}

	ds().Delete(o.key, zone, id)
	return nil
}

// Boot is fake implementation
func (o *NFSOp) Boot(ctx context.Context, zone string, id types.ID) error {
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
func (o *NFSOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *iaas.ShutdownOption) error {
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
func (o *NFSOp) Reset(ctx context.Context, zone string, id types.ID) error {
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

// MonitorFreeDiskSize is fake implementation
func (o *NFSOp) MonitorFreeDiskSize(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.FreeDiskSizeActivity, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &iaas.FreeDiskSizeActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &iaas.MonitorFreeDiskSizeValue{
			Time:         now.Add(time.Duration(i*-5) * time.Minute),
			FreeDiskSize: float64(random(1000)),
		})
	}

	return res, nil
}

// MonitorInterface is fake implementation
func (o *NFSOp) MonitorInterface(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.InterfaceActivity, error) {
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

// MonitorCPU is fake implementation
func (o *NFSOp) MonitorCPU(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.CPUTimeActivity, error) {
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
