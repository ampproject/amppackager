// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/marketplace/licensemanager/v1/lock.proto

package licensemanager

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Lock_State int32

const (
	Lock_STATE_UNSPECIFIED Lock_State = 0
	// Subscription unlocked.
	Lock_UNLOCKED Lock_State = 1
	// Subscription locked to the resource.
	Lock_LOCKED Lock_State = 2
	// Subscription lock deleted.
	Lock_DELETED Lock_State = 3
)

// Enum value maps for Lock_State.
var (
	Lock_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "UNLOCKED",
		2: "LOCKED",
		3: "DELETED",
	}
	Lock_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"UNLOCKED":          1,
		"LOCKED":            2,
		"DELETED":           3,
	}
)

func (x Lock_State) Enum() *Lock_State {
	p := new(Lock_State)
	*p = x
	return p
}

func (x Lock_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Lock_State) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_enumTypes[0].Descriptor()
}

func (Lock_State) Type() protoreflect.EnumType {
	return &file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_enumTypes[0]
}

func (x Lock_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Lock_State.Descriptor instead.
func (Lock_State) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDescGZIP(), []int{0, 0}
}

type Lock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the subscription lock.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the subscription instance.
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// ID of the resource.
	ResourceId string `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// Timestamp of the start of the subscription lock.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Timestamp of the end of the subscription lock.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Update timestamp.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Subscription lock state.
	State Lock_State `protobuf:"varint,8,opt,name=state,proto3,enum=yandex.cloud.marketplace.licensemanager.v1.Lock_State" json:"state,omitempty"`
	// ID of the subscription template.
	TemplateId string `protobuf:"bytes,9,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
}

func (x *Lock) Reset() {
	*x = Lock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lock) ProtoMessage() {}

func (x *Lock) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lock.ProtoReflect.Descriptor instead.
func (*Lock) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDescGZIP(), []int{0}
}

func (x *Lock) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Lock) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *Lock) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *Lock) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Lock) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Lock) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Lock) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Lock) GetState() Lock_State {
	if x != nil {
		return x.State
	}
	return Lock_STATE_UNSPECIFIED
}

func (x *Lock) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

var File_yandex_cloud_marketplace_licensemanager_v1_lock_proto protoreflect.FileDescriptor

var file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDesc = []byte{
	0x0a, 0x35, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x03, 0x0a, 0x04, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x8f, 0x01,
	0x0a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDescOnce sync.Once
	file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDescData = file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDesc
)

func file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDescGZIP() []byte {
	file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDescData)
	})
	return file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDescData
}

var file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_goTypes = []interface{}{
	(Lock_State)(0),               // 0: yandex.cloud.marketplace.licensemanager.v1.Lock.State
	(*Lock)(nil),                  // 1: yandex.cloud.marketplace.licensemanager.v1.Lock
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.marketplace.licensemanager.v1.Lock.start_time:type_name -> google.protobuf.Timestamp
	2, // 1: yandex.cloud.marketplace.licensemanager.v1.Lock.end_time:type_name -> google.protobuf.Timestamp
	2, // 2: yandex.cloud.marketplace.licensemanager.v1.Lock.created_at:type_name -> google.protobuf.Timestamp
	2, // 3: yandex.cloud.marketplace.licensemanager.v1.Lock.updated_at:type_name -> google.protobuf.Timestamp
	0, // 4: yandex.cloud.marketplace.licensemanager.v1.Lock.state:type_name -> yandex.cloud.marketplace.licensemanager.v1.Lock.State
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_init() }
func file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_init() {
	if File_yandex_cloud_marketplace_licensemanager_v1_lock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lock); i {
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
			RawDescriptor: file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_msgTypes,
	}.Build()
	File_yandex_cloud_marketplace_licensemanager_v1_lock_proto = out.File
	file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_rawDesc = nil
	file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_goTypes = nil
	file_yandex_cloud_marketplace_licensemanager_v1_lock_proto_depIdxs = nil
}
