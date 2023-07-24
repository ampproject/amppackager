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
func (o *EnhancedDBOp) Find(ctx context.Context, conditions *iaas.FindCondition) (*iaas.EnhancedDBFindResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, conditions)
	var values []*iaas.EnhancedDB
	for _, res := range results {
		dest := &iaas.EnhancedDB{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.EnhancedDBFindResult{
		Total:       len(results),
		Count:       len(results),
		From:        0,
		EnhancedDBs: values,
	}, nil
}

// Create is fake implementation
func (o *EnhancedDBOp) Create(ctx context.Context, param *iaas.EnhancedDBCreateRequest) (*iaas.EnhancedDB, error) {
	result := &iaas.EnhancedDB{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	if result.DatabaseType == "" {
		result.DatabaseType = "tidb"
	}
	if result.Region == "" {
		result.Region = "is1"
	}
	result.Port = 3306
	result.HostName = fmt.Sprintf("%s.%s-%s.db.sakurausercontent.com", result.DatabaseName, result.DatabaseType, result.Region)
	result.Availability = types.Availabilities.Available

	putEnhancedDB(iaas.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *EnhancedDBOp) Read(ctx context.Context, id types.ID) (*iaas.EnhancedDB, error) {
	value := getEnhancedDBByID(iaas.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.EnhancedDB{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *EnhancedDBOp) Update(ctx context.Context, id types.ID, param *iaas.EnhancedDBUpdateRequest) (*iaas.EnhancedDB, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putEnhancedDB(iaas.APIDefaultZone, value)
	return value, nil
}

// Delete is fake implementation
func (o *EnhancedDBOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, iaas.APIDefaultZone, id)
	return nil
}

// SetPassword is fake implementation
func (o *EnhancedDBOp) SetPassword(ctx context.Context, id types.ID, param *iaas.EnhancedDBSetPasswordRequest) error {
	_, err := o.Read(ctx, id)
	return err
}

func (o *EnhancedDBOp) GetConfig(ctx context.Context, id types.ID) (*iaas.EnhancedDBConfig, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return &iaas.EnhancedDBConfig{MaxConnections: 50}, nil
}

func (o *EnhancedDBOp) SetConfig(ctx context.Context, id types.ID, param *iaas.EnhancedDBSetConfigRequest) error {
	_, err := o.Read(ctx, id)
	return err
}
