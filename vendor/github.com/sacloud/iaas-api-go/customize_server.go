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

package iaas

import "github.com/sacloud/iaas-api-go/types"

// BandWidthAt 指定インデックスのNICの帯域幅を算出
//
// 不明な場合は-1を、制限なしの場合は0を、以外の場合はMbps単位で返す
func (o *Server) BandWidthAt(index int) int {
	if len(o.Interfaces) <= index {
		return -1
	}

	nic := o.Interfaces[index]

	switch nic.UpstreamType {
	case types.UpstreamNetworkTypes.None:
		return -1
	case types.UpstreamNetworkTypes.Shared:
		return 100
	case types.UpstreamNetworkTypes.Switch, types.UpstreamNetworkTypes.Router:
		//
		// 上流ネットワークがスイッチだった場合の帯域制限
		// https://manual.sakura.ad.jp/cloud/support/technical/network.html#support-network-03
		//

		// 専有ホストの場合は制限なし
		if !o.PrivateHostID.IsEmpty() {
			return 0
		}

		// メモリに応じた制限
		memory := o.GetMemoryGB()
		switch {
		case memory < 32:
			return 1000
		case 32 <= memory && memory < 128:
			return 2000
		case 128 <= memory && memory < 224:
			return 5000
		case 224 <= memory:
			return 10000
		default:
			return -1
		}
	default:
		return -1
	}
}

// GetInstanceStatus データベース(サービス)ステータスを返すためのアダプター実装
// PostgreSQLまたはMariaDBのステータス(詳細は以下)をInstanceStatusにラップして返す
//
//	ステータス: GET /appliance/:id/status -> Appliance.ResponseStatus.DBConf.{MariaDB | postgres}.status
//
// 主にStateWaiterで利用する。
func (o *DatabaseStatus) GetInstanceStatus() types.EServerInstanceStatus {
	if o.MariaDBStatus == "running" || o.PostgresStatus == "running" {
		return types.ServerInstanceStatuses.Up
	}
	return types.ServerInstanceStatuses.Unknown
}

// SetInstanceStatus データベース(サービス)ステータスを返すためのアダプター実装
// accessor.InstanceStatusを満たすためのスタブ実装
func (o *DatabaseStatus) SetInstanceStatus(types.EServerInstanceStatus) {
	// noop
}
