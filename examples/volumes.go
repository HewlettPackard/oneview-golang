package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func main() {

	var (
		ClientOV        *ov.OVClient
		st_vol_template = "Auto-VolumeTemplate"
		root_volume     = "volume_without_template"
		new_volume      = "TestVolume"
		name_to_update  = "UpdatedName"
	)

	apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
		"*")

	st_pool, _ := ovc.GetStoragePoolByName("CPG-SSD")
	properties_root := &ov.Properties{
		Name:             root_volume,
		Storagepool:      st_pool.URI,
		Size:             268435456,
		ProvisioningType: "Thin",
	}

	properties_auto := &ov.Properties{
		Name:             "Auto-Volume",
		Storagepool:      st_pool.URI,
		Size:             268435456,
		ProvisioningType: "Thin",
	}

	//Create Storage volume with root template

	root_vol_template, err := ovc.GetRootStorageVolumeTemplate()
	if err != nil {
		fmt.Println(err)
	}
	storageVolume_with_root_volume_template := ov.StorageVolume{TemplateURI: root_vol_template.URI, Properties: properties_root}

	err = ovc.CreateStorageVolume(storageVolume_with_root_volume_template)

	if err != nil {
		fmt.Println("Could not create the volume", err)
	}
	// Create storage volume with name <new_volume>
	properties := &ov.Properties{
		Name:             new_volume,
		Storagepool:      st_pool.URI,
		Size:             268435456,
		ProvisioningType: "Thin",
	}
	vol_template, err := ovc.GetStorageVolumeTemplateByName(st_vol_template)
	if err != nil {
		fmt.Println(err)
	}
	storageVolume := ov.StorageVolume{TemplateURI: vol_template.URI, Properties: properties}
	storageVolume_auto := ov.StorageVolume{TemplateURI: vol_template.URI, Properties: properties_auto}
	err = ovc.CreateStorageVolume(storageVolume)
	err = ovc.CreateStorageVolume(storageVolume_auto)
	if err != nil {
		fmt.Println("Could not create the volume", err)
	}

	// Update the given storage volume
	update_vol, _ := ovc.GetStorageVolumeByName(new_volume)

	updated_storage_volume := ov.StorageVolume{
		ProvisioningTypeForUpdate: update_vol.ProvisioningTypeForUpdate,
		IsPermanent:               update_vol.IsPermanent,
		IsShareable:               update_vol.IsShareable,
		Name:                      name_to_update,
		ProvisionedCapacity:       "107374741824",
		DeviceSpecificAttributes:  update_vol.DeviceSpecificAttributes,
		URI:                       update_vol.URI,
		ETAG:                      update_vol.ETAG,
		Description:               "empty",
		TemplateVersion:           "1.1",
		VolumeTemplateUri:         update_vol.VolumeTemplateUri,
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

	// Delete the created volume without template
	fmt.Println("\nDeleting the volume with name : root_volume")
	err = ovc.DeleteStorageVolume(root_volume)
	if err != nil {
		fmt.Println("Delete Unsuccessful", err)
	}
}
