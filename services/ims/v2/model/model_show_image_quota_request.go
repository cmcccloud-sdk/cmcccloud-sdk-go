package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ShowImageQuotaRequest struct {
}

func (o ShowImageQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowImageQuotaRequest", string(data)}, " ")
}
