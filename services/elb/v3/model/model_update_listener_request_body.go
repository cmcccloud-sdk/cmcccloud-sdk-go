package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type UpdateListenerRequestBody struct {
	Listener *UpdateListenerOption `json:"listener"`
}

func (o UpdateListenerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateListenerRequestBody", string(data)}, " ")
}
