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

	// Crete user

	snmpv3User := ov.SNMPv3User{
		UserName:                 "user0",
		SecurityLevel:            "Authentication and privacy",
		AuthenticationProtocol:   "SHA1",
		AuthenticationPassphrase: "authPass",
		PrivacyProtocol:          "AES-128",
		PrivacyPassphrase:        "1234567812345678",
	}
	//Creating an SNMPv3 User
	_, err := ovc.CreateSNMPv3Users(snmpv3User)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#-----Created SNMPv3 User-----#")
	}
	snmpv3UserUserName, err3 := ovc.GetSNMPv3UserByUserName(snmpv3User.UserName)
	if err3 != nil {
		panic(err3)
	} else {
		fmt.Printf("\n#-----Got SNMPv3 User of UserName %v", snmpv3User.UserName)
		fmt.Println(snmpv3UserUserName)
	}
	snmpUser := ov.SNMPv3Trap{
		DestinationAddress: "1.1.1.1",
		Port:               162,
		UserID:             snmpv3UserUserName.Id,
	}

	//Creating an SNMPv3 Trap Destinations
	_, err = ovc.CreateSNMPv3TrapDestinations(snmpUser)
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

	//Delete SNMPv3 User by UserName

	err6 := ovc.DeleteSNMPv3UserByName(snmpv3User.UserName)
	if err6 != nil {
		panic(err6)
	} else {
		fmt.Printf("\n#-----Deleted SNMPv3 User of Username %v\n", snmpv3User.UserName)
	}

}
