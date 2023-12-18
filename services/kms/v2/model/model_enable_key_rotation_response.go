package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type EnableKeyRotationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableKeyRotationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableKeyRotationResponse struct{}"
	}

	return strings.Join([]string{"EnableKeyRotationResponse", string(data)}, " ")
}
