package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ListenerRef struct {
	Id string `json:"id"`
}

func (o ListenerRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerRef struct{}"
	}

	return strings.Join([]string{"ListenerRef", string(data)}, " ")
}
