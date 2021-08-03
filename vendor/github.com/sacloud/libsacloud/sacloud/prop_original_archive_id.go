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

// propOriginalArchiveID オリジナルアーカイブID内包型
type propOriginalArchiveID struct {
	OriginalArchive *Resource `json:",omitempty"` // オリジナルアーカイブ
}

// GetOriginalArchiveID プランID 取得
func (p *propOriginalArchiveID) GetOriginalArchiveID() ID {
	if p.OriginalArchive == nil {
		return EmptyID
	}
	return p.OriginalArchive.ID
}

// GetStrOriginalArchiveID プランID(文字列) 取得
func (p *propOriginalArchiveID) GetStrOriginalArchiveID() string {
	if p.OriginalArchive == nil {
		return ""
	}
	return p.OriginalArchive.ID.String()
}
