package main

import (
	"encoding/json"
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
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

	spt_name := config.ServerProfileTemplateConfig.ServerPrpofileTemplateName

	// Gets all labels
	responseAllLabels, err := ovc.GetAllLabels("", "", "", "", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get all labels----------------#")
		jsonResponse, _ := json.MarshalIndent(responseAllLabels, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}

	labels := []ov.Label{}
	labels = append(labels, ov.Label{
		Name: "NewestLabel",
	})
	spt_resource, _ := ovc.GetProfileTemplateByName(spt_name)
	// Creates new labels
	label := ov.AssignedLabel{
		ResourceUri: spt_resource.URI,
		Labels:      labels,
	}
	responseLabel, err := ovc.CreateLabel(label)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Labels created----------------#")
		jsonResponse, _ := json.MarshalIndent(responseLabel, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}

	// Gets Assigned Labels
	responseLabel, err = ovc.GetAssignedLabels(responseLabel.ResourceUri)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Assigned labels----------------#")
		jsonResponse, _ := json.MarshalIndent(responseLabel, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}

	// Updates Assigned Labels
	responseLabel.Labels[0].Name = "UpdatedLabel"
	responseLabel, err = ovc.UpdateAssignedLabels(responseLabel)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Updated assigned labels----------------#")
		jsonResponse, _ := json.MarshalIndent(responseLabel, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}

	// Gets a Label By Uri
	response, err := ovc.GetLabelByURI(responseLabel.Labels[0].Uri)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get label by uri----------------#")
		jsonResponse, _ := json.MarshalIndent(response, "", "  ")
		fmt.Print(string(jsonResponse), "\n\n")
	}

	// Deletes all labels from the resource
	err = ovc.DeleteAssignedLabel(string(responseLabel.ResourceUri))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------All labels to assigned recource are deleted----------------#")
	}
}
