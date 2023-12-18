package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type CreateRoutetableReqBody struct {
	Routetable *CreateRouteTableReq `json:"routetable"`
}

func (o CreateRoutetableReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoutetableReqBody struct{}"
	}

	return strings.Join([]string{"CreateRoutetableReqBody", string(data)}, " ")
}
