package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {

	var (
		ClientOV       *ov.OVClient
		new_name       = "TestVolume"
		name_to_update = "UpdatedName"
	)

	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		500,
		"*")

	// Create storage volume with name <new_name>
	properties := &ov.Properties{new_name, utils.NewNstring("/rest/storage-pools/AAA05D5E-BDB5-4FBB-8E65-A8D400A6A8AF"), 107374741824, "Thin"}

	storageVolume := ov.StorageVolumeV3{TemplateURI: utils.NewNstring("/rest/storage-volume-templates/c93ef008-d8f0-40a5-b2d1-a8d400a6a8b7"), Properties: properties, IsPermanent: false}

	err := ovc.CreateStorageVolume(storageVolume)
	if err != nil {
		fmt.Println("Could not create the volume", err)
	}

	// Update the given storage volume
	update_vol, _ := ovc.GetStorageVolumeByName(new_name)

	updated_storage_volume := ov.StorageVolumeV3{
		ProvisioningTypeForUpdate: update_vol.ProvisioningTypeForUpdate,
		IsPermanent:               update_vol.IsPermanent,
		IsShareable:               update_vol.IsShareable,
		Name:                      name_to_update,
		ProvisionedCapacity:       "107374182400",
		DeviceSpecificAttributes:  update_vol.DeviceSpecificAttributes,
		URI:                       update_vol.URI,
		ETAG:                      update_vol.ETAG,
		Description:               "empty",
	}

	err = ovc.UpdateStorageVolume(updated_storage_volume)
	if err != nil {
		fmt.Println("Could not update the volume", err)
	}

	// Get All the volumes present
	fmt.Println("\nGetting all the storage volumes present in the system: \n")
	sort := "name:desc"
	vol_list, err := ovc.GetStorageVolumes("", sort)
	if err != nil {
		fmt.Println("Error Getting the storage volumes ", err)
	}
	for i := 0; i < len(vol_list.Members); i++ {
		fmt.Println(vol_list.Members[i].Name)
	}

	// Get volume by name
	fmt.Println("\nGetting details of volume with name: ", name_to_update)
	vol_by_name, _ := ovc.GetStorageVolumeByName(name_to_update)
	fmt.Println(vol_by_name)

	// Delete the created volume
	fmt.Println("\nDeleting the volume with name : UpdatedName")
	err = ovc.DeleteStorageVolume(name_to_update)
	if err != nil {
		fmt.Println("Delete Unsuccessful", err)
	}
}
