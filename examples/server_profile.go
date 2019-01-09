package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
)

func main() {
	var (
		clientOV            *ov.OVClient
		spt_name 			= "Ansible_demo"
		sp_name             = "test"
//		sp_sn 	            = "VCGRE1S007"
//		new_enclosure_name = "RenamedEnclosure"
//		path               = "/name"
//		op                 = "replace"
	)
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		800,
		"*")

//	server_profile_create_map := ov.ServerProfile{
//		Type: 				"ServerProfileV9",
//		Name: 				"test",
//		ServerHardwareURI: 	""
//	}
	//server_hardware_map := ov.ServerHardware{}
	templates, err := ovc.GetProfileTemplates(fmt.Sprintf("name matches '%s'", spt_name), "name:asc")

	blade, err := ovc.GetServerHardwareByName("SYN03_Frame1, bay 12")

	err = ovc.CreateProfileFromTemplate(sp_name, templates.Members[0], blade)
	if err != nil {
		fmt.Println("Server Profile Create Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Created---------------#")
	}

    sort := ""

    sp_list, err := ovc.GetProfiles("", sort)
    if err != nil {
        fmt.Println("Server Profile Retrieval Failed: ", err)
    } else {
        fmt.Println("#----------------Server Profile List---------------#")

        for i := 0; i < len(sp_list.Members); i++ {
            fmt.Println(sp_list.Members[i].Name)
        }
    }

    sp_del, err := ovc.GetProfileByName(sp_name)
    if err != nil {
        fmt.Println("Server Profile Retrieval By Name Failed: ", err)
    } else {
        fmt.Println("#----------------Server Profile by Name---------------#")
        fmt.Println(sp_del.Name)
    }

    sp, err := ovc.GetProfileBySN(sp_sn)
    if err != nil {
        fmt.Println("Server Profile Retrieval By Serial Number Failed: ", err)
    } else {
        fmt.Println("#----------------Server Profile by Serial Number---------------#")
        fmt.Println(sp.Name)
    }

    sp, err = ovc.GetProfileByURI(sp.URI)
    if err != nil {
        fmt.Println("Server Profile Retrieval By URI Failed: ", err)
    } else {
        fmt.Println("#----------------Server Profile by URI---------------#")
        fmt.Println(sp.Name)
    }

    task, err := ovc.SubmitDeleteProfile(sp_del)
    if err != nil {
        fmt.Println("Server Profile Delete Request Failed: ", err)
    } else {
        fmt.Println("#----------------Server Profile Delete---------------#")
        fmt.Println("Task URI: ", task.URI)
    }

    err = ovc.DeleteProfile("test")
    if err != nil {
        fmt.Println("Server Profile Delete Failed: ", err)
    } else {
        fmt.Println("#----------------Server Profile Deleted---------------#")
    }

}