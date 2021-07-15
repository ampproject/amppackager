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

// ProductServer サーバープラン
type ProductServer struct {
	*Resource                        // ID
	propName                         // 名称
	propDescription                  // 説明
	propAvailability                 // 有功状態
	propCPU                          // CPUコア数
	propMemoryMB                     // メモリサイズ(MB単位)
	propServiceClass                 // サービスクラス
	Generation       PlanGenerations `json:",omitempty"` // 世代
	Commitment       ECommitment     `json:",omitempty"`
}
