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

// propSizeMB サイズ(MB)内包型
type propSizeMB struct {
	SizeMB int `json:",omitempty"` // サイズ(MB単位)
}

// GetSizeMB サイズ(MB単位) 取得
func (p *propSizeMB) GetSizeMB() int {
	return p.SizeMB
}

// SetSizeMB サイズ(MB単位) 設定
func (p *propSizeMB) SetSizeMB(size int) {
	p.SizeMB = size
}

// GetSizeGB サイズ(GB単位) 取得
func (p *propSizeMB) GetSizeGB() int {
	if p.SizeMB <= 0 {
		return 0
	}
	return p.SizeMB / 1024
}

// SetSizeGB サイズ(GB単位) 設定
func (p *propSizeMB) SetSizeGB(size int) {
	if size <= 0 {
		p.SizeMB = 0
	} else {
		p.SizeMB = size * 1024
	}
}

// propMigratedMB コピー済みデータサイズ(MB単位)内包型
type propMigratedMB struct {
	MigratedMB int `json:",omitempty"` // コピー済みデータサイズ(MB単位)
}

// GetMigratedMB サイズ(MB単位) 取得
func (p *propMigratedMB) GetMigratedMB() int {
	return p.MigratedMB
}

// SetMigratedMB サイズ(MB単位) 設定
func (p *propMigratedMB) SetMigratedMB(size int) {
	p.MigratedMB = size
}

// GetMigratedGB サイズ(GB単位) 取得
func (p *propMigratedMB) GetMigratedGB() int {
	if p.MigratedMB <= 0 {
		return 0
	}
	return p.MigratedMB / 1024
}

// SetMigratedGB サイズ(GB単位) 設定
func (p *propMigratedMB) SetMigratedGB(size int) {
	if size <= 0 {
		p.MigratedMB = 0
	} else {
		p.MigratedMB = size * 1024
	}
}
