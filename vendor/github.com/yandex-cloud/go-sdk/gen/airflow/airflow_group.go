// Code generated by sdkgen. DO NOT EDIT.

package airflow

import (
	"context"

	"google.golang.org/grpc"
)

// Airflow provides access to "airflow" component of Yandex.Cloud
type Airflow struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewAirflow creates instance of Airflow
func NewAirflow(g func(ctx context.Context) (*grpc.ClientConn, error)) *Airflow {
	return &Airflow{g}
}

// Cluster gets ClusterService client
func (a *Airflow) Cluster() *ClusterServiceClient {
	return &ClusterServiceClient{getConn: a.getConn}
}