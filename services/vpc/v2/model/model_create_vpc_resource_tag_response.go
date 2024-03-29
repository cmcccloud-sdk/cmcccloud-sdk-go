package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type CreateVpcResourceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateVpcResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcResourceTagResponse struct{}"
	}

	return strings.Join([]string{"CreateVpcResourceTagResponse", string(data)}, " ")
}
