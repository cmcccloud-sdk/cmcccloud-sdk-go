package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateMembersRequest struct {
	PoolId string `json:"pool_id"`

	Body *BatchUpdateMembersRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersRequest", string(data)}, " ")
}
