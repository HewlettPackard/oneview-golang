package main

import (
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func main() {

	var (
		ClientOV     *ov.OVClient
		storage_pool = "CPG-SSD"
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

	//Get storage pool by name to update
	update_pool, _ := ovc.GetStoragePoolByName(storage_pool)

	// Update the given storage pool
	// This API can be used to manage/unmanage a storage pool, update storage pool attributes or to request a refresh of a storage pool.
	// To manage/unmanage a storage pool, issue a PUT with the isManaged attribute set as true to manage or false to unmanage.
	// Attempting to unmanage a StoreVirtual pool is not allowed and the attempt will return a task error.
	// To request a refresh of a storage pool the user must set the "requestingRefresh" attribute to true. The user cannot perform any other attribute update to the storage pool while also requesting a refresh of the pool.
	update_pool.IsManaged = true
	err := ovc.UpdateStoragePool(update_pool)
	if err != nil {
		fmt.Println("Could not update the pool", err)
	}

	// Get All the pools present
	fmt.Println("\nGetting all the storage pools present in the system: \n")
	sort := "name:desc"
	pool_list, err := ovc.GetStoragePools("", sort, "", "")
	if err != nil {
		fmt.Println("Error Getting the storage pools ", err)
	}
	for i := 0; i < len(pool_list.Members); i++ {
		fmt.Println(pool_list.Members[i].Name)
	}
}
