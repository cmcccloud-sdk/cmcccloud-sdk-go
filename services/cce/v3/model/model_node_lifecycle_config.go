package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type NodeLifecycleConfig struct {
	PreInstall *string `json:"preInstall,omitempty"`

	PostInstall *string `json:"postInstall,omitempty"`
}

func (o NodeLifecycleConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeLifecycleConfig struct{}"
	}

	return strings.Join([]string{"NodeLifecycleConfig", string(data)}, " ")
}
