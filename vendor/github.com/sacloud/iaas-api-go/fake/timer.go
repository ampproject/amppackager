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

package fake

import (
	"fmt"
	"time"

	"github.com/sacloud/iaas-api-go/accessor"
	"github.com/sacloud/iaas-api-go/types"
)

var (
	// DiskCopyDuration ディスクコピー処理のtickerで利用するduration
	DiskCopyDuration = 10 * time.Millisecond
	// PowerOnDuration 電源On処理のtickerで利用するduration
	PowerOnDuration = 10 * time.Millisecond
	// PowerOffDuration 電源Off処理のtickerで利用するduration
	PowerOffDuration = 10 * time.Millisecond
)

func startDiskCopy(resourceKey, zone string, readFunc func() (interface{}, error)) {
	counter := 0
	ticker := time.NewTicker(DiskCopyDuration)
	go func() {
		defer ticker.Stop()
		for {
			<-ticker.C

			raw, err := readFunc()
			if raw == nil || err != nil {
				return
			}
			target, ok := raw.(accessor.DiskMigratable)
			if !ok {
				return
			}

			if counter < 3 {
				target.SetAvailability(types.Availabilities.Migrating)
				if counter == 0 {
					target.SetMigratedMB(0)
				} else {
					target.SetMigratedMB(target.GetSizeMB() / counter)
				}
			} else {
				target.SetAvailability(types.Availabilities.Available)
				target.SetMigratedMB(target.GetSizeMB())
				ds().Put(resourceKey, zone, target.(accessor.ID).GetID(), target)
				return
			}
			ds().Put(resourceKey, zone, target.(accessor.ID).GetID(), target)
			counter++
		}
	}()
}

func startMigration(resourceKey, zone string, readFunc func() (interface{}, error)) {
	counter := 0
	ticker := time.NewTicker(DiskCopyDuration)
	go func() {
		defer ticker.Stop()
		for {
			<-ticker.C

			raw, err := readFunc()
			if raw == nil || err != nil {
				return
			}
			target, ok := raw.(accessor.Availability)
			if !ok {
				return
			}

			if counter < 3 {
				target.SetAvailability(types.Availabilities.Migrating)
			} else {
				target.SetAvailability(types.Availabilities.Available)
				ds().Put(resourceKey, zone, target.(accessor.ID).GetID(), target)
				return
			}
			ds().Put(resourceKey, zone, target.(accessor.ID).GetID(), target)
			counter++
		}
	}()
}

func startPowerOn(resourceKey, zone string, readFunc func() (interface{}, error)) {
	counter := 0
	ticker := time.NewTicker(PowerOnDuration)
	go func() {
		defer ticker.Stop()
		for {
			<-ticker.C

			raw, err := readFunc()
			if raw == nil || err != nil {
				return
			}
			target, ok := raw.(accessor.InstanceStatus)
			if !ok {
				return
			}

			if counter < 3 {
				target.SetInstanceStatus(types.ServerInstanceStatuses.Down)
			} else {
				target.SetInstanceStatus(types.ServerInstanceStatuses.Up)
				if status, ok := target.(accessor.Instance); ok {
					status.SetInstanceHostName(fmt.Sprintf("sac-%s-svXXX", zone))
					status.SetInstanceHostInfoURL("")
					status.SetInstanceStatusChangedAt(time.Now())
				}
				if available, ok := target.(accessor.Availability); ok {
					available.SetAvailability(types.Availabilities.Available)
				}
				ds().Put(resourceKey, zone, target.(accessor.ID).GetID(), target)
				return
			}
			ds().Put(resourceKey, zone, target.(accessor.ID).GetID(), target)
			counter++
		}
	}()
}

func startPowerOff(resourceKey, zone string, readFunc func() (interface{}, error)) {
	counter := 0
	ticker := time.NewTicker(PowerOffDuration)
	go func() {
		defer ticker.Stop()
		for {
			<-ticker.C

			raw, err := readFunc()
			if raw == nil || err != nil {
				return
			}
			target, ok := raw.(accessor.InstanceStatus)
			if !ok {
				return
			}

			if status, ok := target.(accessor.Instance); ok {
				status.SetInstanceHostName(fmt.Sprintf("sac-%s-svXXX", zone))
				status.SetInstanceHostInfoURL("")
				status.SetInstanceStatusChangedAt(time.Now())
			}

			if counter < 3 {
				target.SetInstanceStatus(types.ServerInstanceStatuses.Cleaning)
			} else {
				target.SetInstanceStatus(types.ServerInstanceStatuses.Down)
				ds().Put(resourceKey, zone, target.(accessor.ID).GetID(), target)
				return
			}

			ds().Put(resourceKey, zone, target.(accessor.ID).GetID(), target)
			counter++
		}
	}()
}
