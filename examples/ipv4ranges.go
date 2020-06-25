package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
	"strconv"
)

func main() {
	var (
		ClientOV    *ov.OVClient
		id        = "TestFCNetworkGOsdk"

	)
	apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
		"*")
	ipV4Range := ov.ipV4Range{
		type:         "Range",
		name:         "IPv4",
		startAddress: "10.10.1.1",
		endAddress:   "10.10.1.254",
		subnetUri:    "/rest/id-pools/ipv4/subnets/daa36872-03b1-463b-aaf7-09d58b650142"
	}

	fmt.Println(ipV4Range)
	err := ovc.CreateIPv4Range(ipV4Range)
	if err != nil {
		fmt.Println("IPv4 Range Creation Failed: ", err)
	} else {
		fmt.Println("IPv4 Range created successfully...")
	}

	IpV4RangeSchema, err := ovc.GetIpV4RangeSchema()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get IpV4RangeSchema ----#")
		fmt.Println(IpV4RangeSchema)
	}
	ipv4Range_by_id, err := ovc.GetIPv4RangebyId(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get FCNetworks by name----------------#")
		fmt.Println(ipv4Range_by_id)
	}
	ipv4Range_by_id.type = "Range"
	ipv4Range_by_id.enabled = "true"
	err = ovc.UpdateIpv4Range(ipv4Range_by_id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Ipv4Range has been updated with type " + ipv4Range_by_id.type)
	}
	err = ovc.DeleteIpv4Range(ipv4Range_by_id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleted Ipv4Range successfully...")
	}
	allocatedFragments, err := ovc.GetAllocatedFragments("", "", "", "", ipv4Range_by_id.id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get allocatedFragments ----#")
		for i := range allocatedFragments.Members {
			fmt.Println(allocatedFragments.Members[i].Name)
		}
	}
	freeFragments, err := ovc.GetFreeFragments("", "", "", "", ipv4Range_by_id.id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get freeGegments ----#")
		for i := range freeFragments.Members {
			fmt.Println(freeFragments.Members[i].Name)
		}
	}
    idList = ["",""]
	err = ovc.UpdateAllocator(ipv4Range_by_id.id, idList)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Allocated "+ ipv4Range_by_id.ids + "successfully")
	}
	idList = ["",""]
	err = ovc.UpdateCollector(ipv4Range_by_id.id, idList)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Collected "+ ipv4Range_by_id.ids + "successfully")
	}
}
