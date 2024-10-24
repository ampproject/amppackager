// Code generated by protoc-gen-goext. DO NOT EDIT.

package licensemanager

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Instance) SetId(v string) {
	m.Id = v
}

func (m *Instance) SetCloudId(v string) {
	m.CloudId = v
}

func (m *Instance) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Instance) SetTemplateId(v string) {
	m.TemplateId = v
}

func (m *Instance) SetTemplateVersionId(v string) {
	m.TemplateVersionId = v
}

func (m *Instance) SetDescription(v string) {
	m.Description = v
}

func (m *Instance) SetStartTime(v *timestamppb.Timestamp) {
	m.StartTime = v
}

func (m *Instance) SetEndTime(v *timestamppb.Timestamp) {
	m.EndTime = v
}

func (m *Instance) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Instance) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *Instance) SetState(v Instance_State) {
	m.State = v
}

func (m *Instance) SetLocks(v []*Lock) {
	m.Locks = v
}

func (m *Instance) SetLicenseTemplate(v *Template) {
	m.LicenseTemplate = v
}