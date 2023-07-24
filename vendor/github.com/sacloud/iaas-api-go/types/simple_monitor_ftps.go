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

// ESimpleMonitorFTPS シンプル監視FTPSパラメータ
type ESimpleMonitorFTPS string

// String ESimpleMonitorFTPSの文字列表現
func (p ESimpleMonitorFTPS) String() string {
	return string(p)
}

// SimpleMonitorFTPSValues ESimpleMonitorFTPSがとりうる値
var SimpleMonitorFTPSValues = struct {
	Default  ESimpleMonitorFTPS
	Implicit ESimpleMonitorFTPS
	Explicit ESimpleMonitorFTPS
}{
	Default:  ESimpleMonitorFTPS(""),
	Implicit: ESimpleMonitorFTPS("implicit"),
	Explicit: ESimpleMonitorFTPS("explicit"),
}

// SimpleMonitorFTPSStrings x
var SimpleMonitorFTPSStrings = []string{
	SimpleMonitorFTPSValues.Default.String(),
	SimpleMonitorFTPSValues.Implicit.String(),
	SimpleMonitorFTPSValues.Explicit.String(),
}
