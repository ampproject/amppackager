package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPtrRecordsRequest Request Object
type ListPtrRecordsRequest struct {

	// 分页查询起始的资源ID，为空时为查询第一页。  默认值为空。
	Marker *string `json:"marker,omitempty"`

	// 每页返回的资源个数。  取值范围：0~500  取值一般为10，20，50。默认值为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。  取值范围：0~2147483647  默认值为0。  当前设置marker不为空时，以marker为分页起始标识。
	Offset *int32 `json:"offset,omitempty"`

	// 反向解析关联的企业项目ID，长度不超过36个字符。  默认值为0。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源标签。  取值格式：key1,value1|key2,value2  多个标签之间用\"|\"分开，每个标签的键值用英文逗号\",\"相隔。
	Tags *string `json:"tags,omitempty"`

	// 资源状态。
	Status *string `json:"status,omitempty"`
}

func (o ListPtrRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPtrRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListPtrRecordsRequest", string(data)}, " ")
}
