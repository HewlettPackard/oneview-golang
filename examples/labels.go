package main

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
	"strconv"
)

func main() {
	var (
		ClientOV *ov.OVClient
	)
	apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
		"*")

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

	// Creates new labels
	label := ov.AssignedLabel{
		ResourceUri: utils.Nstring("/rest/server-profile-templates/b6777c57-34f1-4491-93c4-8bec773f286c"),
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
