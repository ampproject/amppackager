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

import (
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/packages-go/wait"
)

// WaiterForUp 起動完了まで待つためのStateWaiterを返す
func WaiterForUp(readFunc wait.StateReadFunc) wait.StateWaiter {
	return &StatePollingWaiter{
		ReadFunc: readFunc,
		TargetAvailability: []types.EAvailability{
			types.Availabilities.Available,
		},
		PendingAvailability: []types.EAvailability{
			types.Availabilities.Unknown,
			types.Availabilities.Migrating,
			types.Availabilities.Uploading,
			types.Availabilities.Transferring,
			types.Availabilities.Discontinued,
		},
		TargetInstanceStatus: []types.EServerInstanceStatus{
			types.ServerInstanceStatuses.Up,
		},
		PendingInstanceStatus: []types.EServerInstanceStatus{
			types.ServerInstanceStatuses.Unknown,
			types.ServerInstanceStatuses.Cleaning,
			types.ServerInstanceStatuses.Down,
		},
	}
}

// WaiterForApplianceUp 起動完了まで待つためのStateWaiterを返す
//
// アプライアンス向けに404発生時のリトライを設定可能
func WaiterForApplianceUp(readFunc wait.StateReadFunc, notFoundRetry int) wait.StateWaiter {
	return &StatePollingWaiter{
		ReadFunc: readFunc,
		TargetAvailability: []types.EAvailability{
			types.Availabilities.Available,
		},
		PendingAvailability: []types.EAvailability{
			types.Availabilities.Unknown,
			types.Availabilities.Migrating,
			types.Availabilities.Uploading,
			types.Availabilities.Transferring,
			types.Availabilities.Discontinued,
		},
		TargetInstanceStatus: []types.EServerInstanceStatus{
			types.ServerInstanceStatuses.Up,
		},
		PendingInstanceStatus: []types.EServerInstanceStatus{
			types.ServerInstanceStatuses.Unknown,
			types.ServerInstanceStatuses.Cleaning,
			types.ServerInstanceStatuses.Down,
		},
		NotFoundRetry: notFoundRetry,
	}
}

// WaiterForDown シャットダウン完了まで待つためのStateWaiterを返す
func WaiterForDown(readFunc wait.StateReadFunc) wait.StateWaiter {
	return &StatePollingWaiter{
		ReadFunc: readFunc,
		TargetAvailability: []types.EAvailability{
			types.Availabilities.Available,
		},
		PendingAvailability: []types.EAvailability{
			types.Availabilities.Unknown,
		},
		TargetInstanceStatus: []types.EServerInstanceStatus{
			types.ServerInstanceStatuses.Down,
		},
		PendingInstanceStatus: []types.EServerInstanceStatus{
			types.ServerInstanceStatuses.Up,
			types.ServerInstanceStatuses.Cleaning,
			types.ServerInstanceStatuses.Unknown,
		},
	}
}

// WaiterForReady リソースの利用準備完了まで待つためのStateWaiterを返す
func WaiterForReady(readFunc wait.StateReadFunc) wait.StateWaiter {
	return &StatePollingWaiter{
		ReadFunc: readFunc,
		TargetAvailability: []types.EAvailability{
			types.Availabilities.Available,
		},
		PendingAvailability: []types.EAvailability{
			types.Availabilities.Unknown,
			types.Availabilities.Migrating,
			types.Availabilities.Uploading,
			types.Availabilities.Transferring,
			types.Availabilities.Discontinued,
		},
	}
}
