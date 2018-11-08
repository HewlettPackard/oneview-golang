package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
)

func main() {
	var (
		ClientOV *ov.OVClient
		upset_name = "upset1"
		new_name = "upset77"
		upd_name = "upset88"
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
	upset, _ := ovc.GetUplinkSetByName(upset_name)
	fmt.Println(upset)

	sort := "name:desc"
	upset_list, err := ovc.GetUplinkSets("", sort)
	if err != nil {
		panic(err)
	}
	fmt.Println("# ................... Uplink-Set List .................#")
	for i := 0; i < len(upset_list.Members); i++ {
		fmt.Println(upset_list.Members[i].Name)
	}

/*	upset_id := "02bbab66-4f23-4297-88fa-5420294ec552"
	fmt.Println("#................... GetAssociatedProfiles ....................#")
	up_af, err := ovc.GetAssociatedProfile(upset_id)
	if err != nil {
		panic(err)
	}
	fmt.Println(up_af)

	fmt.Println("#................... GetAssociatedUplinkGroups ...............#")
	eth_up, err := ovc.GetAssociatedUplinkGroup(eth_id)
	if err != nil {
		panic(err)
	}
	fmt.Println(eth_up)

	bandwidth := ov.Bandwidth{MaximumBandwidth: 10000, TypicalBandwidth: 2000}
*/
uplinkSet := ov.UplinkSet{Name: "upset77", LogicalInterconnectURI: "/rest/logical-interconnects/", NetworkURIs: "/rest/uplink-sets/", FcNetworkURIs: "[]", FcoeNetworkURIs: "[]", PortConfigInfos: "[]", ConnectionMode: "Auto", NetworkType: "Ethernet", ManualLoginRedistributionState: "NotSupported"}

//	bulkEthernetNetwork := ov.BulkEthernetNetwork{VlanIdRange: "2-4", Purpose: "General", NamePrefix: "Test_eth", SmartLink: false, PrivateNetwork: false, Bandwidth: bandwidth, Type: "bulk-ethernet-networkV1"}

	er := ovc.CreateUplinkSet(uplinkSet)
	if er != nil {
		fmt.Println("............... UplinkSet Creation Failed:", err)
	}
	fmt.Println(".... Uplink Set Created Success")

/*	err = ovc.CreateBulkEthernetNetwork(bulkEthernetNetwork)
	if err != nil {
		fmt.Println("............. Bulk Ethernet Network Creation Failed:", err)
	}
	fmt.Println(".... Bulk Ethernet Network Created Success")

	bulk_list, err := ovc.GetEthernetNetworks("", sort)
	for i := 0; i < len(bulk_list.Members); i++ {
		fmt.Println(bulk_list.Members[i].Name)
	}
*/
	new_upset, _ := ovc.GetUplinkSetByName(new_name)
	new_upset.Name = upd_name
	err = ovc.UpdateUplinkSet(new_upset)
	if err != nil {
		panic(err)
	}
	fmt.Println("#.................... Uplink-Set after Updating ...........#")
	up_list, err := ovc.GetUplinkSets("", sort)
	for i := 0; i < len(up_list.Members); i++ {
		fmt.Println(up_list.Members[i].Name)
	}

	uplink_del := "ppp"
	err = ovc.DeleteUplinkSet(uplink_del)
	if err != nil {
		panic(err)
	}
	fmt.Println("#...................... Deleted Uplink Set Successfully .....#")

}
