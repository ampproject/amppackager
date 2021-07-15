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

// SSHKeyAPI 公開鍵API
type SSHKeyAPI struct {
	*baseAPI
}

// NewSSHKeyAPI 公開鍵API作成
func NewSSHKeyAPI(client *Client) *SSHKeyAPI {
	return &SSHKeyAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "sshkey"
			},
		},
	}
}

// Generate 公開鍵の作成
func (api *SSHKeyAPI) Generate(name string, passPhrase string, desc string) (*sacloud.SSHKeyGenerated, error) {

	var (
		method = "POST"
		uri    = fmt.Sprintf("%s/generate", api.getResourceURL())
	)

	type genRequest struct {
		Name           string
		GenerateFormat string
		Description    string
		PassPhrase     string
	}

	type request struct {
		SSHKey genRequest
	}
	type response struct {
		*sacloud.ResultFlagValue
		SSHKey *sacloud.SSHKeyGenerated
	}

	body := &request{
		SSHKey: genRequest{
			Name:           name,
			GenerateFormat: "openssh",
			PassPhrase:     passPhrase,
			Description:    desc,
		},
	}

	res := &response{}

	_, err := api.action(method, uri, body, res)
	if err != nil {
		return nil, fmt.Errorf("SSHKeyAPI: generate SSHKey is failed: %s", err)
	}
	return res.SSHKey, nil
}
