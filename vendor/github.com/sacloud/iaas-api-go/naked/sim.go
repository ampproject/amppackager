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
	"encoding/json"
	"strings"
	"time"

	"github.com/sacloud/iaas-api-go/types"
)

// SIM SIM
type SIM struct {
	ID           types.ID     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string       `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string       `yaml:"description"`
	Tags         types.Tags   `yaml:"tags"`
	Status       *SIMStatus   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	ServiceClass string       `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Availability string       `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	CreatedAt    time.Time    `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	ModifiedAt   time.Time    `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Provider     *SIMProvider `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Icon         *Icon        `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Remark       *SIMRemark   `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // Remark
}

// SIMStatus SIMステータス
type SIMStatus struct {
	ICCID   string   `json:",omitempty" yaml:"iccid,omitempty" structs:",omitempty"`  // ICCID
	SIMInfo *SIMInfo `json:"sim,omitempty" yaml:"sim,omitempty" structs:",omitempty"` // SIM詳細情報
}

// SIMInfo SIM詳細情報
type SIMInfo struct {
	ICCID                      string           `json:"iccid,omitempty" yaml:"iccid,omitempty" structs:",omitempty"`
	IMSI                       []string         `json:"imsi,omitempty" yaml:"imsi,omitempty" structs:",omitempty"`
	IMEI                       string           `json:"imei,omitempty" yaml:"imei,omitempty" structs:",omitempty"`
	IP                         string           `json:"ip,omitempty" yaml:"ip,omitempty" structs:",omitempty"`
	SessionStatus              string           `json:"session_status,omitempty" yaml:"session_status,omitempty" structs:",omitempty"`
	IMEILock                   bool             `json:"imei_lock" yaml:"imei_lock"`
	Registered                 bool             `json:"registered" yaml:"registered"`
	Activated                  bool             `json:"activated" yaml:"activated"`
	ResourceID                 string           `json:"resource_id,omitempty" yaml:"resource_id,omitempty" structs:",omitempty"`
	RegisteredDate             time.Time        `json:"registered_date,omitempty" yaml:"registered_date,omitempty" structs:",omitempty"`
	ActivatedDate              time.Time        `json:"activated_date,omitempty" yaml:"activated_date,omitempty" structs:",omitempty"`
	DeactivatedDate            time.Time        `json:"deactivated_date,omitempty" yaml:"deactivated_date,omitempty" structs:",omitempty"`
	SIMGroupID                 string           `json:"simgroup_id,omitempty" yaml:"simgroup_id,omitempty" structs:",omitempty"`
	TrafficBytesOfCurrentMonth *SIMTrafficBytes `json:"traffic_bytes_of_current_month,omitempty" yaml:"traffic_bytes_of_current_month,omitempty" structs:",omitempty"`
	ConnectedIMEI              string           `json:"connected_imei,omitempty" yaml:"connected_imei,omitempty" structs:",omitempty"`
}

// SIMTrafficBytes 当月通信量
type SIMTrafficBytes struct {
	UplinkBytes   types.StringNumber `json:"uplink_bytes,omitempty" yaml:"uplink_bytes,omitempty" structs:",omitempty"`
	DownlinkBytes types.StringNumber `json:"downlink_bytes,omitempty" yaml:"downlink_bytes,omitempty" structs:",omitempty"`
}

// SIMProvider SIMプロバイダー
type SIMProvider struct {
	ID           int    `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Class        string `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	Name         string `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	ServiceClass string `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
}

// SIMRemark remark
type SIMRemark struct {
	PassCode string `json:",omitempty" yaml:"pass_code,omitempty" structs:",omitempty"`
}

// UnmarshalJSON JSONアンマーシャル(配列、オブジェクトが混在するためここで対応)
func (s *SIMTrafficBytes) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}
	type alias SIMTrafficBytes
	tmp := alias{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*s = SIMTrafficBytes(tmp)
	return nil
}

// SIMLog SIMログ
type SIMLog struct {
	Date          *time.Time `json:"date,omitempty" yaml:"date,omitempty" structs:",omitempty"`
	SessionStatus string     `json:"session_status,omitempty" yaml:"session_status,omitempty" structs:",omitempty"`
	ResourceID    string     `json:"resource_id,omitempty" yaml:"resource_id,omitempty" structs:",omitempty"`
	IMEI          string     `json:"imei,omitempty" yaml:"imei,omitempty" structs:",omitempty"`
	IMSI          string     `json:"imsi,omitempty" yaml:"imsi,omitempty" structs:",omitempty"`
}

// SIMNetworkOperatorConfig SIM通信キャリア設定
type SIMNetworkOperatorConfig struct {
	Allow       bool   `json:"allow" yaml:"allow"`
	CountryCode string `json:"country_code,omitempty" yaml:"country_code,omitempty" structs:",omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty" structs:",omitempty"`
}

// SIMNetworkOperatorConfigs SIM通信キャリア設定 リクエストパラメータ
type SIMNetworkOperatorConfigs struct {
	NetworkOperatorConfigs []*SIMNetworkOperatorConfig `json:"network_operator_config,omitempty" yaml:"network_operator_config,omitempty" structs:",omitempty"`
}

// SIMAssignIPRequest IPアドレスアサイン リクエストパラメータ
type SIMAssignIPRequest struct {
	IP string `json:"ip"`
}

// SIMIMEILockRequest IMEIロック リクエストパラメータ
type SIMIMEILockRequest struct {
	IMEI string `json:"imei"`
}
