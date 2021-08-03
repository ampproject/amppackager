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

import "strings"

// BundleInfo バンドル情報
type BundleInfo struct {
	HostClass    string `json:",omitempty"`
	ServiceClass string `json:",omitempty"`
}

// propBundleInfo バンドル情報内包型
type propBundleInfo struct {
	BundleInfo *BundleInfo `json:",omitempty"` // バンドル情報
}

// GetBundleInfo バンドル情報 取得
func (p *propBundleInfo) GetBundleInfo() *BundleInfo {
	return p.BundleInfo
}

func (p *propBundleInfo) IsSophosUTM() bool {
	// SophosUTMであれば編集不可
	if p.BundleInfo != nil && strings.Contains(strings.ToLower(p.BundleInfo.ServiceClass), "sophosutm") {
		return true
	}
	return false
}
