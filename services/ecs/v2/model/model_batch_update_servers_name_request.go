package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateServersNameRequest struct {
	Body *BatchUpdateServersNameRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateServersNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateServersNameRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateServersNameRequest", string(data)}, " ")
}
