package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type AllowedAddressPair struct {
	IpAddress string `json:"ip_address"`

	MacAddress *string `json:"mac_address,omitempty"`
}

func (o AllowedAddressPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowedAddressPair struct{}"
	}

	return strings.Join([]string{"AllowedAddressPair", string(data)}, " ")
}
