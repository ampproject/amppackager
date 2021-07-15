// Copyright 2016-2020 The Libsacloud Authors
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

package api

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
)

// ProductServerAPI サーバープランAPI
type ProductServerAPI struct {
	*baseAPI
}

// NewProductServerAPI サーバープランAPI作成
func NewProductServerAPI(client *Client) *ProductServerAPI {
	return &ProductServerAPI{
		&baseAPI{
			client: client,
			// FuncGetResourceURL
			FuncGetResourceURL: func() string {
				return "product/server"
			},
		},
	}
}

// GetBySpec 指定のコア数/メモリサイズ/世代のプランを取得
func (api *ProductServerAPI) GetBySpec(core, memGB int, gen sacloud.PlanGenerations) (*sacloud.ProductServer, error) {
	return api.GetBySpecCommitment(core, memGB, gen, sacloud.ECommitmentStandard)
}

// GetBySpecCommitment 指定のコア数/メモリサイズ/世代のプランを取得
func (api *ProductServerAPI) GetBySpecCommitment(core, memGB int, gen sacloud.PlanGenerations, commitment sacloud.ECommitment) (*sacloud.ProductServer, error) {
	plans, err := api.Reset().Find()
	if err != nil {
		return nil, err
	}
	var res sacloud.ProductServer
	var found bool
	for _, plan := range plans.ServerPlans {
		if plan.CPU == core && plan.GetMemoryGB() == memGB && plan.Commitment == commitment {
			if gen == sacloud.PlanDefault || gen == plan.Generation {
				// PlanDefaultの場合は複数ヒットしうる。
				// この場合より新しい世代を優先する。
				if found && plan.Generation <= res.Generation {
					continue
				}
				res = plan
				found = true
			}
		}
	}

	if !found {
		return nil, fmt.Errorf("Server Plan[core:%d, memory:%d, gen:%d] is not found", core, memGB, gen)
	}
	return &res, nil
}

// IsValidPlan 指定のコア数/メモリサイズ/世代のプランが存在し、有効であるか判定
func (api *ProductServerAPI) IsValidPlan(core int, memGB int, gen sacloud.PlanGenerations) (bool, error) {

	productServer, err := api.GetBySpec(core, memGB, gen)

	if err != nil {
		return false, err
	}

	if productServer == nil {
		return false, fmt.Errorf("Server Plan[core:%d, memory:%d, gen:%d] is not found", core, memGB, gen)
	}

	if productServer.Availability != sacloud.EAAvailable {
		return false, fmt.Errorf("Server Plan[core:%d, memory:%d, gen:%d] is not available", core, memGB, gen)
	}

	return true, nil
}
