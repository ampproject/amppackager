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

package types

// EUpstreamNetworkType 上流ネットワーク種別
type EUpstreamNetworkType string

// String EUpstreamNetworkTypeの文字列表現
func (t EUpstreamNetworkType) String() string {
	return string(t)
}

var (
	// UpstreamNetworkTypes 上流ネットワーク種別
	UpstreamNetworkTypes = struct {
		Unknown EUpstreamNetworkType
		// Shared 共有セグメント
		Shared EUpstreamNetworkType
		// Switch スイッチ
		Switch EUpstreamNetworkType
		// Router ルータ
		Router EUpstreamNetworkType
		// None 接続なし
		None EUpstreamNetworkType
	}{
		Unknown: EUpstreamNetworkType("unknown"),
		Shared:  EUpstreamNetworkType("shared"),
		Switch:  EUpstreamNetworkType("switch"),
		Router:  EUpstreamNetworkType("router"),
		None:    EUpstreamNetworkType("none"),
	}

	// UpstreamNetworkTypeMap 文字列とEUpstreamNetworkTypeのマッピング
	UpstreamNetworkTypeMap = map[string]EUpstreamNetworkType{
		"unknown": UpstreamNetworkTypes.Unknown,
		"shared":  UpstreamNetworkTypes.Shared,
		"switch":  UpstreamNetworkTypes.Switch,
		"router":  UpstreamNetworkTypes.Router,
		"none":    UpstreamNetworkTypes.None,
	}
)
