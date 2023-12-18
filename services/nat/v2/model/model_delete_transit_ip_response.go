package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteTransitIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTransitIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransitIpResponse struct{}"
	}

	return strings.Join([]string{"DeleteTransitIpResponse", string(data)}, " ")
}
