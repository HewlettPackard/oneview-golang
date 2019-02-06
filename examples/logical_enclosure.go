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
		logical_enclosure   = "sam" //"log_enc_66"
//		logical_enclosure_1 = "log_enclosure77"
//		logical_enclosure_2 = "log_enclosure88"
//		logical_enclosure_3 = "log_enclosure88"
		scope_name = "new-SD5"
	)
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		800,
		"")

	fmt.Println("#................... Logical Enclosure by Name ...............#")
	log_en, err := ovc.GetLogicalEnclosureByName(logical_enclosure)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(log_en)
	}

	scope1, err := ovc.GetScopeByName(scope_name)
	fmt.Println(scope1)
	fmt.Println(scope1.Uri)
	sort := "name:desc"
	log_en_list, err := ovc.GetLogicalEnclosures("", "", "", utils.NewNstring(scope1.Uri), sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... Logical Enclosures List .................#")
		for i := 0; i < len(log_en_list.Members); i++ {
			fmt.Println(log_en_list.Members[i].Name)
		}
	}
/*
	enclosureUris := new([]utils.Nstring)

	*enclosureUris = append(*enclosureUris, utils.NewNstring("/rest/enclosures/0000000000A66101"))
	*enclosureUris = append(*enclosureUris, utils.NewNstring("/rest/enclosures/0000000000A66102"))
	*enclosureUris = append(*enclosureUris, utils.NewNstring("/rest/enclosures/0000000000A66103"))

	logicalEnclosure := ov.LogicalEnclosure{Name: "log_enclosure77",
		EnclosureUris:     *enclosureUris,
		EnclosureGroupUri: utils.NewNstring("/rest/enclosure-groups/e48e8024-5e35-48ea-9bb9-0e4b3c69fb91")}

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
		log_en_after_update, err := ovc.GetLogicalEnclosures("", "", "", "", sort)
		if err != nil {
			fmt.Println(err)
		} else {
			for i := 0; i < len(log_en_after_update.Members); i++ {
				fmt.Println(log_en_after_update.Members[i].Name)
			}
		}
	}

	err = ovc.DeleteLogicalEnclosure(logical_enclosure_3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#...................... Deleted Logical Enclosure Successfully .....#")
	}
*/
}
