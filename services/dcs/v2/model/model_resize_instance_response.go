package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ResizeInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResizeInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceResponse struct{}"
	}

	return strings.Join([]string{"ResizeInstanceResponse", string(data)}, " ")
}
