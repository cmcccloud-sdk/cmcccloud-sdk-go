package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type GlanceShowImageRequest struct {
	ImageId string `json:"image_id"`
}

func (o GlanceShowImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageRequest struct{}"
	}

	return strings.Join([]string{"GlanceShowImageRequest", string(data)}, " ")
}
