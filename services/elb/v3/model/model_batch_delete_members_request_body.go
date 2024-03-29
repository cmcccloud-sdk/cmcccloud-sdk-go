package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteMembersRequestBody struct {
	Members []BatchDeleteMembersOption `json:"members"`
}

func (o BatchDeleteMembersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersRequestBody", string(data)}, " ")
}
