// Code generated by protoc-gen-goext. DO NOT EDIT.

package workload

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *FederatedCredential) SetId(v string) {
	m.Id = v
}

func (m *FederatedCredential) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *FederatedCredential) SetFederationId(v string) {
	m.FederationId = v
}

func (m *FederatedCredential) SetExternalSubjectId(v string) {
	m.ExternalSubjectId = v
}

func (m *FederatedCredential) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}