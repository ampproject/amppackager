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
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
)

// WebAccelAPI ウェブアクセラレータAPI
type WebAccelAPI struct {
	*baseAPI
}

// NewWebAccelAPI ウェブアクセラレータAPI
func NewWebAccelAPI(client *Client) *WebAccelAPI {
	return &WebAccelAPI{
		&baseAPI{
			client:        client,
			apiRootSuffix: sakuraWebAccelAPIRootSuffix,
			FuncGetResourceURL: func() string {
				return ""
			},
		},
	}
}

// WebAccelDeleteCacheResponse ウェブアクセラレータ キャッシュ削除レスポンス
type WebAccelDeleteCacheResponse struct {
	*sacloud.ResultFlagValue
	Results []*sacloud.DeleteCacheResult
}

// Read サイト情報取得
func (api *WebAccelAPI) Read(id string) (*sacloud.WebAccelSite, error) {

	uri := fmt.Sprintf("%s/site/%s", api.getResourceURL(), id)

	data, err := api.client.newRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		Site sacloud.WebAccelSite
	}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res.Site, nil
}

// ReadCertificate 証明書 参照
func (api *WebAccelAPI) ReadCertificate(id sacloud.ID) (*sacloud.WebAccelCertResponseBody, error) {
	uri := fmt.Sprintf("%s/site/%s/certificate", api.getResourceURL(), id)

	data, err := api.client.newRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	var res sacloud.WebAccelCertResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res.Certificate, nil
}

// CreateCertificate 証明書 更新
func (api *WebAccelAPI) CreateCertificate(id sacloud.ID, request *sacloud.WebAccelCertRequest) (*sacloud.WebAccelCertResponse, error) {
	uri := fmt.Sprintf("%s/site/%s/certificate", api.getResourceURL(), id)

	if request.CertificateChain != "" {
		request.CertificateChain = strings.TrimRight(request.CertificateChain, "\n")
	}
	if request.Key != "" {
		request.Key = strings.TrimRight(request.Key, "\n")
	}

	data, err := api.client.newRequest("POST", uri, map[string]interface{}{
		"Certificate": request,
	})
	if err != nil {
		return nil, err
	}

	var res sacloud.WebAccelCertResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateCertificate 証明書 更新
func (api *WebAccelAPI) UpdateCertificate(id sacloud.ID, request *sacloud.WebAccelCertRequest) (*sacloud.WebAccelCertResponse, error) {
	uri := fmt.Sprintf("%s/site/%s/certificate", api.getResourceURL(), id)

	if request.CertificateChain != "" {
		request.CertificateChain = strings.TrimRight(request.CertificateChain, "\n")
	}
	if request.Key != "" {
		request.Key = strings.TrimRight(request.Key, "\n")
	}

	data, err := api.client.newRequest("PUT", uri, map[string]interface{}{
		"Certificate": request,
	})
	if err != nil {
		return nil, err
	}

	var res sacloud.WebAccelCertResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteCertificate 証明書 削除
func (api *WebAccelAPI) DeleteCertificate(id string) (*sacloud.WebAccelCertResponse, error) {
	uri := fmt.Sprintf("%s/site/%s/certificate", api.getResourceURL(), id)

	data, err := api.client.newRequest("DELETE", uri, nil)
	if err != nil {
		return nil, err
	}

	var res sacloud.WebAccelCertResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteCache キャッシュ削除
func (api *WebAccelAPI) DeleteCache(urls ...string) (*WebAccelDeleteCacheResponse, error) {

	type request struct {
		// URL
		URL []string
	}

	uri := fmt.Sprintf("%s/deletecache", api.getResourceURL())

	data, err := api.client.newRequest("POST", uri, &request{URL: urls})
	if err != nil {
		return nil, err
	}

	var res WebAccelDeleteCacheResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
