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
		scope                           = "<scope_name>"
		hypervisor_manager_ip           = "172.18.13.11"
		hypervisor_manager_display_name = "HM2"
	)
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		2200,
		"")

	scp, _ := ovc.GetScopeByName(scope)
	initialScopeUris := &[]utils.Nstring{scp.URI}

	// Adding Hypervisor Manager Server Certificate to Oneview for Secure conection
	server_cert, err := ovc.GetServerCertificateByIp(hypervisor_manager_ip)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Fetched Hypervisor Manager Server Certificate.")
	}
	server_cert.CertificateDetails[0].AliasName = "Hypervisor Manager Server Certificate"
	server_cert.Type = ""
	er := ovc.CreateServerCertificate(server_cert)
	if er != nil {
		fmt.Println("............... Adding Server Certificate Failed: ", er)
	} else {
		fmt.Println("Imported Hypervisor Manager Server Certificate to Oneview for secure connection successfully.")
	}

	hypervisorManager := ov.HypervisorManager{DisplayName: "HM1",
		Name:             "172.18.13.11",
		Username:         "dcs",
		Password:         "dcs",
		Port:             443,
		InitialScopeUris: *initialScopeUris,
		Type:             "HypervisorManagerV2"}

	err = ovc.CreateHypervisorManager(hypervisorManager)
	if err != nil {
		fmt.Println("............... Create Hypervisor Manager Failed:", err)
	} else {
		fmt.Println(".... Create Hypervisor Manager Success")
	}

	fmt.Println("#................... Hypervisor Manager by Name ...............#")
	hypervisor_mgr, err := ovc.GetHypervisorManagerByName(hypervisor_manager_ip)
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

	hypervisor_mgr.DisplayName = hypervisor_manager_display_name
	err = ovc.UpdateHypervisorManager(hypervisor_mgr)
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

	err = ovc.DeleteHypervisorManager(hypervisor_manager_ip)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#...................... Deleted Hypervisor Manager Successfully .....#")
	}

}
