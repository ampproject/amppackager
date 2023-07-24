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

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// Find is fake implementation
func (o *AutoBackupOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.AutoBackupFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.AutoBackup
	for _, res := range results {
		dest := &iaas.AutoBackup{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.AutoBackupFindResult{
		Total:       len(results),
		Count:       len(results),
		From:        0,
		AutoBackups: values,
	}, nil
}

// Create is fake implementation
func (o *AutoBackupOp) Create(ctx context.Context, zone string, param *iaas.AutoBackupCreateRequest) (*iaas.AutoBackup, error) {
	result := &iaas.AutoBackup{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Availability = types.Availabilities.Available
	result.SettingsHash = "settingshash"
	result.AccountID = accountID
	result.ZoneID = zoneIDs[zone]
	result.ZoneName = zone

	putAutoBackup(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *AutoBackupOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.AutoBackup, error) {
	value := getAutoBackupByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.AutoBackup{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *AutoBackupOp) Update(ctx context.Context, zone string, id types.ID, param *iaas.AutoBackupUpdateRequest) (*iaas.AutoBackup, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putAutoBackup(zone, value)
	return value, nil
}

// UpdateSettings is fake implementation
func (o *AutoBackupOp) UpdateSettings(ctx context.Context, zone string, id types.ID, param *iaas.AutoBackupUpdateSettingsRequest) (*iaas.AutoBackup, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putAutoBackup(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *AutoBackupOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, zone, id)
	return nil
}
