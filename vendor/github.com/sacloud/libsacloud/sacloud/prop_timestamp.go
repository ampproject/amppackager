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

import "time"

// propCreatedAt 作成日時内包型
type propCreatedAt struct {
	CreatedAt *time.Time `json:",omitempty"` // 作成日時
}

// GetCreatedAt 作成日時 取得
func (p *propCreatedAt) GetCreatedAt() *time.Time {
	return p.CreatedAt
}

// propModifiedAt 変更日時内包型
type propModifiedAt struct {
	// ModifiedAt 変更日時
	ModifiedAt *time.Time `json:",omitempty"`
}

// GetModifiedAt 変更日時 取得
func (p *propModifiedAt) GetModifiedAt() *time.Time {
	return p.ModifiedAt
}

// propUpdatedAt 変更日時内包型
type propUpdatedAt struct {
	// UpdatedAt 変更日時
	UpdatedAt *time.Time `json:",omitempty"`
}

// GetModifiedAt 変更日時 取得
func (p *propUpdatedAt) GetModifiedAt() *time.Time {
	return p.UpdatedAt
}
