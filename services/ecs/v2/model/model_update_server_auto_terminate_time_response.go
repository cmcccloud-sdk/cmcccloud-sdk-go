package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type UpdateServerAutoTerminateTimeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServerAutoTerminateTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerAutoTerminateTimeResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerAutoTerminateTimeResponse", string(data)}, " ")
}
