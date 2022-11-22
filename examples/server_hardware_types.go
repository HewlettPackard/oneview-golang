package main

import (
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func main() {
	var (
		ClientOV *ov.OVClient
		sort     = "name:desc"
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

	fmt.Println("#---------------------Server Hardware list by sort--------------------------#")
	server_hardware_type_list, err := ovc.GetServerHardwareTypes(0, 0, "", sort)
	if err != nil {
		fmt.Println("Error in getting the server hardware types ", err)
	}
	for i := 0; i < len(server_hardware_type_list.Members); i++ {
		fmt.Println(server_hardware_type_list.Members[i])
	}

	fmt.Println("#-----------------------Server Hardware Type by name-------------------------#")
	server_hardware_type, err := ovc.GetServerHardwareTypeByName(server_hardware_type_list.Members[0].Name)
	if err != nil {
		fmt.Println("Error while getting server hardware ttype by name ", server_hardware_type_list.Members[0].Name, ":", err)
	}
	fmt.Println(server_hardware_type)

	fmt.Println("#----------------------Server Hardware Type by uri---------------------------#")
	server_hardware_type, err = ovc.GetServerHardwareTypeByUri(server_hardware_type.URI)
	if err != nil {
		fmt.Println("Error while getting server hardware type by uri ", server_hardware_type.URI, ":", err)
	}
	fmt.Println(server_hardware_type)

	fmt.Println("#---------------------Server Hardware list by count--------------------------#")
	server_hardware_type_list, err = ovc.GetServerHardwareTypes(0, 3, "", "")
	if err != nil {
		fmt.Println("Error in getting the server hardware types ", err)
	}
	for i := 0; i < len(server_hardware_type_list.Members); i++ {
		fmt.Println(server_hardware_type_list.Members[i].Name)
	}
}
