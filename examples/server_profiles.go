package main

import (
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
)

func main() {
	var (
		ClientOV *ov.OVClient
	)

	config, config_err := ov.LoadConfigFile("config.json")
	if config_err != nil {
		fmt.Println(config_err)
	}
	var (
		sp_name         = config.ServerProfileConfig.ServerProfileName
		sp_by_spt       = "sp_from_spt"
		new_sp_name     = "Renamed Server Profile"
		server_hardware = config.ServerProfileConfig.ServerHardwareName
		scopeName       = "SP-Scope"
		spt_name        = config.ServerProfileConfig.OvTemplatestring

		// you can optionally ignore certain warnings when creatig a server profile
		// omit flag parameter in SubmitNewProfile() or CreateProfileFromTemplate() for default value (none)
		ignoreFlags = []ov.ForceFlag{ov.ForceIgnoreServerHealth}
	)
	ovc := ClientOV.NewOVClient(
		config.OVCred.UserName,
		config.OVCred.Password,
		config.OVCred.Domain,
		config.OVCred.Endpoint,
		config.OVCred.SslVerify,
		config.OVCred.ApiVersion,
		config.OVCred.IfMatch)

	spScope := ov.Scope{Name: scopeName, Description: "Test from script", Type: "ScopeV3"}

	errSP := ovc.CreateScope(spScope)

	if errSP != nil {
		fmt.Println("Error Creating Scope: ", errSP)
	}
	scope, _ := ovc.GetScopeByName(scopeName)

	initialScopeUris := &[]utils.Nstring{scope.URI}
	serverName, err := ovc.GetServerHardwareByName(server_hardware)
	fmt.Println(serverName.URI)
	server_profile_create_map := ov.ServerProfile{
		Type:              "ServerProfileV12",
		Name:              sp_name,
		ServerHardwareURI: serverName.URI,
		InitialScopeUris:  *initialScopeUris,
	}

	err = ovc.SubmitNewProfile(server_profile_create_map, ignoreFlags...)
	if err != nil {
		fmt.Println("Server Profile Create Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Created---------------#")
	}

	// Example: Submit a new Server Profile asynchronously
	//
	// This example demonstrates how to use the `SubmitNewProfileAsync` function to create
	// a new server profile in OneView. The call is asynchronous, meaning it immediately
	// returns a Task object instead of waiting for the operation to finish.
	//
	// The returned Task contains the URI that can be polled or monitored until the
	// profile creation completes. If the server hardware already has a profile assigned
	// or if the hardware is powered on, an error will be returned.
	task, err := ovc.SubmitNewProfileAsync(server_profile_create_map, ignoreFlags...)
	if err != nil {
		fmt.Println("Server Profile Create Failed: ", err)
	} else {

		if err != nil {
			fmt.Println("Error marshaling task:", err)
		} else {
			fmt.Println("Task URI:")
			fmt.Println(string(task.URI))
		}
		fmt.Println("#----------------Server Profile Creation Task Initiated---------------#")
	}

	sort := ""
	server_name := ov.ServerHardwareType{}
	spt, err := ovc.GetProfileTemplateByName(spt_name)
	if err != nil {
		fmt.Println("Server Profile Template Retrieval By Name Failed: ", err)
	} else {
		ServerList, err := ovc.GetServerHardwareList([]string{""}, "", "", "", "")
		hw_name, _ := ovc.GetServerHardwareTypeByUri(spt.ServerHardwareTypeURI)
		for i := 0; i < len(ServerList.Members); i++ {
			server_name, _ = ovc.GetServerHardwareTypeByUri(ServerList.Members[i].ServerHardwareTypeURI)
			if server_name.Name == hw_name.Name {
				serverName = ServerList.Members[i]
			}
		}
		if err != nil {
			fmt.Println("Failed to fetch server hardware name: ", err)
		} else {
			err = ovc.CreateProfileFromTemplate(sp_by_spt, spt, serverName, ignoreFlags...)
			if err != nil {
				fmt.Println("Server Profile Create Failed: ", err)
			} else {
				fmt.Println("#----------------Server Profile Created---------------#")
			}
		}
	}

	sp_list, err := ovc.GetProfiles("", "", "", sort, "")
	if err != nil {
		fmt.Println("Server Profile Retrieval Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile List---------------#")

		for i := 0; i < len(sp_list.Members); i++ {
			fmt.Println(sp_list.Members[i].Name)
		}
	}

	sp1, err := ovc.GetProfileByName(sp_name)
	if err != nil {
		fmt.Println("Server Profile Retrieval By Name Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile by Name---------------#")
		fmt.Println(sp1.Name)
	}

	sp2, err := ovc.GetProfileBySN(string(sp1.SerialNumber))
	if err != nil {
		fmt.Println("Server Profile Retrieval By Serial Number Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile by Serial Number---------------#")
		fmt.Println(sp2.Name)
	}

	sp, err := ovc.GetProfileByURI(sp2.URI)
	if err != nil {
		fmt.Println("Server Profile Retrieval By URI Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile by URI---------------#")
		fmt.Println(sp.Name)
	}

	fmt.Println("Server Profile refresh using PATCH request")
	options := new([]ov.Options)
	*options = append(*options, ov.Options{"replace", "/refreshState", "RefreshPending"})

	err = ovc.PatchServerProfile(sp, *options) //patchRequest)
	if err != nil {
		fmt.Println("Refresh failed", err)
	}

	sp_update_clone := ov.ServerProfile{
		Name:                  new_sp_name,
		URI:                   sp1.URI,
		Type:                  sp1.Type,
		ETAG:                  sp1.ETAG,
		Affinity:              sp1.Affinity,
		ServerHardwareTypeURI: sp1.ServerHardwareTypeURI,
		EnclosureGroupURI:     sp1.EnclosureGroupURI,
	}

	err = ovc.UpdateServerProfile(sp_update_clone)
	if err != nil {
		fmt.Println("Server Profile Create Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Created---------------#")
	}

	sp_list, err = ovc.GetProfiles("", "", "", sort, "")
	if err != nil {
		fmt.Println("Server Profile Retrieval Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile List---------------#")

		for i := 0; i < len(sp_list.Members); i++ {
			fmt.Println(sp_list.Members[i].Name)
		}
	}

	err = ovc.DeleteProfile(sp1.Name)
	if err != nil {
		fmt.Println("Server Profile Delete Request Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Delete---------------#")
		//fmt.Println("Task URI: ", task.URI)
	}

	err = ovc.DeleteProfile(new_sp_name)
	err = ovc.DeleteProfile(sp_by_spt)
	if err != nil {
		fmt.Println("Server Profile Delete Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Deleted---------------#")
	}
	err = ovc.DeleteScope(scopeName)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleted scope successfully...")
	}

}
