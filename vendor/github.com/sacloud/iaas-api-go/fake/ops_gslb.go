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
func (o *GSLBOp) Find(ctx context.Context, conditions *iaas.FindCondition) (*iaas.GSLBFindResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, conditions)
	var values []*iaas.GSLB
	for _, res := range results {
		dest := &iaas.GSLB{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.GSLBFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		GSLBs: values,
	}, nil
}

// Create is fake implementation
func (o *GSLBOp) Create(ctx context.Context, param *iaas.GSLBCreateRequest) (*iaas.GSLB, error) {
	result := &iaas.GSLB{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillAvailability)

	result.FQDN = fmt.Sprintf("site-%d.gslb7.example.ne.jp", result.ID)
	result.SettingsHash = "settingshash"

	putGSLB(iaas.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *GSLBOp) Read(ctx context.Context, id types.ID) (*iaas.GSLB, error) {
	value := getGSLBByID(iaas.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &iaas.GSLB{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *GSLBOp) Update(ctx context.Context, id types.ID, param *iaas.GSLBUpdateRequest) (*iaas.GSLB, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	if param.DelayLoop == 0 {
		param.DelayLoop = 10 // default value
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putGSLB(iaas.APIDefaultZone, value)
	return value, nil
}

// UpdateSettings is fake implementation
func (o *GSLBOp) UpdateSettings(ctx context.Context, id types.ID, param *iaas.GSLBUpdateSettingsRequest) (*iaas.GSLB, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putGSLB(iaas.APIDefaultZone, value)
	return value, nil
}

// Delete is fake implementation
func (o *GSLBOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, iaas.APIDefaultZone, id)
	return nil
}
