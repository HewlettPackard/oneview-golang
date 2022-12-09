package main

import (
	"fmt"
	"strconv"
	"strings"

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

	// Get all the connection template available
	connTemplate, err := ovc.GetConnectionTemplate("", "", "", "")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#-----Got Connection Templates-----#")
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
		fmt.Println("#-----Got Connection Template by name-----#")
		fmt.Println(connTemplate2)
	}

	templatebyUri, err := ovc.GetConnectionTemplateByURI(connTemplate2.URI)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-----Got Connection Template by URI-----#")
		fmt.Println(templatebyUri)
	}

	// Get the default connection template
	default_connection, err := ovc.GetDefaultConnectionTemplate()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-----Got Default Connection Template-----#")
		fmt.Println(default_connection)
	}

	Bandwidthoptions := ov.BandwidthType{
		MaximumBandwidth: 8000,
		TypicalBandwidth: 2500,
	}

	templateoptions := ov.ConnectionTemplate{
		Bandwidth: Bandwidthoptions,
		Name:      testName,
		Type:      "connection-template",
	}
	fmt.Println(templateoptions)
	// updating Connection Template
	// specific id can be given for update.
	id := strings.Split(connTemplate2.URI.String(), "/")[3] // eg: id = 063f055c-2cda-4420-be9d-024d609bc534
	template, err := ovc.UpdateConnectionTemplate(id, templateoptions)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection Template has been updated with maximum bandwidth: " + strconv.Itoa(template.Bandwidth.MaximumBandwidth))
	}

	// revert back the changes
	Bandwidthoption := ov.BandwidthType{
		MaximumBandwidth: 10000,
		TypicalBandwidth: 2500,
	}

	templateoption := ov.ConnectionTemplate{
		Bandwidth: Bandwidthoption,
		Name:      testName,
		Type:      "connection-template",
	}
	fmt.Println(templateoption)
	template2, err := ovc.UpdateConnectionTemplate(id, templateoption)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection Template has been updated with maximum bandwidth: " + strconv.Itoa(template2.Bandwidth.MaximumBandwidth))
	}

}
