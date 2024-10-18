package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NameServersResp struct {

	// 待查询名称服务器的类型。  取值范围: public, private。  如果为空，表示查询所有类型的名称服务器。 如果为public，表示查询公网的名称服务器。  如果为private，表示查询内网的名称服务器。
	Type *string `json:"type,omitempty"`

	// 待查询的region ID。  当查询公网的名称服务器时，此处不填。
	Region *string `json:"region,omitempty"`

	// 名称服务器列表。
	NsRecords *[]NsRecords `json:"ns_records,omitempty"`
}

func (o NameServersResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NameServersResp struct{}"
	}

	return strings.Join([]string{"NameServersResp", string(data)}, " ")
}
