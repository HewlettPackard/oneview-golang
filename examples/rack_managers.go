package main

import (
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func main() {
	var (
		ClientOV        *ov.OVClient
		RackManagerName = "5UF7201002" //"TestRackManagerGOsdk"
		// new_fc_name = "RenamedFCNetwork"
		// falseVar    = false
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

	// scope := ov.Scope{Name: "ScopTest", Description: "Test from script", Type: "ScopeV3"}
	// _ = ovc.CreateScope(scope)
	// scp, _ := ovc.GetScopeByName("ScopTest")

	sort := "name:desc"

	rManagers, err := ovc.GetRackManagerList("", "", "", sort, "")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get Rack Managers sorted by name in descending order----#")
		for i := range rManagers.Members {
			fmt.Println(rManagers.Members[i].Name)
		}
	}
	rm1, err := ovc.GetRackManagerByName(RackManagerName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get Rack Manager by name----------------#")
		fmt.Println(rm1)
	}

	rackManager := ov.RackManager{
		Hostname: "172.18.9.3",
		UserName: "dcs",
		Password: "dcs",

		//InitialScopeUris: *initialScopeUris,
	}

	rmUri, err := ovc.AddRackManager(rackManager)
	if err != nil {
		fmt.Println("............... Add Rack Manager Failed:", err)
	} else {
		fmt.Println(".... Create Rack Manager Success", rmUri)
	}
	rmCreated, err := ovc.GetRackManagerByUri(rmUri)
	fmt.Println(rmCreated.Name)
	rmcName := rmCreated.Name

	err = ovc.DeleteRackManager(rmcName)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleted Rack manager successfully...")
	}

	// fcNetwork01 := ov.FCNetwork{
	// 	AutoLoginRedistribution: falseVar,
	// 	Description:             "Test FC Network 1",
	// 	LinkStabilityTime:       30,
	// 	FabricType:              "FabricAttach",
	// 	Name:                    "testName1",
	// 	Type:                    "fc-networkV4",
	// }
	// err = ovc.CreateFCNetwork(fcNetwork01)

	// fcNetwork02 := ov.FCNetwork{
	// 	AutoLoginRedistribution: falseVar,
	// 	Description:             "Test FC Network 2",
	// 	LinkStabilityTime:       30,
	// 	FabricType:              "FabricAttach",
	// 	Name:                    "testName2",
	// 	Type:                    "fc-networkV4",
	// }
	// err = ovc.CreateFCNetwork(fcNetwork02)

	// fcNetwork1, err := ovc.GetFCNetworkByName("testName1")
	// fcNetwork2, err = ovc.GetFCNetworkByName("testName2")

	// network_uris := &[]utils.Nstring{fcNetwork1.URI, fcNetwork2.URI}
	// bulkDeleteFCNetwork := ov.FCNetworkBulkDelete{FCNetworkUris: *network_uris}
	// err = ovc.DeleteScope("ScopTest")
	// err = ovc.DeleteBulkFcNetwork(bulkDeleteFCNetwork)

	// if err != nil {
	// 	fmt.Println("............. FC Network Bulk-Deletion Failed:", err)
	// } else {
	// 	fmt.Println(".... FC Network Bulk-Delete is Successful")
	// }

}
