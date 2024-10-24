// Code generated by sdkgen. DO NOT EDIT.

// nolint
package api

import (
	"context"

	"google.golang.org/grpc"

	api "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1"
)

//revive:disable

// ReportServiceClient is a api.ReportServiceClient with
// lazy GRPC connection initialization.
type ReportServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// CalculateKpiValues implements api.ReportServiceClient
func (c *ReportServiceClient) CalculateKpiValues(ctx context.Context, in *api.CalculateReportKpiValuesRequest, opts ...grpc.CallOption) (*api.CalculateReportKpiValuesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewReportServiceClient(conn).CalculateKpiValues(ctx, in, opts...)
}

// GetTable implements api.ReportServiceClient
func (c *ReportServiceClient) GetTable(ctx context.Context, in *api.GetTableReportRequest, opts ...grpc.CallOption) (*api.GetTableReportResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewReportServiceClient(conn).GetTable(ctx, in, opts...)
}