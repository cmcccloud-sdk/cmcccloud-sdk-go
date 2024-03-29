package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchQueryRpoAndRtoReq struct {
	Jobs []string `json:"jobs"`
}

func (o BatchQueryRpoAndRtoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryRpoAndRtoReq struct{}"
	}

	return strings.Join([]string{"BatchQueryRpoAndRtoReq", string(data)}, " ")
}
