package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteServerMetadataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerMetadataResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerMetadataResponse", string(data)}, " ")
}
