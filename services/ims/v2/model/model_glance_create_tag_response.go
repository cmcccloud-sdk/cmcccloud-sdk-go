package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type GlanceCreateTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GlanceCreateTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceCreateTagResponse struct{}"
	}

	return strings.Join([]string{"GlanceCreateTagResponse", string(data)}, " ")
}
