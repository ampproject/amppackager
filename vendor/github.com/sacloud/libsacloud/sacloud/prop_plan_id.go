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

// propPlanID プランID内包型
type propPlanID struct {
	Plan *Resource `json:",omitempty"` // プラン
}

// GetPlanID プランID 取得
func (p *propPlanID) GetPlanID() ID {
	if p.Plan == nil {
		return -1
	}
	return p.Plan.GetID()
}

// GetStrPlanID プランID(文字列) 取得
func (p *propPlanID) GetStrPlanID() string {
	if p.Plan == nil {
		return ""
	}
	return p.Plan.GetStrID()
}
