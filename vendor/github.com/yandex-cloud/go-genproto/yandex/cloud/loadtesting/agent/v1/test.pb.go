// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/loadtesting/agent/v1/test.proto

package agent

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

type Test_Status int32

const (
	Test_STATUS_UNSPECIFIED Test_Status = 0
	Test_CREATED            Test_Status = 1
	Test_INITIATED          Test_Status = 2
	Test_PREPARING          Test_Status = 3
	Test_RUNNING            Test_Status = 4
	Test_FINISHING          Test_Status = 5
	Test_DONE               Test_Status = 6
	Test_POST_PROCESSING    Test_Status = 7
	Test_FAILED             Test_Status = 8
	Test_STOPPING           Test_Status = 9
	Test_STOPPED            Test_Status = 10
	Test_AUTOSTOPPED        Test_Status = 11
	Test_WAITING            Test_Status = 12
	Test_DELETING           Test_Status = 13
	Test_LOST               Test_Status = 14
)

// Enum value maps for Test_Status.
var (
	Test_Status_name = map[int32]string{
		0:  "STATUS_UNSPECIFIED",
		1:  "CREATED",
		2:  "INITIATED",
		3:  "PREPARING",
		4:  "RUNNING",
		5:  "FINISHING",
		6:  "DONE",
		7:  "POST_PROCESSING",
		8:  "FAILED",
		9:  "STOPPING",
		10: "STOPPED",
		11: "AUTOSTOPPED",
		12: "WAITING",
		13: "DELETING",
		14: "LOST",
	}
	Test_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATED":            1,
		"INITIATED":          2,
		"PREPARING":          3,
		"RUNNING":            4,
		"FINISHING":          5,
		"DONE":               6,
		"POST_PROCESSING":    7,
		"FAILED":             8,
		"STOPPING":           9,
		"STOPPED":            10,
		"AUTOSTOPPED":        11,
		"WAITING":            12,
		"DELETING":           13,
		"LOST":               14,
	}
)

func (x Test_Status) Enum() *Test_Status {
	p := new(Test_Status)
	*p = x
	return p
}

func (x Test_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Test_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_loadtesting_agent_v1_test_proto_enumTypes[0].Descriptor()
}

func (Test_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_loadtesting_agent_v1_test_proto_enumTypes[0]
}

func (x Test_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Test_Status.Descriptor instead.
func (Test_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDescGZIP(), []int{0, 0}
}

type Test_Generator int32

const (
	Test_GENERATOR_UNSPECIFIED Test_Generator = 0
	Test_PANDORA               Test_Generator = 1
	Test_PHANTOM               Test_Generator = 2
	Test_JMETER                Test_Generator = 3
)

// Enum value maps for Test_Generator.
var (
	Test_Generator_name = map[int32]string{
		0: "GENERATOR_UNSPECIFIED",
		1: "PANDORA",
		2: "PHANTOM",
		3: "JMETER",
	}
	Test_Generator_value = map[string]int32{
		"GENERATOR_UNSPECIFIED": 0,
		"PANDORA":               1,
		"PHANTOM":               2,
		"JMETER":                3,
	}
)

func (x Test_Generator) Enum() *Test_Generator {
	p := new(Test_Generator)
	*p = x
	return p
}

func (x Test_Generator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Test_Generator) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_loadtesting_agent_v1_test_proto_enumTypes[1].Descriptor()
}

func (Test_Generator) Type() protoreflect.EnumType {
	return &file_yandex_cloud_loadtesting_agent_v1_test_proto_enumTypes[1]
}

func (x Test_Generator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Test_Generator.Descriptor instead.
func (Test_Generator) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDescGZIP(), []int{0, 1}
}

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FolderId    string                 `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Labels      map[string]string      `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	StartedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Generator   Test_Generator         `protobuf:"varint,10,opt,name=generator,proto3,enum=yandex.cloud.loadtesting.agent.v1.Test_Generator" json:"generator,omitempty"`
	// AgentInstance ID where Test is running.
	AgentInstanceId string `protobuf:"bytes,11,opt,name=agent_instance_id,json=agentInstanceId,proto3" json:"agent_instance_id,omitempty"`
	// Target VM.
	TargetAddress string `protobuf:"bytes,12,opt,name=target_address,json=targetAddress,proto3" json:"target_address,omitempty"`
	TargetPort    int64  `protobuf:"varint,13,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
	// Version of object under test.
	TargetVersion string `protobuf:"bytes,14,opt,name=target_version,json=targetVersion,proto3" json:"target_version,omitempty"`
	// Test details
	Config string `protobuf:"bytes,15,opt,name=config,proto3" json:"config,omitempty"`
	// Types that are assignable to Ammo:
	//
	//	*Test_AmmoUrls
	//	*Test_AmmoId
	Ammo     isTest_Ammo `protobuf_oneof:"ammo"`
	Cases    []string    `protobuf:"bytes,18,rep,name=cases,proto3" json:"cases,omitempty"`
	Status   Test_Status `protobuf:"varint,19,opt,name=status,proto3,enum=yandex.cloud.loadtesting.agent.v1.Test_Status" json:"status,omitempty"`
	Errors   []string    `protobuf:"bytes,20,rep,name=errors,proto3" json:"errors,omitempty"`
	Favorite bool        `protobuf:"varint,21,opt,name=favorite,proto3" json:"favorite,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_agent_v1_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_agent_v1_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Test) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Test) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Test) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Test) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Test) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Test) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Test) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *Test) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Test) GetGenerator() Test_Generator {
	if x != nil {
		return x.Generator
	}
	return Test_GENERATOR_UNSPECIFIED
}

func (x *Test) GetAgentInstanceId() string {
	if x != nil {
		return x.AgentInstanceId
	}
	return ""
}

func (x *Test) GetTargetAddress() string {
	if x != nil {
		return x.TargetAddress
	}
	return ""
}

func (x *Test) GetTargetPort() int64 {
	if x != nil {
		return x.TargetPort
	}
	return 0
}

func (x *Test) GetTargetVersion() string {
	if x != nil {
		return x.TargetVersion
	}
	return ""
}

func (x *Test) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (m *Test) GetAmmo() isTest_Ammo {
	if m != nil {
		return m.Ammo
	}
	return nil
}

func (x *Test) GetAmmoUrls() string {
	if x, ok := x.GetAmmo().(*Test_AmmoUrls); ok {
		return x.AmmoUrls
	}
	return ""
}

func (x *Test) GetAmmoId() string {
	if x, ok := x.GetAmmo().(*Test_AmmoId); ok {
		return x.AmmoId
	}
	return ""
}

func (x *Test) GetCases() []string {
	if x != nil {
		return x.Cases
	}
	return nil
}

func (x *Test) GetStatus() Test_Status {
	if x != nil {
		return x.Status
	}
	return Test_STATUS_UNSPECIFIED
}

func (x *Test) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *Test) GetFavorite() bool {
	if x != nil {
		return x.Favorite
	}
	return false
}

type isTest_Ammo interface {
	isTest_Ammo()
}

type Test_AmmoUrls struct {
	AmmoUrls string `protobuf:"bytes,16,opt,name=ammo_urls,json=ammoUrls,proto3,oneof"`
}

type Test_AmmoId struct {
	AmmoId string `protobuf:"bytes,17,opt,name=ammo_id,json=ammoId,proto3,oneof"`
}

func (*Test_AmmoUrls) isTest_Ammo() {}

func (*Test_AmmoId) isTest_Ammo() {}

var File_yandex_cloud_loadtesting_agent_v1_test_proto protoreflect.FileDescriptor

var file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61,
	0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xeb, 0x09, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f,
	0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4f, 0x0a, 0x09, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a,
	0x09, 0x61, 0x6d, 0x6d, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x61, 0x6d, 0x6d, 0x6f, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x19, 0x0a, 0x07,
	0x61, 0x6d, 0x6d, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x61, 0x6d, 0x6d, 0x6f, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x12, 0x46, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61,
	0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xe3, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x45, 0x50, 0x41, 0x52, 0x49, 0x4e, 0x47,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12,
	0x0d, 0x0a, 0x09, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4f, 0x53, 0x54,
	0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f,
	0x50, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50,
	0x45, 0x44, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x55, 0x54, 0x4f, 0x53, 0x54, 0x4f, 0x50,
	0x50, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x0d,
	0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x53, 0x54, 0x10, 0x0e, 0x22, 0x4c, 0x0a, 0x09, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x45, 0x4e, 0x45, 0x52,
	0x41, 0x54, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41, 0x4e, 0x44, 0x4f, 0x52, 0x41, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x48, 0x41, 0x4e, 0x54, 0x4f, 0x4d, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x4a, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x03, 0x42, 0x06, 0x0a, 0x04, 0x61, 0x6d, 0x6d, 0x6f,
	0x42, 0x74, 0x0a, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x61, 0x64,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDescOnce sync.Once
	file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDescData = file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDesc
)

func file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDescGZIP() []byte {
	file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDescData)
	})
	return file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDescData
}

var file_yandex_cloud_loadtesting_agent_v1_test_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_loadtesting_agent_v1_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_loadtesting_agent_v1_test_proto_goTypes = []interface{}{
	(Test_Status)(0),              // 0: yandex.cloud.loadtesting.agent.v1.Test.Status
	(Test_Generator)(0),           // 1: yandex.cloud.loadtesting.agent.v1.Test.Generator
	(*Test)(nil),                  // 2: yandex.cloud.loadtesting.agent.v1.Test
	nil,                           // 3: yandex.cloud.loadtesting.agent.v1.Test.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_yandex_cloud_loadtesting_agent_v1_test_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.loadtesting.agent.v1.Test.labels:type_name -> yandex.cloud.loadtesting.agent.v1.Test.LabelsEntry
	4, // 1: yandex.cloud.loadtesting.agent.v1.Test.created_at:type_name -> google.protobuf.Timestamp
	4, // 2: yandex.cloud.loadtesting.agent.v1.Test.started_at:type_name -> google.protobuf.Timestamp
	4, // 3: yandex.cloud.loadtesting.agent.v1.Test.finished_at:type_name -> google.protobuf.Timestamp
	4, // 4: yandex.cloud.loadtesting.agent.v1.Test.updated_at:type_name -> google.protobuf.Timestamp
	1, // 5: yandex.cloud.loadtesting.agent.v1.Test.generator:type_name -> yandex.cloud.loadtesting.agent.v1.Test.Generator
	0, // 6: yandex.cloud.loadtesting.agent.v1.Test.status:type_name -> yandex.cloud.loadtesting.agent.v1.Test.Status
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_yandex_cloud_loadtesting_agent_v1_test_proto_init() }
func file_yandex_cloud_loadtesting_agent_v1_test_proto_init() {
	if File_yandex_cloud_loadtesting_agent_v1_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_loadtesting_agent_v1_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
	file_yandex_cloud_loadtesting_agent_v1_test_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Test_AmmoUrls)(nil),
		(*Test_AmmoId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_loadtesting_agent_v1_test_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_loadtesting_agent_v1_test_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_loadtesting_agent_v1_test_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_loadtesting_agent_v1_test_proto_msgTypes,
	}.Build()
	File_yandex_cloud_loadtesting_agent_v1_test_proto = out.File
	file_yandex_cloud_loadtesting_agent_v1_test_proto_rawDesc = nil
	file_yandex_cloud_loadtesting_agent_v1_test_proto_goTypes = nil
	file_yandex_cloud_loadtesting_agent_v1_test_proto_depIdxs = nil
}
