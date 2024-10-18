package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisassociaterouterRequestBody struct {
	Router *Router `json:"router"`
}

func (o DisassociaterouterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociaterouterRequestBody struct{}"
	}

	return strings.Join([]string{"DisassociaterouterRequestBody", string(data)}, " ")
}
