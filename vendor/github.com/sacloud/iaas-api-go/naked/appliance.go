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

	"github.com/sacloud/iaas-api-go/types"
)

// ApplianceRemark アプライアンスの設定/ステータスなど
//
// Appliance.Remarkを表現する
type ApplianceRemark struct {
	Zone            *ApplianceRemarkZone          `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
	Switch          *ApplianceRemarkSwitch        `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	VRRP            *ApplianceVRRP                `json:",omitempty" yaml:"vrrp,omitempty" structs:",omitempty"`
	Network         *ApplianceRemarkNetwork       `json:",omitempty" yaml:"network,omitempty" structs:",omitempty"`
	Servers         ApplianceRemarkServers        `yaml:"servers"`
	Plan            *AppliancePlan                `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	DBConf          *ApplianceRemarkDBConf        `json:",omitempty" yaml:"db_conf,omitempty" structs:",omitempty"`        // for database
	SourceAppliance *ApplianceSource              `json:",omitempty" yaml:"db_conf,omitempty" structs:",omitempty"`        // for database
	MobileGateway   *ApplianceRemarkMobileGateway `json:",omitempty" yaml:"mobile_gateway,omitempty" structs:",omitempty"` // for mobile gateway
	Router          *ApplianceRemarkRouter        `json:",omitempty" yaml:"router,omitempty" structs:",omitempty"`         // for vpc router
}

// ApplianceRemarkMobileGateway モバイルゲートウェイのグローバルIP
type ApplianceRemarkMobileGateway struct {
	GlobalAddress string
}

// ApplianceRemarkRouter VPCルータのバージョンなど
type ApplianceRemarkRouter struct {
	VPCRouterVersion int `json:",omitempty" yaml:"vpc_router_version,omitempty" structs:",omitempty"`
}

// ApplianceSource クローン元アプライアンス データベースのクローン時に利用
type ApplianceSource struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}

// UnmarshalJSON 配列/オブジェクトが混在することへの対応
func (s *ApplianceSource) UnmarshalJSON(b []byte) error {
	if string(b) == "[]" {
		return nil
	}
	type alias ApplianceSource

	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	*s = ApplianceSource(a)
	return nil
}

// AppliancePlan アプライアンスプラン
type AppliancePlan struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}

// ApplianceVRRP アプライアンスのVRRPの設定
type ApplianceVRRP struct {
	VRID int `json:",omitempty" yaml:"vrid,omitempty" structs:",omitempty"`
}

// ApplianceRemarkNetwork Appliance ネットワーク設定
type ApplianceRemarkNetwork struct {
	DefaultRoute   string `json:",omitempty" yaml:"default_route,omitempty" structs:",omitempty"`
	NetworkMaskLen int    `json:",omitempty" yaml:"network_mask_len,omitempty" structs:",omitempty"`
}

// ApplianceRemarkServers Applianceの稼働している仮想サーバのIPアドレス
type ApplianceRemarkServers []*ApplianceRemarkServer

// ApplianceRemarkServer Applianceの稼働している仮想サーバのIPアドレス
type ApplianceRemarkServer struct {
	IPAddress string `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
}

// ApplianceRemarkSwitch Applianceに接続されているスイッチのID
type ApplianceRemarkSwitch struct {
	ID    types.ID     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Scope types.EScope `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
}

// ApplianceRemarkZone Applianceの属するゾーンのID
type ApplianceRemarkZone struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}

// UnmarshalJSON 配列/オブジェクトが混在することへの対応
func (s *ApplianceRemarkNetwork) UnmarshalJSON(b []byte) error {
	if string(b) == "[]" {
		return nil
	}
	type alias ApplianceRemarkNetwork

	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	*s = ApplianceRemarkNetwork(a)
	return nil
}

// UnmarshalJSON 配列/オブジェクトが混在することへの対応
func (s *ApplianceRemarkServer) UnmarshalJSON(b []byte) error {
	if string(b) == "[]" {
		return nil
	}
	type alias ApplianceRemarkServer

	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	*s = ApplianceRemarkServer(a)
	return nil
}

// UnmarshalJSON 配列/オブジェクトが混在することへの対応
func (s *ApplianceRemarkServers) UnmarshalJSON(b []byte) error {
	if string(b) == "[[]]" {
		return nil
	}
	if string(b) == `[""]` {
		return nil
	}
	type alias ApplianceRemarkServers

	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	*s = ApplianceRemarkServers(a)
	return nil
}

// MarshalJSON APIの要求するJSONフォーマットへの変換
//
// 値がからの場合に配列、かつ内部に空オブジェクトを指定する。(主にVPCルータへの対応)
func (s *ApplianceRemarkServers) MarshalJSON() ([]byte, error) {
	if s == nil || len(*s) == 0 {
		return []byte("[{}]"), nil
	}

	type alias ApplianceRemarkServers

	a := alias(*s)
	return json.Marshal(a)
}

// ApplianceRemarkDBConf データベース設定
type ApplianceRemarkDBConf struct {
	Common *ApplianceRemarkDBConfCommon `json:",omitempty" yaml:"common,omitempty" structs:",omitempty"`
}

// ApplianceRemarkDBConfCommon データベース設定
type ApplianceRemarkDBConfCommon struct {
	DatabaseName     string `json:",omitempty" yaml:"database_name,omitempty" structs:",omitempty"`
	DatabaseVersion  string `json:",omitempty" yaml:"database_version,omitempty" structs:",omitempty"`
	DatabaseRevision string `json:",omitempty" yaml:"database_revision,omitempty" structs:",omitempty"`
	DefaultUser      string `json:",omitempty" yaml:"default_user,omitempty" structs:",omitempty"`
	UserPassword     string `json:",omitempty" yaml:"user_password,omitempty" structs:",omitempty"`
}
