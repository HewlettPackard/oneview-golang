package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		ClientOV     *ov.OVClient
		networkset   = "test"
		networkset_2 = "updatednetworkset"
		networkset_3 = "creatednetworkset"
	)
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		2000,
		"*")
	ovVer, _ := ovc.GetAPIVersion()
	fmt.Println(ovVer)

	fmt.Println("#...................NetworkSet by Name ...............#")
	net_set, err := ovc.GetNetworkSetByName(networkset)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(net_set)
	}

	sort := "name:desc"
	networkset_list, err := ovc.GetNetworkSets("", sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... NetworkSet List .................#")
		for i := 0; i < len(networkset_list.Members); i++ {
			fmt.Println(networkset_list.Members[i].Name)
		}
	}

	networkUris := new([]utils.Nstring)
	//append all your network and fc network uri to networkUris
	*networkUris = append(*networkUris, utils.NewNstring("/rest/ethernet-networks/fa01374d-0e50-41fc-bcd7-3dfc21776b6c"))
	*networkUris = append(*networkUris, utils.NewNstring("/rest/ethernet-networks/d085fa2d-41f8-44e1-950b-3f5076371300"))

	NetworkSet := ov.NetworkSet{Name: networkset_3,
		NativeNetworkUri:      "",
		NetworkUris:           *networkUris,
		ConnectionTemplateUri: "",
		Type:                  "network-setV5",
		NetworkSetType:        "Large",
	}
	err = ovc.CreateNetworkSet(NetworkSet)
	if err != nil {
		fmt.Println("............... NetworkSet Creation Failed:", err)
	} else {
		fmt.Println(".... NetworkSet Created Success.......")
	}

	net_set, err = ovc.GetNetworkSetByName(networkset_3)
	net_set.Name = networkset_2
	err = ovc.UpdateNetworkSet(net_set)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#.................... NetworkSet after Updating ...........#")
		networksets_after_update, err := ovc.GetNetworkSets("", sort)
		if err != nil {
			fmt.Println(err)
		} else {
			for i := 0; i < len(networksets_after_update.Members); i++ {
				fmt.Println(networksets_after_update.Members[i].Name)
			}
		}
	}

	err = ovc.DeleteNetworkSet(networkset_2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#...................... Deleted Ethernet Network Successfully .....#")
	}

}
