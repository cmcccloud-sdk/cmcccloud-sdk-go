package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ModifiyInstanceNameRequest struct {
	Name string `json:"name"`
}

func (o ModifiyInstanceNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifiyInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"ModifiyInstanceNameRequest", string(data)}, " ")
}
