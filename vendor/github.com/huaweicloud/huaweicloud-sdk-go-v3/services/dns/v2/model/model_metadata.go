package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Metadata 返回满足过滤条件的资源总数。
type Metadata struct {

	// 满足查询条件的资源总数，不受分页（即limit、offset参数）影响。
	TotalCount *int32 `json:"total_count,omitempty"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}
