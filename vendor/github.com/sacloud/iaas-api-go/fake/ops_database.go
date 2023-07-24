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
func (o *DatabaseOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.DatabaseFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.Database
	for _, res := range results {
		dest := &iaas.Database{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.DatabaseFindResult{
		Total:     len(results),
		Count:     len(results),
		From:      0,
		Databases: values,
	}, nil
}

// Create is fake implementation
func (o *DatabaseOp) Create(ctx context.Context, zone string, param *iaas.DatabaseCreateRequest) (*iaas.Database, error) {
	result := &iaas.Database{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Class = "database"
	result.Availability = types.Availabilities.Available
	if result.Conf != nil {
		if result.Conf.DatabaseVersion == "" {
			result.Conf.DatabaseVersion = "1"
		}
		if result.Conf.DatabaseRevision == "" {
			result.Conf.DatabaseRevision = "1"
		}
	}

	putDatabase(zone, result)

	id := result.ID
	startPowerOn(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})
	return result, nil
}

// Read is fake implementation
func (o *DatabaseOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.Database, error) {
	value := getDatabaseByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.Database{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *DatabaseOp) Update(ctx context.Context, zone string, id types.ID, param *iaas.DatabaseUpdateRequest) (*iaas.Database, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putDatabase(zone, value)
	return value, nil
}

// UpdateSettings is fake implementation
func (o *DatabaseOp) UpdateSettings(ctx context.Context, zone string, id types.ID, param *iaas.DatabaseUpdateSettingsRequest) (*iaas.Database, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putDatabase(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *DatabaseOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, zone, id)
	return nil
}

// Config is fake implementation
func (o *DatabaseOp) Config(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	return nil
}

// Boot is fake implementation
func (o *DatabaseOp) Boot(ctx context.Context, zone string, id types.ID) error {
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
func (o *DatabaseOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *iaas.ShutdownOption) error {
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
func (o *DatabaseOp) Reset(ctx context.Context, zone string, id types.ID) error {
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

// MonitorCPU is fake implementation
func (o *DatabaseOp) MonitorCPU(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.CPUTimeActivity, error) {
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

// MonitorDisk is fake implementation
func (o *DatabaseOp) MonitorDisk(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.DiskActivity, error) {
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

// MonitorInterface is fake implementation
func (o *DatabaseOp) MonitorInterface(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.InterfaceActivity, error) {
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
			Receive: float64(random(1000)),
			Send:    float64(random(1000)),
		})
	}

	return res, nil
}

// MonitorDatabase is fake implementation
func (o *DatabaseOp) MonitorDatabase(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.DatabaseActivity, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &iaas.DatabaseActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &iaas.MonitorDatabaseValue{
			Time:              now.Add(time.Duration(i*-5) * time.Minute),
			TotalMemorySize:   float64(random(1000)),
			UsedMemorySize:    float64(random(1000)),
			TotalDisk1Size:    float64(random(1000)),
			UsedDisk1Size:     float64(random(1000)),
			TotalDisk2Size:    float64(random(1000)),
			UsedDisk2Size:     float64(random(1000)),
			BinlogUsedSizeKiB: float64(random(1000)),
			DelayTimeSec:      float64(random(1000)),
		})
	}

	return res, nil
}

// Status is fake implementation
func (o *DatabaseOp) Status(ctx context.Context, zone string, id types.ID) (*iaas.DatabaseStatus, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	return &iaas.DatabaseStatus{
		Status:        value.InstanceStatus,
		IsFatal:       false,
		MariaDBStatus: "running",
		Version: &iaas.DatabaseVersionInfo{
			LastModified: value.ModifiedAt.String(),
			CommitHash:   "foobar",
			Status:       "up",
			Tag:          "stable",
		},
	}, nil
}

func (o *DatabaseOp) GetParameter(ctx context.Context, zone string, id types.ID) (*iaas.DatabaseParameter, error) {
	v, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	var settings map[string]interface{}
	raw := ds().Get(ResourceDatabase+"Parameter", zone, id)
	if raw != nil {
		settings = raw.(map[string]interface{})
	}

	meta := fakeDatabaseParameterMetaForMariaDB
	if v.Conf.DatabaseName == "postgres" {
		meta = fakeDatabaseParameterMetaForPostgreSQL
	}
	return &iaas.DatabaseParameter{
		Settings: settings,
		MetaInfo: meta,
	}, nil
}

var (
	fakeDatabaseParameterMetaForMariaDB = []*iaas.DatabaseParameterMeta{
		{
			Type:    "number",
			Name:    "MariaDB/server.cnf/mysqld/max_connections",
			Label:   "max_connections",
			Text:    "同時クライアント接続の最大数を設定します。",
			Example: "100",
			Min:     10,
			Max:     1000,
			MaxLen:  0,
			Reboot:  "static",
		},
		{
			Type:    "string",
			Name:    "MariaDB/server.cnf/mysqld/event_scheduler",
			Label:   "event_scheduler",
			Text:    "イベントスケジュールの有効無効を設定します。",
			Example: "ON",
			Min:     0,
			Max:     0,
			MaxLen:  0,
			Reboot:  "dynamic",
		},
	}
	fakeDatabaseParameterMetaForPostgreSQL = []*iaas.DatabaseParameterMeta{
		{
			Type:    "number",
			Name:    "postgres/postgresql.conf/max_connections",
			Label:   "max_connections",
			Text:    "同時クライアント接続の最大数を設定します。",
			Example: "100",
			Min:     10,
			Max:     1000,
			MaxLen:  0,
			Reboot:  "static",
		},
		{
			Type:    "number",
			Name:    "postgres/postgresql.conf/work_mem",
			Label:   "work_mem",
			Text:    "クエリワークスペースに使用するメモリの最大量を設定します。",
			Example: "4096",
			Min:     64,
			Max:     2147483647,
			MaxLen:  10,
			Reboot:  "dynamic",
		},
	}
)

func (o *DatabaseOp) SetParameter(ctx context.Context, zone string, id types.ID, param map[string]interface{}) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	var settings map[string]interface{}
	raw := ds().Get(ResourceDatabase+"Parameter", zone, id)
	if raw != nil {
		settings = raw.(map[string]interface{})
	} else {
		settings = make(map[string]interface{})
	}
	for k, v := range param {
		if v == nil {
			delete(settings, k)
		} else {
			switch v := v.(type) {
			case int:
				settings[k] = float64(v)
			default:
				settings[k] = v
			}
		}
	}

	ds().Put(ResourceDatabase+"Parameter", zone, id, settings)
	return nil
}
