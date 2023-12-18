package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type UpdateKeyDescriptionResponse struct {
	KeyInfo        *KeyDescriptionInfo `json:"key_info,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateKeyDescriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyDescriptionResponse struct{}"
	}

	return strings.Join([]string{"UpdateKeyDescriptionResponse", string(data)}, " ")
}
