package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteSingleInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSingleInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSingleInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteSingleInstanceResponse", string(data)}, " ")
}
