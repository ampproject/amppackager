package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecordSetsResponse Response Object
type DeleteRecordSetsResponse struct {

	// Record Set的ID。
	Id *string `json:"id,omitempty"`

	// Record Set的名称。
	Name *string `json:"name,omitempty"`

	// Record Set的描述信息。
	Description *string `json:"description,omitempty"`

	// 托管该记录的zone_id。
	ZoneId *string `json:"zone_id,omitempty"`

	// 托管该记录的zone_name。
	ZoneName *string `json:"zone_name,omitempty"`

	// 记录类型。
	Type *string `json:"type,omitempty"`

	// 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty"`

	// 域名解析后的值。
	Records *[]string `json:"records,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 资源状态。
	Status *string `json:"status,omitempty"`

	// 标识是否由系统默认生成，系统默认生成的Record Set不能删除。
	Default *bool `json:"default,omitempty"`

	// 该Record Set所属的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	// 解析线路ID。
	Line *string `json:"line,omitempty"`

	// 解析记录的权重。
	Weight *int32 `json:"weight,omitempty"`

	// 健康检查ID。
	HealthCheckId *string `json:"health_check_id,omitempty"`

	AliasTarget    *AliasTarget `json:"alias_target,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteRecordSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordSetsResponse struct{}"
	}

	return strings.Join([]string{"DeleteRecordSetsResponse", string(data)}, " ")
}
