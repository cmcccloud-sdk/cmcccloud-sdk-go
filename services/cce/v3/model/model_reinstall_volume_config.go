package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ReinstallVolumeConfig struct {
	LvmConfig *string `json:"lvmConfig,omitempty"`

	Storage *Storage `json:"storage,omitempty"`
}

func (o ReinstallVolumeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallVolumeConfig struct{}"
	}

	return strings.Join([]string{"ReinstallVolumeConfig", string(data)}, " ")
}
