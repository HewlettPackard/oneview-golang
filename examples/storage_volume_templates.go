package main

import (
	"fmt"
	"time"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func main() {

	var (
		ClientOV            *ov.OVClient
		name_to_create      = "VolumeTemplateExample"
		name_to_create_auto = "Auto-VolumeTemplate"
		name_to_update      = "VolumeTemplateExample- updated"
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

	name_properties := ov.TemplatePropertyDatatypeStructString{
		Required:    true,
		Type:        "string",
		Title:       "Volume name",
		Description: "A volume name between 1 and 100 characters",
		Maxlength:   100,
		Minlength:   1,
		Meta: &ov.Meta{
			Locked: false,
		},
	}
	st_pool, err := ovc.GetStoragePoolByName("CPG-SSD")
	if err != nil {
		fmt.Println(err)
	}
	storage_pool_properties := ov.TemplatePropertyDatatypeStructString{
		Required:    true,
		Type:        "string",
		Title:       "Storage Pool",
		Description: "A common provisioning group URI reference",
		Default:     string(st_pool.URI),
		Meta: &ov.Meta{
			Locked:       false,
			CreateOnly:   true,
			SemanticType: "device-storage-pool",
		},
		Format: "x-uri-reference",
	}
	snapshot_pool := ov.TemplatePropertyDatatypeStructString{
		Required:    false,
		Type:        "string",
		Title:       "Snapshot Pool",
		Description: "A URI reference to the common provisioning group used to create snapshots",
		Default:     string(st_pool.URI),
		Meta: &ov.Meta{
			Locked:       false,
			SemanticType: "device-snapshot-storage-pool",
		},
		Format: "x-uri-reference",
	}
	size_properties := ov.TemplatePropertyDatatypeStructInt{
		Required:    true,
		Type:        "integer",
		Title:       "Capacity",
		Default:     268435456,
		Minimum:     268435456,
		Maximum:     17592186044416,
		Description: "The capacity of the volume in bytes",
		Meta: &ov.Meta{
			Locked:       false,
			SemanticType: "capacity",
		},
	}
	/*
		dataProtectionLevel_properties := ov.TemplatePropertyDatatypeStructString{
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
			Description: "Indicates the number and configuration of data copies in the Storage Pool",
			Meta: &ov.Meta{
				Locked:       false,
				SemanticType: "device-dataProtectionLevel",
			},
		}
	*/

	template_version_properties := ov.TemplatePropertyDatatypeStructString{
		Required:    true,
		Type:        "string",
		Title:       "Template version",
		Description: "Version of the template",
		Default:     "1.1",
		Meta: &ov.Meta{
			Locked: true,
		},
	}

	description_properties := ov.TemplatePropertyDatatypeStructString{
		Required:    false,
		Type:        "string",
		Title:       "Description",
		Description: "A description for the volume",
		Default:     "A description for the volume",
		Maxlength:   2000,
		Minlength:   1,
		Meta: &ov.Meta{
			Locked: false,
		},
	}

	provisioning_type_properties := ov.TemplatePropertyDatatypeStructString{
		Required:    false,
		Title:       "Provisioning Type",
		Type:        "string",
		Description: "The provisioning type for the volume",
		Default:     "Thin",
		Enum:        []string{"Thin", "Full"},
		Meta: &ov.Meta{
			Locked:     true,
			CreateOnly: true,
			/*			SemanticType: "device-provisioningType",*/
		},
	}

	/*	adaptive_optimization_properties := ov.TemplatePropertyDatatypeStructBool{
			Meta: &ov.Meta{
				Locked: true,
			},
			Type:        "boolean",
			Description: "",
			Default:     true,
			Required:    false,
			Title:       "Adaptive Optimization",
		}
	*/

	isDeduplicated := ov.TemplatePropertyDatatypeStructBool{
		Meta: &ov.Meta{
			Locked: true,
		},
		Type:        "boolean",
		Description: "Enables or disables deduplication of the volume",
		Default:     true,
		Required:    false,
		Title:       "Is Deduplicated",
	}
	is_shareable_properties := ov.TemplatePropertyDatatypeStructBool{
		Meta: &ov.Meta{
			Locked: true,
		},
		Type:        "boolean",
		Description: "The shareability of the volume",
		Default:     true,
		Required:    false,
		Title:       "Is Shareable",
	}

	Properties := ov.TemplateProperties{
		Name:        &name_properties,
		StoragePool: &storage_pool_properties,
		Size:        &size_properties,
		/*		DataProtectionLevel:           &dataProtectionLevel_properties,*/
		SnapshotPool:     &snapshot_pool,
		IsDeduplicated:   &isDeduplicated,
		TemplateVersion:  &template_version_properties,
		Description:      &description_properties,
		ProvisioningType: &provisioning_type_properties,
		/*		IsAdaptiveOptimizationEnabled: &adaptive_optimization_properties,*/
		IsShareable: &is_shareable_properties,
	}

	storageVolumeTemplate := ov.StorageVolumeTemplate{
		TemplateProperties: &Properties,
		Name:               name_to_create,
		Description:        "Volume template Example",
	}

	storageVolumeTemplateAuto := ov.StorageVolumeTemplate{
		TemplateProperties: &Properties,
		Name:               name_to_create_auto,
		Description:        "Volume template Example",
	}

	err = ovc.CreateStorageVolumeTemplate(storageVolumeTemplate)
	err = ovc.CreateStorageVolumeTemplate(storageVolumeTemplateAuto)
	if err != nil {
		fmt.Println("Could not create the volume Template", err)
	} else {
		fmt.Println("Volume Template created successfully", storageVolumeTemplate.Name)
	}

	// Get the volume template by name
	update_vol_template, err := ovc.GetStorageVolumeTemplateByName(name_to_create)

	if err != nil {
		fmt.Println(err)
	} else {
		update_vol_template.Name = name_to_update
		update_vol_template.Description = "Updating description"

		// Update the previously created storage volume template
		err = ovc.UpdateStorageVolumeTemplate(update_vol_template)
		if err != nil {
			fmt.Println("Could not update the volume template", err)
		} else {
			fmt.Println("Volume template updated")
		}
	}
	time.Sleep(2 * time.Second)
	// Get Root template
	fmt.Println("\nGetting root volume templates present in the system: \n")
	st_vol_root, err := ovc.GetRootStorageVolumeTemplate(string(st_pool.URI))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(st_vol_root)

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

	// Delete the created volume template
	fmt.Println("\nDeleting the volume with name : ", name_to_update)
	err = ovc.DeleteStorageVolumeTemplate(name_to_update)
	if err != nil {
		fmt.Println("Delete Unsuccessful", err)
	}

}
