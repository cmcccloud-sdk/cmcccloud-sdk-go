package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ReinstallK8sOptionsConfig struct {
	Labels map[string]string `json:"labels,omitempty"`

	Taints *[]Taint `json:"taints,omitempty"`

	MaxPods *int32 `json:"maxPods,omitempty"`

	NicMultiqueue *string `json:"nicMultiqueue,omitempty"`

	NicThreshold *string `json:"nicThreshold,omitempty"`
}

func (o ReinstallK8sOptionsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallK8sOptionsConfig struct{}"
	}

	return strings.Join([]string{"ReinstallK8sOptionsConfig", string(data)}, " ")
}
