package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ResetNodeRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *ResetNodeList `json:"body,omitempty"`
}

func (o ResetNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetNodeRequest struct{}"
	}

	return strings.Join([]string{"ResetNodeRequest", string(data)}, " ")
}
