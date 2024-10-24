// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/datatransfer/v1/endpoint/serializers.proto

package endpoint

import (
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

type SerializerAuto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SerializerAuto) Reset() {
	*x = SerializerAuto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerializerAuto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializerAuto) ProtoMessage() {}

func (x *SerializerAuto) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializerAuto.ProtoReflect.Descriptor instead.
func (*SerializerAuto) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDescGZIP(), []int{0}
}

type SerializerJSON struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SerializerJSON) Reset() {
	*x = SerializerJSON{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerializerJSON) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializerJSON) ProtoMessage() {}

func (x *SerializerJSON) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializerJSON.ProtoReflect.Descriptor instead.
func (*SerializerJSON) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDescGZIP(), []int{1}
}

type DebeziumSerializerParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the serializer parameter
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Value of the serializer parameter
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DebeziumSerializerParameter) Reset() {
	*x = DebeziumSerializerParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebeziumSerializerParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebeziumSerializerParameter) ProtoMessage() {}

func (x *DebeziumSerializerParameter) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebeziumSerializerParameter.ProtoReflect.Descriptor instead.
func (*DebeziumSerializerParameter) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDescGZIP(), []int{2}
}

func (x *DebeziumSerializerParameter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DebeziumSerializerParameter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SerializerDebezium struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Settings of sterilization parameters as key-value pairs
	SerializerParameters []*DebeziumSerializerParameter `protobuf:"bytes,1,rep,name=serializer_parameters,json=serializerParameters,proto3" json:"serializer_parameters,omitempty"`
}

func (x *SerializerDebezium) Reset() {
	*x = SerializerDebezium{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerializerDebezium) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializerDebezium) ProtoMessage() {}

func (x *SerializerDebezium) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializerDebezium.ProtoReflect.Descriptor instead.
func (*SerializerDebezium) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDescGZIP(), []int{3}
}

func (x *SerializerDebezium) GetSerializerParameters() []*DebeziumSerializerParameter {
	if x != nil {
		return x.SerializerParameters
	}
	return nil
}

// Data serialization format
type Serializer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Serializer:
	//
	//	*Serializer_SerializerAuto
	//	*Serializer_SerializerJson
	//	*Serializer_SerializerDebezium
	Serializer isSerializer_Serializer `protobuf_oneof:"serializer"`
}

func (x *Serializer) Reset() {
	*x = Serializer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Serializer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Serializer) ProtoMessage() {}

func (x *Serializer) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Serializer.ProtoReflect.Descriptor instead.
func (*Serializer) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDescGZIP(), []int{4}
}

func (m *Serializer) GetSerializer() isSerializer_Serializer {
	if m != nil {
		return m.Serializer
	}
	return nil
}

func (x *Serializer) GetSerializerAuto() *SerializerAuto {
	if x, ok := x.GetSerializer().(*Serializer_SerializerAuto); ok {
		return x.SerializerAuto
	}
	return nil
}

func (x *Serializer) GetSerializerJson() *SerializerJSON {
	if x, ok := x.GetSerializer().(*Serializer_SerializerJson); ok {
		return x.SerializerJson
	}
	return nil
}

func (x *Serializer) GetSerializerDebezium() *SerializerDebezium {
	if x, ok := x.GetSerializer().(*Serializer_SerializerDebezium); ok {
		return x.SerializerDebezium
	}
	return nil
}

type isSerializer_Serializer interface {
	isSerializer_Serializer()
}

type Serializer_SerializerAuto struct {
	// Select the serialization format automatically
	SerializerAuto *SerializerAuto `protobuf:"bytes,1,opt,name=serializer_auto,json=serializerAuto,proto3,oneof"`
}

type Serializer_SerializerJson struct {
	// Serialize data in json format
	SerializerJson *SerializerJSON `protobuf:"bytes,2,opt,name=serializer_json,json=serializerJson,proto3,oneof"`
}

type Serializer_SerializerDebezium struct {
	// Serialize data in debezium format
	SerializerDebezium *SerializerDebezium `protobuf:"bytes,3,opt,name=serializer_debezium,json=serializerDebezium,proto3,oneof"`
}

func (*Serializer_SerializerAuto) isSerializer_Serializer() {}

func (*Serializer_SerializerJson) isSerializer_Serializer() {}

func (*Serializer_SerializerDebezium) isSerializer_Serializer() {}

var File_yandex_cloud_datatransfer_v1_endpoint_serializers_proto protoreflect.FileDescriptor

var file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDesc = []byte{
	0x0a, 0x37, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72,
	0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x45, 0x0a, 0x1b, 0x44, 0x65, 0x62, 0x65, 0x7a, 0x69, 0x75, 0x6d,
	0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x12,
	0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x44, 0x65, 0x62, 0x65, 0x7a, 0x69,
	0x75, 0x6d, 0x12, 0x77, 0x0a, 0x15, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x42, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x62, 0x65, 0x7a, 0x69,
	0x75, 0x6d, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x14, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0xcc, 0x02, 0x0a, 0x0a,
	0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x60, 0x0a, 0x0f, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x41, 0x75, 0x74, 0x6f, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x41, 0x75, 0x74, 0x6f, 0x12, 0x60, 0x0a, 0x0f,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x4a, 0x53, 0x4f, 0x4e, 0x48, 0x00, 0x52, 0x0e,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x6c,
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x62,
	0x65, 0x7a, 0x69, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x44, 0x65,
	0x62, 0x65, 0x7a, 0x69, 0x75, 0x6d, 0x48, 0x00, 0x52, 0x12, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x72, 0x44, 0x65, 0x62, 0x65, 0x7a, 0x69, 0x75, 0x6d, 0x42, 0x0c, 0x0a, 0x0a,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x42, 0xa7, 0x01, 0x0a, 0x29, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0xaa, 0x02, 0x25, 0x59,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDescOnce sync.Once
	file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDescData = file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDesc
)

func file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDescData)
	})
	return file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDescData
}

var file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_goTypes = []interface{}{
	(*SerializerAuto)(nil),              // 0: yandex.cloud.datatransfer.v1.endpoint.SerializerAuto
	(*SerializerJSON)(nil),              // 1: yandex.cloud.datatransfer.v1.endpoint.SerializerJSON
	(*DebeziumSerializerParameter)(nil), // 2: yandex.cloud.datatransfer.v1.endpoint.DebeziumSerializerParameter
	(*SerializerDebezium)(nil),          // 3: yandex.cloud.datatransfer.v1.endpoint.SerializerDebezium
	(*Serializer)(nil),                  // 4: yandex.cloud.datatransfer.v1.endpoint.Serializer
}
var file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.datatransfer.v1.endpoint.SerializerDebezium.serializer_parameters:type_name -> yandex.cloud.datatransfer.v1.endpoint.DebeziumSerializerParameter
	0, // 1: yandex.cloud.datatransfer.v1.endpoint.Serializer.serializer_auto:type_name -> yandex.cloud.datatransfer.v1.endpoint.SerializerAuto
	1, // 2: yandex.cloud.datatransfer.v1.endpoint.Serializer.serializer_json:type_name -> yandex.cloud.datatransfer.v1.endpoint.SerializerJSON
	3, // 3: yandex.cloud.datatransfer.v1.endpoint.Serializer.serializer_debezium:type_name -> yandex.cloud.datatransfer.v1.endpoint.SerializerDebezium
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_init() }
func file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_init() {
	if File_yandex_cloud_datatransfer_v1_endpoint_serializers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerializerAuto); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerializerJSON); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebeziumSerializerParameter); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerializerDebezium); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Serializer); i {
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
	file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Serializer_SerializerAuto)(nil),
		(*Serializer_SerializerJson)(nil),
		(*Serializer_SerializerDebezium)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datatransfer_v1_endpoint_serializers_proto = out.File
	file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_rawDesc = nil
	file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_goTypes = nil
	file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_depIdxs = nil
}