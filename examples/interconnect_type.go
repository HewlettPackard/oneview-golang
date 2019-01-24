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
		existing_interconnect = "Synergy10GB"
	)
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		800,
		"*")

	fmt.Println("#................... Interconnect Type by Name ...............#")
	interconnect, err := ovc.GetInterconnectTypeByName(existing_interconnect)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(interconnect)
	}

	interconnect_type := utils.NewNstring("rest/interconnect-types/ed8fbf3b-1962-4602-8eeb-6ef2f303a67f")
	fmt.Println("#................... Interconnect Type by Uri ....................#")
	int_uri, err := ovc.GetInterconnectTypeByUri(interconnect_type)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(int_uri)
	}

	sort := "name:desc"
	interconnect_type_list, err := ovc.GetInterconnectTypes("","","",sort)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("# ................... Interconnect Type List .................#")
		for i := 0; i < len(interconnect_type_list.Members); i++ {
			fmt.Println(interconnect_type_list.Members[i].Name)
		}
	}
}
