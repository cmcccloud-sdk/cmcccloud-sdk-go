package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type UpdatePoolRequestBody struct {
	Pool *UpdatePoolOption `json:"pool"`
}

func (o UpdatePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePoolRequestBody", string(data)}, " ")
}
