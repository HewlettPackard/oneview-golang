package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func main() {
	var (
		ClientOV *ov.OVClient
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

	// Creates an IPv4 Subnet
	subnet := ov.Ipv4Subnet{
		Name:       "SubnetTF",
		NetworkId:  "<networkId>",
		SubnetMask: "<subnetMask>",
		Gateway:    "<gateway>",
		Domain:     "Terraform.com",
	}

	fmt.Println("#-----------------Creating Subnet----------------#")
	err := ovc.CreateIPv4Subnet(subnet)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("IPv4 Subnet created successfully...")
	}

	// Gets all subnets
	fmt.Println("#--------Subnet List--------------#")

	allSubnets, err := ovc.GetIPv4Subnets("", "", "", "")

	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(allSubnets.Members); i++ {
			fmt.Println(allSubnets.Members[i].NetworkId)
		}
	}

	// Gets an IPv4 subnet by Id
	fmt.Println("#-------------Get Ipv4Range by id----------------#")
	subnetName, err := ovc.GetSubnetByNetworkId(subnet.NetworkId)
	fmt.Println(subnetName)
	fmt.Println(subnet.Name)
	id := strings.Split(subnetName.URI.String(), "/")[5]

	subnetById, err := ovc.GetIPv4SubnetbyId(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(subnetById.URI.String(), subnetById.Domain)
	}

	// Create Ethernet network and associate resource to the subnet
	// Allocator and Collector needs a subnet associated with any resource
	ethernetNetwork := ov.EthernetNetwork{
		Name:                "SubnetNetwork",
		VlanId:              9,
		Purpose:             "General",
		EthernetNetworkType: "Tagged",
		Type:                "ethernet-networkV4",
		SubnetUri:           subnetById.URI,
	}
	err = ovc.CreateEthernetNetwork(ethernetNetwork)

	//Creates Range for the above subnet
	fragments := new([]ov.StartStopFragments)
	fragment1 := ov.StartStopFragments{EndAddress: "<endAddress>", StartAddress: "<startAddress>"}
	*fragments = append(*fragments, fragment1)
	ipV4Range := ov.CreateIpv4Range{
		Type:               "Range",
		Name:               "SubnetRange",
		StartStopFragments: *fragments,
		SubnetUri:          subnetById.URI,
	}

	err = ovc.CreateIPv4Range(ipV4Range)

	// Allocates IPv4 Ips from subnet associated with a resource
	fmt.Println("#--------IPv4 Allocation from Subnet--------------#")
	allocator := ov.SubnetAllocatorList{
		Count: 3,
	}
	allocatedIds, err := ovc.AllocateIpv4Subnet(id, allocator)
	if err != nil {
		panic(err)
	} else {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(allocatedIds.IdList)
		}
	}

	// Collects allocated IPv4 Ips
	fmt.Println("#--------Collect IPv4 Ids allocated--------------#")
	collect := ov.SubnetCollectorList{
		IdList: allocatedIds.IdList,
	}
	collectedIds, err := ovc.CollectIpv4Subnet(id, collect)
	if err != nil {
		panic(err)
	} else {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(collectedIds.IdList)
		}
	}

	// Updates Name and NetworkId of the existing subnet
	fmt.Println("#-----------Updates Subnet-------------#")

	updateSubnet := ov.Ipv4Subnet{
		Name:       "SubnetRenamed",
		NetworkId:  "<networkId>",
		SubnetMask: "<subnetMask>",
		Gateway:    "<gateway>",
	}

	err = ovc.UpdateIpv4Subnet(id, updateSubnet)
	if err != nil {
		panic(err)
	} else {
		updatedSubnet, err := ovc.GetIPv4SubnetbyId(id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(updatedSubnet.Name)
		}
	}

	//Unassociate the resource from subnet before deletion
	err = ovc.DeleteEthernetNetwork(ethernetNetwork.Name)

	// Deletes an IPv4 subnet.
	fmt.Println("#----------Delete Subnet---------#")

	err = ovc.DeleteIpv4Subnet(id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleted Ipv4 Subnet successfully")
	}
}
