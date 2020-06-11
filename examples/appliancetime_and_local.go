package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	// "os"
	// "strconv"
)

func main() {
	var (
		ClientOV *ov.OVClient
	)
	// apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	ovc := ClientOV.NewOVClient(
		"Administrator",
		"admin123",
		"LOCAL",
		"https://10.50.9.90/",
		false,
		1600,
		"*")
	ntpServers := new([]utils.Nstring)
	*ntpServers = append(*ntpServers, utils.NewNstring("16.110.135.123"))
	//	*ntpServers = append(*ntpServers, utils.NewNstring("16.85.40.52"))
	applianceTimeandLocal := ov.ApplianceTimeandLocal{
		Locale:     "en_US.UTF-8",
		DateTime:   "2014-09-11T12:10:33",
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
