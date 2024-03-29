package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteServerNicOption struct {
	Id string `json:"id"`
}

func (o BatchDeleteServerNicOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerNicOption struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerNicOption", string(data)}, " ")
}
