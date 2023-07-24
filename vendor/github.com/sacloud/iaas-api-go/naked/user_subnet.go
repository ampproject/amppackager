// Copyright 2022-2023 The sacloud/iaas-api-go Authors
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

package naked

import "encoding/json"

// UserSubnet ユーザーサブネット
type UserSubnet struct {
	DefaultRoute   string `yaml:"default_route"`
	NetworkMaskLen int    `yaml:"network_mask_len"`
}

// UnmarshalJSON DefaultRouteがからの場合に"0.0.0.0"となることへの対応
func (s *UserSubnet) UnmarshalJSON(data []byte) error {
	type alias UserSubnet
	var tmp alias
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp.DefaultRoute == "0.0.0.0" {
		tmp.DefaultRoute = ""
	}
	*s = UserSubnet(tmp)
	return nil
}
