package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		clientOV    *ov.OVClient
		eg_name     = "TestEG"
		new_eg_name = "RenamedEnclosureGroup"
		script      = "#TEST COMMAND"
	)
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		1800,
		"*")

	ibMappings := new([]ov.InterconnectBayMap)
	interconnectBay1 := ov.InterconnectBayMap{InterconnectBay: 1, LogicalInterconnectGroupUri: utils.NewNstring("/rest/logical-interconnect-groups/4b59f8cf-e222-4fe1-9983-43aa836c5e31")}
	interconnectBay2 := ov.InterconnectBayMap{InterconnectBay: 4, LogicalInterconnectGroupUri: utils.NewNstring("/rest/logical-interconnect-groups/8d9fd7b1-f59f-44c6-b9da-429d68c79f6b")}
	*ibMappings = append(*ibMappings, interconnectBay1)
	*ibMappings = append(*ibMappings, interconnectBay2)
	initialScopeUris := new([]utils.Nstring)
	*initialScopeUris = append(*initialScopeUris, utils.NewNstring("/rest/scopes/7fe26585-b7a1-497e-992e-90908f70dfaf"))

	enclosureGroup := ov.EnclosureGroup{Name: eg_name, InterconnectBayMappings: *ibMappings, InitialScopeUris: *initialScopeUris, IpAddressingMode: "External", EnclosureCount: 1}
	/*
		 This is used for C7000 only
		enclosureGroup := ov.EnclosureGroup{Name: eg_name, InitialScopeUris: *initialScopeUris, InterconnectBayMappings: *ibMappings}
	*/

	err := ovc.CreateEnclosureGroup(enclosureGroup)
	if err != nil {
		fmt.Println("Enclosure Group Creation Failed: ", err)
	} else {
		fmt.Println("Enclosure Group created successfully...")
	}

	sort := "name:desc"

	enc_grp_list, err := ovc.GetEnclosureGroups("", "", "", sort, "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#----------------Enclosure Group List---------------#")

	for i := 0; i < len(enc_grp_list.Members); i++ {
		fmt.Println(enc_grp_list.Members[i].Name)
	}

	if ovc.APIVersion > 500 {
		scope_uri := "'/rest/scopes/7fe26585-b7a1-497e-992e-90908f70dfaf'"
		enc_grp_list1, err := ovc.GetEnclosureGroups("", "", "", "", scope_uri)
		if err != nil {
			fmt.Println("Error in getting EnclosureGroups by scope URIs:", err)
		}
		fmt.Println("#-----------Enclosure Groups based on Scope URIs----------#")
		for i := 0; i < len(enc_grp_list1.Members); i++ {
			fmt.Println(enc_grp_list1.Members[i].Name)
		}
	}

	enc_grp, err := ovc.GetEnclosureGroupByName(eg_name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#-------------Enclosure Group by name----------------#")
	fmt.Println(enc_grp)

	uri := enc_grp.URI
	enc_grp, err = ovc.GetEnclosureGroupByUri(uri)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#----------------Enclosure Group by URI--------------#")
	fmt.Println(enc_grp)

	enc_grp.Name = new_eg_name
	err = ovc.UpdateEnclosureGroup(enc_grp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Enclosure Group got updated")
	}

	enc_grp_list, err = ovc.GetEnclosureGroups("", "", "", sort, "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#----------------EnclosureList after updating---------#")
	for i := 0; i < len(enc_grp_list.Members); i++ {
		fmt.Println(enc_grp_list.Members[i].Name)
	}

	// This method is only available on C7000
	update_script, err := ovc.UpdateConfigurationScript(enc_grp.URI, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Update Configuration Script result:", update_script)

	// This method is only available on C700
	conf_script, err := ovc.GetConfigurationScript(enc_grp.URI)
	if err != nil {
		fmt.Println("Error in getting configuration Script: ", err)
	}
	fmt.Println("Configuation Script: ", conf_script)

	fmt.Println(script)
	err = ovc.DeleteEnclosureGroup(new_eg_name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted EnclosureGroup successfully...")
}
