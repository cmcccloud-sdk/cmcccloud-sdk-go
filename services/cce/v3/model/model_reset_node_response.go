package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ResetNodeResponse struct {
	Jobid          *string `json:"jobid,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetNodeResponse struct{}"
	}

	return strings.Join([]string{"ResetNodeResponse", string(data)}, " ")
}
