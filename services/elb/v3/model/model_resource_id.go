package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ResourceId struct {
	Id string `json:"id"`
}

func (o ResourceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceId struct{}"
	}

	return strings.Join([]string{"ResourceId", string(data)}, " ")
}
