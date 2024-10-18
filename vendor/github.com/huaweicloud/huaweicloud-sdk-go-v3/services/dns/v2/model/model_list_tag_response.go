package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagResponse Response Object
type ListTagResponse struct {

	// 标签资源信息列表。
	Resources *[]ResourceItem `json:"resources,omitempty"`

	// 资源总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagResponse struct{}"
	}

	return strings.Join([]string{"ListTagResponse", string(data)}, " ")
}
