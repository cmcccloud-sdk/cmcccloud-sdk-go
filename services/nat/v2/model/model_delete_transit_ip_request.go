package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteTransitIpRequest struct {
	TransitIpId string `json:"transit_ip_id"`
}

func (o DeleteTransitIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransitIpRequest struct{}"
	}

	return strings.Join([]string{"DeleteTransitIpRequest", string(data)}, " ")
}
