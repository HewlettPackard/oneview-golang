package main

import (
	"fmt"
	"os"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	)

func main() {
	var (
		clientOV *ov.OVClient
		eg_name = "DemoEnclosureGroup"
		new_eg_name = "RenamedEnclosureGroup"
		script = "#TEST COMMAND"
	)
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		800)

	ibMappings := new([]ov.InterconnectBayMap)
	interconnectBay1 := ov.InterconnectBayMap {1, utils.NewNstring("/rest/logical-interconnect-groups/65245305-c8e9-4b28-9bec-c5f697dfa1db")}
	*ibMappings = append(*ibMappings,interconnectBay1)

	enclosureGroup := ov.EnclosureGroup{Name: eg_name, InterconnectBayMappings: *ibMappings}

	err := ovc.CreateEnclosureGroup(enclosureGroup)
	if err != nil {
		fmt.Println("Enclosure Group Creation Failed: ",err)
	}
	fmt.Println("Enclosure Group created successfully...")

	sort := "name:desc"

	enc_grp_list,err := ovc.GetEnclosureGroups("","","",sort,"")
	if err != nil {
		panic(err)
	}
	fmt.Println("#----------------Enclosure Group List---------------#")

	for i:=0;i<len(enc_grp_list.Members);i++ {
		fmt.Println(enc_grp_list.Members[i].Name)
	}

	if ovc.APIVersion > 500 {
		scope_uri := "'/rest/scopes/63d1ca81-95b3-41f1-a1ee-f9e1bc2d635f'"
		enc_grp_list1,err := ovc.GetEnclosureGroups("","","","",scope_uri)
		if err != nil {
			fmt.Println("Error in getting EnclosureGroups by scope URIs:",err)
		}
		fmt.Println("#-----------Enclosure Groups based on Scope URIs----------#")
		for i:=0;i<len(enc_grp_list1.Members);i++ {
			fmt.Println(enc_grp_list1.Members[i].Name)
		}
	}

	enc_grp,err := ovc.GetEnclosureGroupByName(eg_name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#-------------Enclosure Group by name----------------#")
	fmt.Println(enc_grp)

	uri := enc_grp.URI
	enc_grp,err = ovc.GetEnclosureGroupByUri(uri)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#----------------Enclosure Group by URI--------------#")
	fmt.Println(enc_grp)

	enc_grp.Name = new_eg_name
	err = ovc.UpdateEnclosureGroup(enc_grp)
	if err != nil {
		panic(err)
	}

	enc_grp_list,err = ovc.GetEnclosureGroups("","","",sort,"")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#----------------EnclosureList after updating---------#")
	for i:=0;i<len(enc_grp_list.Members);i++ {
		fmt.Println(enc_grp_list.Members[i].Name)
	}

	update_script,err := ovc.UpdateConfigurationScript(enc_grp.URI, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Update Configuration Script result:",update_script)

	conf_script,err := ovc.GetConfigurationScript(enc_grp.URI)
	if err != nil {
		fmt.Println("Error in getting configuration Script: ",err)
	}
	fmt.Println("Configuation Script: ",conf_script)

	fmt.Println(script)
	err = ovc.DeleteEnclosureGroup(new_eg_name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted EnclosureGroup successfully...")
}
