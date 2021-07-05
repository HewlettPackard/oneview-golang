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
		id       = "ebb4dac4-6f9f-48b4-b356-fd371a026899"
		eg_name  = "EG_For_IPV4Range"
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
	fragment1 := ov.StartStopFragments{EndAddress: "<ip_address>", StartAddress: "<ip_address>"}
	*fragments = append(*fragments, fragment1)

	ipV4Range := ov.CreateIpv4Range{
		Type:               "Range",
		Name:               "test",
		StartStopFragments: *fragments,
		SubnetUri:          "/rest/id-pools/ipv4/subnets/c04b28fe-f6fe-4a40-bc23-c17ccdf4c777",
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

	iprange, err := ovc.CreateIPv4Range(ipV4Range)
	fmt.Println(iprange.URI)
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
	//Create a resource which can assigned the ipv4range
	scp, _ := ovc.GetScopeByName("Auto-Scope")
	initialScopeUris := new([]utils.Nstring)
	*initialScopeUris = append(*initialScopeUris, scp.URI)
	iprange1 := new([]utils.Nstring)
	*iprange1 = append(*iprange1, iprange.URI)

	enclosureGroup := ov.EnclosureGroup{Name: eg_name, IpAddressingMode: "ipPool", IpRangeUris: *iprange1, InitialScopeUris: *initialScopeUris, EnclosureCount: 3}
	err = ovc.CreateEnclosureGroup(enclosureGroup)
	if err != nil {
		panic(err)
	}

	// Allocates a set of IDs from an IPv4 range.
	// A set of IDs can be allocated through count parameter also.
	idlist := new([]utils.Nstring)
	*idlist = append(*idlist, utils.NewNstring("<ip_address>"))
	*idlist = append(*idlist, utils.NewNstring("<ip_address>"))
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

	id = strings.Split(iprange.URI.String(), "/")[5] //id="a257c58c-bbe9-4174-b2a3-eada622fc555
	// Perform either of the following operations on a Range i.e., Enable Range or Edit Range
	// Performing Enable Range.
	updateIpv4Range := ov.Ipv4Range{Type: "Range", Enabled: false}
	resp, err := ovc.UpdateIpv4Range(id, updateIpv4Range)
	if err != nil {
		panic(err)
	} else {
		if resp.Enabled == false {
			fmt.Println("Ipv4Range has disabled successfully ")
		} else {
			fmt.Println("Ipv4Range has enabled successfully")
		}
	}

	// Performing Edit Range
	fragments_2 := new([]ov.StartStopFragments)
	fragment_1 := ov.StartStopFragments{EndAddress: "<ip_address>", StartAddress: "<ip_address>"}
	*fragments_2 = append(*fragments_2, fragment_1)
	updateRange := ov.Ipv4Range{
		Type:               "Range",
		StartStopFragments: *fragments_2,
		Name:               "Renamed-Range",
	}
	resp_2, err := ovc.UpdateIpv4Range(id, updateRange)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Ipv4Range has updated successfully")
		jsonResponse, _ := json.MarshalIndent(resp_2, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}
	//Delete Enclosure group
	err = ovc.DeleteEnclosureGroup(eg_name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted EnclosureGroup successfully...")

	// Deletes an IPv4 range.
	err = ovc.DeleteIpv4Range(id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Deleted Ipv4Range successfully---#")
	}

}
