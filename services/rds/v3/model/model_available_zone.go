package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type AvailableZone struct {
	Code string `json:"code"`

	Description string `json:"description"`
}

func (o AvailableZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableZone struct{}"
	}

	return strings.Join([]string{"AvailableZone", string(data)}, " ")
}
