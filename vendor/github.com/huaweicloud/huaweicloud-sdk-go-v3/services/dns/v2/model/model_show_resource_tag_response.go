package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceTagResponse Response Object
type ShowResourceTagResponse struct {

	// 指定实例的标签列表。
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceTagResponse", string(data)}, " ")
}
