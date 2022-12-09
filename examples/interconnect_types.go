package main

import (
	"encoding/json"
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
	//	"encoding/json"
)

func main() {
	var (
		ClientOV *ov.OVClient
	)
	// Use configuratin file to set the ip and  credentails
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
	interconnect_type_list, err := ovc.GetInterconnectTypes("", "", "", sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... Interconnect Type List .................#")
		for i := 0; i < len(interconnect_type_list.Members); i++ {
			fmt.Println(interconnect_type_list.Members[i].Name)
			fmt.Println(interconnect_type_list.Members[i].URI)
		}
	}

	fmt.Println("#................... Interconnect Type by Name ...............#")
	interconnect, err := ovc.GetInterconnectTypeByName(string(interconnect_type_list.Members[0].Name))
	if err != nil {
		fmt.Println(err)
	} else {
		interconnect, _ := json.MarshalIndent(interconnect, "", "\t")
		fmt.Print(string(interconnect))
	}

	fmt.Println("#................... Interconnect Type by Uri ....................#")
	int_uri, err := ovc.GetInterconnectTypeByUri(interconnect.URI)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(int_uri)
	}
}
