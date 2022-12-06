package main

import (
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
)

func main() {
	var (
		ClientOV *ov.OVClient
	)
	// Use configuratin file to set the ip and  credentails
	config, config_err := ov.LoadConfigFile("config.json")
	if config_err != nil {
		fmt.Println(config_err)
	}
	ovc := ClientOV.NewOVClient(
		config.OVCred.UserName,
		config.OVCred.Password,
		config.OVCred.Domain,
		config.OVCred.Endpoint,
		config.OVCred.SslVerify,
		config.OVCred.ApiVersion,
		config.OVCred.IfMatch)

	var (
		hcp_name                     = "test"
		new_hcp                      = "test_new"
		scope                        = "ScopeTest"
		server_profile_template_name = "Auto-SPT"
		hypervisor_manager_ip        = config.HypervisorManagerConfig.IpAddress
	)

	initialScopeUris := new([]utils.Nstring)
	scp, scperr := ovc.GetScopeByName(scope)
	if scperr != nil {
		fmt.Println("Error fetching scope: ", scperr)
	}
	*initialScopeUris = append(*initialScopeUris, scp.URI)

	hypervisor_manager, err := ovc.GetHypervisorManagerByName(hypervisor_manager_ip)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hypervisor Manager: ", hypervisor_manager.URI)
	}

	server_profile_template, err := ovc.GetProfileTemplateByName(server_profile_template_name)
	if err != nil {
		fmt.Println("Server Profile Template Retrieval Failed: ", err)
	}

	hypervisorHostProfileTemplate := ov.HypervisorHostProfileTemplate{
		ServerProfileTemplateUri: server_profile_template.URI,
		DeploymentManagerType:    "UserManaged",
		Hostprefix:               "test"}
	hypervisorclustprof := ov.HypervisorClusterProfile{
		Type:                          "HypervisorClusterProfileV6",
		Name:                          hcp_name,
		Description:                   "",
		HypervisorType:                "Vmware",
		HypervisorManagerUri:          hypervisor_manager.URI,
		Path:                          "DC1",
		HypervisorHostProfileTemplate: &hypervisorHostProfileTemplate}
	fmt.Println(hypervisorclustprof)

	err = ovc.CreateHypervisorClusterProfile(hypervisorclustprof)
	if err != nil {
		fmt.Println("Server HypervisorClusterProfile Create Failed: ", err)
	} else {
		fmt.Println("#----------------Server HypervisorClusterProfile Created---------------#")
	}

	sort := ""
	id := ""
	hcp_list, err := ovc.GetHypervisorClusterProfiles("", "", "", sort)
	if err != nil {
		fmt.Println("HypervisorClusterProfile Retrieval Failed: ", err)
	} else {
		fmt.Println("#----------------HypervisorClusterProfile List---------------#")

		for i := 0; i < len(hcp_list.Members); i++ {
			fmt.Println(hcp_list.Members[i].Name)
			if hcp_list.Members[i].Name == hcp_name {
				hcp_uri := hcp_list.Members[i].URI
				id = string(hcp_uri[len("/rest/hypervisor-cluster-profiles/"):])
				fmt.Println(id)
			}

		}
	}

	hcp1, err := ovc.GetHypervisorClusterProfileById(id)
	if err != nil {
		fmt.Println("HypervisorClusterProfile Retrieval By Id Failed: ", err)
	} else {
		fmt.Println("#----------------HypervisorClusterProfile by Id---------------#")
		fmt.Println(hcp1.Name)
	}

	hcp2, err := ovc.GetHypervisorClusterProfileCompliancePreview(id)
	if err != nil {
		fmt.Println("HypervisorClusterProfile Compliance Preview retrieval Failed: ", err)
	} else {
		fmt.Println("#----------------HypervisorClusterProfile Compliance Preview---------------#")
		fmt.Println(hcp2)
	}

	vswitchlayout := ov.VirtualSwitchLayout{
		HypervisorManagerUri:     hypervisor_manager.URI,
		ServerProfileTemplateUri: server_profile_template.URI}
	err = ovc.CreateVirtualSwitchLayout(vswitchlayout)
	if err != nil {
		fmt.Println("Create VirtualSwitchLayou Failed: ", err)
	} else {
		fmt.Println("#----------------Virtual Switch LayoutCreated---------------#")
	}
	hcp1.Name = new_hcp

	err = ovc.UpdateHypervisorClusterProfile(hcp1)
	if err != nil {
		fmt.Println("HypervisorClusterProfile Create Failed: ", err)
	} else {
		fmt.Println("#----------------HypervisorClusterProfile updated---------------#")
	}

	//Delete function accepts 2 optional arguments - softDelete(boolean) and force(boolean) till API1200
	//softDelete is mandatory argument for delete function from API 1600
	if ovc.APIVersion > 1200 {
		err = ovc.DeleteHypervisorClusterProfileSoftDelete(new_hcp, false)
		if err != nil {
			fmt.Println("HypervisorClusterProfile Delete Failed: ", err)
		} else {
			fmt.Println("#---------------HypervisorClusterProfile Deleted---------------#")
		}
	} else {
		err = ovc.DeleteHypervisorClusterProfile(new_hcp)
		if err != nil {
			fmt.Println("HypervisorClusterProfile Delete Failed: ", err)
		} else {
			fmt.Println("#---------------HypervisorClusterProfile Deleted---------------#")
		}
	}

}
