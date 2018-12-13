package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
//	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {

	var (
		ClientOV       *ov.OVClient
//		new_system    = "TestSystem"
		name_to_update = "ThreePAR-1"
	)

	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		500,
		"*")

	// Create storage system with name <new_system>
	storageSystem := ov.StorageSystemV4{Hostname: "172.18.11.11", Username:"dcs", Password:"dcs", Family: "StoreServ"}

	err := ovc.CreateStorageSystem(storageSystem)
	if err != nil {
		fmt.Println("Could not create the system", err)
	}
/*
	// Update the given storage system
	update_system, _ := ovc.GetStorageSystemByName(new_system)

	updated_storage_system := ov.StorageSystemV3{
		ProvisioningTypeForUpdate: update_system.ProvisioningTypeForUpdate,
		IsPermanent:               update_system.IsPermanent,
		IsShareable:               update_system.IsShareable,
		Name:                      name_to_update,
		ProvisionedCapacity:       "107374182400",
		DeviceSpecificAttributes:  update_system.DeviceSpecificAttributes,
		URI:                       update_system.URI,
		ETAG:                      update_system.ETAG,
		Description:               "empty",
	}

	err = ovc.UpdateStorageSystem(updated_storage_system)
	if err != nil {
		fmt.Println("Could not update the system", err)
	}
*/

	// Get All the systems present
	fmt.Println("\nGetting all the storage systems present in the system: \n")
	sort := "name:desc"
	system_list, err := ovc.GetStorageSystems("", sort)
	if err != nil {
		fmt.Println("Error Getting the storage systems ", err)
	} else {
		for i := 0; i < len(system_list.Members); i++ {
			fmt.Println(system_list.Members[i].Name)
		}
	}

	// Get system by name
	fmt.Println("\nGetting details of system with name: ", name_to_update)
	system_by_name, _ := ovc.GetStorageSystemByName(name_to_update)
	fmt.Println(system_by_name)

	// Delete the created system
	fmt.Println("\nDeleting the system with name : UpdatedName")
	//err = ovc.DeleteStorageSystem(name_to_update)
	if err != nil {
		fmt.Println("Delete Unsuccessful", err)
	}
}
