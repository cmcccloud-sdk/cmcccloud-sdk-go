package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateNatGatewayDnatRulesRequest struct {
	Body *BatchCreateNatGatewayDnatRulesRequestBody `json:"body,omitempty"`
}

func (o BatchCreateNatGatewayDnatRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateNatGatewayDnatRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateNatGatewayDnatRulesRequest", string(data)}, " ")
}
