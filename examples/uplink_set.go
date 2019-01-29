package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		ClientOV        *ov.OVClient
		existing_uplink = "upset66"
		new_uplink      = "upset77"
		upd_uplink      = "upset88"
		del_uplink      = "upset88"
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

	fmt.Println("#................... Uplink-Set by Name ...............#")
	uplink_set, err := ovc.GetUplinkSetByName(existing_uplink)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(uplink_set)
	}

	sort := "name:desc"
	uplinkset_list, err := ovc.GetUplinkSets("", "", "", sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... Uplink-Set List .................#")
		for i := 0; i < len(uplinkset_list.Members); i++ {
			fmt.Println(uplinkset_list.Members[i].Name)
		}
	}

	networkUris := new([]utils.Nstring)
	*networkUris = append(*networkUris, utils.NewNstring("/rest/ethernet-networks/cbde97d0-c8f1-4aba-aa86-2b4e5d080401"))
	//      Using make as the [] is going as null
	fcNetworkUris := make([]utils.Nstring, 0)
	fcoeNetworkUris := make([]utils.Nstring, 0)
	portConfigInfos := make([]ov.PortConfigInfos, 0)

	uplinkSet := ov.UplinkSet{Name: "upset77",
		LogicalInterconnectURI:         utils.NewNstring("/rest/logical-interconnects/d4468f89-4442-4324-9c01-624c7382db2d"),
		NetworkURIs:                    *networkUris,
		FcNetworkURIs:                  fcNetworkUris,
		FcoeNetworkURIs:                fcoeNetworkUris,
		PortConfigInfos:                portConfigInfos,
		ConnectionMode:                 "Auto",
		NetworkType:                    "Ethernet",
		EthernetNetworkType:            "Tagged",
		Type:                           "uplink-setV4",
		ManualLoginRedistributionState: "NotSupported"}

	er := ovc.CreateUplinkSet(uplinkSet)
	if er != nil {
		fmt.Println("............... UplinkSet Creation Failed:", er)
	} else {
		fmt.Println(".... Uplink Set Created Successfully")
	}

	new_uplinkset, _ := ovc.GetUplinkSetByName(new_uplink)
	fmt.Println(new_uplinkset)
	new_uplinkset.Name = upd_uplink
	fmt.Println(new_uplinkset)
	err1 := ovc.UpdateUplinkSet(new_uplinkset)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("#.................... Uplink-Set after Updating ...........#")
		uplinkset_after_update, ere := ovc.GetUplinkSets("", "", "", sort)
		if ere != nil {
			fmt.Println(ere)
		} else {
			for i := 0; i < len(uplinkset_after_update.Members); i++ {
				fmt.Println(uplinkset_after_update.Members[i].Name)
			}
		}
	}

	ero := ovc.DeleteUplinkSet(del_uplink)
	if ero != nil {
		fmt.Println(ero)
	} else {
		fmt.Println("#...................... Deleted Uplink Set Successfully .....#")
	}
}
