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

// RDBMSType データベースアプライアンスでのRDBMS種別
type RDBMSType string

// String RDBMSTypeの文字列表現
func (t RDBMSType) String() string {
	return string(t)
}

const (
	// RDBMSTypesMariaDB MariaDB
	RDBMSTypesMariaDB = RDBMSType("MariaDB")
	// RDBMSTypesPostgreSQL PostgreSQL
	RDBMSTypesPostgreSQL = RDBMSType("postgres")
)

// RDBMSTypeStrings 有効なRDBMS種別を示す文字列
var RDBMSTypeStrings = []string{
	strings.ToLower(RDBMSTypesMariaDB.String()),
	strings.ToLower(RDBMSTypesPostgreSQL.String()),
}

// RDBMSTypeFromString 文字列からRDBMSTypeを取得
func RDBMSTypeFromString(s string) RDBMSType {
	switch {
	case s == strings.ToLower(RDBMSTypesMariaDB.String()):
		return RDBMSTypesMariaDB
	case strings.ToLower(s) == "postgresql":
		return RDBMSTypesPostgreSQL
	default:
		return RDBMSType(s)
	}
}
