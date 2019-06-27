
package main

import (
        "fmt"
        "github.com/HewlettPackard/oneview-golang/ov"
        "os"
)

func main() {
        var (
                clientOV        *ov.OVClient
                id              = "fe7dec60-896a-4418-9b3d-24ae9ba6fa6f"
	)

	ovc := clientOV.NewOVClient(
                os.Getenv("ONEVIEW_OV_USER"),
                os.Getenv("ONEVIEW_OV_PASSWORD"),
                os.Getenv("ONEVIEW_OV_DOMAIN"),
                os.Getenv("ONEVIEW_OV_ENDPOINT"),
                false,
                800,
                "*")

	fmt.Println("....  Logical Interconnects Collection .....")
        logicalInterconnectList, _ := ovc.GetLogicalInterconnects("", "0", "10")
        fmt.Println(logicalInterconnectList)

        fmt.Println("....  Logical Interconnect by Id.....")
        lig, _ := ovc.GetLogicalInterconnectById(id)
	fmt.Println("LI by ID")
        fmt.Println(lig)

	err_compliance := ovc.UpdateLogicalInterconnectConsistentStateById(id)

	if err_compliance != nil {
                fmt.Println("Could not update ConsistentState of Logical Interconnect", err_compliance)
	}
}
