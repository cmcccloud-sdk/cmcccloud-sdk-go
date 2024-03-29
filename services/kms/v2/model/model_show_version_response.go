package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ShowVersionResponse struct {
	Version        *interface{} `json:"version,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowVersionResponse", string(data)}, " ")
}
