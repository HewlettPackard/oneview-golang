package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	//	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {

	var (
		ClientOV     *ov.OVClient
		storage_pool = "CPG-SSD"
	)

	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		600,
		"*")

	//Get storage pool by name to update
	update_pool, _ := ovc.GetStoragePoolByName(storage_pool)

	// Update the given storage pool
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
