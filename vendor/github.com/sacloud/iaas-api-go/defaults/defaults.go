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

package defaults

import "time"

// for setup package
var (
	// DefaultStatePollingTimeout StatePollWaiterでのデフォルトタイムアウト
	DefaultStatePollingTimeout = 20 * time.Minute
	// DefaultStatePollingInterval StatePollWaiterでのデフォルトポーリング間隔
	DefaultStatePollingInterval = 5 * time.Second

	// DefaultPowerHelperBootRetrySpan helper/powerでの起動リクエストリトライ間隔
	DefaultPowerHelperBootRetrySpan = 20 * time.Second
	// DefaultPowerHelperShutdownRetrySpan helper/powerでのシャットダウンリクエストリトライ間隔
	DefaultPowerHelperShutdownRetrySpan = 20 * time.Second
	// DefaultPowerHelperInitialRequestTimeout helper/powerでの初回電源リクエスト成功までのタイムアウト
	DefaultPowerHelperInitialRequestTimeout = 30 * time.Minute
	// DefaultPowerHelperInitialRequestRetrySpan helper/powerでの初回リクエスト409+still_creating時のリトライ間隔
	DefaultPowerHelperInitialRequestRetrySpan = 20 * time.Second
)

// for builder package
var (
	// DefaultNICUpdateWaitDuration NIC切断/削除後の待ち時間デフォルト値
	DefaultNICUpdateWaitDuration = 5 * time.Second
)

// for cleanup package
var (
	// DefaultDBStatusPollingInterval データベースアプライアンスのステータス確認ポーリングの間隔
	DefaultDBStatusPollingInterval = 30 * time.Second
)
