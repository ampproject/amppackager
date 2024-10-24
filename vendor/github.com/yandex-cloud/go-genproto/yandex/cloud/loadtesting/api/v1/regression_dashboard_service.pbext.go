// Code generated by protoc-gen-goext. DO NOT EDIT.

package loadtesting

import (
	regression "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/regression"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *CreateRegressionDashboardRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateRegressionDashboardRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateRegressionDashboardRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateRegressionDashboardRequest) SetContent(v *regression.Dashboard_Content) {
	m.Content = v
}

func (m *CreateRegressionDashboardMetadata) SetDashboardId(v string) {
	m.DashboardId = v
}

func (m *GetRegressionDashboardRequest) SetDashboardId(v string) {
	m.DashboardId = v
}

func (m *DeleteRegressionDashboardRequest) SetDashboardId(v string) {
	m.DashboardId = v
}

func (m *DeleteRegressionDashboardRequest) SetEtag(v string) {
	m.Etag = v
}

func (m *DeleteRegressionDashboardMetadata) SetDashboardId(v string) {
	m.DashboardId = v
}

func (m *ListRegressionDashboardsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListRegressionDashboardsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListRegressionDashboardsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListRegressionDashboardsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListRegressionDashboardsResponse) SetDashboards(v []*regression.Dashboard) {
	m.Dashboards = v
}

func (m *ListRegressionDashboardsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *UpdateRegressionDashboardRequest) SetDashboardId(v string) {
	m.DashboardId = v
}

func (m *UpdateRegressionDashboardRequest) SetEtag(v string) {
	m.Etag = v
}

func (m *UpdateRegressionDashboardRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateRegressionDashboardRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateRegressionDashboardRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateRegressionDashboardRequest) SetContent(v *regression.Dashboard_Content) {
	m.Content = v
}

func (m *UpdateRegressionDashboardMetadata) SetDashboardId(v string) {
	m.DashboardId = v
}
