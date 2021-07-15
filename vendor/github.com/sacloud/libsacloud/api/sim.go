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

	"github.com/sacloud/libsacloud/sacloud"
)

// SearchSIMResponse SIM検索レスポンス
type SearchSIMResponse struct {
	// Total 総件数
	Total int `json:",omitempty"`
	// From ページング開始位置
	From int `json:",omitempty"`
	// Count 件数
	Count int `json:",omitempty"`
	// CommonServiceSIMItems SIMリスト
	CommonServiceSIMItems []sacloud.SIM `json:"CommonServiceItems,omitempty"`
}

type simRequest struct {
	CommonServiceSIMItem *sacloud.SIM           `json:"CommonServiceItem,omitempty"`
	From                 int                    `json:",omitempty"`
	Count                int                    `json:",omitempty"`
	Sort                 []string               `json:",omitempty"`
	Filter               map[string]interface{} `json:",omitempty"`
	Exclude              []string               `json:",omitempty"`
	Include              []string               `json:",omitempty"`
}

type simResponse struct {
	*sacloud.ResultFlagValue
	*sacloud.SIM `json:"CommonServiceItem,omitempty"`
}

type simLogResponse struct {
	Logs []sacloud.SIMLog `json:"logs,omitempty"`
	IsOk bool             `json:"is_ok,omitempty"`
}

// SIMAPI SIM API
type SIMAPI struct {
	*baseAPI
}

// NewSIMAPI SIM API作成
func NewSIMAPI(client *Client) *SIMAPI {
	return &SIMAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "commonserviceitem"
			},
			FuncBaseSearchCondition: func() *sacloud.Request {
				res := &sacloud.Request{}
				res.AddFilter("Provider.Class", "sim")
				return res
			},
		},
	}
}

// Find 検索
func (api *SIMAPI) Find() (*SearchSIMResponse, error) {

	data, err := api.client.newRequest("GET", api.getResourceURL(), api.getSearchState())
	if err != nil {
		return nil, err
	}
	var res SearchSIMResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (api *SIMAPI) request(f func(*simResponse) error) (*sacloud.SIM, error) {
	res := &simResponse{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.SIM, nil
}

func (api *SIMAPI) createRequest(value *sacloud.SIM) *simRequest {
	req := &simRequest{}
	req.CommonServiceSIMItem = value
	return req
}

// Create 新規作成
func (api *SIMAPI) Create(value *sacloud.SIM) (*sacloud.SIM, error) {
	return api.request(func(res *simResponse) error {
		return api.create(api.createRequest(value), res)
	})
}

// New 新規作成用パラメーター作成
func (api *SIMAPI) New(name, iccID, passcode string) *sacloud.SIM {
	return sacloud.CreateNewSIM(name, iccID, passcode)
}

// Read 読み取り
func (api *SIMAPI) Read(id sacloud.ID) (*sacloud.SIM, error) {
	return api.request(func(res *simResponse) error {
		return api.read(id, nil, res)
	})
}

// Update 更新
func (api *SIMAPI) Update(id sacloud.ID, value *sacloud.SIM) (*sacloud.SIM, error) {
	return api.request(func(res *simResponse) error {
		return api.update(id, api.createRequest(value), res)
	})
}

// Delete 削除
func (api *SIMAPI) Delete(id sacloud.ID) (*sacloud.SIM, error) {
	return api.request(func(res *simResponse) error {
		return api.delete(id, nil, res)
	})
}

// Activate SIM有効化
func (api *SIMAPI) Activate(id sacloud.ID) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/sim/activate", api.getResourceURL(), id)
	)

	return api.modify(method, uri, nil)
}

// Deactivate SIM無効化
func (api *SIMAPI) Deactivate(id sacloud.ID) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/sim/deactivate", api.getResourceURL(), id)
	)

	return api.modify(method, uri, nil)
}

// AssignIP SIMへのIP割り当て
func (api *SIMAPI) AssignIP(id sacloud.ID, ip string) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/sim/ip", api.getResourceURL(), id)
	)

	return api.modify(method, uri, map[string]interface{}{
		"sim": map[string]interface{}{
			"ip": ip,
		},
	})
}

// ClearIP SIMからのIP割り当て解除
func (api *SIMAPI) ClearIP(id sacloud.ID) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/sim/ip", api.getResourceURL(), id)
	)
	return api.modify(method, uri, nil)
}

// IMEILock IMEIロック
func (api *SIMAPI) IMEILock(id sacloud.ID, imei string) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/sim/imeilock", api.getResourceURL(), id)
	)

	return api.modify(method, uri, map[string]interface{}{
		"sim": map[string]interface{}{
			"imei": imei,
		},
	})
}

// IMEIUnlock IMEIアンロック
func (api *SIMAPI) IMEIUnlock(id sacloud.ID) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/sim/imeilock", api.getResourceURL(), id)
	)
	return api.modify(method, uri, nil)
}

// Logs セッションログ取得
func (api *SIMAPI) Logs(id sacloud.ID, body interface{}) ([]sacloud.SIMLog, error) {
	var (
		method = "GET"
		uri    = fmt.Sprintf("%s/%d/sim/sessionlog", api.getResourceURL(), id)
	)

	res := &simLogResponse{}
	err := api.baseAPI.request(method, uri, body, res)
	if err != nil {
		return nil, err
	}
	return res.Logs, nil
}

// GetNetworkOperator 通信キャリア 取得
func (api *SIMAPI) GetNetworkOperator(id sacloud.ID) (*sacloud.SIMNetworkOperatorConfigs, error) {

	var (
		method = "GET"
		uri    = fmt.Sprintf("%s/%d/sim/network_operator_config", api.getResourceURL(), id)
	)

	res := &sacloud.SIMNetworkOperatorConfigs{}
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SetNetworkOperator 通信キャリア 設定
func (api *SIMAPI) SetNetworkOperator(id sacloud.ID, opConfig ...*sacloud.SIMNetworkOperatorConfig) (bool, error) {

	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/sim/network_operator_config", api.getResourceURL(), id)
	)

	err := api.baseAPI.request(method, uri, &sacloud.SIMNetworkOperatorConfigs{NetworkOperatorConfigs: opConfig}, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Monitor アクティビティーモニター(Up/Down link BPS)取得
func (api *SIMAPI) Monitor(id sacloud.ID, body *sacloud.ResourceMonitorRequest) (*sacloud.MonitorValues, error) {
	var (
		method = "GET"
		uri    = fmt.Sprintf("%s/%d/sim/metrics", api.getResourceURL(), id)
	)
	res := &sacloud.ResourceMonitorResponse{}
	err := api.baseAPI.request(method, uri, body, res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}
