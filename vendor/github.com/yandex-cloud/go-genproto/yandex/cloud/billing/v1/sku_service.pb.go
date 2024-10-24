// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/billing/v1/sku_service.proto

package billing

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type GetSkuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the SKU to return.
	// To get the SKU ID, use [SkuService.List] request.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Currency of the SKU.
	// Can be one of the following:
	// * `RUB`
	// * `USD`
	// * `KZT`
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	// Optional ID of the billing account.
	// If specified, contract prices for a particular billing account are included in the response.
	// To get the billing account ID, use [BillingAccountService.List] request.
	BillingAccountId string `protobuf:"bytes,3,opt,name=billing_account_id,json=billingAccountId,proto3" json:"billing_account_id,omitempty"`
}

func (x *GetSkuRequest) Reset() {
	*x = GetSkuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_sku_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSkuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSkuRequest) ProtoMessage() {}

func (x *GetSkuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_sku_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSkuRequest.ProtoReflect.Descriptor instead.
func (*GetSkuRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_sku_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetSkuRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetSkuRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *GetSkuRequest) GetBillingAccountId() string {
	if x != nil {
		return x.BillingAccountId
	}
	return ""
}

type ListSkusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Currency of the prices.
	// Can be one of the following:
	// * `RUB`
	// * `USD`
	// * `KZT`
	Currency string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	// Optional ID of the billing account.
	// If specified, contract prices for a particular billing account are included in the response.
	// To get the billing account ID, use [BillingAccountService.List] request.
	BillingAccountId string `protobuf:"bytes,2,opt,name=billing_account_id,json=billingAccountId,proto3" json:"billing_account_id,omitempty"`
	// A filter expression that filters resources listed in the response.
	// The expression must specify:
	// 1. The field name. Currently you can use filtering only on the [yandex.cloud.billing.v1.Sku.id] and [yandex.cloud.billing.v1.Sku.service_id] field.
	// 2. An `=` operator.
	// 3. The value in double quotes (`"`). Must be 3-63 characters long and match the regular expression `[a-z][-a-z0-9]{1,61}[a-z0-9]`.
	// E.g. `filter=serviceId="dn28hpu6268356q0j8mk"`.
	Filter string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListSkusResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results,
	// set [page_token] to the [ListSkusResponse.next_page_token]
	// returned by a previous list request.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListSkusRequest) Reset() {
	*x = ListSkusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_sku_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSkusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSkusRequest) ProtoMessage() {}

func (x *ListSkusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_sku_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSkusRequest.ProtoReflect.Descriptor instead.
func (*ListSkusRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_sku_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListSkusRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ListSkusRequest) GetBillingAccountId() string {
	if x != nil {
		return x.BillingAccountId
	}
	return ""
}

func (x *ListSkusRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListSkusRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSkusRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListSkusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of skus.
	Skus []*Sku `protobuf:"bytes,1,rep,name=skus,proto3" json:"skus,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListSkusRequest.page_size], use
	// [next_page_token] as the value
	// for the [ListSkusRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListSkusResponse) Reset() {
	*x = ListSkusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_sku_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSkusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSkusResponse) ProtoMessage() {}

func (x *ListSkusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_sku_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSkusResponse.ProtoReflect.Descriptor instead.
func (*ListSkusResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_sku_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListSkusResponse) GetSkus() []*Sku {
	if x != nil {
		return x.Skus
	}
	return nil
}

func (x *ListSkusResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_billing_v1_sku_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_billing_v1_sku_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6b, 0x75, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6b, 0x75, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x75, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0xd8, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x8a, 0xc8, 0x31, 0x06, 0x3c, 0x3d, 0x31,
	0x30, 0x30, 0x30, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a,
	0xfa, 0xc7, 0x31, 0x06, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x30, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d,
	0x31, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6c,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x6b, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6b, 0x75, 0x52, 0x04,
	0x73, 0x6b, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xef, 0x01, 0x0a,
	0x0a, 0x53, 0x6b, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x26, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x6b, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6b, 0x75, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x12, 0x15, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6b,
	0x75, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x75, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6b, 0x75, 0x73, 0x42, 0x62,
	0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_billing_v1_sku_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_billing_v1_sku_service_proto_rawDescData = file_yandex_cloud_billing_v1_sku_service_proto_rawDesc
)

func file_yandex_cloud_billing_v1_sku_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_billing_v1_sku_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_billing_v1_sku_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_billing_v1_sku_service_proto_rawDescData)
	})
	return file_yandex_cloud_billing_v1_sku_service_proto_rawDescData
}

var file_yandex_cloud_billing_v1_sku_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_billing_v1_sku_service_proto_goTypes = []interface{}{
	(*GetSkuRequest)(nil),    // 0: yandex.cloud.billing.v1.GetSkuRequest
	(*ListSkusRequest)(nil),  // 1: yandex.cloud.billing.v1.ListSkusRequest
	(*ListSkusResponse)(nil), // 2: yandex.cloud.billing.v1.ListSkusResponse
	(*Sku)(nil),              // 3: yandex.cloud.billing.v1.Sku
}
var file_yandex_cloud_billing_v1_sku_service_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.billing.v1.ListSkusResponse.skus:type_name -> yandex.cloud.billing.v1.Sku
	0, // 1: yandex.cloud.billing.v1.SkuService.Get:input_type -> yandex.cloud.billing.v1.GetSkuRequest
	1, // 2: yandex.cloud.billing.v1.SkuService.List:input_type -> yandex.cloud.billing.v1.ListSkusRequest
	3, // 3: yandex.cloud.billing.v1.SkuService.Get:output_type -> yandex.cloud.billing.v1.Sku
	2, // 4: yandex.cloud.billing.v1.SkuService.List:output_type -> yandex.cloud.billing.v1.ListSkusResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_billing_v1_sku_service_proto_init() }
func file_yandex_cloud_billing_v1_sku_service_proto_init() {
	if File_yandex_cloud_billing_v1_sku_service_proto != nil {
		return
	}
	file_yandex_cloud_billing_v1_sku_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_billing_v1_sku_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSkuRequest); i {
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
		file_yandex_cloud_billing_v1_sku_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSkusRequest); i {
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
		file_yandex_cloud_billing_v1_sku_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSkusResponse); i {
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
			RawDescriptor: file_yandex_cloud_billing_v1_sku_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_billing_v1_sku_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_billing_v1_sku_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_billing_v1_sku_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_billing_v1_sku_service_proto = out.File
	file_yandex_cloud_billing_v1_sku_service_proto_rawDesc = nil
	file_yandex_cloud_billing_v1_sku_service_proto_goTypes = nil
	file_yandex_cloud_billing_v1_sku_service_proto_depIdxs = nil
}
