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
	"time"

	"github.com/sacloud/iaas-api-go/types"
)

// LoadBalancer ロードバランサ
type LoadBalancer struct {
	ID           types.ID              `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string                `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string                `yaml:"description"`
	Tags         types.Tags            `yaml:"tags"`
	Icon         *Icon                 `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time            `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time            `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability   `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Class        string                `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	ServiceClass string                `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Plan         *AppliancePlan        `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	Instance     *Instance             `json:",omitempty" yaml:"instance,omitempty" structs:",omitempty"`
	Interfaces   []*Interface          `json:",omitempty" yaml:"interfaces,omitempty" structs:",omitempty"`
	Switch       *Switch               `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	Settings     *LoadBalancerSettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string                `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Remark       *ApplianceRemark      `json:",omitempty" yaml:"remark,omitempty" structs:",omitempty"`
}

// LoadBalancerSettingsUpdate ロードバランサ
type LoadBalancerSettingsUpdate struct {
	Settings     *LoadBalancerSettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string                `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
}

// LoadBalancerSettings ロードバランサの設定
type LoadBalancerSettings struct {
	LoadBalancer []*LoadBalancerSetting `yaml:"load_balancer"`
}

// MarshalJSON nullの場合に空配列を出力するための実装
func (s LoadBalancerSettings) MarshalJSON() ([]byte, error) {
	if s.LoadBalancer == nil {
		s.LoadBalancer = make([]*LoadBalancerSetting, 0)
	}
	type alias LoadBalancerSettings
	tmp := alias(s)
	return json.Marshal(&tmp)
}

// LoadBalancerSetting ロードバランサの設定
type LoadBalancerSetting struct {
	VirtualIPAddress string                           `json:",omitempty" yaml:"virtual_ip_address,omitempty" structs:",omitempty"`
	Port             types.StringNumber               `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`
	DelayLoop        types.StringNumber               `json:",omitempty" yaml:"delay_loop,omitempty" structs:",omitempty"`
	SorryServer      string                           `json:",omitempty" yaml:"sorry_server,omitempty" structs:",omitempty"`
	Description      string                           `yaml:"description"`
	Servers          []*LoadBalancerDestinationServer `yaml:"servers"`
}

// MarshalJSON nullの場合に空配列を出力するための実装
func (s LoadBalancerSetting) MarshalJSON() ([]byte, error) {
	if s.Servers == nil {
		s.Servers = make([]*LoadBalancerDestinationServer, 0)
	}

	type alias LoadBalancerSetting
	tmp := alias(s)
	return json.Marshal(&tmp)
}

// LoadBalancerDestinationServer ロードバランサ配下の実サーバ
type LoadBalancerDestinationServer struct {
	IPAddress   string                   `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
	Port        types.StringNumber       `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`
	Enabled     types.StringFlag         `yaml:"enabled"`
	HealthCheck *LoadBalancerHealthCheck `json:",omitempty" yaml:"health_check,omitempty" structs:",omitempty"`
}

// LoadBalancerHealthCheck ヘルスチェック
type LoadBalancerHealthCheck struct {
	Protocol       types.Protocol     `json:",omitempty" yaml:"protocol,omitempty" structs:""`                 // プロトコル
	Host           string             `json:",omitempty" yaml:"host,omitempty" structs:""`                     // 対象ホスト
	Path           string             `json:",omitempty" yaml:"path,omitempty" structs:""`                     // HTTP/HTTPSの場合のリクエストパス
	Status         types.StringNumber `json:",omitempty" yaml:"status,omitempty" structs:""`                   // 期待するステータスコード
	Port           types.StringNumber `json:",omitempty" yaml:"port,omitempty" structs:""`                     // ポート番号
	Retry          types.StringNumber `json:",omitempty" yaml:"retry,omitempty" struct:",omitempty"`           // リトライ回数
	ConnectTimeout types.StringNumber `json:",omitempty" yaml:"connect_timeout,omitempty" struct:",omitempty"` // タイムアウト
}

// LoadBalancerStatus ロードバランサのステータス
type LoadBalancerStatus struct {
	VirtualIPAddress string                      `json:",omitempty" yaml:"virtual_ip_address,omitempty" structs:",omitempty"`
	Port             types.StringNumber          `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`
	CPS              types.StringNumber          `json:",omitempty" yaml:"cps,omitempty" structs:",omitempty"`
	Servers          []*LoadBalancerServerStatus `json:",omitempty" yaml:"servers,omitempty" structs:",omitempty"`
}

// LoadBalancerServerStatus ロードバランサの実サーバのステータス
type LoadBalancerServerStatus struct {
	ActiveConn types.StringNumber          `json:",omitempty" yaml:"active_conn,omitempty" structs:",omitempty"`
	Status     types.EServerInstanceStatus `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
	IPAddress  string                      `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
	Port       types.StringNumber          `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`
	CPS        types.StringNumber          `json:",omitempty" yaml:"cps,omitempty" structs:",omitempty"`
}
