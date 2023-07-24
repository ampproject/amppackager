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

// EContainerRegistryAccessLevel コンテナレジストリへアクセスレベル
type EContainerRegistryAccessLevel string

// String EContainerRegistryVisibilityの文字列表現
func (v EContainerRegistryAccessLevel) String() string {
	return string(v)
}

// ContainerRegistryAccessLevels コンテナレジストリのアクセス範囲
var ContainerRegistryAccessLevels = struct {
	ReadWrite EContainerRegistryAccessLevel
	ReadOnly  EContainerRegistryAccessLevel
	None      EContainerRegistryAccessLevel
}{
	ReadWrite: "readwrite",
	ReadOnly:  "readonly",
	None:      "none",
}

// ContainerRegistryAccessLevelStrings アクセス範囲に指定可能な文字列
var ContainerRegistryAccessLevelStrings = []string{
	ContainerRegistryAccessLevels.ReadWrite.String(),
	ContainerRegistryAccessLevels.ReadOnly.String(),
	ContainerRegistryAccessLevels.None.String(),
}

// ContainerRegistryAccessLevelMap 文字列とEContainerRegistryVisibilityのマップ
var ContainerRegistryAccessLevelMap = map[string]EContainerRegistryAccessLevel{
	ContainerRegistryAccessLevels.ReadWrite.String(): ContainerRegistryAccessLevels.ReadWrite,
	ContainerRegistryAccessLevels.ReadOnly.String():  ContainerRegistryAccessLevels.ReadOnly,
	ContainerRegistryAccessLevels.None.String():      ContainerRegistryAccessLevels.None,
}
