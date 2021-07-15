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

// propAvailability 有効状態内包型
type propAvailability struct {
	Availability EAvailability `json:",omitempty"` // 有効状態
}

// IsAvailable 有効状態が"有効"か判定
func (p *propAvailability) IsAvailable() bool {
	return p.Availability.IsAvailable()
}

// IsUploading 有効状態が"アップロード中"か判定
func (p *propAvailability) IsUploading() bool {
	return p.Availability.IsUploading()
}

// IsFailed 有効状態が"失敗"か判定
func (p *propAvailability) IsFailed() bool {
	return p.Availability.IsFailed()
}

// IsMigrating 有効状態が"マイグレーション中"か判定
func (p *propAvailability) IsMigrating() bool {
	return p.Availability.IsMigrating()
}
