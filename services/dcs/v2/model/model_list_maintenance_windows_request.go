package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ListMaintenanceWindowsRequest struct {
}

func (o ListMaintenanceWindowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMaintenanceWindowsRequest struct{}"
	}

	return strings.Join([]string{"ListMaintenanceWindowsRequest", string(data)}, " ")
}
