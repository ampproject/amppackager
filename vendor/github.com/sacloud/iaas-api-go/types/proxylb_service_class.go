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
	"strconv"
	"strings"
)

const (
	proxyLBServiceClassPrefix               = "cloud/proxylb/plain/"
	proxyLBServiceClassAnycastPrefix        = "cloud/proxylb/anycast/"
	proxyLBServiceClassPrefixEscaped        = "cloud\\/proxylb\\/plain\\/"
	proxyLBServiceClassAnycastPrefixEscaped = "cloud\\/proxylb\\/anycast\\/"
)

// ProxyLBServiceClass プランとリージョンからサービスクラスを算出
func ProxyLBServiceClass(plan EProxyLBPlan, region EProxyLBRegion) string {
	switch region {
	case ProxyLBRegions.Anycast:
		return proxyLBServiceClassAnycastPrefix + plan.String()
	default:
		return proxyLBServiceClassPrefix + plan.String()
	}
}

// ProxyLBPlanFromServiceClass サービスクラスからプランを算出
func ProxyLBPlanFromServiceClass(serviceClass string) EProxyLBPlan {
	strPlan := serviceClass
	strPlan = strings.ReplaceAll(strPlan, `"`, "")
	strPlan = strings.ReplaceAll(strPlan, proxyLBServiceClassPrefix, "")
	strPlan = strings.ReplaceAll(strPlan, proxyLBServiceClassAnycastPrefix, "")
	strPlan = strings.ReplaceAll(strPlan, proxyLBServiceClassPrefixEscaped, "")
	strPlan = strings.ReplaceAll(strPlan, proxyLBServiceClassAnycastPrefixEscaped, "")

	plan, err := strconv.Atoi(strPlan)
	if err != nil {
		plan = 0
	}

	return EProxyLBPlan(plan)
}
