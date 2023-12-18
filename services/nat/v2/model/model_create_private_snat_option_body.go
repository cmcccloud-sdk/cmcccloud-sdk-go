package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type CreatePrivateSnatOptionBody struct {
	SnatRule *CreatePrivateSnatOption `json:"snat_rule"`
}

func (o CreatePrivateSnatOptionBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateSnatOptionBody struct{}"
	}

	return strings.Join([]string{"CreatePrivateSnatOptionBody", string(data)}, " ")
}
