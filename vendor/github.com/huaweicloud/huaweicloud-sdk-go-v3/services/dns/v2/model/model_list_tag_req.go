package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTagReq struct {

	// 包含标签。 最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。
	Tags *[]TagValues `json:"tags,omitempty"`

	// 最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。
	TagsAny *[]TagValues `json:"tags_any,omitempty"`

	// 最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。
	NotTags *[]TagValues `json:"not_tags,omitempty"`

	// 最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。
	NotTagsAny *[]TagValues `json:"not_tags_any,omitempty"`

	// 每页返回的资源个数。  取值范围：1~1000  参数取值说明：  如果action为filter时，默认为1000。 如果action为count时，无此参数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。  取值范围：0~2147483647  默认值为0。  参数取值说明： 查询第一页数据时，不需要传入此参数。 查询后续页码数据时，将查询前一页数据时响应体中的值带入此参数。 如果action为filter时，默认为0，必须为数字，不能为负数。 如果action为count时，无此参数。
	Offset *int32 `json:"offset,omitempty"`

	// 操作标识（区分大小写）。  取值范围：  filter：分页过滤查询 count：查询总条数
	Action string `json:"action"`

	// key为要匹配的字段，value为匹配的值。  如果value为空字符串则精确匹配，否则模糊匹配。
	Matches *[]Match `json:"matches,omitempty"`
}

func (o ListTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagReq struct{}"
	}

	return strings.Join([]string{"ListTagReq", string(data)}, " ")
}
