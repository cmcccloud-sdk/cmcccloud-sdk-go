package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteSubnetTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubnetTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubnetTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubnetTagResponse", string(data)}, " ")
}
