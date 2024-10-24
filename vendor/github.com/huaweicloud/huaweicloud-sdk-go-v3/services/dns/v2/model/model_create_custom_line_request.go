package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomLineRequest Request Object
type CreateCustomLineRequest struct {
	Body *CreateCustomLines `json:"body,omitempty"`
}

func (o CreateCustomLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomLineRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomLineRequest", string(data)}, " ")
}
