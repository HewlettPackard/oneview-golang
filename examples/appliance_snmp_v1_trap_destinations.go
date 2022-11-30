package main

import (
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func main() {
	var (
		ClientOV *ov.OVClient
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

	snmpUser := ov.SNMPv1Trap{
		Destination:     "192.0.1.14",
		Port:            162,
		CommunityString: "Test1",
	}

	id := "4"
	//Creating an SNMPv1 Trap Destinations
	err := ovc.CreateSNMPv1TrapDestinations(snmpUser, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#-----Created SNMPv1 Trap Destinations-----#")
	}

	//Get all the SNMPv1 Trap Destinations
	allTraps, err := ovc.GetSNMPv1TrapDestinations("", "", "", "")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#-----Got All SNMPv1 Trap Destinations-----#")
		for i := range allTraps.Members {
			fmt.Println(allTraps.Members[i])
		}
	}

	//Get SNMPv1 Trap Destinations by ID
	trapId, err := ovc.GetSNMPv1TrapDestinationsById(id)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\n#-----Got SNMPv1 Trap Destinations of ID %v", id)
		fmt.Println(trapId)
	}

	//updating SNMPv1 Trap Destination
	//Destination Address cannot be changed.
	update := ov.SNMPv1Trap{
		CommunityString: "Test3",
		Port:            190,
		Destination:     "192.0.1.14",
	}
	response, err := ovc.UpdateSNMPv1TrapDestinations(update, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\n#-----Updated SNMPv1 Trap Destinations port and community string: %v\n", response.URI)
	}

	//Delete SNMPv1 Trap Destinations by ID
	err = ovc.DeleteSNMPv1TrapDestinations(id)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\n#-----Deleted SNMPv1 Trap Destinations of ID %v\n", id)
	}
}
