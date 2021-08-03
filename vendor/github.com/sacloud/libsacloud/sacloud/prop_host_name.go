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

package sacloud

// propHostName ホスト名内包型
type propHostName struct {
	HostName string `json:",omitempty"` // ホスト名 (ディスクの修正実施時に指定した初期ホスト名)
}

// GetHostName (初期)ホスト名 取得
func (p *propHostName) GetHostName() string {
	return p.HostName
}
