package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		clientOV                  *ov.OVClient
		server_hardware_type_name = "BL460c G7 1"
		uri                       = "/rest/server-hardware-types/" + "70BDABAA-87D1-4A39-9270-047D08B5C447"
		sort                      = "name:desc"
		//filter = "created equals '2018-01-24T16:24:14.330Z'"
	)
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		800)

	fmt.Println("#-----------------------Server Hardware Type by name-------------------------#")
	server_hardware_type, err := ovc.GetServerHardwareTypeByName(server_hardware_type_name)
	if err != nil {
		fmt.Println("Error while getting server hardware ttype by name ", server_hardware_type_name, ":", err)
	}
	fmt.Println(server_hardware_type)

	fmt.Println("#----------------------Server Hardware Type by uri---------------------------#")
	server_hardware_type, err = ovc.GetServerHardwareTypeByUri(utils.NewNstring(uri))
	if err != nil {
		fmt.Println("Error while getting server hardware type by uri ", uri, ":", err)
	}
	fmt.Println(server_hardware_type)

	fmt.Println("#---------------------Server Hardware list by count--------------------------#")
	server_hardware_type_list, err := ovc.GetServerHardwareTypes(0, 3, "", "")
	if err != nil {
		fmt.Println("Error in getting the server hardware types ", err)
	}
	for i := 0; i < len(server_hardware_type_list.Members); i++ {
		fmt.Println(server_hardware_type_list.Members[i].Name)
	}

	fmt.Println("#---------------------Server Hardware list by sort--------------------------#")
	server_hardware_type_list, err = ovc.GetServerHardwareTypes(0, 0, "", sort)
	if err != nil {
		fmt.Println("Error in getting the server hardware types ", err)
	}
	for i := 0; i < len(server_hardware_type_list.Members); i++ {
		fmt.Println(server_hardware_type_list.Members[i].Name)
	}
}
