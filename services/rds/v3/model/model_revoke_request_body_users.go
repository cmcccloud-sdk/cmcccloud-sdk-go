package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type RevokeRequestBodyUsers struct {
	Name string `json:"name"`
}

func (o RevokeRequestBodyUsers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeRequestBodyUsers struct{}"
	}

	return strings.Join([]string{"RevokeRequestBodyUsers", string(data)}, " ")
}
