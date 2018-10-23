package main
import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {

	var(
		ClientOV *ov.OVClient
		get_By_Name = "3par_1"
		new_name = "3par_2"
	)

	ovc := ClientOV.NewOVClient(
			os.Getenv("ONEVIEW_OV_USER"),
			os.Getenv("ONEVIEW_OV_PASSWORD"),
			os.Getenv("ONEVIEW_OV_DOMAIN"),
			os.Getenv("ONEVIEW_OV_ENDPOINT"),
			false,
			600)

	properties := ov.Properties{new_name,utils.NewNstring("/rest/storage-pools/97BC0F4F-3706-496E-B7A1-A8D90065D7E0"),107374741824,"Thin"}

	storage_volume := ov.StorageVolumeV3{TemplateURI: "/rest/storage-volume-templates/cacb36c7-42d2-430f-9fc9-a8d90065d80e", Properties: properties, IsPermanent: false}
	err := ovc.CreateStorageVolume(storage_volume)
	if err != nil {
	fmt.Println("nOT CREATED", err)
	}


	//DeviceSpecificPropertirs := ov.DeviceSpecificPropertiesStoreServ("NotCopying",utils.NewNstring("/rest/storage-pools/97BC0F4F-3706-496E-B7A1-A8D90065D7E0"))
	//updated_storage_volume := ov.StorageVolumeV3{ ProvisioningType : "Thin", IsPermanent: true, IsShareab


	fmt.Println("Getting all the storage volumes present in the system: \n")
	vol,_ := ovc.GetStorageVolumes("","")
	fmt.Println(vol)

	fmt.Println("Getting details of volume with name: ",storage_volume.Name)
	vol_by_name,_ := ovc.GetStorageVolumeByName(get_By_Name)
	fmt.Println(vol_by_name)
	}
