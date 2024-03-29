package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type RuleRef struct {
	Id string `json:"id"`
}

func (o RuleRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleRef struct{}"
	}

	return strings.Join([]string{"RuleRef", string(data)}, " ")
}
