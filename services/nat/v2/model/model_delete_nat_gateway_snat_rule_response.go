package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteNatGatewaySnatRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNatGatewaySnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewaySnatRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewaySnatRuleResponse", string(data)}, " ")
}
