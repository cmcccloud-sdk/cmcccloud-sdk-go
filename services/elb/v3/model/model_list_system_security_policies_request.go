package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type ListSystemSecurityPoliciesRequest struct {
}

func (o ListSystemSecurityPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemSecurityPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListSystemSecurityPoliciesRequest", string(data)}, " ")
}
