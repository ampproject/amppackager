package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestorePtrReq struct {

	// PTR记录对应的域名。  此处值为null。
	Ptrdname *interface{} `json:"ptrdname"`
}

func (o RestorePtrReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorePtrReq struct{}"
	}

	return strings.Join([]string{"RestorePtrReq", string(data)}, " ")
}
