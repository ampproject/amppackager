package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateRecordSet struct {

	// RecordSet资源ID。
	Id string `json:"id"`

	// RecordSet资源描述。
	Description *string `json:"description,omitempty"`

	// Record Set的有效缓存时间，以秒为单位。 取值范围：300-2147483647。 默认值为300s。
	Ttl *int32 `json:"ttl,omitempty"`

	// 解析记录的权重，默认为null。 当weight=null时，表示该解析记录不设置权重。 当weight=0，表示备用域名解析记录。 当weight>0，表示主用域名解析记录。 取值范围：0~100 在相同域名、类型、线路下的解析记录，规则如下： 全部设置权重，或全部不设置权重。 当不设置权重时，只能创建一个解析记录。 当设置权重时，最多能创建20个解析记录。
	Weight *int32 `json:"weight,omitempty"`

	// 解析记录的值。不同类型解析记录对应的值的规则不同。
	Records []string `json:"records"`
}

func (o BatchUpdateRecordSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateRecordSet struct{}"
	}

	return strings.Join([]string{"BatchUpdateRecordSet", string(data)}, " ")
}
