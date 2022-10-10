// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: yandex/cloud/mdb/elasticsearch/v1/auth.proto

package elasticsearch

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type AuthProvider_Type int32

const (
	AuthProvider_TYPE_UNSPECIFIED AuthProvider_Type = 0
	AuthProvider_NATIVE           AuthProvider_Type = 1
	AuthProvider_SAML             AuthProvider_Type = 2
)

// Enum value maps for AuthProvider_Type.
var (
	AuthProvider_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "NATIVE",
		2: "SAML",
	}
	AuthProvider_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"NATIVE":           1,
		"SAML":             2,
	}
)

func (x AuthProvider_Type) Enum() *AuthProvider_Type {
	p := new(AuthProvider_Type)
	*p = x
	return p
}

func (x AuthProvider_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthProvider_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_enumTypes[0].Descriptor()
}

func (AuthProvider_Type) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_enumTypes[0]
}

func (x AuthProvider_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthProvider_Type.Descriptor instead.
func (AuthProvider_Type) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDescGZIP(), []int{1, 0}
}

type AuthProviders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Providers []*AuthProvider `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (x *AuthProviders) Reset() {
	*x = AuthProviders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthProviders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthProviders) ProtoMessage() {}

func (x *AuthProviders) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthProviders.ProtoReflect.Descriptor instead.
func (*AuthProviders) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthProviders) GetProviders() []*AuthProvider {
	if x != nil {
		return x.Providers
	}
	return nil
}

type AuthProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    AuthProvider_Type `protobuf:"varint,1,opt,name=type,proto3,enum=yandex.cloud.mdb.elasticsearch.v1.AuthProvider_Type" json:"type,omitempty"`
	Name    string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Order   int64             `protobuf:"varint,3,opt,name=order,proto3" json:"order,omitempty"`
	Enabled bool              `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// selector ui settings
	Hidden      bool   `protobuf:"varint,5,opt,name=hidden,proto3" json:"hidden,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Hint        string `protobuf:"bytes,7,opt,name=hint,proto3" json:"hint,omitempty"`
	Icon        string `protobuf:"bytes,8,opt,name=icon,proto3" json:"icon,omitempty"`
	// Types that are assignable to Settings:
	//	*AuthProvider_Saml
	Settings isAuthProvider_Settings `protobuf_oneof:"settings"`
}

func (x *AuthProvider) Reset() {
	*x = AuthProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthProvider) ProtoMessage() {}

func (x *AuthProvider) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthProvider.ProtoReflect.Descriptor instead.
func (*AuthProvider) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthProvider) GetType() AuthProvider_Type {
	if x != nil {
		return x.Type
	}
	return AuthProvider_TYPE_UNSPECIFIED
}

func (x *AuthProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthProvider) GetOrder() int64 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *AuthProvider) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *AuthProvider) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *AuthProvider) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AuthProvider) GetHint() string {
	if x != nil {
		return x.Hint
	}
	return ""
}

func (x *AuthProvider) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (m *AuthProvider) GetSettings() isAuthProvider_Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (x *AuthProvider) GetSaml() *SamlSettings {
	if x, ok := x.GetSettings().(*AuthProvider_Saml); ok {
		return x.Saml
	}
	return nil
}

type isAuthProvider_Settings interface {
	isAuthProvider_Settings()
}

type AuthProvider_Saml struct {
	Saml *SamlSettings `protobuf:"bytes,9,opt,name=saml,proto3,oneof"`
}

func (*AuthProvider_Saml) isAuthProvider_Settings() {}

type SamlSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdpEntityId        string `protobuf:"bytes,1,opt,name=idp_entity_id,json=idpEntityId,proto3" json:"idp_entity_id,omitempty"`
	IdpMetadataFile    []byte `protobuf:"bytes,2,opt,name=idp_metadata_file,json=idpMetadataFile,proto3" json:"idp_metadata_file,omitempty"`
	SpEntityId         string `protobuf:"bytes,3,opt,name=sp_entity_id,json=spEntityId,proto3" json:"sp_entity_id,omitempty"`
	KibanaUrl          string `protobuf:"bytes,4,opt,name=kibana_url,json=kibanaUrl,proto3" json:"kibana_url,omitempty"`
	AttributePrincipal string `protobuf:"bytes,5,opt,name=attribute_principal,json=attributePrincipal,proto3" json:"attribute_principal,omitempty"`
	AttributeGroups    string `protobuf:"bytes,6,opt,name=attribute_groups,json=attributeGroups,proto3" json:"attribute_groups,omitempty"`
	AttributeName      string `protobuf:"bytes,7,opt,name=attribute_name,json=attributeName,proto3" json:"attribute_name,omitempty"`
	AttributeEmail     string `protobuf:"bytes,8,opt,name=attribute_email,json=attributeEmail,proto3" json:"attribute_email,omitempty"`
	AttributeDn        string `protobuf:"bytes,9,opt,name=attribute_dn,json=attributeDn,proto3" json:"attribute_dn,omitempty"`
}

func (x *SamlSettings) Reset() {
	*x = SamlSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamlSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamlSettings) ProtoMessage() {}

func (x *SamlSettings) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamlSettings.ProtoReflect.Descriptor instead.
func (*SamlSettings) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SamlSettings) GetIdpEntityId() string {
	if x != nil {
		return x.IdpEntityId
	}
	return ""
}

func (x *SamlSettings) GetIdpMetadataFile() []byte {
	if x != nil {
		return x.IdpMetadataFile
	}
	return nil
}

func (x *SamlSettings) GetSpEntityId() string {
	if x != nil {
		return x.SpEntityId
	}
	return ""
}

func (x *SamlSettings) GetKibanaUrl() string {
	if x != nil {
		return x.KibanaUrl
	}
	return ""
}

func (x *SamlSettings) GetAttributePrincipal() string {
	if x != nil {
		return x.AttributePrincipal
	}
	return ""
}

func (x *SamlSettings) GetAttributeGroups() string {
	if x != nil {
		return x.AttributeGroups
	}
	return ""
}

func (x *SamlSettings) GetAttributeName() string {
	if x != nil {
		return x.AttributeName
	}
	return ""
}

func (x *SamlSettings) GetAttributeEmail() string {
	if x != nil {
		return x.AttributeEmail
	}
	return ""
}

func (x *SamlSettings) GetAttributeDn() string {
	if x != nil {
		return x.AttributeDn
	}
	return ""
}

var File_yandex_cloud_mdb_elasticsearch_v1_auth_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62,
	0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5e, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x4d, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x22, 0xc3, 0x03, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x48, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d,
	0x64, 0x62, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0x8a, 0xc8, 0x31, 0x04, 0x3c,
	0x3d, 0x35, 0x30, 0xf2, 0xc7, 0x31, 0x10, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a,
	0x30, 0x2d, 0x39, 0x5f, 0x2d, 0x5d, 0x2a, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68,
	0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04,
	0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x32, 0x35, 0x30, 0x52, 0x04, 0x68, 0x69, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x32, 0x35, 0x30, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12,
	0x45, 0x0a, 0x04, 0x73, 0x61, 0x6d, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62,
	0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x61, 0x6d, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x48, 0x00,
	0x52, 0x04, 0x73, 0x61, 0x6d, 0x6c, 0x22, 0x32, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x41, 0x4d, 0x4c, 0x10, 0x02, 0x42, 0x0a, 0x0a, 0x08, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xce, 0x03, 0x0a, 0x0c, 0x53, 0x61, 0x6d, 0x6c, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2d, 0x0a, 0x0d, 0x69, 0x64, 0x70, 0x5f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x32, 0x35, 0x30, 0x52, 0x0b, 0x69, 0x64, 0x70, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x11, 0x69, 0x64, 0x70, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x0b, 0x8a, 0xc8, 0x31, 0x07, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x30, 0x30, 0x52, 0x0f,
	0x69, 0x64, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x2b, 0x0a, 0x0c, 0x73, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x32, 0x35, 0x30,
	0x52, 0x0a, 0x73, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0a,
	0x6b, 0x69, 0x62, 0x61, 0x6e, 0x61, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x32, 0x35, 0x30, 0x52, 0x09, 0x6b, 0x69, 0x62,
	0x61, 0x6e, 0x61, 0x55, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x13, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x12, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x12, 0x33, 0x0a, 0x10, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31,
	0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2f, 0x0a, 0x0e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0d, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2b, 0x0a, 0x0c, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x64, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x44, 0x6e, 0x42, 0x7c, 0x0a, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64, 0x62, 0x2e,
	0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31,
	0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDescData = file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDesc
)

func file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDescData
}

var file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_goTypes = []interface{}{
	(AuthProvider_Type)(0), // 0: yandex.cloud.mdb.elasticsearch.v1.AuthProvider.Type
	(*AuthProviders)(nil),  // 1: yandex.cloud.mdb.elasticsearch.v1.AuthProviders
	(*AuthProvider)(nil),   // 2: yandex.cloud.mdb.elasticsearch.v1.AuthProvider
	(*SamlSettings)(nil),   // 3: yandex.cloud.mdb.elasticsearch.v1.SamlSettings
}
var file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.mdb.elasticsearch.v1.AuthProviders.providers:type_name -> yandex.cloud.mdb.elasticsearch.v1.AuthProvider
	0, // 1: yandex.cloud.mdb.elasticsearch.v1.AuthProvider.type:type_name -> yandex.cloud.mdb.elasticsearch.v1.AuthProvider.Type
	3, // 2: yandex.cloud.mdb.elasticsearch.v1.AuthProvider.saml:type_name -> yandex.cloud.mdb.elasticsearch.v1.SamlSettings
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_init() }
func file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_init() {
	if File_yandex_cloud_mdb_elasticsearch_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthProviders); i {
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
		file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthProvider); i {
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
		file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SamlSettings); i {
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
	file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AuthProvider_Saml)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_elasticsearch_v1_auth_proto = out.File
	file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_rawDesc = nil
	file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_goTypes = nil
	file_yandex_cloud_mdb_elasticsearch_v1_auth_proto_depIdxs = nil
}
