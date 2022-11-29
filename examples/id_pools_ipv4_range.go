package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"path"
	"strings"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
)

func main() {
	var (
		ClientOV *ov.OVClient
		eg_name  = "EG_For_IPV4Range"
	)
	config, config_err := ov.LoadConfigFile("config.json")
	if config_err != nil {
		fmt.Println(config_err)
	}

	ovc := ClientOV.NewOVClient(
		config.OVCred.UserName,
		config.OVCred.Password,
		config.OVCred.Domain,
		config.OVCred.Endpoint,
		config.OVCred.SslVerify,
		config.OVCred.ApiVersion,
		config.OVCred.IfMatch)

	// Creates an IPv4 Subnet
	subnet := ov.Ipv4Subnet{

		Name:       config.IdPoolsIpv4SubnetRange.SubnetName,
		NetworkId:  config.IdPoolsIpv4SubnetRange.NetworkId,
		SubnetMask: config.IdPoolsIpv4SubnetRange.SubnetMask,
		Gateway:    config.IdPoolsIpv4SubnetRange.Gateway,
		Domain:     config.IdPoolsIpv4SubnetRange.Domain,
	}

	fmt.Println("#-----------------Creating Subnet----------------#")
	err := ovc.CreateIPv4Subnet(subnet)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("IPv4 Subnet created successfully...")
	}
	//Gets an IPv4 subnet by Id
	fmt.Println("#-------------Get Ipv4Range by id----------------#")
	subnetName, err := ovc.GetSubnetByNetworkId(subnet.NetworkId)

	subnet_id := strings.Split(subnetName.URI.String(), "/")[5]

	subnetById, err := ovc.GetIPv4SubnetbyId(subnet_id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(subnetById.URI.String(), subnetById.Domain)
	}

	// Create ip4 range
	fragments := new([]ov.StartStopFragments)
	fragment1 := ov.StartStopFragments{EndAddress: config.IdPoolsIpv4SubnetRange.EndAddress1, StartAddress: config.IdPoolsIpv4SubnetRange.StartAddress}
	*fragments = append(*fragments, fragment1)

	ipV4Range := ov.CreateIpv4Range{
		Type:               "Range",
		Name:               "test",
		StartStopFragments: *fragments,
		SubnetUri:          subnetById.URI,
	}

	// // // Creates an IPv4 range

	iprange, err := ovc.CreateIPv4Range(ipV4Range)
	fmt.Println(iprange.URI)
	if err != nil {
		fmt.Println("IPv4 Range Creation Failed: ", err)
	} else {
		fmt.Println("IPv4 Range created successfully...")
		jsonResponse, _ := json.MarshalIndent(ipV4Range, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}

	myUrl, err := url.Parse(string(iprange.URI))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path.Base(myUrl.Path))
	id := path.Base(myUrl.Path)

	// // Gets an IPv4 range

	ipv4Range_by_uri, err := ovc.GetIPv4RangebyId("", id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get Ipv4Range by uri----------------#")
		jsonResponse, _ := json.MarshalIndent(ipv4Range_by_uri, "", "  ")
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
	*iprange1 = append(*iprange1, utils.Nstring(iprange.URI))

	enclosureGroup := ov.EnclosureGroup{Name: eg_name, IpAddressingMode: "ipPool", IpRangeUris: *iprange1, InitialScopeUris: *initialScopeUris, EnclosureCount: 3}
	err = ovc.CreateEnclosureGroup(enclosureGroup)
	if err != nil {
		panic(err)
	}

	// Allocates a set of IDs from an IPv4 range.
	// A set of IDs can be allocated through count parameter also.
	idlist := new([]utils.Nstring)
	*idlist = append(*idlist, config.IdPoolsIpv4SubnetRange.IdList[0])
	*idlist = append(*idlist, config.IdPoolsIpv4SubnetRange.IdList[1])
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
	// // Perform either of the following operations on a Range i.e., Enable Range or Edit Range
	// // Performing Enable Range.
	updateIpv4Range := ov.Ipv4Range{Type: "Range", Enabled: true}
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
	fragment_1 := ov.StartStopFragments{EndAddress: config.IdPoolsIpv4SubnetRange.EndAddress2, StartAddress: config.IdPoolsIpv4SubnetRange.StartAddress}
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

	err = ovc.DeleteIpv4Subnet(subnet_id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Deleted Ipv4Range successfully---#")
	}

	//******************* Create iscsi network for automation*****************************
	err1 := ovc.CreateIPv4Subnet(subnet)
	if err1 != nil {
		panic(err)
	} else {
		fmt.Println("IPv4 Subnet created successfully...")
	}
	///
	fmt.Println("#-------------Get Ipv4Range by id----------------#")
	subnetName1, err := ovc.GetSubnetByNetworkId(subnet.NetworkId)
	fmt.Println(subnetName1)
	fmt.Println(subnet.Name)
	subnet_id = strings.Split(subnetName1.URI.String(), "/")[5]

	//
	subnetById1, err2 := ovc.GetIPv4SubnetbyId(subnet_id)
	if err1 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(subnetById.URI.String(), subnetById1.Domain)
	}

	// Create ip4 range
	fragments = new([]ov.StartStopFragments)
	fragment1 = ov.StartStopFragments{EndAddress: config.IdPoolsIpv4SubnetRange.EndAddress1, StartAddress: config.IdPoolsIpv4SubnetRange.StartAddress}
	*fragments = append(*fragments, fragment1)

	ipV4Range_iscsi := ov.CreateIpv4Range{
		Type:               "Range",
		Name:               "iscsi",
		StartStopFragments: *fragments,
		SubnetUri:          subnetById1.URI,
	}

	// // // Creates an IPv4 range

	iprange, err3 := ovc.CreateIPv4Range(ipV4Range_iscsi)
	fmt.Println(iprange.URI)
	if err3 != nil {
		fmt.Println("IPv4 Range Creation Failed: ", err3)
	} else {
		fmt.Println("IPv4 Range created successfully...")
		jsonResponse, _ := json.MarshalIndent(ipV4Range, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}
}
