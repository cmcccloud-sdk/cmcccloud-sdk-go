package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type AddImageTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o AddImageTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageTagRequestBody struct{}"
	}

	return strings.Join([]string{"AddImageTagRequestBody", string(data)}, " ")
}
