package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchAddOrDeleteTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddOrDeleteTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrDeleteTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddOrDeleteTagsResponse", string(data)}, " ")
}
