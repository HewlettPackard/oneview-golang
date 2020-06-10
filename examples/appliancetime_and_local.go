package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
	"strconv"
)

func main() {
	var (
		ClientOV    *ov.OVClient
		falseVar    = false
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
	ntpServers := new([]utils.Nstring)
	*ntpServers = append(*ntpServers, utils.NewNstring("16.110.135.123"))
	*ntpServers = append(*ntpServers, utils.NewNstring("16.85.40.52"))
	applianceTimeandLocal := ov.ApplianceTimeandLocal
		locale:     "en_US.UTF-8",
		dateTime:   "2014-09-11T12:10:33",
		timezone:   "UTC",
		ntpServers: *ntpServers,
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
		for i := range applianceTimeandLocals.Members {
			fmt.Println(applianceTimeandLocals.Members[i])
		}
	}
}
