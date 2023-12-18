package region

import (
	"fmt"
	"github.com/cmcccloud-sdk/cmcccloud-sdk-go/core/region"
)

var (
	CIDC_RP_11 = region.NewRegion("CIDC-RP-11",
		"https://vpc.cidc-rp-11.joint.cmecloud.cn")
	CIDC_RP_12 = region.NewRegion("CIDC-RP-12",
		"https://vpc.cidc-rp-12.joint.cmecloud.cn")
	CIDC_RP_13 = region.NewRegion("CIDC-RP-13",
		"https://vpc.cidc-rp-13.joint.cmecloud.cn")
	CIDC_RP_19 = region.NewRegion("CIDC-RP-19",
		"https://vpc.cidc-rp-19.joint.cmecloud.cn")
	CIDC_RP_2000 = region.NewRegion("CIDC-RP-2000",
		"https://vpc.cidc-rp-2000.joint.cmecloud.cn")
)

var staticFields = map[string]*region.Region{
	"CIDC-RP-11":   CIDC_RP_11,
	"CIDC-RP-12":   CIDC_RP_12,
	"CIDC-RP-13":   CIDC_RP_13,
	"CIDC-RP-19":   CIDC_RP_19,
	"CIDC-RP-2000": CIDC_RP_2000,
}

var provider = region.DefaultProviderChain("VPC")

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
