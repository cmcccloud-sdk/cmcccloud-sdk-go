package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ConclusionItem struct {
	Id int32 `json:"id"`

	Params map[string]string `json:"params,omitempty"`
}

func (o ConclusionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConclusionItem struct{}"
	}

	return strings.Join([]string{"ConclusionItem", string(data)}, " ")
}
