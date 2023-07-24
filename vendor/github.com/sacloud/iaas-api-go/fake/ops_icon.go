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
func (o *IconOp) Find(ctx context.Context, conditions *iaas.FindCondition) (*iaas.IconFindResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, conditions)
	var values []*iaas.Icon
	for _, res := range results {
		dest := &iaas.Icon{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.IconFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		Icons: values,
	}, nil
}

// Create is fake implementation
func (o *IconOp) Create(ctx context.Context, param *iaas.IconCreateRequest) (*iaas.Icon, error) {
	result := &iaas.Icon{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillModifiedAt)

	result.Availability = types.Availabilities.Available
	result.Scope = types.Scopes.User
	result.URL = fmt.Sprintf("https://secure.sakura.ad.jp/cloud/zone/is1a/api/cloud/1.1/icon/%d.png", result.ID)

	putIcon(iaas.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *IconOp) Read(ctx context.Context, id types.ID) (*iaas.Icon, error) {
	value := getIconByID(iaas.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.Icon{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *IconOp) Update(ctx context.Context, id types.ID, param *iaas.IconUpdateRequest) (*iaas.Icon, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putIcon(iaas.APIDefaultZone, value)
	return value, nil
}

// Delete is fake implementation
func (o *IconOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, iaas.APIDefaultZone, id)
	return nil
}
