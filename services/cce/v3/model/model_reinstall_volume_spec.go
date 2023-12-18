package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ReinstallVolumeSpec struct {
	ImageID *string `json:"imageID,omitempty"`

	CmkID *string `json:"cmkID,omitempty"`
}

func (o ReinstallVolumeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallVolumeSpec struct{}"
	}

	return strings.Join([]string{"ReinstallVolumeSpec", string(data)}, " ")
}
