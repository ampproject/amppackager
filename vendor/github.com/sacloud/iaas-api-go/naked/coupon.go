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

package naked

import (
	"time"

	"github.com/sacloud/iaas-api-go/types"
)

// Coupon クーポン情報
type Coupon struct {
	ID             types.ID   `json:"CouponID,omitempty" yaml:",omitempty" structs:",omitempty"`         // クーポンID
	MemberID       string     `json:",omitempty" yaml:"member_id,omitempty" structs:",omitempty"`        // メンバーID
	ContractID     types.ID   `json:",omitempty" yaml:"contract_id,omitempty" structs:",omitempty"`      // 契約ID
	ServiceClassID types.ID   `json:",omitempty" yaml:"service_class_id,omitempty" structs:",omitempty"` // サービスクラスID
	Discount       int64      `json:",omitempty" yaml:"discount,omitempty" structs:",omitempty"`         // クーポン残高
	AppliedAt      *time.Time `json:",omitempty" yaml:"applied_at,omitempty" structs:",omitempty"`       // 適用開始日
	UntilAt        *time.Time `json:",omitempty" yaml:"until_at,omitempty" structs:",omitempty"`         // 有効期限
}
