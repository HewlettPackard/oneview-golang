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
		UserName:                 "user47",
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

	//Get all the SNMPv3 User
	allsnmpv3user, err1 := ovc.GetSNMPv3Users("", "", "", "")
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

	//Get SNMPv3 User by username
	username := "user47"
	snmpv3User, err3 := ovc.GetSNMPv3UserByUserName(username)
	if err3 != nil {
		panic(err3)
	} else {
		fmt.Printf("\n#-----Got SNMPv3 User of UserName %v", username)
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
	response, err4 := ovc.UpdateSNMPv3User(updated_snmpv3user, id)
	if err4 != nil {
		panic(err4)
	} else {
		fmt.Printf("\n#-----Updated SNMPv1 Trap Destinations port and community string: %v\n", response.URI)
	}

	// //Delete SNMPv3 User by ID
	err5 := ovc.DeleteSNMPv3UserById(id)
	if err5 != nil {
		panic(err5)
	} else {
		fmt.Printf("\n#-----Deleted SNMPv3 User of ID %v\n", id)
	}
	//Delete SNMPv3 User by UserName
	username1 := "user3"

	err6 := ovc.DeleteSNMPv3UserByName(username1)
	if err5 != nil {
		panic(err6)
	} else {
		fmt.Printf("\n#-----Deleted SNMPv3 User of Username %v\n", username1)
	}
}
