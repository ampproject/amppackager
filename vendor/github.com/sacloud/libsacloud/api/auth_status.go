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

	"github.com/sacloud/libsacloud/sacloud"
)

// AuthStatusAPI 認証状態API
type AuthStatusAPI struct {
	*baseAPI
}

// NewAuthStatusAPI 認証状態API作成
func NewAuthStatusAPI(client *Client) *AuthStatusAPI {
	return &AuthStatusAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "auth-status"
			},
		},
	}
}

// Read 読み取り
func (api *AuthStatusAPI) Read() (*sacloud.AuthStatus, error) {

	data, err := api.client.newRequest("GET", api.getResourceURL(), nil)
	if err != nil {
		return nil, err
	}
	var res sacloud.AuthStatus
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Find 検索
func (api *AuthStatusAPI) Find() (*sacloud.AuthStatus, error) {
	return api.Read()
}
