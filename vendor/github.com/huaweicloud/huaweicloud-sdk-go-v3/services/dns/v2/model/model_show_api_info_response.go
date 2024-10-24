package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApiInfoResponse Response Object
type ShowApiInfoResponse struct {
	Version        *VersionItem `json:"version,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowApiInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowApiInfoResponse", string(data)}, " ")
}
