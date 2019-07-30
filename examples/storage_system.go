package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
)

func main() {

	var (
		ClientOV       *ov.OVClient
		name_to_create = "Cluster-1"
		managed_domain = "TestDomain" //Variable to update the managedDomain
	)

	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		1000,
		"*")

	// Create storage system
	storageSystem := ov.StorageSystemV4{Hostname: "<hostname>", Username: "<username>", Password: "<password>", Family: "<family>", Description: "<description>"}

	err := ovc.CreateStorageSystem(storageSystem)
	if err != nil {
		fmt.Println("Could not create the system", err)
	}

	//The below example is to update a storeServ system.
	//Please refer the API reference for fields to update a storeVirtual system.

	// Get the storage system to be updated
	update_system, _ := ovc.GetStorageSystemByName(name_to_create)

	// Update the given storage system
	//Managed domain is mandatory attribute for update
	DeviceSpecificAttributesForUpdate := update_system.StorageSystemDeviceSpecificAttributes
	DeviceSpecificAttributesForUpdate.ManagedDomain = managed_domain

	updated_storage_system := ov.StorageSystemV4{
		Name: name_to_create,
		StorageSystemDeviceSpecificAttributes: DeviceSpecificAttributesForUpdate,
		URI:         update_system.URI,
		ETAG:        update_system.ETAG,
		Description: "Updated the storage system",
		Credentials: update_system.Credentials,
		Hostname:    update_system.Hostname,
		Ports:       update_system.Ports,
	}

	err = ovc.UpdateStorageSystem(updated_storage_system)
	if err != nil {
		fmt.Println("Could not update the system", err)
	}

	// Get All the systems present
	fmt.Println("\nGetting all the storage systems present in the appliance: \n")
	sort := "name:desc"
	system_list, err := ovc.GetStorageSystems("", sort)
	if err != nil {
		fmt.Println("Error Getting the storage systems ", err)
	} else {
		for i := 0; i < len(system_list.Members); i++ {
			fmt.Println(system_list.Members[i].Name)
		}
	}

	// Get reachable ports
	fmt.Println("\n Getting rechable ports of:", name_to_create)
	reachable_ports, _ := ovc.GetReachablePorts(update_system.URI)
	fmt.Println(reachable_ports)

	// Get volume sets
	fmt.Println("\n Getting volume sets of:", name_to_create)
	volume_sets, _ := ovc.GetVolumeSets(update_system.URI)
	fmt.Println(volume_sets)

	// Delete the created system
	fmt.Println("\nDeleting the system with name : ", name_to_create)
	err = ovc.DeleteStorageSystem(name_to_create)
	if err != nil {
		fmt.Println("Delete Unsuccessful", err)
	}
}
