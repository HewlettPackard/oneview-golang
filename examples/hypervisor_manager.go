package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		ClientOV                        *ov.OVClient
		hypervisor_manager              = "172.18.13.11"
		hypervisor_manager_display_name = "HM2"
		//hypervisor_manager_2 = "eth88"
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
	initialScopeUris := &[]utils.Nstring{utils.NewNstring("/rest/scopes/03beb5a0-bf48-4c43-94a5-74b7b5de1255")}
	hypervisorManager := ov.HypervisorManager{DisplayName: "HM1", Name: "172.18.13.11", Username: "dcs", Password: "dcs", Port: 443, InitialScopeUris: *initialScopeUris, Type: "HypervisorManagerV2"}

	err := ovc.AddHypervisorManager(hypervisorManager)
	if err != nil {
		fmt.Println("............... Adding Hypervisor Manager Failed:", err)
	} else {
		fmt.Println(".... Adding Hypervisor Man Success")
	}

	fmt.Println("#................... Hypervisor Manager by Name ...............#")
	hypervisor_mgr, err := ovc.GetHypervisorManagerByName(hypervisor_manager)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hypervisor_mgr)
	}

	sort := "name:desc"
	hypervisor_mgr_list, err := ovc.GetHypervisorManagers("", "", "", sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... Hypervisor Managers List .................#")
		for i := 0; i < len(hypervisor_mgr_list.Members); i++ {
			fmt.Println(hypervisor_mgr_list.Members[i].Name)
		}
	}

	hypervisor_mgnr, _ := ovc.GetHypervisorManagerByName(hypervisor_manager)
	hypervisor_mgnr.DisplayName = hypervisor_manager_display_name
	err = ovc.UpdateHypervisorManager(hypervisor_mgnr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#.................... Hypervisor Manager after Updating ...........#")
		hypervisor_mgr_after_update, err := ovc.GetHypervisorManagers("", "", "", sort)
		if err != nil {
			fmt.Println(err)
		} else {
			for i := 0; i < len(hypervisor_mgr_after_update.Members); i++ {
				fmt.Println(hypervisor_mgr_after_update.Members[i].Name)
			}
		}
	}

	err = ovc.DeleteHypervisorManager(hypervisor_manager)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#...................... Deleteed Hypervisor Manager Successfully .....#")
	}

}
