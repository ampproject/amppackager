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

// Bill 請求情報
type Bill struct {
	Amount         int64      `json:",omitempty"` // 金額
	BillID         ID         `json:",omitempty"` // 請求ID
	Date           *time.Time `json:",omitempty"` // 請求日
	MemberID       string     `json:",omitempty"` // 会員ID
	Paid           bool       `json:",omitempty"` // 支払済フラグ
	PayLimit       *time.Time `json:",omitempty"` // 支払い期限
	PaymentClassID ID         `json:",omitempty"` // 支払いクラスID

}

// BillDetail 支払い明細情報
type BillDetail struct {
	ContractID     ID         `json:",omitempty"` // 契約ID
	Amount         int64      `json:",omitempty"` // 金額
	Description    string     `json:",omitempty"` // 説明
	Index          int        `json:",omitempty"` // インデックス
	ServiceClassID ID         `json:",omitempty"` // サービスクラスID
	Usage          int64      `json:",omitempty"` // 秒数
	Zone           string     `json:",omitempty"` // ゾーン
	ContractEndAt  *time.Time `json:",omitempty"` // 契約終了日時
}

// IsContractEnded 支払済か判定
func (d *BillDetail) IsContractEnded(t time.Time) bool {
	return d.ContractEndAt != nil && d.ContractEndAt.Before(t)
}
