package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchModifyPwdReq struct {
	Jobs []ModifyPwdEndPoint `json:"jobs"`
}

func (o BatchModifyPwdReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyPwdReq struct{}"
	}

	return strings.Join([]string{"BatchModifyPwdReq", string(data)}, " ")
}
