package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePublicZoneRequest Request Object
type CreatePublicZoneRequest struct {
	Body *CreatePublicZoneReq `json:"body,omitempty"`
}

func (o CreatePublicZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicZoneRequest struct{}"
	}

	return strings.Join([]string{"CreatePublicZoneRequest", string(data)}, " ")
}
