package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ShowQuotaOfTenantRequest struct {
}

func (o ShowQuotaOfTenantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaOfTenantRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotaOfTenantRequest", string(data)}, " ")
}
