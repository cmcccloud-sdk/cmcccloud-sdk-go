package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteKeyStoreResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteKeyStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeyStoreResponse struct{}"
	}

	return strings.Join([]string{"DeleteKeyStoreResponse", string(data)}, " ")
}
