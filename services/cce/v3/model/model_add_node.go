package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type AddNode struct {
	ServerID string `json:"serverID"`

	Spec *ReinstallNodeSpec `json:"spec"`
}

func (o AddNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNode struct{}"
	}

	return strings.Join([]string{"AddNode", string(data)}, " ")
}
