package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type RestoreExistingInstanceRequestBody struct {
	Source *RestoreExistingInstanceRequestBodySource `json:"source"`

	Target *TargetInstanceRequest `json:"target"`
}

func (o RestoreExistingInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreExistingInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreExistingInstanceRequestBody", string(data)}, " ")
}
