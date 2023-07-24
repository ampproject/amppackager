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

package types

import (
	"fmt"
	"strings"
)

// EProxyLBPlan エンハンスドロードバランサのプラン
//
// エンハンスドロードバランサではプランはIDを受け取る形(Plan.ID)ではなく、
// ServiceClassに"cloud/proxylb/plain/100"のような形で文字列で指定する。
// このままでは扱いにくいためEProxyLBPlan型を設け、この型でjson.Marshaler/Unmarshalerを実装し
// プラン名とServiceClassでの文字列表現とで相互変換可能とする。
type EProxyLBPlan int

// Int EProxyLBPlanのint表現
func (p EProxyLBPlan) Int() int {
	return int(p)
}

// String EProxyLBPlanの文字列表現
func (p EProxyLBPlan) String() string {
	return fmt.Sprintf("%d", p)
}

// ProxyLBPlans エンハンスドロードバランサのプラン
var ProxyLBPlans = struct {
	CPS100    EProxyLBPlan
	CPS500    EProxyLBPlan
	CPS1000   EProxyLBPlan
	CPS5000   EProxyLBPlan
	CPS10000  EProxyLBPlan
	CPS50000  EProxyLBPlan
	CPS100000 EProxyLBPlan
	CPS400000 EProxyLBPlan
}{
	CPS100:    EProxyLBPlan(100),
	CPS500:    EProxyLBPlan(500),
	CPS1000:   EProxyLBPlan(1_000),
	CPS5000:   EProxyLBPlan(5_000),
	CPS10000:  EProxyLBPlan(10_000),
	CPS50000:  EProxyLBPlan(50_000),
	CPS100000: EProxyLBPlan(100_000),
	CPS400000: EProxyLBPlan(400_000),
}

// ProxyLBPlanValues 有効なプランを表すint値
var ProxyLBPlanValues = []int{
	int(ProxyLBPlans.CPS100),
	int(ProxyLBPlans.CPS500),
	int(ProxyLBPlans.CPS1000),
	int(ProxyLBPlans.CPS5000),
	int(ProxyLBPlans.CPS10000),
	int(ProxyLBPlans.CPS50000),
	int(ProxyLBPlans.CPS100000),
	int(ProxyLBPlans.CPS400000),
}

// ProxyLBPlanString 有効なプランを表す文字列(スペース区切り)
var ProxyLBPlanString = strings.Join([]string{
	ProxyLBPlans.CPS100.String(),
	ProxyLBPlans.CPS500.String(),
	ProxyLBPlans.CPS1000.String(),
	ProxyLBPlans.CPS5000.String(),
	ProxyLBPlans.CPS10000.String(),
	ProxyLBPlans.CPS50000.String(),
	ProxyLBPlans.CPS100000.String(),
	ProxyLBPlans.CPS400000.String(),
}, " ")
