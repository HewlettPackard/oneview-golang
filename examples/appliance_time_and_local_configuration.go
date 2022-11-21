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
	ntpServers := new([]utils.Nstring)
	*ntpServers = append(*ntpServers, utils.NewNstring("127.0.0.1"))
	applianceTimeandLocal := ov.ApplianceTimeandLocal{
		Locale: "en_US.UTF-8",
		// DateTime:   "2014-09-11T12:10:33",
		Timezone:   "UTC",
		NtpServers: *ntpServers,
	}
	fmt.Println(applianceTimeandLocal)
	err := ovc.CreateApplianceTimeandLocal(applianceTimeandLocal)
	if err != nil {
		fmt.Println("ApplianceTime and Local Creation Failed: ", err)
	} else {
		fmt.Println("ApplianceTime and Local created successfully...")
	}

	applianceTimeandLocals, err := ovc.GetApplianceTimeandLocals("", "", "", "")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get ApplianceTime and Local ----#")
		fmt.Println(applianceTimeandLocals)
	}
}
