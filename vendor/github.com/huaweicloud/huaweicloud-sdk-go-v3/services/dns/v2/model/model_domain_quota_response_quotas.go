package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainQuotaResponseQuotas struct {

	// 资源类型。
	QuotaKey string `json:"quota_key"`

	// 资源配额的最大值。
	QuotaLimit int32 `json:"quota_limit"`

	// 配额已使用数量。
	Used int32 `json:"used"`

	// 配额统计单位，取固定值“count”。
	Unit string `json:"unit"`
}

func (o DomainQuotaResponseQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainQuotaResponseQuotas struct{}"
	}

	return strings.Join([]string{"DomainQuotaResponseQuotas", string(data)}, " ")
}
