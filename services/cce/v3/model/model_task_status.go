package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type TaskStatus struct {
	JobID *string `json:"jobID,omitempty"`
}

func (o TaskStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskStatus struct{}"
	}

	return strings.Join([]string{"TaskStatus", string(data)}, " ")
}
