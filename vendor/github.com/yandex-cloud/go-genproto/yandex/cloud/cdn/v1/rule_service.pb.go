// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/cdn/v1/rule_service.proto

package cdn

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

type ListResourceRulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of resource.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ListResourceRulesRequest) Reset() {
	*x = ListResourceRulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourceRulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourceRulesRequest) ProtoMessage() {}

func (x *ListResourceRulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourceRulesRequest.ProtoReflect.Descriptor instead.
func (*ListResourceRulesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_rule_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListResourceRulesRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type ListResourceRulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of the resource rules.
	Rules []*Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *ListResourceRulesResponse) Reset() {
	*x = ListResourceRulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourceRulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourceRulesResponse) ProtoMessage() {}

func (x *ListResourceRulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourceRulesResponse.ProtoReflect.Descriptor instead.
func (*ListResourceRulesResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_rule_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListResourceRulesResponse) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type CreateResourceRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of resource.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// Name of created resource rule.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Resource rule pattern.
	RulePattern string           `protobuf:"bytes,3,opt,name=rule_pattern,json=rulePattern,proto3" json:"rule_pattern,omitempty"`
	Options     *ResourceOptions `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *CreateResourceRuleRequest) Reset() {
	*x = CreateResourceRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceRuleRequest) ProtoMessage() {}

func (x *CreateResourceRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceRuleRequest.ProtoReflect.Descriptor instead.
func (*CreateResourceRuleRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_rule_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResourceRuleRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *CreateResourceRuleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateResourceRuleRequest) GetRulePattern() string {
	if x != nil {
		return x.RulePattern
	}
	return ""
}

func (x *CreateResourceRuleRequest) GetOptions() *ResourceOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type CreateResourceRuleMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of resource.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// ID of created resource rule.
	RuleId int64 `protobuf:"varint,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
}

func (x *CreateResourceRuleMetadata) Reset() {
	*x = CreateResourceRuleMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResourceRuleMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResourceRuleMetadata) ProtoMessage() {}

func (x *CreateResourceRuleMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResourceRuleMetadata.ProtoReflect.Descriptor instead.
func (*CreateResourceRuleMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_rule_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateResourceRuleMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *CreateResourceRuleMetadata) GetRuleId() int64 {
	if x != nil {
		return x.RuleId
	}
	return 0
}

type GetResourceRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of resource.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// ID of the requested resource rule.
	RuleId int64 `protobuf:"varint,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
}

func (x *GetResourceRuleRequest) Reset() {
	*x = GetResourceRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceRuleRequest) ProtoMessage() {}

func (x *GetResourceRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceRuleRequest.ProtoReflect.Descriptor instead.
func (*GetResourceRuleRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_rule_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetResourceRuleRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *GetResourceRuleRequest) GetRuleId() int64 {
	if x != nil {
		return x.RuleId
	}
	return 0
}

type UpdateResourceRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of resource.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// ID of updated resource rule.
	RuleId int64 `protobuf:"varint,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	// Name of updated resource rule.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Resource rule pattern.
	RulePattern string           `protobuf:"bytes,4,opt,name=rule_pattern,json=rulePattern,proto3" json:"rule_pattern,omitempty"`
	Options     *ResourceOptions `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *UpdateResourceRuleRequest) Reset() {
	*x = UpdateResourceRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceRuleRequest) ProtoMessage() {}

func (x *UpdateResourceRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceRuleRequest.ProtoReflect.Descriptor instead.
func (*UpdateResourceRuleRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_rule_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateResourceRuleRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdateResourceRuleRequest) GetRuleId() int64 {
	if x != nil {
		return x.RuleId
	}
	return 0
}

func (x *UpdateResourceRuleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateResourceRuleRequest) GetRulePattern() string {
	if x != nil {
		return x.RulePattern
	}
	return ""
}

func (x *UpdateResourceRuleRequest) GetOptions() *ResourceOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type UpdateResourceRuleMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of resource.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// ID of updated resource rule.
	RuleId int64 `protobuf:"varint,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
}

func (x *UpdateResourceRuleMetadata) Reset() {
	*x = UpdateResourceRuleMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceRuleMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceRuleMetadata) ProtoMessage() {}

func (x *UpdateResourceRuleMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceRuleMetadata.ProtoReflect.Descriptor instead.
func (*UpdateResourceRuleMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_rule_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateResourceRuleMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdateResourceRuleMetadata) GetRuleId() int64 {
	if x != nil {
		return x.RuleId
	}
	return 0
}

type DeleteResourceRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of resource.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// ID of deleted resource rule.
	RuleId int64 `protobuf:"varint,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
}

func (x *DeleteResourceRuleRequest) Reset() {
	*x = DeleteResourceRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResourceRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceRuleRequest) ProtoMessage() {}

func (x *DeleteResourceRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceRuleRequest.ProtoReflect.Descriptor instead.
func (*DeleteResourceRuleRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_rule_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteResourceRuleRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *DeleteResourceRuleRequest) GetRuleId() int64 {
	if x != nil {
		return x.RuleId
	}
	return 0
}

type DeleteResourceRuleMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of resource.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// ID of deleted resource rule.
	RuleId int64 `protobuf:"varint,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
}

func (x *DeleteResourceRuleMetadata) Reset() {
	*x = DeleteResourceRuleMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResourceRuleMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceRuleMetadata) ProtoMessage() {}

func (x *DeleteResourceRuleMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceRuleMetadata.ProtoReflect.Descriptor instead.
func (*DeleteResourceRuleMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_rule_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteResourceRuleMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *DeleteResourceRuleMetadata) GetRuleId() int64 {
	if x != nil {
		return x.RuleId
	}
	return 0
}

var File_yandex_cloud_cdn_v1_rule_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_cdn_v1_rule_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x64, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a,
	0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x22, 0xde, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d,
	0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7,
	0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x30, 0x0a, 0x0c, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x05,
	0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x0b, 0x72, 0x75, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x12, 0x3e, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x6c, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c,
	0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x07, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64,
	0x22, 0x68, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x07, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02,
	0x3e, 0x30, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x22, 0xf7, 0x01, 0x0a, 0x19, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8,
	0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x07, 0x72, 0x75, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02, 0x3e, 0x30,
	0x52, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x0c, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8,
	0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x0b, 0x72, 0x75, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x6e, 0x12, 0x3e, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x6c, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31,
	0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x07, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65,
	0x49, 0x64, 0x22, 0x6b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d,
	0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x07, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x06, 0xfa, 0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0x6c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x75, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x07,
	0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa,
	0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x32, 0x80, 0x06,
	0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x9b, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x3e, 0xb2, 0xd2, 0x2a, 0x22, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x12, 0x6e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c,
	0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x63, 0x64, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0xa5, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x48, 0xb2, 0xd2, 0x2a, 0x22, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01,
	0x2a, 0x32, 0x17, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2f, 0x7b, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xb3, 0x01, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x56, 0xb2, 0xd2, 0x2a, 0x33, 0x0a, 0x1a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x7d,
	0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x64,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x64, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_cdn_v1_rule_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_cdn_v1_rule_service_proto_rawDescData = file_yandex_cloud_cdn_v1_rule_service_proto_rawDesc
)

func file_yandex_cloud_cdn_v1_rule_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_cdn_v1_rule_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_cdn_v1_rule_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_cdn_v1_rule_service_proto_rawDescData)
	})
	return file_yandex_cloud_cdn_v1_rule_service_proto_rawDescData
}

var file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_yandex_cloud_cdn_v1_rule_service_proto_goTypes = []interface{}{
	(*ListResourceRulesRequest)(nil),   // 0: yandex.cloud.cdn.v1.ListResourceRulesRequest
	(*ListResourceRulesResponse)(nil),  // 1: yandex.cloud.cdn.v1.ListResourceRulesResponse
	(*CreateResourceRuleRequest)(nil),  // 2: yandex.cloud.cdn.v1.CreateResourceRuleRequest
	(*CreateResourceRuleMetadata)(nil), // 3: yandex.cloud.cdn.v1.CreateResourceRuleMetadata
	(*GetResourceRuleRequest)(nil),     // 4: yandex.cloud.cdn.v1.GetResourceRuleRequest
	(*UpdateResourceRuleRequest)(nil),  // 5: yandex.cloud.cdn.v1.UpdateResourceRuleRequest
	(*UpdateResourceRuleMetadata)(nil), // 6: yandex.cloud.cdn.v1.UpdateResourceRuleMetadata
	(*DeleteResourceRuleRequest)(nil),  // 7: yandex.cloud.cdn.v1.DeleteResourceRuleRequest
	(*DeleteResourceRuleMetadata)(nil), // 8: yandex.cloud.cdn.v1.DeleteResourceRuleMetadata
	(*Rule)(nil),                       // 9: yandex.cloud.cdn.v1.Rule
	(*ResourceOptions)(nil),            // 10: yandex.cloud.cdn.v1.ResourceOptions
	(*operation.Operation)(nil),        // 11: yandex.cloud.operation.Operation
}
var file_yandex_cloud_cdn_v1_rule_service_proto_depIdxs = []int32{
	9,  // 0: yandex.cloud.cdn.v1.ListResourceRulesResponse.rules:type_name -> yandex.cloud.cdn.v1.Rule
	10, // 1: yandex.cloud.cdn.v1.CreateResourceRuleRequest.options:type_name -> yandex.cloud.cdn.v1.ResourceOptions
	10, // 2: yandex.cloud.cdn.v1.UpdateResourceRuleRequest.options:type_name -> yandex.cloud.cdn.v1.ResourceOptions
	0,  // 3: yandex.cloud.cdn.v1.ResourceRulesService.List:input_type -> yandex.cloud.cdn.v1.ListResourceRulesRequest
	2,  // 4: yandex.cloud.cdn.v1.ResourceRulesService.Create:input_type -> yandex.cloud.cdn.v1.CreateResourceRuleRequest
	4,  // 5: yandex.cloud.cdn.v1.ResourceRulesService.Get:input_type -> yandex.cloud.cdn.v1.GetResourceRuleRequest
	5,  // 6: yandex.cloud.cdn.v1.ResourceRulesService.Update:input_type -> yandex.cloud.cdn.v1.UpdateResourceRuleRequest
	7,  // 7: yandex.cloud.cdn.v1.ResourceRulesService.Delete:input_type -> yandex.cloud.cdn.v1.DeleteResourceRuleRequest
	1,  // 8: yandex.cloud.cdn.v1.ResourceRulesService.List:output_type -> yandex.cloud.cdn.v1.ListResourceRulesResponse
	11, // 9: yandex.cloud.cdn.v1.ResourceRulesService.Create:output_type -> yandex.cloud.operation.Operation
	9,  // 10: yandex.cloud.cdn.v1.ResourceRulesService.Get:output_type -> yandex.cloud.cdn.v1.Rule
	11, // 11: yandex.cloud.cdn.v1.ResourceRulesService.Update:output_type -> yandex.cloud.operation.Operation
	11, // 12: yandex.cloud.cdn.v1.ResourceRulesService.Delete:output_type -> yandex.cloud.operation.Operation
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_cdn_v1_rule_service_proto_init() }
func file_yandex_cloud_cdn_v1_rule_service_proto_init() {
	if File_yandex_cloud_cdn_v1_rule_service_proto != nil {
		return
	}
	file_yandex_cloud_cdn_v1_resource_proto_init()
	file_yandex_cloud_cdn_v1_rule_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourceRulesRequest); i {
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
		file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourceRulesResponse); i {
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
		file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceRuleRequest); i {
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
		file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResourceRuleMetadata); i {
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
		file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceRuleRequest); i {
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
		file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResourceRuleRequest); i {
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
		file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResourceRuleMetadata); i {
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
		file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResourceRuleRequest); i {
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
		file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResourceRuleMetadata); i {
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
			RawDescriptor: file_yandex_cloud_cdn_v1_rule_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_cdn_v1_rule_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_cdn_v1_rule_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_cdn_v1_rule_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_cdn_v1_rule_service_proto = out.File
	file_yandex_cloud_cdn_v1_rule_service_proto_rawDesc = nil
	file_yandex_cloud_cdn_v1_rule_service_proto_goTypes = nil
	file_yandex_cloud_cdn_v1_rule_service_proto_depIdxs = nil
}
