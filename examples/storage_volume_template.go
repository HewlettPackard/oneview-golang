package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	//	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {

	var (
		ClientOV *ov.OVClient
		//		new_volume     = "TestVolume"
		name_to_update = "voltempl1"
	)

	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		1000,
		"*")

	name_properties := ov.TemplatePropertyDatatypeStruct{
		Required:    true,
		Type:        "string",
		Title:       "Volume name",
		Description: "Volume Name",
		Maxlength:   100,
		Minlength:   1,
		Meta: ov.Meta{
			Locked: false,
		},
	}

	storage_pool_properties := ov.TemplatePropertyDatatypeStruct{
		Required:    true,
		Type:        "string",
		Title:       "Storage pool",
		Description: " ",
		Default:     "/rest/storage-pools/F693B0B6-AD80-40C0-935D-AA99009ED046",
		Meta: ov.Meta{
			Locked:       false,
			CreateOnly:   true,
			SemanticType: "device-storage-pool",
		},
		Format: "x-uri-reference",
	}

	size_properties := ov.TemplatePropertyDatatypeStructInt{
		Required:    true,
		Type:        "integer",
		Title:       "Capacity",
		Default:     1073741824,
		Minimum:     4194304,
		Description: " ",
		Meta: ov.Meta{
			Locked:       false,
			SemanticType: "capacity",
		},
	}

	dataProtectionLevel_properties := ov.TemplatePropertyDatatypeStruct{
		Required: true,
		Type:     "string",
		Enum: []string{"NetworkRaid0None",
			"NetworkRaid5SingleParity",
			"NetworkRaid10Mirror2Way",
			"NetworkRaid10Mirror3Way",
			"NetworkRaid10Mirror4Way",
			"NetworkRaid6DualParity",
		},
		Title:       "Data Protection Level",
		Default:     "NetworkRaid10Mirror2Way",
		Description: " ",
		Meta: ov.Meta{
			Locked:       false,
			SemanticType: "device-dataProtectionLevel",
		},
	}

	template_version_properties := ov.TemplatePropertyDatatypeStruct{
		Required:    true,
		Type:        "string",
		Title:       "template version",
		Description: "version of the template",
		Default:     "1.1",
		Meta: ov.Meta{
			Locked: true,
		},
	}

	description_properties := ov.TemplatePropertyDatatypeStruct{
		Required:    false,
		Type:        "string",
		Title:       "Description",
		Description: " ",
		Default:     " ",
		Maxlength:   2000,
		Minlength:   1,
		Meta: ov.Meta{
			Locked: false,
		},
	}

	provisioning_type_properties := ov.TemplatePropertyDatatypeStruct{
		Required:    false,
		Title:       "Provisioning Type",
		Type:        "string",
		Description: " ",
		Default:     "Thin",
		Enum:        []string{"Thin", "Full"},
		Meta: ov.Meta{
			Locked:       true,
			CreateOnly:   true,
			SemanticType: "device-provisioningType",
		},
	}

	adaptive_optimization_properties := ov.TemplatePropertyDatatypeStructBool{
		Meta: ov.Meta{
			Locked: true,
		},
		Type:        "boolean",
		Description: " ",
		Default:     true,
		Required:    false,
		Title:       "Adaptive optimization",
	}

	is_shareable_properties := ov.TemplatePropertyDatatypeStructBool{
		Meta: ov.Meta{
			Locked: true,
		},
		Type:        "boolean",
		Description: " ",
		Default:     true,
		Required:    false,
		Title:       "IsShareable",
	}

	Properties := ov.TemplateProperties{
		Name:                          &name_properties,
		StoragePool:                   &storage_pool_properties,
		Size:                          &size_properties,
		DataProtectionLevel:           &dataProtectionLevel_properties,
		TemplateVersion:               &template_version_properties,
		Description:                   &description_properties,
		ProvisioningType:              &provisioning_type_properties,
		IsAdaptiveOptimizationEnabled: &adaptive_optimization_properties,
		IsShareable:                   &is_shareable_properties,
	}

	storageVolumeTemplate := ov.StorageVolumeTemplate{
		TemplateProperties: Properties,
		Name:               "VolumeTemplateExample",
		Description:        "Volume template Example",
		RootTemplateUri:    "/rest/storage-volume-templates/533c5b9e-26c3-4c2e-af4c-aa99009ed20e",
	}

	err := ovc.CreateStorageVolumeTemplate(storageVolumeTemplate)
	if err != nil {
		fmt.Println("Could not create the volume", err)
	}

	/*
		// Update the given storage volume
		update_vol, _ := ovc.GetStorageVolumeByName(new_volume)

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
	*/
	// Get All the volume templates present
	fmt.Println("\nGetting all the storage volume templates present in the system: \n")
	sort := "name:desc"
	vol_temp_list, err := ovc.GetStorageVolumeTemplates("", sort, "", "")
	if err != nil {
		fmt.Println("Error Getting the storage volume templates ", err)
	}
	for i := 0; i < len(vol_temp_list.Members); i++ {
		fmt.Println(vol_temp_list.Members[i].Name)
	}

	// Get volume by name
	fmt.Println("\nGetting details of volume with name: ", name_to_update)
	vol_by_name, _ := ovc.GetStorageVolumeTemplateByName(name_to_update)
	fmt.Println(vol_by_name.URI)

	/*	// Delete the created volume
		fmt.Println("\nDeleting the volume with name : UpdatedName")
		err = ovc.DeleteStorageVolume(name_to_update)
		if err != nil {
			fmt.Println("Delete Unsuccessful", err)
		}
	*/
}
