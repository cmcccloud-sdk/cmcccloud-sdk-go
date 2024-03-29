package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DisassociateServerVirtualIpRequestBody struct {
	Nic *DisassociateServerVirtualIpOption `json:"nic"`
}

func (o DisassociateServerVirtualIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateServerVirtualIpRequestBody struct{}"
	}

	return strings.Join([]string{"DisassociateServerVirtualIpRequestBody", string(data)}, " ")
}
