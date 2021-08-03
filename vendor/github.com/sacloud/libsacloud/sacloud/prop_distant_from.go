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

// propDistantFrom ストレージ隔離対象ディスク内包型
type propDistantFrom struct {
	DistantFrom []ID `json:",omitempty"` // ストレージ隔離対象ディスク
}

// GetDistantFrom ストレージ隔離対象ディスク 取得
func (p *propDistantFrom) GetDistantFrom() []ID {
	return p.DistantFrom
}

// SetDistantFrom ストレージ隔離対象ディスク 設定
func (p *propDistantFrom) SetDistantFrom(ids []ID) {
	p.DistantFrom = ids
}

// AddDistantFrom ストレージ隔離対象ディスク 追加
func (p *propDistantFrom) AddDistantFrom(id ID) {
	p.DistantFrom = append(p.DistantFrom, id)
}

// ClearDistantFrom ストレージ隔離対象ディスク クリア
func (p *propDistantFrom) ClearDistantFrom() {
	p.DistantFrom = []ID{}
}
