package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type UpdatePortRequestBody struct {
	Port *UpdatePortOption `json:"port"`
}

func (o UpdatePortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePortRequestBody", string(data)}, " ")
}
