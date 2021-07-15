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

import "time"

// Coupon クーポン情報
type Coupon struct {
	CouponID       ID        `json:",omitempty"` // クーポンID
	MemberID       string    `json:",omitempty"` // メンバーID
	ContractID     ID        `json:",omitempty"` // 契約ID
	ServiceClassID ID        `json:",omitempty"` // サービスクラスID
	Discount       int64     `json:",omitempty"` // クーポン残高
	AppliedAt      time.Time `json:",omitempty"` // 適用開始日
	UntilAt        time.Time `json:",omitempty"` // 有効期限
}
