package main

import (
    "fmt"
    "github.com/HewlettPackard/oneview-golang/ov"
    "os"
)

func main() {
    var (
        clientOV            *ov.OVClient
        spt_name            = "test"
        eg_name             = "SYN03_EC"
        server_hardware_type_name = "SY 480 Gen9 1"
    )
    ovc := clientOV.NewOVClient(
        os.Getenv("ONEVIEW_OV_USER"),
        os.Getenv("ONEVIEW_OV_PASSWORD"),
        os.Getenv("ONEVIEW_OV_DOMAIN"),
        os.Getenv("ONEVIEW_OV_ENDPOINT"),
        false,
        800,
        "*")

    server_hardware_type, err := ovc.GetServerHardwareTypeByName(server_hardware_type_name)

    enc_grp, err := ovc.GetEnclosureGroupByName(eg_name)

    conn_settings := ov.ConnectionSettings{
        ManageConnections:true,
    }

    server_profile_template := ov.ServerProfile{
        Type: "ServerProfileTemplateV5",
        Name: "test",
        EnclosureGroupURI: enc_grp.URI,
        ServerHardwareTypeURI: server_hardware_type.URI,
        ConnectionSettings: conn_settings,
    }

    err = ovc.CreateProfileTemplate(server_profile_template)
    if err != nil {
        fmt.Println("Server Profile Template Creation Failed: ", err)
    } else {
        fmt.Println("#----------------Server Profile Template Created---------------#")
    }

    sort := ""

    spt_list, err := ovc.GetProfileTemplates("", sort)
    if err != nil {
        fmt.Println("Server Profile Template Retrieval Failed: ", err)
    } else {
        fmt.Println("#----------------Server Profile Template List---------------#")

        for i := 0; i < len(spt_list.Members); i++ {
            fmt.Println(spt_list.Members[i].Name)
        }
    }

    sp_del, err := ovc.GetProfileTemplateByName(spt_name)
    if err != nil {
        fmt.Println("Server Profile Template Retrieval By Name Failed: ", err)
    } else {
        fmt.Println("#----------------Server Profile Template by Name---------------#")
        fmt.Println(sp_del.Name)
    }

    err = ovc.DeleteProfileTemplate("test")
    if err != nil {
        fmt.Println("Server Profile Template Delete Failed: ", err)
    } else {
        fmt.Println("#----------------Server Profile Template Deleted---------------#")
    }
}
