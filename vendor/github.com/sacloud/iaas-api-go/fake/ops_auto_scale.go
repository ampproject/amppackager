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
func (o *AutoScaleOp) Find(ctx context.Context, conditions *iaas.FindCondition) (*iaas.AutoScaleFindResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, conditions)
	var values []*iaas.AutoScale
	for _, res := range results {
		dest := &iaas.AutoScale{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.AutoScaleFindResult{
		Total:     len(results),
		Count:     len(results),
		From:      0,
		AutoScale: values,
	}, nil
}

// Create is fake implementation
func (o *AutoScaleOp) Create(ctx context.Context, param *iaas.AutoScaleCreateRequest) (*iaas.AutoScale, error) {
	result := &iaas.AutoScale{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)
	result.Availability = types.Availabilities.Available

	// TODO core logic is not implemented

	putAutoScale(iaas.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *AutoScaleOp) Read(ctx context.Context, id types.ID) (*iaas.AutoScale, error) {
	value := getAutoScaleByID(iaas.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.AutoScale{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *AutoScaleOp) Update(ctx context.Context, id types.ID, param *iaas.AutoScaleUpdateRequest) (*iaas.AutoScale, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putAutoScale(iaas.APIDefaultZone, value)
	return value, nil
}

// UpdateSettings is fake implementation
func (o *AutoScaleOp) UpdateSettings(ctx context.Context, id types.ID, param *iaas.AutoScaleUpdateSettingsRequest) (*iaas.AutoScale, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putAutoScale(iaas.APIDefaultZone, value)
	return value, nil
}

// Delete is fake implementation
func (o *AutoScaleOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, iaas.APIDefaultZone, id)
	return nil
}

func (o *AutoScaleOp) Status(ctx context.Context, id types.ID) (*iaas.AutoScaleStatus, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return &iaas.AutoScaleStatus{
		LatestLogs:    []string{"log1", "log2"},
		ResourcesText: "...",
	}, nil
}

func (o *AutoScaleOp) ScaleUp(ctx context.Context, id types.ID) error {
	return nil
}

func (o *AutoScaleOp) ScaleDown(ctx context.Context, id types.ID) error {
	return nil
}
