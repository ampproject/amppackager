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

package sacloud

// propPrivateHostPlan 専有ホストプラン内包型
type propPrivateHostPlan struct {
	Plan *ProductPrivateHost `json:",omitempty"` // 専有ホストプラン
}

// GetPrivateHostPlan 専有ホストプラン取得
func (p *propPrivateHostPlan) GetPrivateHostPlan() *ProductPrivateHost {
	return p.Plan
}

// SetPrivateHostPlan 専有ホストプラン設定
func (p *propPrivateHostPlan) SetPrivateHostPlan(plan *ProductPrivateHost) {
	p.Plan = plan
}

// SetPrivateHostPlanByID 専有ホストプラン設定
func (p *propPrivateHostPlan) SetPrivateHostPlanByID(planID ID) {
	if p.Plan == nil {
		p.Plan = &ProductPrivateHost{}
	}
	p.Plan.Resource = NewResource(planID)
}

// GetCPU CPUコア数 取得
func (p *propPrivateHostPlan) GetCPU() int {
	if p.Plan == nil {
		return -1
	}

	return p.Plan.GetCPU()
}

// GetMemoryMB メモリ(MB) 取得
func (p *propPrivateHostPlan) GetMemoryMB() int {
	if p.Plan == nil {
		return -1
	}

	return p.Plan.GetMemoryMB()
}

// GetMemoryGB メモリ(GB) 取得
func (p *propPrivateHostPlan) GetMemoryGB() int {
	if p.Plan == nil {
		return -1
	}

	return p.Plan.GetMemoryGB()
}
