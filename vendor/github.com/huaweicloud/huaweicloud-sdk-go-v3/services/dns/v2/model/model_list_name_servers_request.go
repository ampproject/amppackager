package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNameServersRequest Request Object
type ListNameServersRequest struct {

	// 待查询名称服务器的类型。 取值范围: public, private。 如果为空，表示查询所有类型的名称服务器。 如果为public，表示查询公网的名称服务器。 如果为private，表示查询内网的名称服务器。 搜索模式精确搜索。 默认值为空。
	Type *string `json:"type,omitempty"`

	// 待查询的region ID。 当查询公网的名称服务器时，此处不填。 搜索模式精确搜索。 默认值为空。
	Region *string `json:"region,omitempty"`
}

func (o ListNameServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNameServersRequest struct{}"
	}

	return strings.Join([]string{"ListNameServersRequest", string(data)}, " ")
}
