package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ResetServerPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetServerPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetServerPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetServerPasswordResponse", string(data)}, " ")
}
