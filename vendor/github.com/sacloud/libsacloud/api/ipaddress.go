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

// IPAddressAPI IPアドレスAPI
type IPAddressAPI struct {
	*baseAPI
}

// NewIPAddressAPI IPアドレスAPI新規作成
func NewIPAddressAPI(client *Client) *IPAddressAPI {
	return &IPAddressAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "ipaddress"
			},
		},
	}
}

// Read 読み取り
func (api *IPAddressAPI) Read(ip string) (*sacloud.IPAddress, error) {
	return api.request(func(res *sacloud.Response) error {
		var (
			method = "GET"
			uri    = fmt.Sprintf("%s/%s", api.getResourceURL(), ip)
		)

		return api.baseAPI.request(method, uri, nil, res)
	})

}

// Update 更新(ホスト名逆引き設定)
func (api *IPAddressAPI) Update(ip string, hostName string) (*sacloud.IPAddress, error) {

	type request struct {
		// IPAddress
		IPAddress map[string]string
	}

	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%s", api.getResourceURL(), ip)
		body   = &request{IPAddress: map[string]string{}}
	)
	body.IPAddress["HostName"] = hostName

	return api.request(func(res *sacloud.Response) error {
		return api.baseAPI.request(method, uri, body, res)
	})
}
