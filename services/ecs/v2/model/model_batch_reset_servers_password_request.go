package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchResetServersPasswordRequest struct {
	Body *BatchResetServersPasswordRequestBody `json:"body,omitempty"`
}

func (o BatchResetServersPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResetServersPasswordRequest struct{}"
	}

	return strings.Join([]string{"BatchResetServersPasswordRequest", string(data)}, " ")
}
