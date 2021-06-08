package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
)

func main() {
	var (
		ClientOV *ov.OVClient
		id       = "b1b869f8-3d5a-4d4a-b0a2-fb6634f045d6"
	)
	apiversion, _ := strconv.Atoi("ONEVIEW_APIVERSION")
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
		"*")
	fragments := new([]ov.StartStopFragments)
	fragment1 := ov.StartStopFragments{EndAddress: "10.16.0.100", StartAddress: "10.16.0.10"}
	*fragments = append(*fragments, fragment1)
	ipV4Range := ov.CreateIpv4Range{
		Type:               "Range",
		Name:               "test",
		StartStopFragments: *fragments,
		SubnetUri:          "/rest/id-pools/ipv4/subnets/40f76df9-1e39-4e5a-81fc-14614efea5e8",
	}

	// Gets an IPv4 range
	ipv4Range_by_id, err := ovc.GetIPv4RangebyId("", id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get Ipv4Range by id----------------#")
		jsonResponse, _ := json.MarshalIndent(ipv4Range_by_id, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}

	// Creates an IPv4 range
	fmt.Println(ipV4Range)
	iprange, err := ovc.CreateIPv4Range(ipV4Range)
	if err != nil {
		fmt.Println("IPv4 Range Creation Failed: ", err)
	} else {
		fmt.Println("IPv4 Range created successfully...")
		jsonResponse, _ := json.MarshalIndent(ipV4Range, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}

	// Gets all allocated fragments in an IPv4 range
	allocatedFragments, err := ovc.GetAllocatedFragments("", "", id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get allocatedFragments ----#")
		for i := range allocatedFragments.Members {
			fmt.Println(allocatedFragments.Members[i])
		}
	}

	// Gets all free fragments in an IPv4 range
	freeFragments, err := ovc.GetFreeFragments("", "", id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get freefragments ----#")
		for i := range freeFragments.Members {
			fmt.Println(freeFragments.Members[i])
		}
	}

	//Allocates a set of IDs from an IPv4 range.
	idlist := new([]utils.Nstring)
	*idlist = append(*idlist, utils.NewNstring("10.1.0.2"))
	*idlist = append(*idlist, utils.NewNstring("10.1.0.3"))
	allocateId := ov.UpdateAllocatorList{
		IdList: *idlist,
	}
	allocate, err := ovc.AllocateId(allocateId, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Allocated ID Successfully---#")
		jsonResponse, _ := json.MarshalIndent(allocate, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}

	//Collects a set of IDs back to an IPv4 range.
	collectId := ov.UpdateCollectorList{
		IdList: *idlist,
	}
	collect, err := ovc.CollectId(collectId, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Collected ID Successfully---#")
		jsonResponse, _ := json.MarshalIndent(collect, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}
	/*
		// Perform either of the following operations on a Range i.e., Enable Range or Edit Range
		updateIpv4Range := ov.Updateipv4{Type: "Range", Enabled: false}
		err = ovc.UpdateIpv4Range("a257c58c-bbe9-4174-b2a3-eada622fc555", updateIpv4Range)
		if err != nil {
			panic(err)
		} else {
			if updateIpv4Range.Enabled == false {
				fmt.Println("Ipv4Range has disabled successfully ")
			} else {
				fmt.Println("Ipv4Range has enabled successfully")
			}
		}
	*/

	id = strings.Split(iprange.URI.String(), "/")[5] //id="a257c58c-bbe9-4174-b2a3-eada622fc555
	// Deletes an IPv4 range.
	err = ovc.DeleteIpv4Range(id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Deleted Ipv4Range successfully---#")
	}
}
