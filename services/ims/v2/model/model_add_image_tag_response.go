package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type AddImageTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddImageTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageTagResponse struct{}"
	}

	return strings.Join([]string{"AddImageTagResponse", string(data)}, " ")
}
