package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type RemoveNodesSpec struct {
	Login *Login `json:"login"`

	Nodes []NodeItem `json:"nodes"`
}

func (o RemoveNodesSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveNodesSpec struct{}"
	}

	return strings.Join([]string{"RemoveNodesSpec", string(data)}, " ")
}
