package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ShowUserInstancesRequest struct {
}

func (o ShowUserInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowUserInstancesRequest", string(data)}, " ")
}
