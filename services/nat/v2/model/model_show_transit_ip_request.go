package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ShowTransitIpRequest struct {
	TransitIpId string `json:"transit_ip_id"`
}

func (o ShowTransitIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransitIpRequest struct{}"
	}

	return strings.Join([]string{"ShowTransitIpRequest", string(data)}, " ")
}
