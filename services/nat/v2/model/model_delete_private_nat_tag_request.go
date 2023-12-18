package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeletePrivateNatTagRequest struct {
	Key string `json:"key"`

	ResourceId string `json:"resource_id"`
}

func (o DeletePrivateNatTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateNatTagRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivateNatTagRequest", string(data)}, " ")
}
