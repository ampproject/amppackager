// Code generated by sdkgen. DO NOT EDIT.

// nolint
package agent

import (
	"context"

	"google.golang.org/grpc"

	agent "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/agent/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// AgentRegistrationServiceClient is a agent.AgentRegistrationServiceClient with
// lazy GRPC connection initialization.
type AgentRegistrationServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// ExternalAgentRegister implements agent.AgentRegistrationServiceClient
func (c *AgentRegistrationServiceClient) ExternalAgentRegister(ctx context.Context, in *agent.ExternalAgentRegisterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return agent.NewAgentRegistrationServiceClient(conn).ExternalAgentRegister(ctx, in, opts...)
}

// Register implements agent.AgentRegistrationServiceClient
func (c *AgentRegistrationServiceClient) Register(ctx context.Context, in *agent.RegisterRequest, opts ...grpc.CallOption) (*agent.RegisterResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return agent.NewAgentRegistrationServiceClient(conn).Register(ctx, in, opts...)
}