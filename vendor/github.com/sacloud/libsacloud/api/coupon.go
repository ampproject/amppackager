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

// CouponAPI クーポン情報API
type CouponAPI struct {
	*baseAPI
}

// NewCouponAPI クーポン情報API作成
func NewCouponAPI(client *Client) *CouponAPI {
	return &CouponAPI{
		&baseAPI{
			client:        client,
			apiRootSuffix: sakuraBillingAPIRootSuffix,
			FuncGetResourceURL: func() string {
				return "coupon"
			},
		},
	}
}

// CouponResponse クーポン情報レスポンス
type CouponResponse struct {
	*sacloud.ResultFlagValue
	// AllCount 件数
	AllCount int `json:",omitempty"`
	// CountPerPage ページあたり件数
	CountPerPage int `json:",omitempty"`
	// Page 現在のページ番号
	Page int `json:",omitempty"`
	// Coupons クーポン情報 リスト
	Coupons []*sacloud.Coupon
}

// Find クーポン情報 全件取得
func (api *CouponAPI) Find() ([]*sacloud.Coupon, error) {
	authStatus, err := api.client.AuthStatus.Read()
	if err != nil {
		return nil, err
	}
	accountID := authStatus.Account.GetStrID()

	uri := fmt.Sprintf("%s/%s", api.getResourceURL(), accountID)
	data, err := api.client.newRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	var res CouponResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res.Coupons, nil
}
