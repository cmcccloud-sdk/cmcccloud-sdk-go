package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ListAvailableZonesRequest struct {
}

func (o ListAvailableZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesRequest", string(data)}, " ")
}
