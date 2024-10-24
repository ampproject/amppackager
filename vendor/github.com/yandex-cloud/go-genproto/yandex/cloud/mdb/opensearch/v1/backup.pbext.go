// Code generated by protoc-gen-goext. DO NOT EDIT.

package opensearch

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Backup) SetId(v string) {
	m.Id = v
}

func (m *Backup) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Backup) SetSourceClusterId(v string) {
	m.SourceClusterId = v
}

func (m *Backup) SetStartedAt(v *timestamppb.Timestamp) {
	m.StartedAt = v
}

func (m *Backup) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Backup) SetIndices(v []string) {
	m.Indices = v
}

func (m *Backup) SetOpensearchVersion(v string) {
	m.OpensearchVersion = v
}

func (m *Backup) SetSizeBytes(v int64) {
	m.SizeBytes = v
}

func (m *Backup) SetIndicesTotal(v int64) {
	m.IndicesTotal = v
}