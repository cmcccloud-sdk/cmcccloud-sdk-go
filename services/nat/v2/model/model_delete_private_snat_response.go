package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeletePrivateSnatResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateSnatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateSnatResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateSnatResponse", string(data)}, " ")
}
