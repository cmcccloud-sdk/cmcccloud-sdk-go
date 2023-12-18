package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeletePrivateDnatResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateDnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateDnatResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateDnatResponse", string(data)}, " ")
}
