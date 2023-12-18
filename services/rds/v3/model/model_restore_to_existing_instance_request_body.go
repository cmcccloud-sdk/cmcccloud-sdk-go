package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type RestoreToExistingInstanceRequestBody struct {
	Source *RestoreToExistingInstanceRequestBodySource `json:"source"`

	Target *RestoreToExistingInstanceRequestBodyTarget `json:"target"`
}

func (o RestoreToExistingInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreToExistingInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreToExistingInstanceRequestBody", string(data)}, " ")
}
