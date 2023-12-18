package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ResetNode struct {
	NodeID string `json:"nodeID"`

	Spec *ReinstallNodeSpec `json:"spec"`
}

func (o ResetNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetNode struct{}"
	}

	return strings.Join([]string{"ResetNode", string(data)}, " ")
}
