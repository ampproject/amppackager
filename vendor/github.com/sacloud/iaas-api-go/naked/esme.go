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

// ESME 2要素認証 SMS送信サービス
type ESME struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `yaml:"description"`
	Tags         types.Tags          `yaml:"tags"`
	Icon         *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	ServiceClass string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Provider     *Provider           `json:",omitempty" yaml:"provider,omitempty" structs:",omitempty"`

	/* Note: 以下3フィールドはCommonServiceItem共通フィールドだがESMEでは常にnullとなっているためここでは定義しない */

	// Settings     *ESMESettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	// SettingsHash string              `json:",omitempty" yaml:"setting_hash,omitempty" structs:",omitempty"`
	// Status       *ESMEStatus `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
}

// ESMESendSMSRequest SMS送信リクエスト
type ESMESendSMSRequest struct {
	Destination  string              `json:"destination,omitempty" yaml:"destination,omitempty" structs:",omitempty"` // 宛先 現在は81(+81)開始固定
	Sender       string              `json:"sender,omitempty" yaml:"sender,omitempty" structs:",omitempty"`           // 送信者名 本文中に反映される
	DomainName   string              `json:"domain_name,omitempty" yaml:"domain_name,omitempty" structs:",omitempty"` // Web OTPを利用する際に本文中に記載されるオリジン(FQDNで指定)
	OTPOperation types.EOTPOperation `json:"otpOperation,omitempty" yaml:"otp_operation,omitempty" structs:",omitempty"`
	OTP          string              `json:"otp,omitempty" yaml:"otp,omitempty" structs:",omitempty"` // ワンタイムパスワード、OTPOperationがinputの場合に使用する
}

// ESMESendSMSResponse SMS送信結果
type ESMESendSMSResponse struct {
	MessageID string `json:"messageId,omitempty" yaml:"message_id,omitempty" structs:",omitempty"`
	Status    string `json:"status,omitempty" yaml:"status,omitempty" structs:",omitempty"` // Accepted/Delivered以外にもエラーなどの認識していないデータがありそうなためtypesで型を定義せず一旦string型としておく
	OTP       string `json:"otp,omitempty" yaml:"otp,omitempty" structs:",omitempty"`
}

type ESMELogs struct {
	Logs []*ESMELog `json:"logs,omitempty" yaml:"logs,omitempty"`
}

type ESMELog struct {
	MessageID   string     `json:"messageId,omitempty" yaml:"message_id,omitempty" structs:",omitempty"`
	Status      string     `json:"status,omitempty" yaml:"status,omitempty" structs:",omitempty"` // Accepted/Delivered以外にもエラーなどの認識していないデータがありそうなためtypesで型を定義せず一旦string型としておく
	OTP         string     `json:"otp,omitempty" yaml:"otp,omitempty" structs:",omitempty"`
	Destination string     `json:"destination,omitempty" yaml:"destination,omitempty" structs:",omitempty"`
	SentAt      *time.Time `json:"sentAt,omitempty" yaml:"sent_at,omitempty" structs:",omitempty"`
	DoneAt      *time.Time `json:"doneAt,omitempty" yaml:"done_at,omitempty" structs:",omitempty"`
	RetryCount  int        `json:"retryCounnt,omitempty" yaml:"retry_count,omitempty" structs:",omitempty"`
}
