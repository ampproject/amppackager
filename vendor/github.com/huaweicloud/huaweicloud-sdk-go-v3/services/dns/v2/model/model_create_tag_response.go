package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTagResponse Response Object
type CreateTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagResponse struct{}"
	}

	return strings.Join([]string{"CreateTagResponse", string(data)}, " ")
}
