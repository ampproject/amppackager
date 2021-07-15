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

package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/sacloud"
)

//HACK: さくらのAPI側仕様: Applianceの内容によってJSONフォーマットが異なるため
//      ロードバランサ/VPCルータそれぞれでリクエスト/レスポンスデータ型を定義する。

// SearchVPCRouterResponse VPCルーター検索レスポンス
type SearchVPCRouterResponse struct {
	// Total 総件数
	Total int `json:",omitempty"`
	// From ページング開始位置
	From int `json:",omitempty"`
	// Count 件数
	Count int `json:",omitempty"`
	// VPCRouters VPCルーター リスト
	VPCRouters []sacloud.VPCRouter `json:"Appliances,omitempty"`
}

type vpcRouterRequest struct {
	VPCRouter *sacloud.VPCRouter     `json:"Appliance,omitempty"`
	From      int                    `json:",omitempty"`
	Count     int                    `json:",omitempty"`
	Sort      []string               `json:",omitempty"`
	Filter    map[string]interface{} `json:",omitempty"`
	Exclude   []string               `json:",omitempty"`
	Include   []string               `json:",omitempty"`
}

type vpcRouterResponse struct {
	*sacloud.ResultFlagValue
	*sacloud.VPCRouter `json:"Appliance,omitempty"`
	Success            interface{} `json:",omitempty"` //HACK: さくらのAPI側仕様: 戻り値:Successがbool値へ変換できないためinterface{}
}

type vpcRouterStatusResponse struct {
	*sacloud.ResultFlagValue
	*sacloud.VPCRouterStatus `json:"Router"`
	Success                  interface{} `json:",omitempty"` //HACK: さくらのAPI側仕様: 戻り値:Successがbool値へ変換できないためinterface{}
}

type vpcRouterS2sConnInfoResponse struct {
	*sacloud.ResultFlagValue
	*sacloud.SiteToSiteConnectionInfo
	Name string
}

// VPCRouterAPI VPCルーターAPI
type VPCRouterAPI struct {
	*baseAPI
}

// NewVPCRouterAPI VPCルーターAPI作成
func NewVPCRouterAPI(client *Client) *VPCRouterAPI {
	return &VPCRouterAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "appliance"
			},
			FuncBaseSearchCondition: func() *sacloud.Request {
				res := &sacloud.Request{}
				res.AddFilter("Class", "vpcrouter")
				return res
			},
		},
	}
}

// Find 検索
func (api *VPCRouterAPI) Find() (*SearchVPCRouterResponse, error) {
	data, err := api.client.newRequest("GET", api.getResourceURL(), api.getSearchState())
	if err != nil {
		return nil, err
	}
	var res SearchVPCRouterResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (api *VPCRouterAPI) request(f func(*vpcRouterResponse) error) (*sacloud.VPCRouter, error) {
	res := &vpcRouterResponse{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.VPCRouter, nil
}

func (api *VPCRouterAPI) createRequest(value *sacloud.VPCRouter) *vpcRouterResponse {
	return &vpcRouterResponse{VPCRouter: value}
}

// New 新規作成用パラメーター作成
func (api *VPCRouterAPI) New() *sacloud.VPCRouter {
	return sacloud.CreateNewVPCRouter()
}

// Create 新規作成
func (api *VPCRouterAPI) Create(value *sacloud.VPCRouter) (*sacloud.VPCRouter, error) {
	return api.request(func(res *vpcRouterResponse) error {
		return api.create(api.createRequest(value), res)
	})
}

// Read 読み取り
func (api *VPCRouterAPI) Read(id sacloud.ID) (*sacloud.VPCRouter, error) {
	return api.request(func(res *vpcRouterResponse) error {
		return api.read(id, nil, res)
	})
}

// Update 更新
func (api *VPCRouterAPI) Update(id sacloud.ID, value *sacloud.VPCRouter) (*sacloud.VPCRouter, error) {
	return api.request(func(res *vpcRouterResponse) error {
		return api.update(id, api.createRequest(value), res)
	})
}

// UpdateSetting 設定更新
func (api *VPCRouterAPI) UpdateSetting(id sacloud.ID, value *sacloud.VPCRouter) (*sacloud.VPCRouter, error) {
	req := &sacloud.VPCRouter{
		// Settings
		Settings: value.Settings,
	}
	return api.request(func(res *vpcRouterResponse) error {
		return api.update(id, api.createRequest(req), res)
	})
}

// Delete 削除
func (api *VPCRouterAPI) Delete(id sacloud.ID) (*sacloud.VPCRouter, error) {
	return api.request(func(res *vpcRouterResponse) error {
		return api.delete(id, nil, res)
	})
}

// Config 設定変更の反映
func (api *VPCRouterAPI) Config(id sacloud.ID) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/config", api.getResourceURL(), id)
	)
	return api.modify(method, uri, nil)
}

// ConnectToSwitch 指定のインデックス位置のNICをスイッチへ接続
func (api *VPCRouterAPI) ConnectToSwitch(id sacloud.ID, switchID sacloud.ID, nicIndex int) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/interface/%d/to/switch/%d", api.getResourceURL(), id, nicIndex, switchID)
	)
	return api.modify(method, uri, nil)
}

// DisconnectFromSwitch 指定のインデックス位置のNICをスイッチから切断
func (api *VPCRouterAPI) DisconnectFromSwitch(id sacloud.ID, nicIndex int) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/interface/%d/to/switch", api.getResourceURL(), id, nicIndex)
	)
	return api.modify(method, uri, nil)
}

// IsUp 起動しているか判定
func (api *VPCRouterAPI) IsUp(id sacloud.ID) (bool, error) {
	router, err := api.Read(id)
	if err != nil {
		return false, err
	}
	return router.Instance.IsUp(), nil
}

// IsDown ダウンしているか判定
func (api *VPCRouterAPI) IsDown(id sacloud.ID) (bool, error) {
	router, err := api.Read(id)
	if err != nil {
		return false, err
	}
	return router.Instance.IsDown(), nil
}

// Boot 起動
func (api *VPCRouterAPI) Boot(id sacloud.ID) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/power", api.getResourceURL(), id)
	)
	return api.modify(method, uri, nil)
}

// Shutdown シャットダウン(graceful)
func (api *VPCRouterAPI) Shutdown(id sacloud.ID) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/power", api.getResourceURL(), id)
	)

	return api.modify(method, uri, nil)
}

// Stop シャットダウン(force)
func (api *VPCRouterAPI) Stop(id sacloud.ID) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/power", api.getResourceURL(), id)
	)

	return api.modify(method, uri, map[string]bool{"Force": true})
}

// RebootForce 再起動
func (api *VPCRouterAPI) RebootForce(id sacloud.ID) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/reset", api.getResourceURL(), id)
	)

	return api.modify(method, uri, nil)
}

// SleepUntilUp 起動するまで待機
func (api *VPCRouterAPI) SleepUntilUp(id sacloud.ID, timeout time.Duration) error {
	handler := waitingForUpFunc(func() (hasUpDown, error) {
		return api.Read(id)
	}, 0)
	return blockingPoll(handler, timeout)
}

// SleepUntilDown ダウンするまで待機
func (api *VPCRouterAPI) SleepUntilDown(id sacloud.ID, timeout time.Duration) error {
	handler := waitingForDownFunc(func() (hasUpDown, error) {
		return api.Read(id)
	}, 0)
	return blockingPoll(handler, timeout)
}

// SleepWhileCopying コピー終了まで待機
//
// maxRetry: リクエストタイミングによって、コピー完了までの間に404エラーとなる場合がある。
// 通常そのまま待てばコピー完了するため、404エラーが発生してもmaxRetryで指定した回数分は待機する。
func (api *VPCRouterAPI) SleepWhileCopying(id sacloud.ID, timeout time.Duration, maxRetry int) error {
	handler := waitingForAvailableFunc(func() (hasAvailable, error) {
		return api.Read(id)
	}, maxRetry)
	return blockingPoll(handler, timeout)
}

// AsyncSleepWhileCopying コピー終了まで待機(非同期)
func (api *VPCRouterAPI) AsyncSleepWhileCopying(id sacloud.ID, timeout time.Duration, maxRetry int) (chan (interface{}), chan (interface{}), chan (error)) {
	handler := waitingForAvailableFunc(func() (hasAvailable, error) {
		return api.Read(id)
	}, maxRetry)
	return poll(handler, timeout)
}

// AddStandardInterface スタンダードプランでのインターフェース追加
func (api *VPCRouterAPI) AddStandardInterface(routerID sacloud.ID, switchID sacloud.ID, ipaddress string, maskLen int) (*sacloud.VPCRouter, error) {
	return api.addInterface(routerID, switchID, &sacloud.VPCRouterInterface{
		IPAddress:        []string{ipaddress},
		NetworkMaskLen:   maskLen,
		VirtualIPAddress: "",
	})
}

// AddPremiumInterface プレミアムプランでのインターフェース追加
func (api *VPCRouterAPI) AddPremiumInterface(routerID sacloud.ID, switchID sacloud.ID, ipaddresses []string, maskLen int, virtualIP string) (*sacloud.VPCRouter, error) {
	return api.addInterface(routerID, switchID, &sacloud.VPCRouterInterface{
		IPAddress:        ipaddresses,
		NetworkMaskLen:   maskLen,
		VirtualIPAddress: virtualIP,
	})
}

func (api *VPCRouterAPI) addInterface(routerID sacloud.ID, switchID sacloud.ID, routerNIC *sacloud.VPCRouterInterface) (*sacloud.VPCRouter, error) {
	router, err := api.Read(routerID)
	if err != nil {
		return nil, err
	}
	req := &sacloud.VPCRouter{Settings: &sacloud.VPCRouterSettings{}}

	if router.Settings == nil {
		req.Settings = &sacloud.VPCRouterSettings{
			Router: &sacloud.VPCRouterSetting{
				Interfaces: []*sacloud.VPCRouterInterface{nil},
			},
		}
	} else {
		req.Settings.Router = router.Settings.Router
	}

	index := len(req.Settings.Router.Interfaces) // add to last
	return api.addInterfaceAt(routerID, switchID, routerNIC, index)
}

// AddStandardInterfaceAt スタンダードプランでの指定位置へのインターフェース追加
func (api *VPCRouterAPI) AddStandardInterfaceAt(routerID sacloud.ID, switchID sacloud.ID, ipaddress string, maskLen int, index int) (*sacloud.VPCRouter, error) {
	return api.addInterfaceAt(routerID, switchID, &sacloud.VPCRouterInterface{
		IPAddress:        []string{ipaddress},
		NetworkMaskLen:   maskLen,
		VirtualIPAddress: "",
	}, index)
}

// AddPremiumInterfaceAt プレミアムプランでの指定位置へのインターフェース追加
func (api *VPCRouterAPI) AddPremiumInterfaceAt(routerID sacloud.ID, switchID sacloud.ID, ipaddresses []string, maskLen int, virtualIP string, index int) (*sacloud.VPCRouter, error) {
	return api.addInterfaceAt(routerID, switchID, &sacloud.VPCRouterInterface{
		IPAddress:        ipaddresses,
		NetworkMaskLen:   maskLen,
		VirtualIPAddress: virtualIP,
	}, index)
}

func (api *VPCRouterAPI) addInterfaceAt(routerID sacloud.ID, switchID sacloud.ID, routerNIC *sacloud.VPCRouterInterface, index int) (*sacloud.VPCRouter, error) {
	router, err := api.Read(routerID)
	if err != nil {
		return nil, err
	}

	req := &sacloud.VPCRouter{Settings: &sacloud.VPCRouterSettings{}}

	if router.Settings == nil {
		req.Settings = &sacloud.VPCRouterSettings{
			Router: &sacloud.VPCRouterSetting{
				Interfaces: []*sacloud.VPCRouterInterface{nil},
			},
		}
	} else {
		req.Settings.Router = router.Settings.Router
	}

	//connect to switch
	_, err = api.ConnectToSwitch(routerID, switchID, index)
	if err != nil {
		return nil, err
	}

	for i := 0; i < index; i++ {
		if len(req.Settings.Router.Interfaces) < index {
			req.Settings.Router.Interfaces = append(req.Settings.Router.Interfaces, nil)
		}
	}

	if len(req.Settings.Router.Interfaces) < index+1 {
		req.Settings.Router.Interfaces = append(req.Settings.Router.Interfaces, routerNIC)
	} else {
		req.Settings.Router.Interfaces[index] = routerNIC
	}

	res, err := api.UpdateSetting(routerID, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

// DeleteInterfaceAt 指定位置のインターフェース削除
func (api *VPCRouterAPI) DeleteInterfaceAt(routerID sacloud.ID, index int) (*sacloud.VPCRouter, error) {
	router, err := api.Read(routerID)
	if err != nil {
		return nil, err
	}

	req := &sacloud.VPCRouter{Settings: &sacloud.VPCRouterSettings{}}

	if router.Settings == nil {
		req.Settings = &sacloud.VPCRouterSettings{
			// Router
			Router: &sacloud.VPCRouterSetting{
				// Interfaces
				Interfaces: []*sacloud.VPCRouterInterface{nil},
			},
		}
	} else {
		req.Settings.Router = router.Settings.Router
	}

	//disconnect to switch
	_, err = api.DisconnectFromSwitch(routerID, index)
	if err != nil {
		return nil, err
	}

	if index < len(req.Settings.Router.Interfaces) {
		req.Settings.Router.Interfaces[index] = nil
	}

	res, err := api.UpdateSetting(routerID, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

// MonitorBy 指定位置のインターフェースのアクティビティーモニター取得
func (api *VPCRouterAPI) MonitorBy(id sacloud.ID, nicIndex int, body *sacloud.ResourceMonitorRequest) (*sacloud.MonitorValues, error) {
	return api.baseAPI.applianceMonitorBy(id, "interface", nicIndex, body)
}

// Status ログなどのステータス情報 取得
func (api *VPCRouterAPI) Status(id sacloud.ID) (*sacloud.VPCRouterStatus, error) {
	var (
		method = "GET"
		uri    = fmt.Sprintf("%s/%d/status", api.getResourceURL(), id)
		res    = &vpcRouterStatusResponse{}
	)
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return nil, err
	}
	return res.VPCRouterStatus, nil
}

// SiteToSiteConnectionDetails サイト間VPN接続情報を取得
func (api *VPCRouterAPI) SiteToSiteConnectionDetails(id sacloud.ID) (*sacloud.SiteToSiteConnectionInfo, error) {
	var (
		method = "GET"
		uri    = fmt.Sprintf("%s/%d/vpcrouter/sitetosite/connectiondetails", api.getResourceURL(), id)
		res    = &vpcRouterS2sConnInfoResponse{}
	)
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return nil, err
	}
	return res.SiteToSiteConnectionInfo, nil
}
