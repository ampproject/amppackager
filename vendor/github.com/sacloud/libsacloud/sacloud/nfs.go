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

import (
	"encoding/json"
)

// NFS NFS
type NFS struct {
	*Appliance // アプライアンス共通属性

	Remark   *NFSRemark   `json:",omitempty"` // リマーク
	Settings *NFSSettings `json:",omitempty"` // NFS設定
}

// NFSRemark リマーク
type NFSRemark struct {
	*ApplianceRemarkBase
	Plan *struct {
		ID json.Number `json:",omitempty"`
	} `json:",omitempty"` // プラン
	// TODO Zone
	//Zone *Resource
	//SourceAppliance *Resource // クローン元DB
}

// SetRemarkPlanID プランID設定
func (n NFSRemark) SetRemarkPlanID(planID ID) {
	if n.Plan == nil {
		n.Plan = &struct {
			ID json.Number `json:",omitempty"`
		}{}
	}
	n.Plan.ID = json.Number(planID)
}

// NFSSettings NFS設定リスト
type NFSSettings struct {
}

// NFSPlan プラン(HDD/SSD)
type NFSPlan int

var (
	// NFSPlanHDD 標準プラン(HDD)
	NFSPlanHDD = NFSPlan(1)
	// NFSPlanSSD SSHプラン
	NFSPlanSSD = NFSPlan(2)
)

// String NFSプランの文字列表現
func (p NFSPlan) String() string {
	switch p {
	case NFSPlanHDD:
		return "HDD"
	case NFSPlanSSD:
		return "SSD"
	default:
		return ""
	}
}

// NFSSize NFSサイズ
type NFSSize int

var (
	// NFSSize20G 20Gプラン
	NFSSize20G = NFSSize(20)
	// NFSSize100G 100Gプラン
	NFSSize100G = NFSSize(100)
	// NFSSize500G 500Gプラン
	NFSSize500G = NFSSize(500)
	// NFSSize1T 1T(1024GB)プラン
	NFSSize1T = NFSSize(1024 * 1)
	// NFSSize2T 2T(2048GB)プラン
	NFSSize2T = NFSSize(1024 * 2)
	// NFSSize4T 4T(4096GB)プラン
	NFSSize4T = NFSSize(1024 * 4)
	// NFSSize8T 8TBプラン
	NFSSize8T = NFSSize(1024 * 8)
	// NFSSize12T 12TBプラン
	NFSSize12T = NFSSize(1024 * 12)
)

// AllowNFSNormalPlanSizes 指定可能なNFSサイズ(標準プラン)
func AllowNFSNormalPlanSizes() []int {
	return []int{
		int(NFSSize100G),
		int(NFSSize500G),
		int(NFSSize1T),
		int(NFSSize2T),
		int(NFSSize4T),
		int(NFSSize8T),
		int(NFSSize12T),
	}
}

// AllowNFSSSDPlanSizes 指定可能なNFSサイズ(SSDプラン)
func AllowNFSSSDPlanSizes() []int {
	return []int{
		int(NFSSize20G),
		int(NFSSize100G),
		int(NFSSize500G),
		int(NFSSize1T),
		int(NFSSize2T),
		int(NFSSize4T),
	}
}

// AllowNFSAllPlanSizes 指定可能なNFSサイズ(HDD/SSD両方)
//
// ディスクプランによって選択できないサイズを含むため、利用する側で適切にバリデーションを行う必要がある
func AllowNFSAllPlanSizes() []int {
	return []int{
		int(NFSSize20G),
		int(NFSSize100G),
		int(NFSSize500G),
		int(NFSSize1T),
		int(NFSSize2T),
		int(NFSSize4T),
		int(NFSSize8T),
		int(NFSSize12T),
	}
}

// CreateNFSValue NFS作成用パラメーター
type CreateNFSValue struct {
	SwitchID        ID        // 接続先スイッチID
	IPAddress       string    // IPアドレス
	MaskLen         int       // ネットワークマスク長
	DefaultRoute    string    // デフォルトルート
	Name            string    // 名称
	Description     string    // 説明
	Tags            []string  // タグ
	Icon            *Resource // アイコン
	SourceAppliance *Resource // クローン元NFS
}

// NewNFS NFS作成(冗長化なし)
func NewNFS(values *CreateNFSValue) *NFS {

	return &NFS{
		Appliance: &Appliance{
			Class:           "nfs",
			propName:        propName{Name: values.Name},
			propDescription: propDescription{Description: values.Description},
			propTags:        propTags{Tags: values.Tags},
			//propPlanID:      propPlanID{Plan: &Resource{ID: int64(values.Plan)}},
			propIcon: propIcon{
				&Icon{
					Resource: values.Icon,
				},
			},
		},
		Remark: &NFSRemark{
			ApplianceRemarkBase: &ApplianceRemarkBase{
				Switch: &ApplianceRemarkSwitch{
					ID: values.SwitchID,
				},
				Network: &ApplianceRemarkNetwork{
					NetworkMaskLen: values.MaskLen,
					DefaultRoute:   values.DefaultRoute,
				},
				Servers: []interface{}{
					map[string]interface{}{"IPAddress": values.IPAddress},
				},
			},
			//propPlanID: propPlanID{Plan: &Resource{ID: int64(values.Plan)}},
		},
	}

}

// IPAddress IPアドレスを取得
func (n *NFS) IPAddress() string {
	if len(n.Remark.Servers) < 1 {
		return ""
	}

	v, ok := n.Remark.Servers[0].(map[string]interface{})
	if !ok {
		return ""
	}

	if ip, ok := v["IPAddress"]; ok {
		return ip.(string)
	}
	return ""
}

// NetworkMaskLen ネットワークマスク長を取得
func (n *NFS) NetworkMaskLen() int {
	if n.Remark.Network == nil {
		return -1
	}
	return n.Remark.Network.NetworkMaskLen
}

// DefaultRoute デフォルトゲートウェイを取得
func (n *NFS) DefaultRoute() string {
	if n.Remark.Network == nil {
		return ""
	}
	return n.Remark.Network.DefaultRoute
}

// NFSPlans NFSプラン
type NFSPlans struct {
	HDD []NFSPlanValue
	SSD []NFSPlanValue
}

// FindPlanID プランとサイズからプランIDを取得
func (p NFSPlans) FindPlanID(plan NFSPlan, size NFSSize) ID {
	var plans []NFSPlanValue
	switch plan {
	case NFSPlanHDD:
		plans = p.HDD
	case NFSPlanSSD:
		plans = p.SSD
	default:
		return -1
	}

	for _, plan := range plans {
		if plan.Availability == "available" && plan.Size == int(size) {
			return plan.PlanID
		}
	}

	return -1
}

// FindByPlanID プランIDから該当プランを取得
func (p NFSPlans) FindByPlanID(planID ID) (NFSPlan, *NFSPlanValue) {

	for _, plan := range p.SSD {
		id := plan.PlanID
		if id == planID {
			return NFSPlanSSD, &plan
		}
	}

	for _, plan := range p.HDD {
		id := plan.PlanID
		if id == planID {
			return NFSPlanHDD, &plan
		}
	}
	return NFSPlan(-1), nil
}

// NFSPlanValue NFSプラン
type NFSPlanValue struct {
	Size         int    `json:"size"`
	Availability string `json:"availability"`
	PlanID       ID     `json:"planId"`
}
