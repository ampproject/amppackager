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
	"time"

	"github.com/sacloud/libsacloud/sacloud"
)

// CDROMAPI ISOイメージAPI
type CDROMAPI struct {
	*baseAPI
}

// NewCDROMAPI ISOイメージAPI新規作成
func NewCDROMAPI(client *Client) *CDROMAPI {
	return &CDROMAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "cdrom"
			},
		},
	}
}

// Create 新規作成
func (api *CDROMAPI) Create(value *sacloud.CDROM) (*sacloud.CDROM, *sacloud.FTPServer, error) {
	f := func(res *sacloud.Response) error {
		return api.create(api.createRequest(value), res)
	}
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, nil, err
	}
	return res.CDROM, res.FTPServer, nil
}

// OpenFTP FTP接続開始
func (api *CDROMAPI) OpenFTP(id sacloud.ID, reset bool) (*sacloud.FTPServer, error) {
	var (
		method = "PUT"
		uri    = fmt.Sprintf("%s/%d/ftp", api.getResourceURL(), id)
		body   = map[string]bool{"ChangePassword": reset}
		res    = &sacloud.Response{}
	)

	result, err := api.action(method, uri, body, res)
	if !result || err != nil {
		return nil, err
	}

	return res.FTPServer, nil
}

// CloseFTP FTP接続終了
func (api *CDROMAPI) CloseFTP(id sacloud.ID) (bool, error) {
	var (
		method = "DELETE"
		uri    = fmt.Sprintf("%s/%d/ftp", api.getResourceURL(), id)
	)
	return api.modify(method, uri, nil)

}

// SleepWhileCopying コピー終了まで待機
func (api *CDROMAPI) SleepWhileCopying(id sacloud.ID, timeout time.Duration) error {
	handler := waitingForAvailableFunc(func() (hasAvailable, error) {
		return api.Read(id)
	}, 0)
	return blockingPoll(handler, timeout)
}

// AsyncSleepWhileCopying コピー終了まで待機(非同期)
func (api *CDROMAPI) AsyncSleepWhileCopying(id sacloud.ID, timeout time.Duration) (chan (interface{}), chan (interface{}), chan (error)) {
	handler := waitingForAvailableFunc(func() (hasAvailable, error) {
		return api.Read(id)
	}, 0)
	return poll(handler, timeout)
}
