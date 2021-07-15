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

// PrivateHost 専有ホスト
type PrivateHost struct {
	*Resource       // ID
	propName        // 名称
	propDescription // 説明

	propPrivateHostPlan  // 専有ホストプラン
	propHost             // ホスト(物理)
	propAssignedCPU      // 割当済みCPUコア数
	propAssignedMemoryMB // 割当済みメモリ(MB)

	propIcon      // アイコン
	propTags      // タグ
	propCreatedAt // 作成日時

}

const (
	// PrivateHostClassDynamic 専有ホストプラン(標準)
	PrivateHostClassDynamic = "dynamic"
	// PrivateHostClassWindows 専有ホストプラン(windows)
	PrivateHostClassWindows = "ms_windows"
)
