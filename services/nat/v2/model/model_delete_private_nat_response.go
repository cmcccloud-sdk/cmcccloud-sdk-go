package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeletePrivateNatResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateNatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateNatResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateNatResponse", string(data)}, " ")
}
