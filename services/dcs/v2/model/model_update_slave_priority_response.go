package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type UpdateSlavePriorityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSlavePriorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlavePriorityResponse struct{}"
	}

	return strings.Join([]string{"UpdateSlavePriorityResponse", string(data)}, " ")
}
