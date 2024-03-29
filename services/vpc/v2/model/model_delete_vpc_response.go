package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteVpcResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpcResponse", string(data)}, " ")
}
