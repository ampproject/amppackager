// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: yandex/cloud/oauth/claims.proto

package oauth

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

// Claims representation, see https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims for details.
type SubjectClaims struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Subject - Identifier for the End-User at the Issuer.
	Sub string `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	// End-User's full name in displayable form including all name parts, possibly including titles and suffixes, ordered according to the End-User's locale and preferences.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Given name(s) or first name(s) of the End-User. Note that in some cultures, people can have multiple given names; all can be present, with the names being separated by space characters.
	GivenName string `protobuf:"bytes,3,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	// Surname(s) or last name(s) of the End-User. Note that in some cultures, people can have multiple family names or no family name; all can be present, with the names being separated by space characters.
	FamilyName string `protobuf:"bytes,4,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	// Shorthand name by which the End-User wishes to be referred to at the RP, such as janedoe or j.doe.
	// This value MAY be any valid JSON string including special characters such as @, /, or whitespace. The RP MUST NOT rely upon this value being unique, as discussed in Section 5.7.
	PreferredUsername string `protobuf:"bytes,7,opt,name=preferred_username,json=preferredUsername,proto3" json:"preferred_username,omitempty"`
	// URL of the End-User's profile picture. This URL MUST refer to an image file (for example, a PNG, JPEG, or GIF image file),
	// rather than to a Web page containing an image. Note that this URL SHOULD specifically reference a profile photo of the End-User suitable for displaying when describing the End-User, rather than an arbitrary photo taken by the End-User.
	Picture string `protobuf:"bytes,9,opt,name=picture,proto3" json:"picture,omitempty"`
	// End-User's preferred e-mail address. Its value MUST conform to the RFC 5322 [RFC5322] addr-spec syntax.
	// The RP MUST NOT rely upon this value being unique, as discussed in Section 5.7.
	Email string `protobuf:"bytes,11,opt,name=email,proto3" json:"email,omitempty"`
	// String from zoneinfo [zoneinfo] time zone database representing the End-User's time zone. For example, Europe/Paris or America/Los_Angeles.
	Zoneinfo string `protobuf:"bytes,15,opt,name=zoneinfo,proto3" json:"zoneinfo,omitempty"`
	// End-User's locale, represented as a BCP47 [RFC5646] language tag. This is typically an ISO 639-1 Alpha-2 [ISO639-1] language code in lowercase and an ISO 3166-1 Alpha-2 [ISO3166-1] country code in uppercase, separated by a dash.
	// For example, en-US or fr-CA. As a compatibility note, some implementations have used an underscore as the separator rather than a dash, for example, en_US; Relying Parties MAY choose to accept this locale syntax as well.
	Locale string `protobuf:"bytes,16,opt,name=locale,proto3" json:"locale,omitempty"`
	// End-User's preferred telephone number. E.164 [E.164] is RECOMMENDED as the format of this Claim, for example, +1 (425) 555-1212 or +56 (2) 687 2400.
	// If the phone number contains an extension, it is RECOMMENDED that the extension be represented using the RFC 3966 [RFC3966] extension syntax, for example, +1 (604) 555-1234;ext=5678.
	PhoneNumber string `protobuf:"bytes,17,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// User federation, non-empty only for federated users.
	Federation *Federation `protobuf:"bytes,100,opt,name=federation,proto3" json:"federation,omitempty"`
}

func (x *SubjectClaims) Reset() {
	*x = SubjectClaims{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_oauth_claims_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubjectClaims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubjectClaims) ProtoMessage() {}

func (x *SubjectClaims) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_oauth_claims_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubjectClaims.ProtoReflect.Descriptor instead.
func (*SubjectClaims) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_oauth_claims_proto_rawDescGZIP(), []int{0}
}

func (x *SubjectClaims) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *SubjectClaims) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubjectClaims) GetGivenName() string {
	if x != nil {
		return x.GivenName
	}
	return ""
}

func (x *SubjectClaims) GetFamilyName() string {
	if x != nil {
		return x.FamilyName
	}
	return ""
}

func (x *SubjectClaims) GetPreferredUsername() string {
	if x != nil {
		return x.PreferredUsername
	}
	return ""
}

func (x *SubjectClaims) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *SubjectClaims) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SubjectClaims) GetZoneinfo() string {
	if x != nil {
		return x.Zoneinfo
	}
	return ""
}

func (x *SubjectClaims) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *SubjectClaims) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *SubjectClaims) GetFederation() *Federation {
	if x != nil {
		return x.Federation
	}
	return nil
}

// Minimalistic analog of yandex.cloud.organizationmanager.v1.saml.Federation
type Federation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the federation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the federation. The name is unique within the cloud or organization
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Federation) Reset() {
	*x = Federation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_oauth_claims_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Federation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Federation) ProtoMessage() {}

func (x *Federation) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_oauth_claims_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Federation.ProtoReflect.Descriptor instead.
func (*Federation) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_oauth_claims_proto_rawDescGZIP(), []int{1}
}

func (x *Federation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Federation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_yandex_cloud_oauth_claims_proto protoreflect.FileDescriptor

var file_yandex_cloud_oauth_claims_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x02, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35,
	0x30, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69,
	0x76, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x67, 0x69, 0x76, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x7a, 0x6f, 0x6e,
	0x65, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x7a, 0x6f, 0x6e,
	0x65, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x3e, 0x0a, 0x0a, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x3e, 0x0a, 0x0a, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01,
	0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x59, 0x0a, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_oauth_claims_proto_rawDescOnce sync.Once
	file_yandex_cloud_oauth_claims_proto_rawDescData = file_yandex_cloud_oauth_claims_proto_rawDesc
)

func file_yandex_cloud_oauth_claims_proto_rawDescGZIP() []byte {
	file_yandex_cloud_oauth_claims_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_oauth_claims_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_oauth_claims_proto_rawDescData)
	})
	return file_yandex_cloud_oauth_claims_proto_rawDescData
}

var file_yandex_cloud_oauth_claims_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_oauth_claims_proto_goTypes = []interface{}{
	(*SubjectClaims)(nil), // 0: yandex.cloud.oauth.SubjectClaims
	(*Federation)(nil),    // 1: yandex.cloud.oauth.Federation
}
var file_yandex_cloud_oauth_claims_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.oauth.SubjectClaims.federation:type_name -> yandex.cloud.oauth.Federation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_oauth_claims_proto_init() }
func file_yandex_cloud_oauth_claims_proto_init() {
	if File_yandex_cloud_oauth_claims_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_oauth_claims_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubjectClaims); i {
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
		file_yandex_cloud_oauth_claims_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Federation); i {
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
			RawDescriptor: file_yandex_cloud_oauth_claims_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_oauth_claims_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_oauth_claims_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_oauth_claims_proto_msgTypes,
	}.Build()
	File_yandex_cloud_oauth_claims_proto = out.File
	file_yandex_cloud_oauth_claims_proto_rawDesc = nil
	file_yandex_cloud_oauth_claims_proto_goTypes = nil
	file_yandex_cloud_oauth_claims_proto_depIdxs = nil
}