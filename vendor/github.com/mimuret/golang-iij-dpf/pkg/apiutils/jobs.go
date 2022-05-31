package apiutils

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"path"
	"strconv"
	"time"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
)

func WaitJob(ctx context.Context, c api.ClientInterface, jobID string, interval time.Duration) (*core.Job, error) {
	job := &core.Job{
		RequestID: jobID,
	}
	if _, err := c.Read(ctx, job); err != nil {
		return nil, fmt.Errorf("failed to read Job: %w", err)
	}
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()
	for job.Status == core.JobStatusRunning {
		job.RequestID = jobID
		if err := c.WatchRead(ctx, interval, job); err != nil {
			return nil, err
		}
	}
	if job.Status == core.JobStatusFailed {
		return job, fmt.Errorf("JobID %s job failed: type: %s msg: %s", jobID, job.ErrorType, job.ErrorMessage)
	}
	return job, nil
}

func ParseeResourceSystemID(job *core.Job) (string, error) {
	u, err := url.Parse(job.ResourceUrl)
	if err != nil {
		return "", fmt.Errorf("failed to parse resource-url: %s , %w", job.ResourceUrl, err)
	}
	_, id := path.Split(u.Path)
	return id, nil
}

func ParseeResourceID(job *core.Job) (int64, error) {
	idStr, err := ParseeResourceSystemID(job)
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert to int64 resource-url: %s id: %s, %w", job.ResourceUrl, idStr, err)
	}
	return id, nil
}
