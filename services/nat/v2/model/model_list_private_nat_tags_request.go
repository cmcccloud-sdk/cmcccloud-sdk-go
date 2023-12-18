package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ListPrivateNatTagsRequest struct {
}

func (o ListPrivateNatTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNatTagsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateNatTagsRequest", string(data)}, " ")
}
