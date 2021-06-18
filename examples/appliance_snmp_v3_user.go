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
		UserName:                 "user0",
		SecurityLevel:            "Authentication and privacy",
		AuthenticationProtocol:   "SHA1",
		AuthenticationPassphrase: "authPass",
		PrivacyProtocol:          "AES-128",
		PrivacyPassphrase:        "1234567812345678",
	}

	snmpv3User1 := ov.SNMPv3User{
		UserName:                 "user1",
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

	//Creating second SNMP User
	_, err0 := ovc.CreateSNMPv3Users(snmpv3User1)
	if err0 != nil {
		panic(err0)
	} else {
		fmt.Println("#-----Created SNMPv3 User-----#")
	}

	//Get all the SNMPv3 User
	allsnmpv3user, err1 := ovc.GetSNMPv3Users("", "", "", "")
	if err1 != nil {
		panic(err1)
	} else {
		fmt.Println("#-----Got All SNMPv3 Users-----#")
		for i := range allsnmpv3user.Members {
			fmt.Println(allsnmpv3user.Members[i])
		}
	}

	//Get SNMPv3 User by username
	snmpv3UserUserName, err3 := ovc.GetSNMPv3UserByUserName(snmpv3User.UserName)
	if err3 != nil {
		panic(err3)
	} else {
		fmt.Printf("\n#-----Got SNMPv3 User of UserName %v", snmpv3User.UserName)
		fmt.Println(snmpv3UserUserName)
	}

	//Get SNMPv3 User by ID

	snmpv3UserID, err2 := ovc.GetSNMPv3UserById(snmpv3UserUserName.Id)
	if err2 != nil {
		panic(err2)
	} else {
		fmt.Printf("\n#-----Got SNMPv3 User of ID %v", snmpv3UserUserName.Id)
		fmt.Println(snmpv3UserID)
	}

	//updating SNMPv3 User

	snmpv3UserID.PrivacyPassphrase = "55555555"
	snmpv3UserID.AuthenticationPassphrase = "aaaaaaaaa"

	response, err4 := ovc.UpdateSNMPv3User(snmpv3UserID, snmpv3UserUserName.Id)
	if err4 != nil {
		panic(err4)
	} else {
		fmt.Printf("\n#-----Updated SNMPv3 User Privacy passphrase and authenction passphrase : %v\n", response.URI)
	}

	// //Delete SNMPv3 User by ID
	err5 := ovc.DeleteSNMPv3UserById(response.Id)
	if err5 != nil {
		panic(err5)
	} else {
		fmt.Printf("\n#-----Deleted SNMPv3 User of ID %v\n", response.Id)
	}
	//Delete SNMPv3 User by UserName

	err6 := ovc.DeleteSNMPv3UserByName(snmpv3User1.UserName)
	if err5 != nil {
		panic(err6)
	} else {
		fmt.Printf("\n#-----Deleted SNMPv3 User of Username %v\n", snmpv3User1.UserName)
	}
}
