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

// propTags タグ内包型
type propTags struct {
	Tags []string // タグ
}

// HasTag 指定のタグを持っているか判定
func (p *propTags) HasTag(target string) bool {

	for _, tag := range p.Tags {
		if target == tag {
			return true
		}
	}

	return false
}

// AppendTag タグを追加
func (p *propTags) AppendTag(target string) {
	if p.HasTag(target) {
		return
	}

	p.Tags = append(p.Tags, target)
}

// RemoveTag 指定のタグを削除
func (p *propTags) RemoveTag(target string) {
	if !p.HasTag(target) {
		return
	}
	res := []string{}
	for _, tag := range p.Tags {
		if tag != target {
			res = append(res, tag)
		}
	}

	p.Tags = res
}

// ClearTags 全タグを削除
func (p *propTags) ClearTags() {
	p.Tags = []string{}
}

// GetTags タグ取得
func (p *propTags) GetTags() []string {
	return p.Tags
}

// SetTags タグを設定
func (p *propTags) SetTags(tags []string) {
	p.Tags = tags
}
