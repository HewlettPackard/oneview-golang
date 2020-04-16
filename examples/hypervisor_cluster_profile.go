package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		clientOV                *ov.OVClient
		hcp_name                = "test"
		new_hcp                 = "test_new"
		server_profile_template = utils.Nstring("/rest/server-profile-templates/278cadfb-2e86-4a05-8932-972553518259")
		hypervisor_manager      = utils.Nstring("/rest/hypervisor-managers/1ded903a-ac66-41cf-ba57-1b9ded9359b6")
		/*		sp_sn       = "VCGRE1S007"
				new_sp_name = "Renamed Server HypervisorClusterProfile"
		*/)
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		1200,
		"*")

	initialScopeUris := new([]utils.Nstring)
	*initialScopeUris = append(*initialScopeUris, utils.NewNstring("/rest/scopes/74877630-9a22-4061-9db4-d12b6c4cfee0"))

	hypervisorHostProfileTemplate := ov.HypervisorHostProfileTemplate{
		ServerProfileTemplateUri: server_profile_template,
		Hostprefix:               "test"}

	hypervisorclustprof := ov.HypervisorClusterProfile{
		Type:                          "HypervisorClusterProfileV3",
		Name:                          hcp_name,
		Description:                   "",
		HypervisorType:                "Vmware",
		HypervisorManagerUri:          hypervisor_manager,
		Path:                          "DC1",
		HypervisorHostProfileTemplate: &hypervisorHostProfileTemplate}
	fmt.Println(hypervisorclustprof)
	err := ovc.CreateHypervisorClusterProfile(hypervisorclustprof)
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
		HypervisorManagerUri:     hypervisor_manager,
		ServerProfileTemplateUri: server_profile_template}
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

	err = ovc.DeleteHypervisorClusterProfile(id)
	if err != nil {
		fmt.Println("HypervisorClusterProfile Delete Failed: ", err)
	} else {
		fmt.Println("#---------------HypervisorClusterProfile Deleted---------------#")
	}

}
