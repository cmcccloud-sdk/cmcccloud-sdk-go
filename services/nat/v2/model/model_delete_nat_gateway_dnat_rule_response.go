package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteNatGatewayDnatRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNatGatewayDnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewayDnatRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewayDnatRuleResponse", string(data)}, " ")
}
