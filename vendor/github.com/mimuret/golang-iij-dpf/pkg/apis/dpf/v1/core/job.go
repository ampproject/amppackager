package core

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

type JobStatus string

var _ apis.Spec = &Job{}

const (
	JobStatusRunning    JobStatus = "RUNNING"
	JobStatusSuccessful JobStatus = "SUCCESSFUL"
	JobStatusFailed     JobStatus = "FAILED"
)

func (c JobStatus) String() string { return string(c) }

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type Job struct {
	AttributeMeta
	RequestID    string    `read:"request_id"`
	Status       JobStatus `read:"status"`
	ResourceUrl  string    `read:"resources_url"`
	ErrorType    string    `read:"error_type"`
	ErrorMessage string    `read:"error_message"`
}

func (c *Job) GetError() error {
	if c.Status == JobStatusFailed {
		return fmt.Errorf("ErrorType: %s Messages %s", c.ErrorType, c.ErrorMessage)
	}
	return nil
}

func (c *Job) GetName() string { return "jobs" }
func (c *Job) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionRead {
		return action.ToMethod(), fmt.Sprintf("/jobs/%s", c.RequestID)
	}
	return "", ""
}

func (c *Job) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.RequestID)
}

func init() {
	register(&Job{})
}
