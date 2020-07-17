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
		server_profile_template = utils.Nstring("/rest/server-profile-templates/fee8d166-ad5b-49a1-a568-131759b7c1cc")
		hypervisor_manager      = utils.Nstring("/rest/hypervisor-managers/ee0e59da-6511-4d7b-bdfc-50445ef3dbcd")
	)
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		1800,
		"*")

	initialScopeUris := new([]utils.Nstring)
	*initialScopeUris = append(*initialScopeUris, utils.NewNstring("/rest/scopes/8ef32b43-2478-4aea-bb68-a65d0fbfea93"))

	deploymentPlan := ov.DeploymentPlan{
		ServerPassword: "dcs"}

	hypervisorHostProfileTemplate := ov.HypervisorHostProfileTemplate{
		ServerProfileTemplateUri: server_profile_template,
		DeploymentPlan:           &deploymentPlan,
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
