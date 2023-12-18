package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type CinderDeleteVolumeTransferResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CinderDeleteVolumeTransferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderDeleteVolumeTransferResponse struct{}"
	}

	return strings.Join([]string{"CinderDeleteVolumeTransferResponse", string(data)}, " ")
}
