package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ImportKeyMaterialResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ImportKeyMaterialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportKeyMaterialResponse struct{}"
	}

	return strings.Join([]string{"ImportKeyMaterialResponse", string(data)}, " ")
}
