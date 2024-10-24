// Code generated by protoc-gen-goext. DO NOT EDIT.

package loadtesting

import (
	agent "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/agent"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *CreateAgentRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateAgentRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateAgentRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateAgentRequest) SetComputeInstanceParams(v *agent.CreateComputeInstance) {
	m.ComputeInstanceParams = v
}

func (m *CreateAgentRequest) SetAgentVersion(v string) {
	m.AgentVersion = v
}

func (m *CreateAgentRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateAgentMetadata) SetAgentId(v string) {
	m.AgentId = v
}

func (m *GetAgentRequest) SetAgentId(v string) {
	m.AgentId = v
}

func (m *DeleteAgentRequest) SetAgentId(v string) {
	m.AgentId = v
}

func (m *DeleteAgentMetadata) SetAgentId(v string) {
	m.AgentId = v
}

func (m *ListAgentsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListAgentsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListAgentsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListAgentsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListAgentsResponse) SetAgents(v []*agent.Agent) {
	m.Agents = v
}

func (m *ListAgentsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *UpdateAgentRequest) SetAgentId(v string) {
	m.AgentId = v
}

func (m *UpdateAgentRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateAgentRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateAgentRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateAgentRequest) SetComputeInstanceParams(v *agent.CreateComputeInstance) {
	m.ComputeInstanceParams = v
}

func (m *UpdateAgentRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateAgentMetadata) SetAgentId(v string) {
	m.AgentId = v
}
