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

// SpecialTag 特殊タグ
type SpecialTag string

// SpecialTags 特殊タグ一覧
var SpecialTags = struct {
	// GroupA サーバをグループ化し起動ホストを分離します(グループA)
	GroupA SpecialTag
	// GroupB サーバをグループ化し起動ホストを分離します(グループB)
	GroupB SpecialTag
	// GroupC サーバをグループ化し起動ホストを分離します(グループC)
	GroupC SpecialTag
	// GroupD サーバをグループ化し起動ホストを分離します(グループD)
	GroupD SpecialTag
	// AutoReboot サーバ停止時に自動起動します
	AutoReboot SpecialTag
	// KeyboardUS リモートスクリーン画面でUSキーボード入力します
	KeyboardUS SpecialTag
	// BootCDROM 優先ブートデバイスをCD-ROMに設定します
	BootCDROM SpecialTag
	// BootNetwork 優先ブートデバイスをPXE bootに設定します
	BootNetwork SpecialTag
	// CPUTopology CPUソケット数を1と認識させる
	CPUTopology SpecialTag
}{
	GroupA:      SpecialTag("@group=a"),
	GroupB:      SpecialTag("@group=b"),
	GroupC:      SpecialTag("@group=c"),
	GroupD:      SpecialTag("@group=d"),
	AutoReboot:  SpecialTag("@auto-reboot"),
	KeyboardUS:  SpecialTag("@keyboard-us"),
	BootCDROM:   SpecialTag("@boot-cdrom"),
	BootNetwork: SpecialTag("@boot-network"),
	CPUTopology: SpecialTag("@cpu-topology"),
}
