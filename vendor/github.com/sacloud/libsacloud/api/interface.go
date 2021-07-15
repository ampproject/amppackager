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

// InterfaceAPI インターフェースAPI
type InterfaceAPI struct {
	*baseAPI
}

// NewInterfaceAPI インターフェースAPI作成
func NewInterfaceAPI(client *Client) *InterfaceAPI {
	return &InterfaceAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "interface"
			},
		},
	}
}

// CreateAndConnectToServer 新規作成しサーバーへ接続する
func (api *InterfaceAPI) CreateAndConnectToServer(serverID sacloud.ID) (*sacloud.Interface, error) {
	iface := api.New()
	iface.Server = &sacloud.Server{
		// Resource
		Resource: &sacloud.Resource{ID: serverID},
	}
	return api.Create(iface)
}

// ConnectToSwitch スイッチへ接続する
func (api *InterfaceAPI) ConnectToSwitch(interfaceID sacloud.ID, switchID sacloud.ID) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/to/switch/%d", api.getResourceURL(), interfaceID, switchID)
	)
	return api.modify(method, uri, nil)
}

// ConnectToSharedSegment 共有セグメントへ接続する
func (api *InterfaceAPI) ConnectToSharedSegment(interfaceID sacloud.ID) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/to/switch/shared", api.getResourceURL(), interfaceID)
	)
	return api.modify(method, uri, nil)
}

// DisconnectFromSwitch スイッチと切断する
func (api *InterfaceAPI) DisconnectFromSwitch(interfaceID sacloud.ID) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/to/switch", api.getResourceURL(), interfaceID)
	)
	return api.modify(method, uri, nil)
}

// Monitor アクティビティーモニター取得
func (api *InterfaceAPI) Monitor(id sacloud.ID, body *sacloud.ResourceMonitorRequest) (*sacloud.MonitorValues, error) {
	return api.baseAPI.monitor(id, body)
}

// ConnectToPacketFilter パケットフィルター適用
func (api *InterfaceAPI) ConnectToPacketFilter(interfaceID sacloud.ID, packetFilterID sacloud.ID) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("/%s/%d/to/packetfilter/%d", api.getResourceURL(), interfaceID, packetFilterID)
	)
	return api.modify(method, uri, nil)
}

// DisconnectFromPacketFilter パケットフィルター切断
func (api *InterfaceAPI) DisconnectFromPacketFilter(interfaceID sacloud.ID) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("/%s/%d/to/packetfilter", api.getResourceURL(), interfaceID)
	)
	return api.modify(method, uri, nil)
}

// SetDisplayIPAddress 表示用IPアドレス 設定
func (api *InterfaceAPI) SetDisplayIPAddress(interfaceID sacloud.ID, ipaddress string) (bool, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("/%s/%d", api.getResourceURL(), interfaceID)
	)
	body := map[string]interface{}{
		"Interface": map[string]string{
			"UserIPAddress": ipaddress,
		},
	}
	return api.modify(method, uri, body)
}

// DeleteDisplayIPAddress 表示用IPアドレス 削除
func (api *InterfaceAPI) DeleteDisplayIPAddress(interfaceID sacloud.ID) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("/%s/%d", api.getResourceURL(), interfaceID)
	)
	return api.modify(method, uri, nil)
}
