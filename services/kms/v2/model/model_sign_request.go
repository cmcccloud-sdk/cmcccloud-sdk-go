package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type SignRequest struct {
	Body *SignRequestBody `json:"body,omitempty"`
}

func (o SignRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignRequest struct{}"
	}

	return strings.Join([]string{"SignRequest", string(data)}, " ")
}
