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
func (o *ServerPlanOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.ServerPlanFindResult, error) {
	results, _ := find(o.key, iaas.APIDefaultZone, conditions)
	var values []*iaas.ServerPlan
	for _, res := range results {
		dest := &iaas.ServerPlan{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.ServerPlanFindResult{
		Total:       len(results),
		Count:       len(results),
		From:        0,
		ServerPlans: values,
	}, nil
}

// Read is fake implementation
func (o *ServerPlanOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.ServerPlan, error) {
	value := getServerPlanByID(iaas.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.ServerPlan{}
	copySameNameField(value, dest)
	return dest, nil
}
