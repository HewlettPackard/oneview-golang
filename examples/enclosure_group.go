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
		2000,
		"*")

	ibMappings := new([]ov.InterconnectBayMap)
	interconnectBay1 := ov.InterconnectBayMap{InterconnectBay: 3, LogicalInterconnectGroupUri: utils.NewNstring("/rest/logical-interconnect-groups/b0fce8d1-4916-4564-8e91-1bd32527aba4")}
	interconnectBay2 := ov.InterconnectBayMap{InterconnectBay: 6, LogicalInterconnectGroupUri: utils.NewNstring("/rest/logical-interconnect-groups/b0fce8d1-4916-4564-8e91-1bd32527aba4")}
	*ibMappings = append(*ibMappings, interconnectBay1)
	*ibMappings = append(*ibMappings, interconnectBay2)
	initialScopeUris := new([]utils.Nstring)
	*initialScopeUris = append(*initialScopeUris, utils.NewNstring("/rest/scopes/94a9804e-8521-4c26-bb00-e4875be53498"))

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
		scope_uri := "'/rest/scopes/94a9804e-8521-4c26-bb00-e4875be53498'"
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
