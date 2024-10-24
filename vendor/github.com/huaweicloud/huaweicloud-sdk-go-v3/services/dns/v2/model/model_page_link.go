package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageLink 指向当前资源或者其他资源的链接。当查询需要分页时，需要包含一个next链接指向下一页。
type PageLink struct {

	// 当前资源的链接。
	Self *string `json:"self,omitempty"`

	// 下一页资源的链接。
	Next *string `json:"next,omitempty"`
}

func (o PageLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageLink struct{}"
	}

	return strings.Join([]string{"PageLink", string(data)}, " ")
}
