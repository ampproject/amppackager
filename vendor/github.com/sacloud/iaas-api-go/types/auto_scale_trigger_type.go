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

type EAutoScaleTriggerType string

// AutoScaleTriggerTypes サーバプランCPUコミットメント
var AutoScaleTriggerTypes = struct {
	// Default CPUトリガーがデフォルト
	Default EAutoScaleTriggerType
	// Standard 通常
	CPU EAutoScaleTriggerType
	// DedicatedCPU コア専有
	Router EAutoScaleTriggerType
	// DedicatedCPU コア専有
	Schedule EAutoScaleTriggerType
}{
	Default:  EAutoScaleTriggerType(""),
	CPU:      EAutoScaleTriggerType("cpu"),
	Router:   EAutoScaleTriggerType("router"),
	Schedule: EAutoScaleTriggerType("schedule"),
}

// String EAutoScaleTriggerTypeの文字列表現
func (c EAutoScaleTriggerType) String() string {
	return string(c)
}

// AutoScaleTriggerTypeStrings サーバプランCPUコミットメントを表す文字列
var AutoScaleTriggerTypeStrings = []string{"cpu", "router", "schedule"}
