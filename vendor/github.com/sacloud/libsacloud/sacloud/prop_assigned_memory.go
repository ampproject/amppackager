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

// propAssignedMemoryMB サイズ(MB)内包型
type propAssignedMemoryMB struct {
	AssignedMemoryMB int `json:",omitempty"` // サイズ(MB単位)
}

// GetAssignedMemoryMB サイズ(MB単位) 取得
func (p *propAssignedMemoryMB) GetAssignedMemoryMB() int {
	return p.AssignedMemoryMB
}

// GetAssignedMemoryGB サイズ(GB単位) 取得
func (p *propAssignedMemoryMB) GetAssignedMemoryGB() int {
	if p.AssignedMemoryMB <= 0 {
		return 0
	}
	return p.AssignedMemoryMB / 1024
}
