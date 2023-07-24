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

import "strings"

// EnhancedDBType エンハンスドデータベースでの種別
type EnhancedDBType string

// String EnhancedDBTypeの文字列表現
func (t EnhancedDBType) String() string {
	return string(t)
}

const (
	// EnhancedDBTypesTiDB TiDB
	EnhancedDBTypesTiDB = EnhancedDBType("tidb")
	// EnhancedDBTypesMariaDB MariaDB
	EnhancedDBTypesMariaDB = EnhancedDBType("mariadb")
)

// EnhancedDBTypeStrings 有効な種別を示す文字列
var EnhancedDBTypeStrings = []string{
	strings.ToLower(EnhancedDBTypesTiDB.String()),
	strings.ToLower(EnhancedDBTypesMariaDB.String()),
}

// EnhancedDBTypeFromString 文字列からEnhancedDBTypeを取得
func EnhancedDBTypeFromString(s string) EnhancedDBType {
	switch {
	case strings.ToLower(s) == strings.ToLower(EnhancedDBTypesTiDB.String()):
		return EnhancedDBTypesTiDB
	case strings.ToLower(s) == strings.ToLower(EnhancedDBTypesMariaDB.String()):
		return EnhancedDBTypesMariaDB
	default:
		return EnhancedDBType(s)
	}
}
