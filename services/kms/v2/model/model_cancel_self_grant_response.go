package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type CancelSelfGrantResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelSelfGrantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSelfGrantResponse struct{}"
	}

	return strings.Join([]string{"CancelSelfGrantResponse", string(data)}, " ")
}
