package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ResetNodeList struct {
	ApiVersion string `json:"apiVersion"`

	Kind string `json:"kind"`

	NodeList []ResetNode `json:"nodeList"`
}

func (o ResetNodeList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetNodeList struct{}"
	}

	return strings.Join([]string{"ResetNodeList", string(data)}, " ")
}
