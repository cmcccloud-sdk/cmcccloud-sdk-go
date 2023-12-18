package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type AddNodeRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *AddNodeList `json:"body,omitempty"`
}

func (o AddNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNodeRequest struct{}"
	}

	return strings.Join([]string{"AddNodeRequest", string(data)}, " ")
}
