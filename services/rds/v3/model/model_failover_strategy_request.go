package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type FailoverStrategyRequest struct {
	RepairStrategy string `json:"repairStrategy"`
}

func (o FailoverStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailoverStrategyRequest struct{}"
	}

	return strings.Join([]string{"FailoverStrategyRequest", string(data)}, " ")
}
