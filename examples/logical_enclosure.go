package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		ClientOV           *ov.OVClient
		logical_enclosure   = "SYN03_LE"
		logical_enclosure_1 = "log_enclosure77"
		logical_enclosure_2 = "log_enclosure88"
		logical_enclosure_3 = "log_enclosure88"
	)
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		800,
		"")
	ovVer, _ := ovc.GetAPIVersion()
	fmt.Println(ovVer)

	fmt.Println("#................... Logical Enclosure by Name ...............#")
	log_en, err := ovc.GetLogicalEnclosureByName(logical_enclosure)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(log_en)
	}

	sort := "name:desc"
	log_en_list, err := ovc.GetLogicalEnclosures("", "", "", "", sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... Logical Enclosures List .................#")
		for i := 0; i < len(log_en_list.Members); i++ {
			fmt.Println(log_en_list.Members[i].Name)
		}
	}

	enclosureUris := new([]utils.Nstring) 

	*enclosureUris = append(*enclosureUris, utils.NewNstring("/rest/enclosures/013645CN759000AC"))
	*enclosureUris = append(*enclosureUris, utils.NewNstring("/rest/enclosures/013645CN759000AD"))


	logicalEnclosure := ov.LogicalEnclosure{Name: "log_enclosure77",
		EnclosureUris: *enclosureUris, 
		EnclosureGroupUri: utils.NewNstring("/rest/enclosure-groups/0f2a3f46-36ad-4c8f-9e88-763c062855d3")}


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

}
