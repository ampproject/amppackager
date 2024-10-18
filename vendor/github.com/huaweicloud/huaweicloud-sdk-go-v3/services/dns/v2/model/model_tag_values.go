package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagValues struct {

	// 键。最大长度36个unicode字符。 key不能为空。(搜索时不对此参数做校验)
	Key *string `json:"key,omitempty"`

	// 值列表。每个值最大长度43个unicode字符。&ast;为系统保留字符。如果里面的value是以&ast;开头时，表示按照&ast;后面的值全模糊匹配。如果values缺失，则表示匹配任意值。value之间为或的关系。
	Values *[]string `json:"values,omitempty"`
}

func (o TagValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagValues struct{}"
	}

	return strings.Join([]string{"TagValues", string(data)}, " ")
}
