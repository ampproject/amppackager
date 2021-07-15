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
	"net/url"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
)

// Reset 検索条件のリセット
func (api *WebAccelAPI) Reset() *WebAccelAPI {
	api.SetEmpty()
	return api
}

// SetEmpty 検索条件のリセット
func (api *WebAccelAPI) SetEmpty() {
	api.reset()
}

// FilterBy 指定キーでのフィルター
func (api *WebAccelAPI) FilterBy(key string, value interface{}) *WebAccelAPI {
	api.filterBy(key, value, false)
	return api
}

// WithNameLike 名称条件
func (api *WebAccelAPI) WithNameLike(name string) *WebAccelAPI {
	return api.FilterBy("Name", name)
}

// SetFilterBy 指定キーでのフィルター
func (api *WebAccelAPI) SetFilterBy(key string, value interface{}) {
	api.filterBy(key, value, false)
}

// SetNameLike 名称条件
func (api *WebAccelAPI) SetNameLike(name string) {
	api.FilterBy("Name", name)
}

// Find サイト一覧取得
func (api *WebAccelAPI) Find() (*sacloud.SearchResponse, error) {

	uri := fmt.Sprintf("%s/site", api.getResourceURL())

	data, err := api.client.newRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	var res sacloud.SearchResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	// handle filter(API側がFilterに対応していないためここでフィルタリング)
	for key, filter := range api.getSearchState().Filter {
		if key != "Name" {
			continue
		}
		strNames, ok := filter.(string)
		if !ok {
			continue
		}

		names := strings.Split(strNames, " ")
		filtered := []sacloud.WebAccelSite{}
		for _, site := range res.WebAccelSites {
			for _, name := range names {

				u, _ := url.Parse(name)

				if strings.Contains(site.Name, u.Path) {
					filtered = append(filtered, site)
					break
				}
			}
		}
		res.WebAccelSites = filtered
	}

	return &res, nil
}
