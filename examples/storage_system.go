package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
)

func main() {

	var (
		ClientOV       *ov.OVClient
		name_to_create = "ThreePAR-1"
		//name_to_create = "ThreePAR-2"
		//Variable to update the managedDomain
		managed_domain = "TestDomain"
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
	storageSystem := ov.StorageSystemV4{Hostname: "172.18.11.11", Username: "dcs", Password: "dcs", Family: "StoreServ"}

	err := ovc.CreateStorageSystem(storageSystem)
	if err != nil {
		fmt.Println("Could not create the system", err)
	}

	// Update the given storage system
	update_system, _ := ovc.GetStorageSystemByName(name_to_create)

	//Managed domain is mandatory attribute for update
	DeviceSpecificAttributesForUpdate := update_system.StorageSystemDeviceSpecificAttributes
	DeviceSpecificAttributesForUpdate.ManagedDomain = managed_domain

	updated_storage_system := ov.StorageSystemV4{
		Name: name_to_create,
		StorageSystemDeviceSpecificAttributes: DeviceSpecificAttributesForUpdate,
		URI:         update_system.URI,
		ETAG:        update_system.ETAG,
		Description: "empty",
		Credentials: update_system.Credentials,
		Hostname:    update_system.Hostname,
		Ports:       update_system.Ports,
	}

	err = ovc.UpdateStorageSystem(updated_storage_system)
	if err != nil {
		fmt.Println("Could not update the system", err)
	}

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
	fmt.Println("\nGetting details of system with name: ", name_to_create)
	system_by_name, _ := ovc.GetStorageSystemByName(name_to_create)
	fmt.Println(system_by_name)

	// Delete the created system
	fmt.Println("\nDeleting the system with name : ", name_to_create)
	err = ovc.DeleteStorageSystem(name_to_create)
	if err != nil {
		fmt.Println("Delete Unsuccessful", err)
	}
}
