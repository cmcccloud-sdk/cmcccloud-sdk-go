package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteIpGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteIpGroupResponse", string(data)}, " ")
}
