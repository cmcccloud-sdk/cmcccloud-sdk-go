package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type DeleteLoadBalancerForceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadBalancerForceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerForceResponse struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerForceResponse", string(data)}, " ")
}
