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

	snmpv3User := ov.SNMPv3User{
		UserName:                 "user7",
		SecurityLevel:            "Authentication and privacy",
		AuthenticationProtocol:   "SHA1",
		AuthenticationPassphrase: "authPass",
		PrivacyProtocol:          "AES-128",
		PrivacyPassphrase:        "1234567812345678",
	}

	//Creating an SNMPv1 Trap Destinations
	_, err := ovc.CreateSNMPv3Users(snmpv3User)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#-----Created SNMPv3 User-----#")
	}

	//Get all the SNMPv3 User
	allsnmpv3user, err := ovc.GetSNMPv3Users("", "", "", "")
	if err1 != nil {
		panic(err1)
	} else {
		fmt.Println("#-----Got All SNMPv1 Trap Destinations-----#")
		for i := range allsnmpv3user.Members {
			fmt.Println(allsnmpv3user.Members[i])
		}
	}

	//Get SNMPv3 User by ID
	id := "49a96a2a-a200-44df-b600-3a4c3d9f267c"
	snmpv3User, err2 := ovc.GetSNMPv3UserById(id)
	if err2 != nil {
		panic(err2)
	} else {
		fmt.Printf("\n#-----Got SNMPv3 User of ID %v", id)
		fmt.Println(snmpv3User)
	}

	//updating SNMPv3 User

	updated_snmpv3user := ov.SNMPv3User{
		Type:                     "Users",
		URI:                      "/rest/appliance/snmpv3-trap-forwarding/users/1e706f76-c585-4fc5-8a2b-86af678e4be1",
		Category:                 "SnmpV3User",
		ETAG:                     "2021-06-14T16:49:41.137Z",
		Id:                       "1e706f76-c585-4fc5-8a2b-86af678e4be1",
		UserName:                 "user2",
		SecurityLevel:            "Authentication and privacy",
		AuthenticationProtocol:   "SHA1",
		AuthenticationPassphrase: "dsfdsfdsfdfdsfds",
		PrivacyProtocol:          "AES-128",
		PrivacyPassphrase:        "dsfdsfdsf545454",
	}
	response, err3 := ovc.UpdateSNMPv3User(updated_snmpv3user, id)
	if err != nil {
		panic(err3)
	} else {
		fmt.Printf("\n#-----Updated SNMPv1 Trap Destinations port and community string: %v\n", response.URI)
	}

	//Delete SNMPv3 User by ID
	err4 := ovc.DeleteSNMPv3UserById(id)
	if err4 != nil {
		panic(err4)
	} else {
		fmt.Printf("\n#-----Deleted SNMPv3 User of ID %v\n", id)
	}
	//Delete SNMPv3 User by ID
	username := "user3"

	err5 := ovc.DeleteSNMPv3UserByName(username)
	if err5 != nil {
		panic(err5)
	} else {
		fmt.Printf("\n#-----Deleted SNMPv3 User of Username %v\n", username)
	}
}
