package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ListTransitIpTagsResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Tags           *[]Tags `json:"tags,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTransitIpTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTransitIpTagsResponse", string(data)}, " ")
}
