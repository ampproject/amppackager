package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateZoneRequest Request Object
type CreatePrivateZoneRequest struct {
	Body *CreatePrivateZoneReq `json:"body,omitempty"`
}

func (o CreatePrivateZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateZoneRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateZoneRequest", string(data)}, " ")
}
