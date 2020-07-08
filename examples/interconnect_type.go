package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		ClientOV              *ov.OVClient
		existing_interconnect = "Virtual Connect SE 40Gb F8 Module for Synergy"
	)
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		1800,
		"*")

	fmt.Println("#................... Interconnect Type by Name ...............#")
	interconnect, err := ovc.GetInterconnectTypeByName(existing_interconnect)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(interconnect)
	}

	interconnect_type := utils.NewNstring("/rest/interconnects/e1242fc7-4e54-488d-90b2-2f8469a2b80a")
	fmt.Println("#................... Interconnect Type by Uri ....................#")
	int_uri, err := ovc.GetInterconnectTypeByUri(interconnect_type)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(int_uri)
	}

	sort := "name:desc"
	interconnect_type_list, err := ovc.GetInterconnectTypes("", "", "", sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... Interconnect Type List .................#")
		for i := 0; i < len(interconnect_type_list.Members); i++ {
			fmt.Println(interconnect_type_list.Members[i].Name)
		}
	}
}
