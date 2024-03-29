package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type CreateSecurityGroupRequest struct {
	Body *CreateSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o CreateSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRequest", string(data)}, " ")
}
