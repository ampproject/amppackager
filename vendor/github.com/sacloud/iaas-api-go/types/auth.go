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

// EAuthClass 認証クラス型
type EAuthClass string

// AuthClasses 認証クラス
var AuthClasses = struct {
	// Unknown 不明
	Unknown EAuthClass
	// Account アカウント認証
	Account EAuthClass
}{
	Unknown: EAuthClass(""),
	Account: EAuthClass("account"),
}

// EAuthMethod 認証メソッド型
type EAuthMethod string

// AuthMethods 認証メソッド
var AuthMethods = struct {
	// Unknown 不明
	Unknown EAuthMethod
	// APIKey APIキーによる認証
	APIKey EAuthMethod
}{
	Unknown: EAuthMethod(""),
	APIKey:  EAuthMethod("apikey"),
}

// EOperationPenalty ペナルティ型
type EOperationPenalty string

// OperationPenalties ペナルティ
var OperationPenalties = struct {
	// Unknown 不明
	Unknown EOperationPenalty
	// None ペナルティなし
	None EOperationPenalty
}{
	Unknown: EOperationPenalty(""),
	None:    EOperationPenalty("none"),
}

// EPermission パーミッション型
type EPermission string

// Permissions パーミッション
var Permissions = struct {
	// Unknown 不明
	Unknown EPermission
	// Create 作成・削除権限
	Create EPermission
	// Arrange 設定変更権限
	Arrange EPermission
	// Power 電源操作権限
	Power EPermission
	// View リソース閲覧権限
	View EPermission
}{
	Unknown: EPermission(""),
	Create:  EPermission("create"),
	Arrange: EPermission("arrange"),
	Power:   EPermission("power"),
	View:    EPermission("view"),
}

// ExternalPermission 他サービスへのアクセス権
//
// 各権限を表す文字列を+区切りで持つ。
// 例: イベントログと請求閲覧権限がある場合: eventlog+bill
type ExternalPermission string

// PermittedEventLog イベントログ権限を持つか
func (p *ExternalPermission) PermittedEventLog() bool {
	return strings.Contains(string(*p), "eventlog")
}

// PermittedObjectStorage オブジェクトストレージの権限を持つか
func (p *ExternalPermission) PermittedObjectStorage() bool {
	return strings.Contains(string(*p), "dstorage")
}

// PermittedBill 請求閲覧権限を持つか
func (p *ExternalPermission) PermittedBill() bool {
	return strings.Contains(string(*p), "bill")
}

// PermittedWebAccel ウェブアクセラレータの権限を持つか
func (p *ExternalPermission) PermittedWebAccel() bool {
	return strings.Contains(string(*p), "cdn")
}

// PermittedPHY PHYの権限を持つか
func (p *ExternalPermission) PermittedPHY() bool {
	return strings.Contains(string(*p), "dedicatedphy")
}
