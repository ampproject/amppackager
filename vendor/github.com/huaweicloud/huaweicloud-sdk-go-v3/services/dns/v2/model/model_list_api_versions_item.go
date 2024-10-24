package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListApiVersionsItem struct {

	// 版本状态，包含：  CURRENT：表示该版本为主推版本。 SUPPORTED：表示为老版本，但是现在还在继续支持。 DEPRECATED：表示为废弃版本，存在后续删除的可能。
	Status *string `json:"status,omitempty"`

	// 版本号。
	Id *string `json:"id,omitempty"`

	// 指向当前版本的url。
	Links *[]LinksItem `json:"links,omitempty"`
}

func (o ListApiVersionsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionsItem struct{}"
	}

	return strings.Join([]string{"ListApiVersionsItem", string(data)}, " ")
}
