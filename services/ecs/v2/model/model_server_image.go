package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ServerImage struct {
	Id string `json:"id"`
}

func (o ServerImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerImage struct{}"
	}

	return strings.Join([]string{"ServerImage", string(data)}, " ")
}
