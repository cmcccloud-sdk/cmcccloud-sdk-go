package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchPrecheckReq struct {
	Jobs []PreCheckInfo `json:"jobs"`
}

func (o BatchPrecheckReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPrecheckReq struct{}"
	}

	return strings.Join([]string{"BatchPrecheckReq", string(data)}, " ")
}
