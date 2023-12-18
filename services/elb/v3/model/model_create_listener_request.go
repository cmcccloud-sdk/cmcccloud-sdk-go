package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type CreateListenerRequest struct {
	Body *CreateListenerRequestBody `json:"body,omitempty"`
}

func (o CreateListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerRequest struct{}"
	}

	return strings.Join([]string{"CreateListenerRequest", string(data)}, " ")
}
