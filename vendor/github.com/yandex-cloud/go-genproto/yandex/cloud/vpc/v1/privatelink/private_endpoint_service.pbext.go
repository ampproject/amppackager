// Code generated by protoc-gen-goext. DO NOT EDIT.

package privatelink

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetPrivateEndpointRequest) SetPrivateEndpointId(v string) {
	m.PrivateEndpointId = v
}

func (m *InternalIpv4AddressSpec) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *InternalIpv4AddressSpec) SetAddress(v string) {
	m.Address = v
}

type AddressSpec_Address = isAddressSpec_Address

func (m *AddressSpec) SetAddress(v AddressSpec_Address) {
	m.Address = v
}

func (m *AddressSpec) SetAddressId(v string) {
	m.Address = &AddressSpec_AddressId{
		AddressId: v,
	}
}

func (m *AddressSpec) SetInternalIpv4AddressSpec(v *InternalIpv4AddressSpec) {
	m.Address = &AddressSpec_InternalIpv4AddressSpec{
		InternalIpv4AddressSpec: v,
	}
}

type CreatePrivateEndpointRequest_Service = isCreatePrivateEndpointRequest_Service

func (m *CreatePrivateEndpointRequest) SetService(v CreatePrivateEndpointRequest_Service) {
	m.Service = v
}

func (m *CreatePrivateEndpointRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreatePrivateEndpointRequest) SetName(v string) {
	m.Name = v
}

func (m *CreatePrivateEndpointRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreatePrivateEndpointRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreatePrivateEndpointRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *CreatePrivateEndpointRequest) SetAddressSpec(v *AddressSpec) {
	m.AddressSpec = v
}

func (m *CreatePrivateEndpointRequest) SetDnsOptions(v *PrivateEndpoint_DnsOptions) {
	m.DnsOptions = v
}

func (m *CreatePrivateEndpointRequest) SetObjectStorage(v *PrivateEndpoint_ObjectStorage) {
	m.Service = &CreatePrivateEndpointRequest_ObjectStorage{
		ObjectStorage: v,
	}
}

func (m *CreatePrivateEndpointMetadata) SetPrivateEndpointId(v string) {
	m.PrivateEndpointId = v
}

func (m *UpdatePrivateEndpointRequest) SetPrivateEndpointId(v string) {
	m.PrivateEndpointId = v
}

func (m *UpdatePrivateEndpointRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdatePrivateEndpointRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdatePrivateEndpointRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdatePrivateEndpointRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdatePrivateEndpointRequest) SetAddressSpec(v *AddressSpec) {
	m.AddressSpec = v
}

func (m *UpdatePrivateEndpointRequest) SetDnsOptions(v *PrivateEndpoint_DnsOptions) {
	m.DnsOptions = v
}

func (m *UpdatePrivateEndpointMetadata) SetPrivateEndpointId(v string) {
	m.PrivateEndpointId = v
}

func (m *DeletePrivateEndpointRequest) SetPrivateEndpointId(v string) {
	m.PrivateEndpointId = v
}

func (m *DeletePrivateEndpointMetadata) SetPrivateEndpointId(v string) {
	m.PrivateEndpointId = v
}

type ListPrivateEndpointsRequest_Container = isListPrivateEndpointsRequest_Container

func (m *ListPrivateEndpointsRequest) SetContainer(v ListPrivateEndpointsRequest_Container) {
	m.Container = v
}

func (m *ListPrivateEndpointsRequest) SetFolderId(v string) {
	m.Container = &ListPrivateEndpointsRequest_FolderId{
		FolderId: v,
	}
}

func (m *ListPrivateEndpointsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListPrivateEndpointsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListPrivateEndpointsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListPrivateEndpointsResponse) SetPrivateEndpoints(v []*PrivateEndpoint) {
	m.PrivateEndpoints = v
}

func (m *ListPrivateEndpointsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListPrivateEndpointOperationsRequest) SetPrivateEndpointId(v string) {
	m.PrivateEndpointId = v
}

func (m *ListPrivateEndpointOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListPrivateEndpointOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListPrivateEndpointOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListPrivateEndpointOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
