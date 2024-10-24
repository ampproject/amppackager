// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/ai/stt/v3/stt_service.proto

package stt

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

type GetRecognitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
}

func (x *GetRecognitionRequest) Reset() {
	*x = GetRecognitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_stt_v3_stt_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecognitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecognitionRequest) ProtoMessage() {}

func (x *GetRecognitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_stt_v3_stt_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecognitionRequest.ProtoReflect.Descriptor instead.
func (*GetRecognitionRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_stt_v3_stt_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRecognitionRequest) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

var File_yandex_cloud_ai_stt_v3_stt_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_stt_v3_stt_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x73, 0x74, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x74, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x20, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x73, 0x74,
	0x74, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d,
	0x35, 0x30, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32,
	0x71, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x63, 0x0a,
	0x12, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x12, 0x22, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x6b, 0x69, 0x74, 0x2e,
	0x73, 0x74, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68,
	0x6b, 0x69, 0x74, 0x2e, 0x73, 0x74, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x32, 0xb3, 0x02, 0x0a, 0x0f, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x9c, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x7a, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x26, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63,
	0x68, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x74, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x7a, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x40, 0xb2, 0xd2, 0x2a, 0x17, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x73, 0x74, 0x74, 0x2f,
	0x76, 0x33, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x41, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x80, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63,
	0x68, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x74, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x74,
	0x74, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16,
	0x2f, 0x73, 0x74, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x42, 0x5c, 0x0a, 0x1a, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x69, 0x2e,
	0x73, 0x74, 0x74, 0x2e, 0x76, 0x33, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x73, 0x74, 0x74, 0x2f,
	0x76, 0x33, 0x3b, 0x73, 0x74, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_stt_v3_stt_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_stt_v3_stt_service_proto_rawDescData = file_yandex_cloud_ai_stt_v3_stt_service_proto_rawDesc
)

func file_yandex_cloud_ai_stt_v3_stt_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_stt_v3_stt_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_stt_v3_stt_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_stt_v3_stt_service_proto_rawDescData)
	})
	return file_yandex_cloud_ai_stt_v3_stt_service_proto_rawDescData
}

var file_yandex_cloud_ai_stt_v3_stt_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_ai_stt_v3_stt_service_proto_goTypes = []interface{}{
	(*GetRecognitionRequest)(nil), // 0: speechkit.stt.v3.GetRecognitionRequest
	(*StreamingRequest)(nil),      // 1: speechkit.stt.v3.StreamingRequest
	(*RecognizeFileRequest)(nil),  // 2: speechkit.stt.v3.RecognizeFileRequest
	(*StreamingResponse)(nil),     // 3: speechkit.stt.v3.StreamingResponse
	(*operation.Operation)(nil),   // 4: yandex.cloud.operation.Operation
}
var file_yandex_cloud_ai_stt_v3_stt_service_proto_depIdxs = []int32{
	1, // 0: speechkit.stt.v3.Recognizer.RecognizeStreaming:input_type -> speechkit.stt.v3.StreamingRequest
	2, // 1: speechkit.stt.v3.AsyncRecognizer.RecognizeFile:input_type -> speechkit.stt.v3.RecognizeFileRequest
	0, // 2: speechkit.stt.v3.AsyncRecognizer.GetRecognition:input_type -> speechkit.stt.v3.GetRecognitionRequest
	3, // 3: speechkit.stt.v3.Recognizer.RecognizeStreaming:output_type -> speechkit.stt.v3.StreamingResponse
	4, // 4: speechkit.stt.v3.AsyncRecognizer.RecognizeFile:output_type -> yandex.cloud.operation.Operation
	3, // 5: speechkit.stt.v3.AsyncRecognizer.GetRecognition:output_type -> speechkit.stt.v3.StreamingResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_stt_v3_stt_service_proto_init() }
func file_yandex_cloud_ai_stt_v3_stt_service_proto_init() {
	if File_yandex_cloud_ai_stt_v3_stt_service_proto != nil {
		return
	}
	file_yandex_cloud_ai_stt_v3_stt_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_stt_v3_stt_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecognitionRequest); i {
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
			RawDescriptor: file_yandex_cloud_ai_stt_v3_stt_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_yandex_cloud_ai_stt_v3_stt_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_stt_v3_stt_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_stt_v3_stt_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_stt_v3_stt_service_proto = out.File
	file_yandex_cloud_ai_stt_v3_stt_service_proto_rawDesc = nil
	file_yandex_cloud_ai_stt_v3_stt_service_proto_goTypes = nil
	file_yandex_cloud_ai_stt_v3_stt_service_proto_depIdxs = nil
}
