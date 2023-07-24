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

// EContainerRegistryPermission コンテナレジストリへアクセスレベル
type EContainerRegistryPermission string

// String EContainerRegistryVisibilityの文字列表現
func (v EContainerRegistryPermission) String() string {
	return string(v)
}

// ContainerRegistryPermissions コンテナレジストリのアクセス範囲
var ContainerRegistryPermissions = struct {
	All       EContainerRegistryPermission
	ReadWrite EContainerRegistryPermission
	ReadOnly  EContainerRegistryPermission
}{
	All:       "all",
	ReadWrite: "readwrite",
	ReadOnly:  "readonly",
}

// ContainerRegistryPermissionStrings アクセス範囲に指定可能な文字列
var ContainerRegistryPermissionStrings = []string{
	ContainerRegistryPermissions.All.String(),
	ContainerRegistryPermissions.ReadWrite.String(),
	ContainerRegistryPermissions.ReadOnly.String(),
}

// ContainerRegistryPermissionMap 文字列とEContainerRegistryPermissionのマップ
var ContainerRegistryPermissionMap = map[string]EContainerRegistryPermission{
	ContainerRegistryPermissions.All.String():       ContainerRegistryPermissions.All,
	ContainerRegistryPermissions.ReadWrite.String(): ContainerRegistryPermissions.ReadWrite,
	ContainerRegistryPermissions.ReadOnly.String():  ContainerRegistryPermissions.ReadOnly,
}
