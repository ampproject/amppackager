package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLineGroupRequest Request Object
type CreateLineGroupRequest struct {
	Body *CreateLineGroupsReq `json:"body,omitempty"`
}

func (o CreateLineGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLineGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateLineGroupRequest", string(data)}, " ")
}
