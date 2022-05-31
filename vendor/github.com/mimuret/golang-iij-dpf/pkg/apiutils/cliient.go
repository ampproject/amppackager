package apiutils

import (
	"context"
	"time"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
)

func SyncUpdate(ctx context.Context, cl api.ClientInterface, s api.Spec, body interface{}) (string, *core.Job, error) {
	requestID, err := cl.Update(ctx, s, body)
	if err != nil {
		return requestID, nil, err
	}
	job, err := WaitJob(ctx, cl, requestID, time.Second)
	return requestID, job, err
}

func SyncCreate(ctx context.Context, cl api.ClientInterface, s api.Spec, body interface{}) (string, *core.Job, error) {
	requestID, err := cl.Create(ctx, s, body)
	if err != nil {
		return requestID, nil, err
	}
	job, err := WaitJob(ctx, cl, requestID, time.Second)
	return requestID, job, err
}

func SyncApply(ctx context.Context, cl api.ClientInterface, s api.Spec, body interface{}) (string, *core.Job, error) {
	requestID, err := cl.Apply(ctx, s, body)
	if err != nil {
		return requestID, nil, err
	}
	job, err := WaitJob(ctx, cl, requestID, time.Second)
	return requestID, job, err
}

func SyncDelete(ctx context.Context, cl api.ClientInterface, s api.Spec) (string, *core.Job, error) {
	requestID, err := cl.Delete(ctx, s)
	if err != nil {
		return requestID, nil, err
	}
	job, err := WaitJob(ctx, cl, requestID, time.Second)
	return requestID, job, err
}

func SyncCancel(ctx context.Context, cl api.ClientInterface, s api.Spec) (string, *core.Job, error) {
	requestID, err := cl.Cancel(ctx, s)
	if err != nil {
		return requestID, nil, err
	}
	job, err := WaitJob(ctx, cl, requestID, time.Second)
	return requestID, job, err
}
