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
func (o *SimpleMonitorOp) Find(ctx context.Context, conditions *iaas.FindCondition) (*iaas.SimpleMonitorFindResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, conditions)
	var values []*iaas.SimpleMonitor
	for _, res := range results {
		dest := &iaas.SimpleMonitor{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.SimpleMonitorFindResult{
		Total:          len(results),
		Count:          len(results),
		From:           0,
		SimpleMonitors: values,
	}, nil
}

// Create is fake implementation
func (o *SimpleMonitorOp) Create(ctx context.Context, param *iaas.SimpleMonitorCreateRequest) (*iaas.SimpleMonitor, error) {
	result := &iaas.SimpleMonitor{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Name = param.Target
	result.Class = "simplemon"
	result.Availability = types.Availabilities.Available
	result.SettingsHash = "settingshash"
	if result.DelayLoop == 0 {
		result.DelayLoop = 60
	}
	if result.MaxCheckAttempts == 0 {
		result.MaxCheckAttempts = 3
	}
	if result.RetryInterval == 0 {
		result.RetryInterval = 10
	}
	if result.NotifyInterval == 0 {
		result.NotifyInterval = 7200
	}

	putSimpleMonitor(iaas.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *SimpleMonitorOp) Read(ctx context.Context, id types.ID) (*iaas.SimpleMonitor, error) {
	value := getSimpleMonitorByID(iaas.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.SimpleMonitor{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *SimpleMonitorOp) Update(ctx context.Context, id types.ID, param *iaas.SimpleMonitorUpdateRequest) (*iaas.SimpleMonitor, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	if value.DelayLoop == 0 {
		value.DelayLoop = 60
	}
	if value.MaxCheckAttempts == 0 {
		value.MaxCheckAttempts = 3
	}
	if value.RetryInterval == 0 {
		value.RetryInterval = 10
	}
	if value.NotifyInterval == 0 {
		value.NotifyInterval = 7200
	}
	putSimpleMonitor(iaas.APIDefaultZone, value)
	return value, nil
}

// UpdateSettings is fake implementation
func (o *SimpleMonitorOp) UpdateSettings(ctx context.Context, id types.ID, param *iaas.SimpleMonitorUpdateSettingsRequest) (*iaas.SimpleMonitor, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	if value.DelayLoop == 0 {
		value.DelayLoop = 60
	}
	if value.NotifyInterval == 0 {
		value.NotifyInterval = 7200
	}
	putSimpleMonitor(iaas.APIDefaultZone, value)
	return value, nil
}

// Delete is fake implementation
func (o *SimpleMonitorOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, iaas.APIDefaultZone, id)
	return nil
}

// MonitorResponseTime is fake implementation
func (o *SimpleMonitorOp) MonitorResponseTime(ctx context.Context, id types.ID, condition *iaas.MonitorCondition) (*iaas.ResponseTimeSecActivity, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &iaas.ResponseTimeSecActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &iaas.MonitorResponseTimeSecValue{
			Time:            now.Add(time.Duration(i*-5) * time.Minute),
			ResponseTimeSec: float64(random(1000)),
		})
	}

	return res, nil
}

// HealthStatus is fake implementation
func (o *SimpleMonitorOp) HealthStatus(ctx context.Context, id types.ID) (*iaas.SimpleMonitorHealthStatus, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return &iaas.SimpleMonitorHealthStatus{
		LastCheckedAt:       time.Now(),
		LastHealthChangedAt: time.Now(),
		Health:              types.SimpleMonitorHealth.Up,
	}, nil
}
