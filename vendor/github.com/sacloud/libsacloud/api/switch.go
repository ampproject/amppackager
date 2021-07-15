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
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
)

// SwitchAPI スイッチAPI
type SwitchAPI struct {
	*baseAPI
}

// NewSwitchAPI スイッチAPI作成
func NewSwitchAPI(client *Client) *SwitchAPI {
	return &SwitchAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "switch"
			},
		},
	}
}

// DisconnectFromBridge ブリッジとの切断
func (api *SwitchAPI) DisconnectFromBridge(switchID sacloud.ID) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/to/bridge", api.getResourceURL(), switchID)
	)
	return api.modify(method, uri, nil)
}

// ConnectToBridge ブリッジとの接続
func (api *SwitchAPI) ConnectToBridge(switchID sacloud.ID, bridgeID sacloud.ID) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/to/bridge/%d", api.getResourceURL(), switchID, bridgeID)
	)
	return api.modify(method, uri, nil)
}

// GetServers スイッチに接続されているサーバー一覧取得
func (api *SwitchAPI) GetServers(switchID sacloud.ID) ([]sacloud.Server, error) {
	var (
		method = "GET"
		uri    = fmt.Sprintf("%s/%d/server", api.getResourceURL(), switchID)
		res    = &sacloud.SearchResponse{}
	)
	err := api.baseAPI.request(method, uri, nil, res)
	if err != nil {
		return nil, err
	}
	return res.Servers, nil
}
