package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type NovaListAvailabilityZonesRequest struct {
}

func (o NovaListAvailabilityZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListAvailabilityZonesRequest struct{}"
	}

	return strings.Join([]string{"NovaListAvailabilityZonesRequest", string(data)}, " ")
}
