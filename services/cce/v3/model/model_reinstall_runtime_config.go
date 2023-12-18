package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ReinstallRuntimeConfig struct {
	DockerBaseSize *int32 `json:"dockerBaseSize,omitempty"`

	Runtime *Runtime `json:"runtime,omitempty"`
}

func (o ReinstallRuntimeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallRuntimeConfig struct{}"
	}

	return strings.Join([]string{"ReinstallRuntimeConfig", string(data)}, " ")
}
