package main

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/i3s"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"io/ioutil"
	"os"
)

func main() {
	var (
		clientOV             *ov.OVClient
		i3sClient            *i3s.I3SClient
		endpoint             = os.Getenv("ONEVIEW_OV_ENDPOINT")
		i3sc_endpoint        = os.Getenv("ONEVIEW_I3S_ENDPOINT")
		username             = os.Getenv("ONEVIEW_OV_USER")
		password             = os.Getenv("ONEVIEW_OV_PASSWORD")
		domain               = os.Getenv("ONEVIEW_OV_DOMAIN")
		api_version          = 1600
		deployment_plan_name = "DemoDeploymentPlan"
		new_name             = "RenamedDeploymentPlan"
	)
	ovc := clientOV.NewOVClient(
		username,
		password,
		domain,
		endpoint,
		false,
		api_version,
		"*")
	ovc.RefreshLogin()
	i3sc := i3sClient.NewI3SClient(i3sc_endpoint, false, api_version, ovc.APIKey)

	customAttributes := new([]i3s.CustomAttribute)
	var ca1, ca5, ca6, ca7, ca8, ca9  i3s.CustomAttribute



	ca1.Constraints = "{\"options\":[\"enabled\",\"disabled\"]}"
	ca1.Editable = true
	ca1.ID = "9303c168-3e23-4025-9e63-2bddd644b461"
	ca1.Name = "SSH"
	ca1.Type = "option"
	ca1.Value = "enabled"
	ca1.Visible = true
	*customAttributes = append(*customAttributes, ca1)
	


	ca5.Constraints = "{\"helpText\":\"\"}"
	ca5.Editable = true
	ca5.ID = "4c7c8229-1dc6-4656-857d-3392a26585ee"
	ca5.Name = "DomainName"
	ca5.Type = "string"
	ca5.Visible = true
	*customAttributes = append(*customAttributes, ca5)

	ca7.Constraints = "{\"helpText\":\"\"}"
	ca7.Editable = true
	ca7.ID = "3c7c8229-1dc6-4656-857d-3392a26585ee"
	ca7.Name = "Hostname"
	ca7.Type = "string"
	ca7.Visible = true
	*customAttributes = append(*customAttributes, ca7)

	ca8.Constraints = "{\"ipv4static\":true,\"ipv4dhcp\":true,\"ipv4disable\":false,\"parameters\":[\"dns1\",\"dns2\",\"gateway\",\"ipaddress\",\"mac\",\"netmask\",\"vlanid\"]}"
	ca8.Editable = true
	ca8.ID = "ec1d95d0-690a-482b-8efd-53bec6e9bfce"
	ca8.Name = "ManagementNIC"
	ca8.Type = "nic"
	ca8.Visible = true
	*customAttributes = append(*customAttributes, ca8)
	
	ca6.Constraints = "{\"ipv4static\":true,\"ipv4dhcp\":false,\"ipv4disable\":false,\"parameters\":[\"mac\",\"vlanid\"]}"
	ca6.Editable = true
	ca6.ID = "ec1d95d0-6h0a-482b-8efd-53bec6e9bfce"
	ca6.Name = "ManagementNIC2"
	ca6.Type = "nic"
	ca6.Visible = true
	*customAttributes = append(*customAttributes, ca6)

	ca9.Constraints = "{\"options\":[\"\"]}"
	ca9.Editable = true
	ca9.Description = "Administrator Password"
	ca9.ID = "a881b1af-9034-4c69-a890-5e8c83a13d25"
	ca9.Name = "Password"
	ca9.Type = "password"
	ca9.Visible = true
	*customAttributes = append(*customAttributes, ca9)



	var deploymentPlan i3s.DeploymentPlan
	deploymentPlan.Name = deployment_plan_name
	deploymentPlan.Type = "OEDeploymentPlanV5"
	deploymentPlan.OEBuildPlanURI = "/rest/build-plans/ff13e3d5-b4c3-47cd-bd99-b6e3dd2c4e66"
	deploymentPlan.CustomAttributes = *customAttributes
	deploymentPlan.HPProvided = false

	fmt.Println("HPProvided:", deploymentPlan.HPProvided)
	file, _ := json.MarshalIndent(deploymentPlan, "", " ")
	ioutil.WriteFile("inut.json", file, 0644)
	fmt.Println("***********Creating Deployment Plan****************")
	err := i3sc.CreateDeploymentPlan(deploymentPlan)
	if err != nil {
		fmt.Println("Deployment Plan Creation Failed: ", err)
	} else {
		fmt.Println("Deployment Plan created successfully...")
	}

	sort := "name:desc"
	count := "5"
	fmt.Println("**************Get Deployment Plans sorted by name in descending order****************")
	dps, err := i3sc.GetDeploymentPlans(count, "", "", sort, "")
	if err != nil {
		fmt.Println("Error while getting deployment plans:", err)
	} else {
		for i := range dps.Members {
			fmt.Println(dps.Members[i].Name)
		}
	}

	fmt.Println("***********Getting Deployment Plan By Name****************")
	deployment_plan, err := i3sc.GetDeploymentPlanByName(deployment_plan_name)
	if err != nil {
		fmt.Println("Error in getting deployment plan ", err)
	}
	fmt.Println(deployment_plan)

	fmt.Println("***********Updating Deployment Plan****************")
	deployment_plan.Name = new_name
	deployment_plan.GoldenImageUri = "/rest/golden-images/7e709af9-5446-426e-9ca1-df06c63df2cd"
	deployment_plan.Description = utils.NewNstring("Testing Deployment plan")
	err = i3sc.UpdateDeploymentPlan(deployment_plan)
	if err != nil {
		//panic(err)
		fmt.Println("Error whilw updating Deployment Plan:", err)
	} else {
		fmt.Println("Deployment Plan has been updated with name: " + deploymentPlan.Name)
	}

	fmt.Println("***********Deleting Deployment Plan****************")
	err = i3sc.DeleteDeploymentPlan(new_name)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleteed Deployment Plan successfully...")
	}
}
