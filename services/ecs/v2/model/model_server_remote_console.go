package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ServerRemoteConsole struct {
	Protocol string `json:"protocol"`

	Type string `json:"type"`

	Url string `json:"url"`
}

func (o ServerRemoteConsole) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerRemoteConsole struct{}"
	}

	return strings.Join([]string{"ServerRemoteConsole", string(data)}, " ")
}
