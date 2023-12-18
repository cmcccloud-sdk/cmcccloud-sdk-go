package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchQueryJobReq struct {
	Jobs []string `json:"jobs"`
}

func (o BatchQueryJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryJobReq struct{}"
	}

	return strings.Join([]string{"BatchQueryJobReq", string(data)}, " ")
}
