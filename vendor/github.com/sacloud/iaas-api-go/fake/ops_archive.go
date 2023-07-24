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
	"context"
	"fmt"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// Find is fake implementation
func (o *ArchiveOp) Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.ArchiveFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*iaas.Archive
	for _, res := range results {
		dest := &iaas.Archive{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &iaas.ArchiveFindResult{
		Total:    len(results),
		Count:    len(results),
		From:     0,
		Archives: values,
	}, nil
}

// Create is fake implementation
func (o *ArchiveOp) Create(ctx context.Context, zone string, param *iaas.ArchiveCreateRequest) (*iaas.Archive, error) {
	result := &iaas.Archive{}

	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillScope)

	if !param.SourceArchiveID.IsEmpty() {
		source, err := o.Read(ctx, zone, param.SourceArchiveID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), "SourceArchive is not found")
		}
		result.SourceArchiveAvailability = source.Availability
	}
	if !param.SourceDiskID.IsEmpty() {
		diskOp := NewDiskOp()
		source, err := diskOp.Read(ctx, zone, param.SourceDiskID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), "SourceDisk is not found")
		}
		result.SourceDiskAvailability = source.Availability
	}

	result.DisplayOrder = int64(random(100))
	result.Availability = types.Availabilities.Migrating
	result.DiskPlanID = types.DiskPlans.HDD
	result.DiskPlanName = "標準プラン"
	result.DiskPlanStorageClass = "iscsi9999"

	putArchive(zone, result)

	id := result.ID
	startDiskCopy(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})

	return result, nil
}

// CreateBlank is fake implementation
func (o *ArchiveOp) CreateBlank(ctx context.Context, zone string, param *iaas.ArchiveCreateBlankRequest) (*iaas.Archive, *iaas.FTPServer, error) {
	result := &iaas.Archive{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillScope)

	result.Availability = types.Availabilities.Uploading

	putArchive(zone, result)

	return result, &iaas.FTPServer{
		HostName:  fmt.Sprintf("sac-%s-ftp.example.jp", zone),
		IPAddress: "192.0.2.1",
		User:      fmt.Sprintf("archive%d", result.ID),
		Password:  "password-is-not-a-password",
	}, nil
}

// Read is fake implementation
func (o *ArchiveOp) Read(ctx context.Context, zone string, id types.ID) (*iaas.Archive, error) {
	value := getArchiveByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &iaas.Archive{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *ArchiveOp) Update(ctx context.Context, zone string, id types.ID, param *iaas.ArchiveUpdateRequest) (*iaas.Archive, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return value, nil
}

// Delete is fake implementation
func (o *ArchiveOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, zone, id)
	return nil
}

// OpenFTP is fake implementation
func (o *ArchiveOp) OpenFTP(ctx context.Context, zone string, id types.ID, openOption *iaas.OpenFTPRequest) (*iaas.FTPServer, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	value.SetAvailability(types.Availabilities.Uploading)
	putArchive(zone, value)

	return &iaas.FTPServer{
		HostName:  fmt.Sprintf("sac-%s-ftp.example.jp", zone),
		IPAddress: "192.0.2.1",
		User:      fmt.Sprintf("archive%d", id),
		Password:  "password-is-not-a-password",
	}, nil
}

// CloseFTP is fake implementation
func (o *ArchiveOp) CloseFTP(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	if !value.Availability.IsUploading() {
		value.SetAvailability(types.Availabilities.Available)
	}
	putArchive(zone, value)
	return nil
}

// Share is fake implementation
func (o *ArchiveOp) Share(ctx context.Context, zone string, id types.ID) (*iaas.ArchiveShareInfo, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	value.SetAvailability(types.Availabilities.Uploading)
	putArchive(zone, value)

	return &iaas.ArchiveShareInfo{
		SharedKey: types.ArchiveShareKey(fmt.Sprintf("%s:%s:%s", zone, id.String(), "xxx")),
	}, nil
}

// CreateFromShared is fake implementation
func (o *ArchiveOp) CreateFromShared(ctx context.Context, zone string, sourceArchiveID types.ID, zoneID types.ID, param *iaas.ArchiveCreateRequestFromShared) (*iaas.Archive, error) {
	result := &iaas.Archive{}

	var destZone string
	for name, id := range zoneIDs {
		if id == zoneID {
			destZone = name
			break
		}
	}

	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillScope)

	result.DisplayOrder = int64(random(100))
	result.Availability = types.Availabilities.Transferring
	result.DiskPlanID = types.DiskPlans.HDD
	result.DiskPlanName = "標準プラン"
	result.DiskPlanStorageClass = "iscsi9999"

	putArchive(destZone, result)

	id := result.ID
	startDiskCopy(o.key, destZone, func() (interface{}, error) {
		return o.Read(context.Background(), destZone, id)
	})

	return result, nil
}

// Transfer is fake implementation
func (o *ArchiveOp) Transfer(ctx context.Context, zone string, sourceArchiveID types.ID, destZoneID types.ID, param *iaas.ArchiveTransferRequest) (*iaas.Archive, error) {
	result := &iaas.Archive{}

	var destZone string
	for name, id := range zoneIDs {
		if id == destZoneID {
			destZone = name
			break
		}
	}

	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillScope)

	result.DisplayOrder = int64(random(100))
	result.Availability = types.Availabilities.Transferring
	result.DiskPlanID = types.DiskPlans.HDD
	result.DiskPlanName = "標準プラン"
	result.DiskPlanStorageClass = "iscsi9999"

	putArchive(destZone, result)

	id := result.ID
	startDiskCopy(o.key, destZone, func() (interface{}, error) {
		return o.Read(context.Background(), destZone, id)
	})

	return result, nil
}
