package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainQuotaResponse Response Object
type ShowDomainQuotaResponse struct {

	// 配额项数据。
	Quotas         *[]DomainQuotaResponseQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowDomainQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainQuotaResponse", string(data)}, " ")
}
