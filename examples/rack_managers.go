package main

import (
	"fmt"
	"path"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func main() {
	var (
		ClientOV        *ov.OVClient
		RackManagerName = "TestRackManagerGOsdk"
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

	sort := "name:desc"

	rManagers, err := ovc.GetRackManagerList("", "", "", sort, "")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get Rack Managers sorted by name in descending order----#")
		for i := range rManagers.Members {
			fmt.Println(rManagers.Members[i].Name)
		}
	}
	rm1, err := ovc.GetRackManagerByName(RackManagerName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get Rack Manager by name----------------#")
		fmt.Println(rm1)
	}

	rackManager := ov.RackManager{
		Hostname: "<rack-manager-hostname>",
		UserName: "<rack-manager-username>",
		Password: "<rack-manager-password>",
	}

	rmUri, err := ovc.AddRackManager(rackManager)
	rmID := path.Base(rmUri)
	if err != nil {
		fmt.Println("............... Add Rack Manager Failed:", err)
	} else {
		fmt.Println(".... Create Rack Manager Success", rmID)
	}
	rmCreated, err := ovc.GetRackManagerById(rmID)
	fmt.Println(rmCreated.Name)
	rmcName := rmCreated.Name

	err = ovc.DeleteRackManager(rmcName)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleted Rack manager successfully...")
	}

}
