package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type GlanceDeleteTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GlanceDeleteTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteTagResponse struct{}"
	}

	return strings.Join([]string{"GlanceDeleteTagResponse", string(data)}, " ")
}
