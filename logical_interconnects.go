package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
//	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
	"strconv"
	"strings"
)

func newTrue() *bool {
	b := true
	return &b
}
func newFalse() *bool {
	b := false
	return &b
}

func main() {
	var (
		clientOV         *ov.OVClient
		//ethernet_network = "Auto-ethernet_network"
		//tcId             = "1"
	)
	apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
		"*")

	logicalInterconnect, err := ovc.GetLogicalInterconnects("", "", "")
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("#-------------All Logical Interconnect Names----------------#")
		fmt.Println(logicalInterconnect.Members[0].Name)
	}
	//interconnectURI := string(logicalInterconnect.Members[0].URI)
	id := strings.Replace(string(logicalInterconnect.Members[0].URI), "/rest/logical-interconnects/", "", 1)

	fmt.Println("....  Logical Interconnects Collection .....")
	logicalInterconnectList, _ := ovc.GetLogicalInterconnects("", "0", "10")
	fmt.Println(logicalInterconnectList)

	fmt.Println("....  Logical Interconnect by Id.....")
	lig, _ := ovc.GetLogicalInterconnectById(id)
	fmt.Printf("%+v\n", lig)
}
