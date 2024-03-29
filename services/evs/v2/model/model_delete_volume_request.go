package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteVolumeRequest struct {
	VolumeId string `json:"volume_id"`
}

func (o DeleteVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVolumeRequest struct{}"
	}

	return strings.Join([]string{"DeleteVolumeRequest", string(data)}, " ")
}
