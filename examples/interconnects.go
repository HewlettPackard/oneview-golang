package main

import (
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
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
	interconnect_list, err := ovc.GetInterconnects("", "", "", sort)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("#----------------Interconnect List---------------#")

		for i := 0; i < len(interconnect_list.Members); i++ {
			fmt.Println(interconnect_list.Members[i].Name)
		}

		interconnect, err := ovc.GetInterconnectByName(interconnect_list.Members[0].Name)
		if err != nil {
			fmt.Println(err)
		} else {

			fmt.Println("#-------------Interconnect by Name----------------#")
			fmt.Println(interconnect.Name)

			uri := interconnect.URI
			interconnect, err = ovc.GetInterconnectByUri(uri)
			if err != nil {
				fmt.Println(err)
			} else {

				fmt.Println("#----------------Interconnect by URI--------------#")
				fmt.Println(interconnect.Name)
			}
		}
	}
}
