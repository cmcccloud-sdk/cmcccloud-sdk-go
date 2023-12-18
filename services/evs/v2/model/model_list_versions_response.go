package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ListVersionsResponse struct {
	Versions       *[]Versions `json:"versions,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListVersionsResponse", string(data)}, " ")
}