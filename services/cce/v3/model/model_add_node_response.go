package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type AddNodeResponse struct {
	Jobid          *string `json:"jobid,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNodeResponse struct{}"
	}

	return strings.Join([]string{"AddNodeResponse", string(data)}, " ")
}
