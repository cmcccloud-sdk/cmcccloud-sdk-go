package model

import (
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/utils"

	"strings"
)

type NovaCreateKeypairRequestBody struct {
	Keypair *NovaCreateKeypairOption `json:"keypair"`
}

func (o NovaCreateKeypairRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateKeypairRequestBody struct{}"
	}

	return strings.Join([]string{"NovaCreateKeypairRequestBody", string(data)}, " ")
}
