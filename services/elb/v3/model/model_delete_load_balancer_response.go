package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteLoadBalancerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadBalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerResponse struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerResponse", string(data)}, " ")
}
