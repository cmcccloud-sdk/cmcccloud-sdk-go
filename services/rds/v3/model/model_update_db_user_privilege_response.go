package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type UpdateDbUserPrivilegeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDbUserPrivilegeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbUserPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"UpdateDbUserPrivilegeResponse", string(data)}, " ")
}
