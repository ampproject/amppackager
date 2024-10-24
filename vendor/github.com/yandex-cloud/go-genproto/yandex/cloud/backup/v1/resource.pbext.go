// Code generated by protoc-gen-goext. DO NOT EDIT.

package backup

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Resource) SetComputeInstanceId(v string) {
	m.ComputeInstanceId = v
}

func (m *Resource) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Resource) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *Resource) SetOnline(v bool) {
	m.Online = v
}

func (m *Resource) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *Resource) SetStatus(v Resource_Status) {
	m.Status = v
}

func (m *Resource) SetStatusDetails(v string) {
	m.StatusDetails = v
}

func (m *Resource) SetStatusProgress(v int64) {
	m.StatusProgress = v
}

func (m *Resource) SetLastBackupTime(v *timestamppb.Timestamp) {
	m.LastBackupTime = v
}

func (m *Resource) SetNextBackupTime(v *timestamppb.Timestamp) {
	m.NextBackupTime = v
}

func (m *Resource) SetResourceId(v string) {
	m.ResourceId = v
}

func (m *Resource) SetIsActive(v bool) {
	m.IsActive = v
}

func (m *Resource) SetInitStatus(v Resource_InitStatus) {
	m.InitStatus = v
}

func (m *Resource) SetMetadata(v string) {
	m.Metadata = v
}

func (m *Resource) SetType(v ResourceType) {
	m.Type = v
}

func (m *Progress) SetCurrent(v int64) {
	m.Current = v
}

func (m *Progress) SetTotal(v int64) {
	m.Total = v
}

func (m *Task) SetId(v int64) {
	m.Id = v
}

func (m *Task) SetCancellable(v bool) {
	m.Cancellable = v
}

func (m *Task) SetPolicyId(v string) {
	m.PolicyId = v
}

func (m *Task) SetType(v Task_Type) {
	m.Type = v
}

func (m *Task) SetProgress(v *Progress) {
	m.Progress = v
}

func (m *Task) SetStatus(v Task_Status) {
	m.Status = v
}

func (m *Task) SetEnqueuedAt(v *timestamppb.Timestamp) {
	m.EnqueuedAt = v
}

func (m *Task) SetStartedAt(v *timestamppb.Timestamp) {
	m.StartedAt = v
}

func (m *Task) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *Task) SetCompletedAt(v *timestamppb.Timestamp) {
	m.CompletedAt = v
}

func (m *Task) SetComputeInstanceId(v string) {
	m.ComputeInstanceId = v
}

func (m *Task) SetResultCode(v Task_Code) {
	m.ResultCode = v
}

func (m *Task) SetError(v string) {
	m.Error = v
}
