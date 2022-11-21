package main

import (
	"encoding/json"
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
	getsshaccess, err := ovc.GetSshAccess()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#--- Got the Appliance SSH access ---#")
		jsonResponse, _ := json.MarshalIndent(getsshaccess, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}

	sshaccess := ov.ApplianceSshAccess{
		AllowSshAccess: false,
	}
	jsonResponse, _ := json.MarshalIndent(sshaccess, "", "  ")
	fmt.Print(string(jsonResponse), "\n\n")
	err = ovc.SetSshAccess(sshaccess)
	if err != nil {
		fmt.Println("Appliance SSH access set failed: ", err)
	} else {
		fmt.Println("Appliance SSH access set successfully...")
	}

}
