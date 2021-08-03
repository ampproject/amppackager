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

// IPv6AddrAPI IPv6アドレスAPI
type IPv6AddrAPI struct {
	*baseAPI
}

// NewIPv6AddrAPI IPv6アドレスAPI新規作成
func NewIPv6AddrAPI(client *Client) *IPv6AddrAPI {
	return &IPv6AddrAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "ipv6addr"
			},
		},
	}
}

// Read 読み取り
func (api *IPv6AddrAPI) Read(ip string) (*sacloud.IPv6Addr, error) {
	return api.request(func(res *sacloud.Response) error {
		var (
			method = "GET"
			uri    = fmt.Sprintf("%s/%s", api.getResourceURL(), ip)
		)

		return api.baseAPI.request(method, uri, nil, res)
	})

}

// Create 新規作成
func (api *IPv6AddrAPI) Create(ip string, hostName string) (*sacloud.IPv6Addr, error) {

	type request struct {
		// IPv6Addr
		IPv6Addr map[string]string
	}

	var (
		method = "POST"
		uri    = api.getResourceURL()
		body   = &request{IPv6Addr: map[string]string{}}
	)
	body.IPv6Addr["IPv6Addr"] = ip
	body.IPv6Addr["HostName"] = hostName

	return api.request(func(res *sacloud.Response) error {
		return api.baseAPI.request(method, uri, body, res)
	})
}

// Update 更新
func (api *IPv6AddrAPI) Update(ip string, hostName string) (*sacloud.IPv6Addr, error) {

	type request struct {
		// IPv6Addr
		IPv6Addr map[string]string
	}

	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%s", api.getResourceURL(), ip)
		body   = &request{IPv6Addr: map[string]string{}}
	)
	body.IPv6Addr["HostName"] = hostName

	return api.request(func(res *sacloud.Response) error {
		return api.baseAPI.request(method, uri, body, res)
	})
}

// Delete 削除
func (api *IPv6AddrAPI) Delete(ip string) (*sacloud.IPv6Addr, error) {

	type request struct {
		// IPv6Addr
		IPv6Addr map[string]string
	}

	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%s", api.getResourceURL(), ip)
		body   = &request{IPv6Addr: map[string]string{}}
	)

	return api.request(func(res *sacloud.Response) error {
		return api.baseAPI.request(method, uri, body, res)
	})
}
