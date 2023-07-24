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
func (o *LicenseOp) Find(ctx context.Context, conditions *iaas.FindCondition) (*iaas.LicenseFindResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, conditions)
	var values []*iaas.License
	for _, res := range results {
		dest := &iaas.License{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.LicenseFindResult{
		Total:    len(results),
		Count:    len(results),
		From:     0,
		Licenses: values,
	}, nil
}

// Create is fake implementation
func (o *LicenseOp) Create(ctx context.Context, param *iaas.LicenseCreateRequest) (*iaas.License, error) {
	result := &iaas.License{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillModifiedAt)
	result.LicenseInfoName = "Windows RDS SAL"
	putLicense(iaas.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *LicenseOp) Read(ctx context.Context, id types.ID) (*iaas.License, error) {
	value := getLicenseByID(iaas.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.License{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *LicenseOp) Update(ctx context.Context, id types.ID, param *iaas.LicenseUpdateRequest) (*iaas.License, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putLicense(iaas.APIDefaultZone, value)
	return value, nil
}

// Delete is fake implementation
func (o *LicenseOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, iaas.APIDefaultZone, id)
	return nil
}
