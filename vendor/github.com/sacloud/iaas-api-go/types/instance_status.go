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

// EServerInstanceStatus サーバーインスタンスステータス
type EServerInstanceStatus string

// ServerInstanceStatuses サーバーインスタンスステータス
var ServerInstanceStatuses = &struct {
	Unknown  EServerInstanceStatus
	Up       EServerInstanceStatus
	Cleaning EServerInstanceStatus
	Down     EServerInstanceStatus
}{
	Unknown:  EServerInstanceStatus(""),
	Up:       EServerInstanceStatus("up"),
	Cleaning: EServerInstanceStatus("cleaning"),
	Down:     EServerInstanceStatus("down"),
}

// IsUp インスタンスが起動しているか判定
func (e EServerInstanceStatus) IsUp() bool {
	return e == ServerInstanceStatuses.Up
}

// IsDown インスタンスがダウンしているか確認
func (e EServerInstanceStatus) IsDown() bool {
	return e == ServerInstanceStatuses.Down
}
