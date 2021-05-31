package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
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

	// Get all the Firmware Baseline available
	firmware, err := ovc.GetFirmwareBaselineList("", "", "")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#-----Got Firmware Baseline-----#")
		for i := range firmware.Members {
			fmt.Println(firmware.Members[i])
		}
	}

	id := strings.Split(firmware.Members[0].Uri.String(), "/")[3] //eg: Synergy_Custom_SPP_2021_02_01_Z7550-97110
	// Get Firmware Baseline by id
	firmware2, err := ovc.GetFirmwareBaselineByUri(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-----Got Firmware Baseline by Uri-----#")
		fmt.Println(firmware2)
	}

	//create custom service pack
	scope := ov.Scope{Name: "ScopTest", Description: "Test from script", Type: "ScopeV3"}
	_ = ovc.CreateScope(scope)
	scp, _ := ovc.GetScopeByName("ScopTest")
	initialScopeUris := &[]utils.Nstring{scp.URI}

	hotfix := &[]utils.Nstring{firmware.Members[1].Uri} // iniialize Hotfix Uri
	customSP := ov.CustomServicePack{
		CustomBaselineName: "Custom Service Pack",
		BaselineUri:        firmware2.Uri.String(), // initialize Service pack Uri
		InitialScopeUris:   *initialScopeUris,
		HotfixUris:         *hotfix,
	}
	err = ovc.CreateCustomServicePack(customSP, "false") // force parameter is set as false
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-----Custom Service Pack Created Successfully-----#")
	}

	//Delete Firmware Baseline
	err = ovc.DeleteFirmwareBaseline(id, "false") // force parameter is set as false
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-----Firmware Baseline Deleted Successfully-----#")
	}

}
