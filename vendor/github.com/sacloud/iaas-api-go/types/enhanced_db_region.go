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

// EnhancedDBRegion エンハンスドDBでのリージョン
type EnhancedDBRegion string

// String RDBMSRegionの文字列表現
func (t EnhancedDBRegion) String() string {
	return string(t)
}

const (
	// EnhancedDBRegionsTiDB TiDB
	EnhancedDBRegionsIs1 = EnhancedDBRegion("is1")
	// EnhancedDBRegionsMariaDB MariaDB
	EnhancedDBRegionsTk1 = EnhancedDBRegion("tk1")
)

// EnhancedDBRegionStrings 有効なリージョンを示す文字列
var EnhancedDBRegionStrings = []string{
	strings.ToLower(EnhancedDBRegionsIs1.String()),
	strings.ToLower(EnhancedDBRegionsTk1.String()),
}

// EnhancedDBRegionFromString 文字列からEnhancedDBRegionを取得
func EnhancedDBRegionFromString(s string) EnhancedDBRegion {
	switch {
	case strings.ToLower(s) == strings.ToLower(EnhancedDBRegionsIs1.String()):
		return EnhancedDBRegionsIs1
	case strings.ToLower(s) == strings.ToLower(EnhancedDBRegionsTk1.String()):
		return EnhancedDBRegionsTk1
	default:
		return EnhancedDBRegion(s)
	}
}
