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
	"fmt"
	"time"

	"github.com/sacloud/iaas-api-go/types"
)

// LocalRouter ローカルルータ
type LocalRouter struct {
	ID           types.ID             `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string               `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string               `yaml:"description"`
	Tags         types.Tags           `yaml:"tags"`
	Icon         *Icon                `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time           `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time           `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability  `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Provider     *Provider            `json:",omitempty" yaml:"provider,omitempty" structs:",omitempty"`
	Settings     *LocalRouterSettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string               `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Status       *LocalRouterStatus   `json:",omitempty" yaml:"status" structs:",omitempty"`
	ServiceClass string               `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
}

// LocalRouterSettings セッティング
type LocalRouterSettings struct {
	LocalRouter *LocalRouterSetting `json:",omitempty" yaml:"local_router,omitempty" structs:",omitempty"`
}

// LocalRouterSettingsUpdate ローカルルータ セッティング更新
type LocalRouterSettingsUpdate struct {
	Settings     *LocalRouterSettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string               `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
}

// LocalRouterSetting セッティング
type LocalRouterSetting struct {
	Switch       *LocalRouterSettingSwitch        `yaml:"switch"`
	Interface    *LocalRouterSettingInterface     `yaml:"interface"`
	Peers        []*LocalRouterSettingPeer        `yaml:"peers"`
	StaticRoutes []*LocalRouterSettingStaticRoute `yaml:"static_routes"`
}

// LocalRouterSettingSwitch ローカルルータのスイッチ設定
type LocalRouterSettingSwitch struct {
	Code     string `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // リソースIDなど
	Category string `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // cloud/vps/専用サーバなどを表す
	ZoneID   string `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // クラウドの場合is1aなど Note: VPSの場合は数値型となる
}

// UnmarshalJSON ZoneIDに数値/文字列が混在する問題への対応
func (l *LocalRouterSettingSwitch) UnmarshalJSON(b []byte) error {
	type alias struct {
		Code     string      `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // リソースIDなど
		Category string      `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // cloud/vps/専用サーバなどを表す
		ZoneID   interface{} `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // クラウドの場合is1aなど Note: VPSの場合は数値型となる
	}

	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	zoneID := ""
	switch v := a.ZoneID.(type) {
	case string:
		zoneID = v
	case int, int8, int16, int32, int64:
		zoneID = fmt.Sprintf("%d", v)
	case float32, float64:
		zoneID = fmt.Sprintf("%d", int(v.(float64)))
	}

	*l = LocalRouterSettingSwitch{
		Code:     a.Code,
		Category: a.Category,
		ZoneID:   zoneID,
	}
	return nil
}

// LocalRouterSettingInterface インターフェース設定
type LocalRouterSettingInterface struct {
	VirtualIPAddress string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	IPAddress        []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	NetworkMaskLen   int      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	VRID             int      `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// LocalRouterSettingPeer ピア設定
type LocalRouterSettingPeer struct {
	ID          string `yaml:"id"` // 文字列でないとエラーになるためtypes.IDではなくstringとする
	SecretKey   string `yaml:"secret_key"`
	Enabled     bool   `yaml:"enabled"`
	Description string `yaml:"description"`
}

// LocalRouterSettingStaticRoute スタティックルート
type LocalRouterSettingStaticRoute struct {
	Prefix  string `yaml:"prefix"`
	NextHop string `yaml:"next_hop"`
}

// LocalRouterStatus ステータス
type LocalRouterStatus struct {
	SecretKeys []string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// LocalRouterHealth ローカルルータのヘルスチェック結果
type LocalRouterHealth struct {
	Peers []*struct {
		ID     types.ID
		Status types.EServerInstanceStatus
		Routes []string
	} `yaml:"peers"`
}
