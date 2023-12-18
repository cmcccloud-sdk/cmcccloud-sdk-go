package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ListTransitIpsByTagsResponse struct {
	Resources *[]Resource `json:"resources,omitempty"`

	RequestId *string `json:"request_id,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTransitIpsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTransitIpsByTagsResponse", string(data)}, " ")
}
