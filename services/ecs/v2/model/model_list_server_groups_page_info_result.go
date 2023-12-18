package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ListServerGroupsPageInfoResult struct {
	NextMarker *string `json:"next_marker,omitempty"`
}

func (o ListServerGroupsPageInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerGroupsPageInfoResult struct{}"
	}

	return strings.Join([]string{"ListServerGroupsPageInfoResult", string(data)}, " ")
}
