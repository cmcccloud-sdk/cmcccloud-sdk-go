package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ReinstallServerConfig struct {
	UserTags *[]UserTag `json:"userTags,omitempty"`

	RootVolume *ReinstallVolumeSpec `json:"rootVolume,omitempty"`
}

func (o ReinstallServerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerConfig struct{}"
	}

	return strings.Join([]string{"ReinstallServerConfig", string(data)}, " ")
}
