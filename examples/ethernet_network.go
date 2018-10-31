package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
)

func main() {
	var (
		ClientOV *ov.OVClient
		eth_name = "eth1"
		new_name = "eth77"
		upd_name = "eth88"
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

	fmt.Println("#................... Ethernet Network by Name ...............#")
	eth, _ := ovc.GetEthernetNetworkByName(eth_name)
	fmt.Println(eth)

	sort := "name:desc"
	eth_list, err := ovc.GetEthernetNetworks("", sort)
	if err != nil {
		panic(err)
	}
	fmt.Println("# ................... Ethernet Networks List .................#")
	for i := 0; i < len(eth_list.Members); i++ {
		fmt.Println(eth_list.Members[i].Name)
	}

	eth_id := "02bbab66-4f23-4297-88fa-5420294ec552"
	fmt.Println("#................... GetAssociatedProfiles ....................#")
	eth_af, err := ovc.GetAssociatedProfile(eth_id)
	if err != nil {
		panic(err)
	}
	fmt.Println(eth_af)

	fmt.Println("#................... GetAssociatedUplinkGroups ...............#")
	eth_up, err := ovc.GetAssociatedUplinkGroup(eth_id)
	if err != nil {
		panic(err)
	}
	fmt.Println(eth_up)

	bandwidth := ov.Bandwidth{MaximumBandwidth: 10000, TypicalBandwidth: 2000}

	ethernetNetwork := ov.EthernetNetwork{Name: "eth77", VlanId: 10, Purpose: "General", SmartLink: false, PrivateNetwork: false, ConnectionTemplateUri: "", EthernetNetworkType: "Tagged", Type: "ethernet-networkV4"}

	bulkEthernetNetwork := ov.BulkEthernetNetwork{VlanIdRange: "2-4", Purpose: "General", NamePrefix: "Test_eth", SmartLink: false, PrivateNetwork: false, Bandwidth: bandwidth, Type: "bulk-ethernet-networkV1"}

	er := ovc.CreateEthernetNetwork(ethernetNetwork)
	if er != nil {
		fmt.Println("............... Ethernet Network Creation Failed:", err)
	}
	fmt.Println(".... Ethernet Network Created Success")

	err = ovc.CreateBulkEthernetNetwork(bulkEthernetNetwork)
	if err != nil {
		fmt.Println("............. Bulk Ethernet Network Creation Failed:", err)
	}
	fmt.Println(".... Bulk Ethernet Network Created Success")

	bulk_list, err := ovc.GetEthernetNetworks("", sort)
	for i := 0; i < len(bulk_list.Members); i++ {
		fmt.Println(bulk_list.Members[i].Name)
	}

	new_eth, _ := ovc.GetEthernetNetworkByName(new_name)
	new_eth.Name = upd_name
	err = ovc.UpdateEthernetNetwork(new_eth)
	if err != nil {
		panic(err)
	}
	fmt.Println("#.................... Ethernet Network after Updating ...........#")
	up_list, err := ovc.GetEthernetNetworks("", sort)
	for i := 0; i < len(up_list.Members); i++ {
		fmt.Println(up_list.Members[i].Name)
	}

	eth_del := "ppp"
	err = ovc.DeleteEthernetNetwork(eth_del)
	if err != nil {
		panic(err)
	}
	fmt.Println("#...................... Deleted Ethernet Network Successfully .....#")

}
