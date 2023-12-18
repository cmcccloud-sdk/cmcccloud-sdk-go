package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchModifyJobReq struct {
	Jobs []ModifyJobReq `json:"jobs"`
}

func (o BatchModifyJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyJobReq struct{}"
	}

	return strings.Join([]string{"BatchModifyJobReq", string(data)}, " ")
}
