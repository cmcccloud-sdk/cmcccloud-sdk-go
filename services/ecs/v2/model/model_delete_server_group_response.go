package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteServerGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupResponse", string(data)}, " ")
}
