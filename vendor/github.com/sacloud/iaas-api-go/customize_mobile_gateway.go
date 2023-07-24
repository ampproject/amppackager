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

package iaas

import (
	"fmt"

	"github.com/sacloud/iaas-api-go/types"
)

type MobileGatewaySIMRoutes []*MobileGatewaySIMRoute

func (o *MobileGatewaySIMRoutes) FindByID(resourceID types.ID) *MobileGatewaySIMRoute {
	for _, r := range *o {
		if r.ResourceID == resourceID.String() {
			return r
		}
	}
	return nil
}

func (o *MobileGatewaySIMRoutes) Exists(resourceID types.ID) bool {
	return o.FindByID(resourceID) != nil
}

func (o *MobileGatewaySIMRoutes) Add(route *MobileGatewaySIMRoute) error {
	if o.Exists(types.StringID(route.ResourceID)) {
		return fmt.Errorf("SIM[%s] already exists in SIMRoutes", route.ResourceID)
	}
	*o = append(*o, route)
	return nil
}

func (o *MobileGatewaySIMRoutes) Update(route *MobileGatewaySIMRoute) error {
	r := o.FindByID(types.StringID(route.ResourceID))
	if r == nil {
		return fmt.Errorf("SIM[%s] not found in SIMRoutes", route.ResourceID)
	}
	r.Prefix = route.Prefix
	r.ICCID = route.ICCID
	return nil
}

func (o *MobileGatewaySIMRoutes) Delete(resourceID types.ID) error {
	if !o.Exists(resourceID) {
		return fmt.Errorf("SIM[%s] not found in SIMRoutes", resourceID)
	}
	var rs []*MobileGatewaySIMRoute
	for _, r := range *o {
		if r.ResourceID != resourceID.String() {
			rs = append(rs, r)
		}
	}
	*o = rs
	return nil
}

func (o *MobileGatewaySIMRoutes) ToRequestParameter() []*MobileGatewaySIMRouteParam {
	var rs []*MobileGatewaySIMRouteParam
	for _, r := range *o {
		rs = append(rs, &MobileGatewaySIMRouteParam{
			ResourceID: r.ResourceID,
			Prefix:     r.Prefix,
		})
	}
	return rs
}

type MobileGatewaySIMs []*MobileGatewaySIMInfo

func (o *MobileGatewaySIMs) FindByID(resourceID types.ID) *MobileGatewaySIMInfo {
	for _, s := range *o {
		if s.ResourceID == resourceID.String() {
			return s
		}
	}
	return nil
}

func (o *MobileGatewaySIMs) Exists(resourceID types.ID) bool {
	return o.FindByID(resourceID) != nil
}
