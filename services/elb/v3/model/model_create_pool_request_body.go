package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type CreatePoolRequestBody struct {
	Pool *CreatePoolOption `json:"pool"`
}

func (o CreatePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePoolRequestBody", string(data)}, " ")
}
