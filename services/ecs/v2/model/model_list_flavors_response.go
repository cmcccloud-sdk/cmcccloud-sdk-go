package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ListFlavorsResponse struct {
	Flavors        *[]Flavor `json:"flavors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}
