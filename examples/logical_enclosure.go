package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		ClientOV            *ov.OVClient
		logical_enclosure   = "log_enc_66"
		logical_enclosure_1 = "log_enclosure77"
		logical_enclosure_2 = "log_enclosure88"
		scope_name          = "updated-SD2"
	)
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		1000,
		"*")

	fmt.Println("#................... Logical Enclosure by Name ...............#")
	log_en, err := ovc.GetLogicalEnclosureByName(logical_enclosure)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(log_en)
	}

	// Update From Group
	/*
	err = ovc.UpdateFromGroupLogicalEnclosure(log_en)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#............. Update From Group Logical Enclosure Successfully .....#")
	}
	*/
	fmt.Println("#................... Create Logical Enclosure ...............#")
	scope1, err := ovc.GetScopeByName(scope_name)
	scope_uri := scope1.URI
	scope_Uris := new([]string)
	*scope_Uris = append(*scope_Uris, scope_uri.String())

	sort := "name:desc"
	log_en_list, err := ovc.GetLogicalEnclosures("", "", "", *scope_Uris, sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... Logical Enclosures List .................#")
		for i := 0; i < len(log_en_list.Members); i++ {
			fmt.Println(log_en_list.Members[i].Name)
		}
	}

	enclosureUris := new([]utils.Nstring)
	*enclosureUris = append(*enclosureUris, utils.NewNstring("/rest/enclosures/0000000000A66101"))
	*enclosureUris = append(*enclosureUris, utils.NewNstring("/rest/enclosures/0000000000A66102"))
	*enclosureUris = append(*enclosureUris, utils.NewNstring("/rest/enclosures/0000000000A66103"))

	logicalEnclosure := ov.LogicalEnclosure{Name: logical_enclosure_1,
		EnclosureUris:     *enclosureUris,
		EnclosureGroupUri: utils.NewNstring("/rest/enclosure-groups/4586b2d6-1a8f-48af-809e-cb05d04a03a8")}

	er := ovc.CreateLogicalEnclosure(logicalEnclosure)
	if er != nil {
		fmt.Println("............... Logical Enclosure Creation Failed:", err)
	} else {
		fmt.Println(".... Logical Enclosure Created Success")
	}

	log_enc, _ := ovc.GetLogicalEnclosureByName(logical_enclosure_1)
	log_enc.Name = logical_enclosure_2
	err = ovc.UpdateLogicalEnclosure(log_enc)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#.................... Logical Enclosure after Updating ...........#")
		log_en_after_update, err := ovc.GetLogicalEnclosures("", "", "", *scope_Uris, sort)
		if err != nil {
			fmt.Println(err)
		} else {
			for i := 0; i < len(log_en_after_update.Members); i++ {
				fmt.Println(log_en_after_update.Members[i].Name)
			}
		}
	}

	err = ovc.DeleteLogicalEnclosure(logical_enclosure_2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#...................... Deleted Logical Enclosure Successfully .....#")
	}

}
