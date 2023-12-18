package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ReinstallNodeSpec struct {
	Os string `json:"os"`

	Login *Login `json:"login"`

	Name *string `json:"name,omitempty"`

	ServerConfig *ReinstallServerConfig `json:"serverConfig,omitempty"`

	VolumeConfig *ReinstallVolumeConfig `json:"volumeConfig,omitempty"`

	RuntimeConfig *ReinstallRuntimeConfig `json:"runtimeConfig,omitempty"`

	K8sOptions *ReinstallK8sOptionsConfig `json:"k8sOptions,omitempty"`

	Lifecycle *NodeLifecycleConfig `json:"lifecycle,omitempty"`

	InitializedConditions *[]string `json:"initializedConditions,omitempty"`

	ExtendParam *ReinstallExtendParam `json:"extendParam,omitempty"`
}

func (o ReinstallNodeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallNodeSpec struct{}"
	}

	return strings.Join([]string{"ReinstallNodeSpec", string(data)}, " ")
}
