package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func main() {
	var (
		ClientOV *ov.OVClient
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

	snmpUser := ov.SNMPv3Trap{
		DestinationAddress: "1.1.1.1",
		Port:               162,
		UserID:             "d1c813f0-04a9-49e2-be31-01541d3a0912",
	}

	//Creating an SNMPv3 Trap Destinations
	err := ovc.CreateSNMPv3TrapDestinations(snmpUser)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#-----Created SNMPv3 Trap Destinations-----#")
	}

	//Get all the SNMPv3 Trap Destinations
	allTraps, err := ovc.GetSNMPv3TrapDestinations("", "", "", "")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#-----Got All SNMPv3 Trap Destinations-----#")
		for i := range allTraps.Members {
			fmt.Println(allTraps.Members[i])
		}
	}

	//Get SNMPv3 Trap Destinations by ID
	id := allTraps.Members[0].ID
	trapId, err := ovc.GetSNMPv3TrapDestinationsById(id)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\n#-----Got SNMPv3 Trap Destinations of ID %v", id)
		fmt.Println(trapId)
	}

	//updating SNMPv3 Trap Destination
	//ID and Destination Address cannot be changed.
	update := ov.SNMPv3Trap{
		ID:                 trapId.ID,
		UserID:             snmpUser.UserID,
		Port:               190,
		DestinationAddress: "1.1.1.1",
	}
	response, err := ovc.UpdateSNMPv3TrapDestinations(update)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\n#-----Updated SNMPv3 Trap Destinations of ID %v", response.ID)
	}

	//Delete SNMPv3 Trap Destinations by ID
	err = ovc.DeleteSNMPv3TrapDestinations(id)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\n#-----Deleted SNMPv3 Trap Destinations of ID %v\n", id)
	}
}
