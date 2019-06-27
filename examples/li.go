package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
)

func main() {
	var (
		clientOV *ov.OVClient
		id       = "d4468f89-4442-4324-9c01-624c7382db2d"
	)

	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		"P@ssw0rd!",
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		800,
		"*")

	fmt.Println("....  Logical Interconnects Collection .....")
	fmt.Println(ovc.GetAPIVersion())
	logicalInterconnectList, _ := ovc.GetLogicalInterconnects("", "", "")
	//        fmt.Println(logicalInterconnectList)
	for i := range logicalInterconnectList.Members {
		fmt.Println(logicalInterconnectList.Members[i].Name)
	}

	fmt.Println("....  Logical Interconnect by Id.....")
	lig, _ := ovc.GetLogicalInterconnectById(id)
	fmt.Println("LI by ID")
	fmt.Println(lig.Name)

	err_update_compliance := ovc.UpdateLogicalInterconnectConsistentStateById(id)

	if err_update_compliance != nil {
		fmt.Println("Could not update ConsistentState of Logical Interconnect", err_update_compliance)
	}
}
