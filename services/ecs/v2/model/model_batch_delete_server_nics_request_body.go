package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteServerNicsRequestBody struct {
	Nics []BatchDeleteServerNicOption `json:"nics"`
}

func (o BatchDeleteServerNicsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerNicsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerNicsRequestBody", string(data)}, " ")
}
