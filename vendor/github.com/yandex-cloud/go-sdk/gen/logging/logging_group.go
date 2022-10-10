// Code generated by sdkgen. DO NOT EDIT.

package logging

import (
	"context"

	"google.golang.org/grpc"
)

// Logging provides access to "logging" component of Yandex.Cloud
type Logging struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewLogging creates instance of Logging
func NewLogging(g func(ctx context.Context) (*grpc.ClientConn, error)) *Logging {
	return &Logging{g}
}

// LogGroup gets LogGroupService client
func (l *Logging) LogGroup() *LogGroupServiceClient {
	return &LogGroupServiceClient{getConn: l.getConn}
}
