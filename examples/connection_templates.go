package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/HewlettPackard/oneview-golang/ov"
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

	// Get all the connection template available
	connTemplate, err := ovc.GetConnectionTemplate("", "", "", "")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Got Connection Templates----#")
		for i := range connTemplate.Members {
			fmt.Println(connTemplate.Members[i])
		}
	}

	testName := connTemplate.Members[0].Name
	// Get connection template by name
	connTemplate2, err := ovc.GetConnectionTemplateByName(testName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Got Connection Template by name----------------#")
		fmt.Println(connTemplate2)
	}

	// Get the default connection template
	default_connection, err := ovc.GetDefaultConnectionTemplate()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#--------Got Default Connection Template-------#")
		fmt.Println(default_connection)
	}

	// updating Connection Template
	connTemplate2.Bandwidth.MaximumBandwidth = 8000
	err = ovc.UpdateConnectionTemplate(connTemplate2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection Template has been updated with maximum bandwidth: " + strconv.Itoa(connTemplate2.Bandwidth.MaximumBandwidth))
	}

	// revert back the changes
	connTemplate2.Bandwidth.MaximumBandwidth = 10000
	err = ovc.UpdateConnectionTemplate(connTemplate2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection Template has been updated with maximum bandwidth: " + strconv.Itoa(connTemplate2.Bandwidth.MaximumBandwidth))
	}

}
