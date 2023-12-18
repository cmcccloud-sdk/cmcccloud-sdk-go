package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateDeleteTransitIpTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateDeleteTransitIpTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteTransitIpTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteTransitIpTagsResponse", string(data)}, " ")
}
