package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
//	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		ClientOV *ov.OVClient
		existing_uplink = "IS3"
		new_uplink = "new_uplinkset"
		upd_uplink = "updated_uplinkset"
	)
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		800)
	ovVer, _ := ovc.GetAPIVersion()
	fmt.Println(ovVer)

	fmt.Println("#................... Uplink-Set by Name ...............#")
	uplink_set, err := ovc.GetUplinkSetByName(existing_uplink)
	if err != nil {
		fmt.Println(err)
	}
	else {
	fmt.Println(uplink_set) 
        }

	sort := "name:desc"
	uplinkset_list, err := ovc.GetUplinkSets("", sort)
	if err != nil {
		panic(err)
	}
	fmt.Println("# ................... Uplink-Set List .................#")
	for i := 0; i < len(uplinkset_list.Members); i++ {
		fmt.Println(uplinkset_list.Members[i].Name)
	} 

	networkUris := new([]utils.Nstring)
	*networkUris = append(*networkUris, utils.NewNstring("/rest/ethernet-networks/cbde97d0-c8f1-4aba-aa86-2b4e5d080401"))
//      Using make as the [] is going as null	
	fcNetworkUris := make([]utils.Nstring,0)
	fcoeNetworkUris := make([]utils.Nstring,0)
        portConfigInfos := make([]utils.Nstring,0)

	uplinkSet := ov.UplinkSet{Name: "upset77", LogicalInterconnectURI: utils.NewNstring("/rest/logical-interconnects/d4468f89-4442-4324-9c01-624c7382db2d"), NetworkURIs: *networkUris, FcNetworkURIs: fcNetworkUris, FcoeNetworkURIs: fcoeNetworkUris, PortConfigInfos: portConfigInfos, ConnectionMode: "Auto", NetworkType: "Ethernet", EthernetNetworkType: "Tagged", Type: "uplink-setV4", ManualLoginRedistributionState: "NotSupported"}
     fmt.Println(uplinkSet)
	er := ovc.CreateUplinkSet(uplinkSet)
	if er != nil {
		fmt.Println("............... UplinkSet Creation Failed:", er)
	}
	fmt.Println(".... Uplink Set Created Success")

	new_uplinkset, _ := ovc.GetUplinkSetByName(new_uplink)

	new_uplinkset.Name = upd_uplink
	err := ovc.UpdateUplinkSet(new_uplinkset)
	if err != nil {
		panic(err)
	}
	fmt.Println("#.................... Uplink-Set after Updating ...........#")
	up_uplink_list, ere := ovc.GetUplinkSets("", sort)
        if ere!= nil {
		panic(ere)
	}

	for i := 0; i < len(up_uplink_list.Members); i++ {
		fmt.Println(up_uplink_list.Members[i].Name)
	 } 
	//uplink_del := "upset88"
	ero := ovc.DeleteUplinkSet(upd_uplink)
	if ero != nil {
		panic(ero)
	}
	fmt.Println("#...................... Deleted Uplink Set Successfully .....#")


}
