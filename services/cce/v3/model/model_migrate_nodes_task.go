package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type MigrateNodesTask struct {
	ApiVersion *string `json:"apiVersion,omitempty"`

	Kind *string `json:"kind,omitempty"`

	Spec *MigrateNodesSpec `json:"spec"`

	Status *TaskStatus `json:"status,omitempty"`
}

func (o MigrateNodesTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateNodesTask struct{}"
	}

	return strings.Join([]string{"MigrateNodesTask", string(data)}, " ")
}
