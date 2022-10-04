package main

import (
	"fmt"
	"strconv"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func main() {
	var (
		ClientOV *ov.OVClient
	)
	// apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	// ovc := ClientOV.NewOVClient(
	// 	os.Getenv("ONEVIEW_OV_USER"),
	// 	os.Getenv("ONEVIEW_OV_PASSWORD"),
	// 	os.Getenv("ONEVIEW_OV_DOMAIN"),
	// 	os.Getenv("ONEVIEW_OV_ENDPOINT"),
	// 	false,
	// 	apiversion,
	// 	"*")

	config, config_err := ov.LoadConfigFile("oneview_config.json")
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

	localelist, err := ovc.GetLocales()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Got the supported appliance locales.----#")
		for i := range localelist.Members {
			fmt.Println("\nLocale " + strconv.Itoa(i+1))
			fmt.Print("DisplayName : " + localelist.Members[i].DisplayName + "\n")
			fmt.Print("Locale : " + localelist.Members[i].Locale + "\n")
			fmt.Print("LocaleName : " + localelist.Members[i].LocaleName + "\n")
		}
	}
}
