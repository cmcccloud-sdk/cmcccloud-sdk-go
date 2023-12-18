package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ReinstallExtendParam struct {
	AlphaCceNodeImageID *string `json:"alpha.cce/NodeImageID,omitempty"`
}

func (o ReinstallExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallExtendParam struct{}"
	}

	return strings.Join([]string{"ReinstallExtendParam", string(data)}, " ")
}
