package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type UpdateKeyRotationIntervalResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateKeyRotationIntervalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyRotationIntervalResponse struct{}"
	}

	return strings.Join([]string{"UpdateKeyRotationIntervalResponse", string(data)}, " ")
}
