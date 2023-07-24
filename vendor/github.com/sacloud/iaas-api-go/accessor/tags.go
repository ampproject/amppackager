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

package accessor

import "github.com/sacloud/iaas-api-go/types"

// Tags Tagsを持つリソース向けのインターフェース
type Tags interface {
	GetTags() types.Tags
	SetTags(v types.Tags)
}

// HasTag 指定のタグが存在する場合trueを返す
func HasTag(target Tags, tag string) bool {
	tags := target.GetTags()
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

// AppendTag 指定のタグを追加
func AppendTag(target Tags, tag string) {
	if HasTag(target, tag) {
		return
	}
	tags := target.GetTags()
	target.SetTags(append(tags, tag))
}

// RemoveTag 指定のタグを削除
func RemoveTag(target Tags, tag string) {
	if !HasTag(target, tag) {
		return
	}

	tags := target.GetTags()
	nt := types.Tags{}
	for _, t := range tags {
		if t != tag {
			nt = append(nt, t)
		}
	}
	target.SetTags(nt)
}

// ClearTags 全タグをクリア
func ClearTags(target Tags) {
	target.SetTags(types.Tags{})
}
