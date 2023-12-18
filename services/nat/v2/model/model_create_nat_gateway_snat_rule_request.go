package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type CreateNatGatewaySnatRuleRequest struct {
	Body *CreateNatGatewaySnatRuleRequestOption `json:"body,omitempty"`
}

func (o CreateNatGatewaySnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewaySnatRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateNatGatewaySnatRuleRequest", string(data)}, " ")
}
