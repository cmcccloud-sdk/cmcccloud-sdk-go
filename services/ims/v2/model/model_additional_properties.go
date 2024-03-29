package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type AdditionalProperties struct {
	Type string `json:"type"`
}

func (o AdditionalProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdditionalProperties struct{}"
	}

	return strings.Join([]string{"AdditionalProperties", string(data)}, " ")
}
