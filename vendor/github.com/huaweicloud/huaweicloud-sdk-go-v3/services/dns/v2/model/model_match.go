package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Match struct {

	// 键。当前值限定为resource_name。
	Key string `json:"key"`

	// 值。每个值最大长度255个unicode字符。不能包含“_”,“%”特殊字符。
	Value *string `json:"value,omitempty"`
}

func (o Match) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}
