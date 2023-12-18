package region

import (
	"fmt"
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/region"
)

var (
	CIDC_RP_12 = region.NewRegion("CIDC-RP-12",
		"https://nat.cidc-rp-12.joint.cmecloud.cn")
)

var staticFields = map[string]*region.Region{
	"CIDC-RP-12": CIDC_RP_12,
}

var provider = region.DefaultProviderChain("NAT")

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
