// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/mdb/greenplum/v1/backup_service.proto

package greenplum

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetBackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the backup to return.
	BackupId string `protobuf:"bytes,1,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
}

func (x *GetBackupRequest) Reset() {
	*x = GetBackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBackupRequest) ProtoMessage() {}

func (x *GetBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBackupRequest.ProtoReflect.Descriptor instead.
func (*GetBackupRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetBackupRequest) GetBackupId() string {
	if x != nil {
		return x.BackupId
	}
	return ""
}

type ListBackupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder to list backups in.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// The maximum number of results per page to return.
	//
	// If the number of available results is larger than [page_size], the service returns a [ListBackupsResponse.next_page_token] that can be used to get the next page of results in subsequent list requests.
	//
	// Default value is 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The page token. To get the next page of results, set [page_token] to the [ListBackupsResponse.next_page_token] returned by the previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListBackupsRequest) Reset() {
	*x = ListBackupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBackupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBackupsRequest) ProtoMessage() {}

func (x *ListBackupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBackupsRequest.ProtoReflect.Descriptor instead.
func (*ListBackupsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListBackupsRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ListBackupsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBackupsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListBackupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested list of backups.
	Backups []*Backup `protobuf:"bytes,1,rep,name=backups,proto3" json:"backups,omitempty"`
	// This token allows you to get the next page of results for a list request.
	//
	// If the number of results is larger than [ListBackupsRequest.page_size] specified in the request, use the [next_page_token] as the value for the [ListBackupsRequest.page_token] parameter in the next list request.
	//
	// Each subsequent ListBackups request has its own [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListBackupsResponse) Reset() {
	*x = ListBackupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBackupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBackupsResponse) ProtoMessage() {}

func (x *ListBackupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBackupsResponse.ProtoReflect.Descriptor instead.
func (*ListBackupsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListBackupsResponse) GetBackups() []*Backup {
	if x != nil {
		return x.Backups
	}
	return nil
}

func (x *ListBackupsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type DeleteBackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the backup to delete.
	BackupId string `protobuf:"bytes,1,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
}

func (x *DeleteBackupRequest) Reset() {
	*x = DeleteBackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBackupRequest) ProtoMessage() {}

func (x *DeleteBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBackupRequest.ProtoReflect.Descriptor instead.
func (*DeleteBackupRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteBackupRequest) GetBackupId() string {
	if x != nil {
		return x.BackupId
	}
	return ""
}

type DeleteBackupMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the Greenplum backup that is currently being deleted.
	BackupId string `protobuf:"bytes,1,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
	// ID of the Greenplum backup that is being deleted.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *DeleteBackupMetadata) Reset() {
	*x = DeleteBackupMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBackupMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBackupMetadata) ProtoMessage() {}

func (x *DeleteBackupMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBackupMetadata.ProtoReflect.Descriptor instead.
func (*DeleteBackupMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBackupMetadata) GetBackupId() string {
	if x != nil {
		return x.BackupId
	}
	return ""
}

func (x *DeleteBackupMetadata) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

var File_yandex_cloud_mdb_greenplum_v1_backup_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x67, 0x72, 0x65,
	0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8,
	0xc7, 0x31, 0x01, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x64, 0x22, 0x92, 0x01,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31,
	0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x30, 0x2d, 0x31, 0x30, 0x30, 0x30, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8,
	0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x7e, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x52, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x38, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7,
	0x31, 0x01, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x64, 0x22, 0x5c, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52,
	0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x32, 0xff, 0x03, 0x0a, 0x0d, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x90, 0x01, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c,
	0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x22, 0x31, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2d, 0x67,
	0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x73, 0x2f, 0x7b, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x94, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x64, 0x2d, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x12, 0xc3, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0xb2, 0xd2, 0x2a, 0x2d, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2b, 0x2a, 0x29, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2d, 0x67, 0x72, 0x65, 0x65,
	0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73,
	0x2f, 0x7b, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x70, 0x0a, 0x21,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2e, 0x76,
	0x31, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75,
	0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDescData = file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDesc
)

func file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDescData
}

var file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_goTypes = []interface{}{
	(*GetBackupRequest)(nil),     // 0: yandex.cloud.mdb.greenplum.v1.GetBackupRequest
	(*ListBackupsRequest)(nil),   // 1: yandex.cloud.mdb.greenplum.v1.ListBackupsRequest
	(*ListBackupsResponse)(nil),  // 2: yandex.cloud.mdb.greenplum.v1.ListBackupsResponse
	(*DeleteBackupRequest)(nil),  // 3: yandex.cloud.mdb.greenplum.v1.DeleteBackupRequest
	(*DeleteBackupMetadata)(nil), // 4: yandex.cloud.mdb.greenplum.v1.DeleteBackupMetadata
	(*Backup)(nil),               // 5: yandex.cloud.mdb.greenplum.v1.Backup
	(*operation.Operation)(nil),  // 6: yandex.cloud.operation.Operation
}
var file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_depIdxs = []int32{
	5, // 0: yandex.cloud.mdb.greenplum.v1.ListBackupsResponse.backups:type_name -> yandex.cloud.mdb.greenplum.v1.Backup
	0, // 1: yandex.cloud.mdb.greenplum.v1.BackupService.Get:input_type -> yandex.cloud.mdb.greenplum.v1.GetBackupRequest
	1, // 2: yandex.cloud.mdb.greenplum.v1.BackupService.List:input_type -> yandex.cloud.mdb.greenplum.v1.ListBackupsRequest
	3, // 3: yandex.cloud.mdb.greenplum.v1.BackupService.Delete:input_type -> yandex.cloud.mdb.greenplum.v1.DeleteBackupRequest
	5, // 4: yandex.cloud.mdb.greenplum.v1.BackupService.Get:output_type -> yandex.cloud.mdb.greenplum.v1.Backup
	2, // 5: yandex.cloud.mdb.greenplum.v1.BackupService.List:output_type -> yandex.cloud.mdb.greenplum.v1.ListBackupsResponse
	6, // 6: yandex.cloud.mdb.greenplum.v1.BackupService.Delete:output_type -> yandex.cloud.operation.Operation
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_init() }
func file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_init() {
	if File_yandex_cloud_mdb_greenplum_v1_backup_service_proto != nil {
		return
	}
	file_yandex_cloud_mdb_greenplum_v1_backup_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBackupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBackupsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBackupsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBackupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBackupMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_greenplum_v1_backup_service_proto = out.File
	file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_rawDesc = nil
	file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_goTypes = nil
	file_yandex_cloud_mdb_greenplum_v1_backup_service_proto_depIdxs = nil
}
