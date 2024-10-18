package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRSetBatchLinesReq struct {

	// 后缀需以Zone Name结束且为FQDN（即以“.”号结束的完整主机名）。
	Name string `json:"name"`

	// 可选配置，对域名的描述。 长度不超过255个字符。
	Description *string `json:"description,omitempty"`

	// Record Set的类型。 取值范围：A,AAAA,MX,CNAME,TXT,NS,SRV,CAA。
	Type string `json:"type"`

	// 解析线路域名参数。最多支持50个。
	Lines []BatchCreateRecordSetWithLine `json:"lines"`
}

func (o CreateRSetBatchLinesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRSetBatchLinesReq struct{}"
	}

	return strings.Join([]string{"CreateRSetBatchLinesReq", string(data)}, " ")
}
