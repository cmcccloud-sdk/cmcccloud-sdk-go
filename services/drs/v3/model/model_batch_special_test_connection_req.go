package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchSpecialTestConnectionReq struct {
	Jobs []BatchJobActionReq `json:"jobs"`
}

func (o BatchSpecialTestConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSpecialTestConnectionReq struct{}"
	}

	return strings.Join([]string{"BatchSpecialTestConnectionReq", string(data)}, " ")
}
