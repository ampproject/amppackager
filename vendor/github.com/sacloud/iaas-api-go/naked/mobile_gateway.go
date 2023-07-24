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

// MobileGateway モバイルゲートウェイ
type MobileGateway struct {
	ID           types.ID                `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Class        string                  `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	Name         string                  `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Tags         types.Tags              `yaml:"tags"`
	Description  string                  `yaml:"description"`
	Plan         *AppliancePlan          `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	Settings     *MobileGatewaySettings  `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string                  `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Remark       *ApplianceRemark        `json:",omitempty" yaml:"remark,omitempty" structs:",omitempty"`
	Availability types.EAvailability     `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Instance     *Instance               `json:",omitempty" yaml:"instance,omitempty" structs:",omitempty"`
	ServiceClass string                  `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time              `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	Icon         *Icon                   `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	Switch       *Switch                 `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	Interfaces   MobileGatewayInterfaces `json:",omitempty" yaml:"interfaces,omitempty" structs:",omitempty"`
}

// MobileGatewaySettingsUpdate モバイルゲートウェイ
type MobileGatewaySettingsUpdate struct {
	Settings     *MobileGatewaySettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string                 `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
}

// MobileGatewayInterfaces 要素がnullにことがある場合に対応するためのtype
//
// 例: モバイルゲートウェイ 作成時、eth0/eth1の2要素が返ってくるがeth1の分はnullとなっている。
type MobileGatewayInterfaces []*Interface

// UnmarshalJSON 配列中にnullが返ってくる(VPCルータなど)への対応
func (i *MobileGatewayInterfaces) UnmarshalJSON(b []byte) error {
	type alias MobileGatewayInterfaces
	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}

	var dest []*Interface
	for i, v := range a {
		if v != nil {
			if v.Index == 0 {
				v.Index = i
			}
			dest = append(dest, v)
		}
	}

	*i = MobileGatewayInterfaces(dest)
	return nil
}

// MobileGatewaySettings モバイルゲートウェイ セッティング
type MobileGatewaySettings struct {
	MobileGateway *MobileGatewaySetting `json:",omitempty" yaml:"mobile_gateway,omitempty" structs:",omitempty"`
}

// MobileGatewaySetting モバイルゲートウェイ セッティング
type MobileGatewaySetting struct {
	Interfaces               MobileGatewayInterfacesSettings        `json:",omitempty" yaml:"interfaces,omitempty" structs:",omitempty"`
	InternetConnection       *MobileGatewayInternetConnection       `json:",omitempty" yaml:"internet_connection,omitempty" structs:",omitempty"`
	StaticRoutes             []*MobileGatewayStaticRoute            `json:",omitempty" yaml:"static_routes,omitempty" structs:",omitempty"`
	InterDeviceCommunication *MobileGatewayInterDeviceCommunication `json:",omitempty" yaml:"inter_device_communication,omitempty" structs:",omitempty"`
}

// MobileGatewayInterDeviceCommunication デバイス間通信
type MobileGatewayInterDeviceCommunication struct {
	Enabled types.StringFlag `yaml:"enabled"`
}

// MobileGatewayInternetConnection インターネット接続
type MobileGatewayInternetConnection struct {
	Enabled types.StringFlag `yaml:"enabled"`
}

// MobileGatewayStaticRoute スタティックルート
type MobileGatewayStaticRoute struct {
	Prefix  string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	NextHop string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// MobileGatewayResolver DNS登録用パラメータ
type MobileGatewayResolver struct {
	SimGroup *MobileGatewaySIMGroup `json:"sim_group,omitempty" yaml:"sim_group,omitempty" structs:",omitempty"`
}

// MobileGatewaySIMGroup DNS登録用SIMグループ値
type MobileGatewaySIMGroup struct {
	DNS1 string `json:"dns_1,omitempty" yaml:"dns_1,omitempty" structs:",omitempty"`
	DNS2 string `json:"dns_2,omitempty" yaml:"dns_2,omitempty" structs:",omitempty"`
}

// UnmarshalJSON JSONアンマーシャル(配列、オブジェクトが混在するためここで対応)
func (m *MobileGatewaySIMGroup) UnmarshalJSON(data []byte) error {
	targetData := strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", "")
	if targetData == `[]` {
		return nil
	}

	type alias MobileGatewaySIMGroup
	tmp := alias{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MobileGatewaySIMGroup(tmp)
	return nil
}

// MobileGatewaySIMRoute SIMルート
type MobileGatewaySIMRoute struct {
	ICCID      string `json:"iccid,omitempty" yaml:"iccid,omitempty" structs:",omitempty"`
	Prefix     string `json:"prefix,omitempty" yaml:"prefix,omitempty" structs:",omitempty"`
	ResourceID string `json:"resource_id,omitempty" yaml:"resource_id,omitempty" structs:",omitempty"`
}

// MobileGatewaySIMRoutes SIMルート一覧
type MobileGatewaySIMRoutes struct {
	SIMRoutes []*MobileGatewaySIMRoute `json:"sim_routes" yaml:"sim_routes,omitempty" structs:",omitempty"`
}

// TrafficStatus トラフィックコントロール 当月通信量
type TrafficStatus struct {
	UplinkBytes    types.StringNumber `json:"uplink_bytes,omitempty" yaml:"uplink_bytes,omitempty" structs:",omitempty"`
	DownlinkBytes  types.StringNumber `json:"downlink_bytes,omitempty" yaml:"downlink_bytes,omitempty" structs:",omitempty"`
	TrafficShaping bool               `json:"traffic_shaping" yaml:"traffic_shaping"` // 帯域制限
}

// UnmarshalJSON JSONアンマーシャル(uint64文字列対応)
func (s *TrafficStatus) UnmarshalJSON(data []byte) error {
	type alias TrafficStatus
	tmp := alias{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = TrafficStatus(tmp)
	return nil
}

// TrafficMonitoringConfig トラフィックコントロール 設定
type TrafficMonitoringConfig struct {
	TrafficQuotaInMB     int                          `json:"traffic_quota_in_mb" yaml:"traffic_quota_in_mb"`
	BandWidthLimitInKbps int                          `json:"bandwidth_limit_in_kbps" yaml:"bandwidth_limit_in_kbps"`
	EMailConfig          TrafficMonitoringNotifyEmail `json:"email_config" yaml:"email_config"`
	SlackConfig          TrafficMonitoringNotifySlack `json:"slack_config" yaml:"slack_config"`
	AutoTrafficShaping   bool                         `json:"auto_traffic_shaping" yaml:"auto_traffic_shaping"`
}

// TrafficMonitoringNotifyEmail トラフィックコントロール通知設定
type TrafficMonitoringNotifyEmail struct {
	Enabled bool `json:"enabled" yaml:"enabled"` // 有効/無効
}

// TrafficMonitoringNotifySlack トラフィックコントロール通知設定
type TrafficMonitoringNotifySlack struct {
	Enabled             bool   `json:"enabled" yaml:"enabled"`                         // 有効/無効
	IncomingWebhooksURL string `json:"slack_url,omitempty" yaml:"slack_url,omitempty"` // Slack通知の場合のWebhook URL
}

// MobileGatewayInterface インターフェース
type MobileGatewayInterface struct {
	IPAddress      []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	NetworkMaskLen int      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	// Index 仮想フィールド、VPCルータなどでInterfaces(実体は[]*Interface)を扱う場合にUnmarshalJSONの中で設定される
	//
	// Findした際のAPIからの応答にも同名のフィールドが含まれるが無関係。
	Index int
}

// MobileGatewayInterfacesSettings Interface配列
//
// 配列中にnullが返ってくる(VPCルータなど)への対応のためのtype
type MobileGatewayInterfacesSettings []*MobileGatewayInterface

// UnmarshalJSON 配列中にnullが返ってくる(VPCルータなど)への対応
func (i *MobileGatewayInterfacesSettings) UnmarshalJSON(b []byte) error {
	type alias MobileGatewayInterfacesSettings
	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}

	var dest []*MobileGatewayInterface
	for i, v := range a {
		if v != nil {
			if v.Index == 0 {
				v.Index = i
			}
			dest = append(dest, v)
		}
	}

	*i = MobileGatewayInterfacesSettings(dest)
	return nil
}

// MarshalJSON 配列中にnullが入る場合(VPCルータなど)への対応
func (i MobileGatewayInterfacesSettings) MarshalJSON() ([]byte, error) {
	max := 0
	for _, iface := range i {
		if max < iface.Index {
			max = iface.Index
		}
	}

	var dest = make([]*MobileGatewayInterface, max+1)
	for _, iface := range i {
		dest[iface.Index] = iface
	}

	return json.Marshal(dest)
}

// MarshalJSON JSON
func (i *MobileGatewayInterface) MarshalJSON() ([]byte, error) {
	type alias struct {
		IPAddress      []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
		NetworkMaskLen int      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	}

	tmp := alias{
		IPAddress:      i.IPAddress,
		NetworkMaskLen: i.NetworkMaskLen,
	}
	return json.Marshal(tmp)
}
