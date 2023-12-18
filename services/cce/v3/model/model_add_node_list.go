package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type AddNodeList struct {
	ApiVersion string `json:"apiVersion"`

	Kind string `json:"kind"`

	NodeList []AddNode `json:"nodeList"`
}

func (o AddNodeList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNodeList struct{}"
	}

	return strings.Join([]string{"AddNodeList", string(data)}, " ")
}
