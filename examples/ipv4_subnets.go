package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"strconv"
	"strings"
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
		NetworkId:  "192.169.1.0",
		SubnetMask: "255.255.255.0",
		Gateway:    "192.169.1.1",
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
	id := strings.Split(allSubnets.Members[0].URI.String(), "/")[5]
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(allSubnets.Members); i++ {
			fmt.Println(allSubnets.Members[i].NetworkId)
			eachSubnet := allSubnets.Members[i]

			// Gets id of the subnet created above for update and delete
			if eachSubnet.Name == subnet.Name {
				id = strings.Split(eachSubnet.URI.String(), "/")[5]
			}
		}

	}

	// Gets an IPv4 subnet by Id
	fmt.Println("#-------------Get Ipv4Range by id----------------#")

	subnetById, err := ovc.GetIPv4SubnetbyId(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(subnetById.URI.String(), subnetById.Domain)
	}

	// Allocates IPv4 Ips from subnet associated with a resource
	fmt.Println("#--------IPv4 Allocation from Subnet--------------#")
	idAllocator := strings.Split(allSubnets.Members[1].URI.String(), "/")[5]
	allocator := ov.SubnetAllocatorList{
		Count: 1,
	}
	allocatedIds, err := ovc.AllocateIpv4Subnet(idAllocator, allocator)
	if err != nil {
		panic(err)
	} else {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(allocatedIds.IdList)
		}
	}

	// Allocates IPv4 Ips from subnet associated with a resource
	fmt.Println("#--------Collect IPv4 Ids allocated--------------#")
	collect := ov.SubnetCollectorList{
		IdList: allocatedIds.IdList,
	}
	collectedIds, err := ovc.CollectIpv4Subnet(idAllocator, collect)
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
		NetworkId:  "192.169.1.0",
		SubnetMask: "255.255.255.0",
		Gateway:    "192.169.1.1",
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

	// Deletes an IPv4 subnet.
	fmt.Println("#----------Delete Subnet---------#")

	err = ovc.DeleteIpv4Subnet(id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleted Ipv4 Subnet successfully")
	}
}
