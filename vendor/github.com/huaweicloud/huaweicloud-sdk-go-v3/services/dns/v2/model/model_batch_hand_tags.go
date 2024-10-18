package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchHandTags struct {

	// 标签列表。删除时tags结构体不能缺失。
	Tags []Tag `json:"tags"`

	// 操作标识（区分大小写）：create（创建）、delete（删除）。
	Action string `json:"action"`
}

func (o BatchHandTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchHandTags struct{}"
	}

	return strings.Join([]string{"BatchHandTags", string(data)}, " ")
}
