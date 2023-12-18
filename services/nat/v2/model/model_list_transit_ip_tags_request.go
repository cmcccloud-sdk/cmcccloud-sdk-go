package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ListTransitIpTagsRequest struct {
}

func (o ListTransitIpTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTransitIpTagsRequest", string(data)}, " ")
}
