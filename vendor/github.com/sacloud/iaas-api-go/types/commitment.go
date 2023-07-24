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

// ECommitment サーバプランCPUコミットメント
//
// 通常 or コア専有
type ECommitment string

// Commitments サーバプランCPUコミットメント
var Commitments = struct {
	// Unknown 不明
	Unknown ECommitment
	// Standard 通常
	Standard ECommitment
	// DedicatedCPU コア専有
	DedicatedCPU ECommitment
}{
	Unknown:      ECommitment(""),
	Standard:     ECommitment("standard"),
	DedicatedCPU: ECommitment("dedicatedcpu"),
}

// IsStandard Standardであるか判定
func (c ECommitment) IsStandard() bool {
	return c == Commitments.Standard
}

// IsDedicatedCPU DedicatedCPUであるか判定
func (c ECommitment) IsDedicatedCPU() bool {
	return c == Commitments.DedicatedCPU
}

// String ECommitmentの文字列表現
func (c ECommitment) String() string {
	return string(c)
}

// CommitmentStrings サーバプランCPUコミットメントを表す文字列
var CommitmentStrings = []string{"standard", "dedicatedcpu"}
