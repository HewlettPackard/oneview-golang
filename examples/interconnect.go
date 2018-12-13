package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	// "github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		clientOV    *ov.OVClient
		// eg_name     = "DemoEnclosureGroup"
		// new_eg_name = "RenamedEnclosureGroup"
		// script      = "#TEST COMMAND"
	)
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		800)

	interconnect_list, err := ovc.GetInterconnects("", "", "", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#----------------Interconnect List---------------#")

		for i := 0; i < len(interconnect_list.Members); i++ {
			fmt.Println(interconnect_list.Members[i].Name)
		}
	}

	interconnect, err := ovc.GetInterconnectByName(interconnect_list.Members[0].Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#-------------Interconnect by name----------------#")
	fmt.Println(interconnect.Name)

	uri := interconnect.URI
	interconnect, err = ovc.GetInterconnectByUri(uri)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#----------------Interconnect by URI--------------#")
	fmt.Println(interconnect.Name)
}
