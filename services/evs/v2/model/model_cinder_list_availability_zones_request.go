package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type CinderListAvailabilityZonesRequest struct {
}

func (o CinderListAvailabilityZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListAvailabilityZonesRequest struct{}"
	}

	return strings.Join([]string{"CinderListAvailabilityZonesRequest", string(data)}, " ")
}
