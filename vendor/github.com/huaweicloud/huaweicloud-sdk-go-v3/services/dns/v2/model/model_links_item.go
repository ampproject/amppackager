package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LinksItem 指向当前资源或者其他资源的链接。当查询需要分页时，需要包含一个next链接指向下一页。
type LinksItem struct {

	// 对应快捷链接。
	Href string `json:"href"`

	// 快捷链接标记名称。
	Rel string `json:"rel"`
}

func (o LinksItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksItem struct{}"
	}

	return strings.Join([]string{"LinksItem", string(data)}, " ")
}
