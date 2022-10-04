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
	ovc := ClientOV.NewOVClient(
		config.OVCred.UserName,
		config.OVCred.Password,
		config.OVCred.Domain,
		config.OVCred.Endpoint,
		config.OVCred.SslVerify,
		config.OVCred.ApiVersion,
		config.OVCred.IfMatch)
	scope := ov.Scope{Name: "ScopeHardware", Description: "Test from script", Type: "ScopeV3"}
	_ = ovc.CreateScope(scope)
	scp, _ := ovc.GetScopeByName("ScopeHardware")
	initialScopeUris := &[]utils.Nstring{scp.URI}

	fmt.Println("-----------------------------")
	fmt.Println("Add Single Rack server to the appliance:")
	rackServer := ov.ServerHardware{
		Hostname:           "172.18.41.1",
		Username:           "dcs",
		Password:           "dcs",
		Force:              false,
		LicensingIntent:    "OneView", //OneView or OneViewNoiLO for Managed
		ConfigurationState: "Managed",
		InitialScopeUris:   *initialScopeUris,
	}

	_, err := ovc.AddRackServer(rackServer)
	fmt.Println("Added rack-server successfully.")

	fmt.Println("-----------------------------")
	fmt.Println("Add multiple Rack servers:")
	hostsAndRanges := &[]utils.Nstring{"172.18.41.2-172.18.41.7"}
	multipleRackServers := ov.ServerHardware{
		MpHostsAndRanges:   *hostsAndRanges,
		Username:           "dcs",
		Password:           "dcs",
		Force:              false,
		LicensingIntent:    "OneView", //OneView or OneViewNoiLO for Managed
		ConfigurationState: "Managed",
		InitialScopeUris:   *initialScopeUris,
	}

	err = ovc.AddMultipleRackServers(multipleRackServers)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Added multiple rack-servers successfully.")
	}

	fmt.Println("-----------------------------")
	fmt.Println("Get Server Hardware List:")
	filters := []string{""}
	ServerList, err := ovc.GetServerHardwareList(filters, "", "", "", "")

	if err == nil {
		for i := 0; i < ServerList.Count; i++ {
			fmt.Println(ServerList.Members[i].Name)
			fmt.Println(ServerList.Members[i].MpModel)
		}

	} else {
		fmt.Println("Failed to fetch server List : ", err)
	}

	fmt.Println("-----------------------------")
	fmt.Println("Get Server Firmware List:")
	filters = []string{"serverModel='ProLiant DL380 Gen10'"}
	FirmwareList, err := ovc.GetServerFirmwareList(filters, "", "", "")
	if err == nil {
		for i := 0; i < FirmwareList.Count; i++ {
			fmt.Println(FirmwareList.Members[i].ServerName)
			fmt.Println(FirmwareList.Members[i].ServerModel)
		}

	} else {
		fmt.Println("Failed to fetch Firmware List : ", err)
	}

	fmt.Println("-----------------------------")
	fmt.Println("Get server hardware list by Name:")
	serverName, err := ovc.GetServerHardwareByName(ServerList.Members[0].Name)
	if err != nil {
		fmt.Println("Failed to fetch server hardware name: ", err)
	} else {

		fmt.Println("Server hardware details ")
		fmt.Println(serverName.Model)
		fmt.Println(serverName.URI)
		fmt.Println(serverName.UUID)

		fmt.Println("----------------------")
		fmt.Println("Server hardware IloIPAddress ")
		fmt.Println(serverName.GetIloIPAddress())

		fmt.Println("----------------------")
		fmt.Println("Server hardware MpIPAddresses ")
		if ovc.IsHardwareSchemaV2() {
			for i := 0; i < len(serverName.MpHostInfo.MpIPAddresses); i++ {
				fmt.Println(serverName.MpHostInfo.MpIPAddresses[i].Address)
			}
		}
	}

	fmt.Println("---------------------------")
	fmt.Println("Get server hardware list by URI:")

	ServerId, err := ovc.GetServerHardwareByUri(serverName.URI)
	if err == nil {
		fmt.Println(ServerId.URI)
	} else {
		fmt.Println("Failed to fetch server hardware : ", err)
	}

	fmt.Println("---------------------")
	fmt.Println("Get firmware inventory for a server-hardware:")
	firmware, err := ovc.GetServerFirmwareByUri(ServerId.URI)

	if err == nil {
		fmt.Println(firmware.ServerName)
	} else {
		fmt.Println("Failed to fetch Firmwares: ", err)
	}

	fmt.Println("----------------------")
	fmt.Println("Get Server Hardwares by SH Type and Enclosure Group:")
	serverHarware, err := ovc.GetAvailableHardware(ServerList.Members[0].ServerHardwareTypeURI, ServerList.Members[0].ServerGroupURI)

	if err == nil {
		fmt.Println(serverHarware.Type)
	} else {
		fmt.Println("Failed to fetch server hardware by URI: ", err)
	}

	fmt.Println("-----------------------------")
	fmt.Println("Refresh Server hardware:")
	refreshState := ov.ServerHardware{
		RefreshState: "RefreshPending",
	}
	err = ovc.RefreshServerHardware(ServerList.Members[0].UUID.String(), refreshState)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Server hardware refreshed successfully")
	}

	fmt.Println("-----------------------------")
	fmt.Println("Update Firmware Version to minimum iLO Firmware Version:")
	err = ovc.UpdateiLOFirmwareVersion(ServerList.Members[0].UUID.String())
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Updated firmware version successfully")
	}

	fmt.Println("-----------------------------")
	fmt.Println("Update oneTimeBoot to Network")
	bootOption := "Network" //bootOption can also be Normal, CDROM, HDD, USB
	err = ovc.SetOneTimeBoot(ServerList.Members[0].UUID.String(), bootOption)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Onetimeboot value changed from",
			ServerList.Members[0].OneTimeBoot, "to", bootOption)
	}

	fmt.Println("-----------------------------")
	fmt.Println("Update uidState to On")
	uidState := "On"
	err = ovc.SetUidState(ServerList.Members[0].UUID.String(), uidState)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("UidState value changed from",
			ServerList.Members[0].UidState, "to", uidState)
	}

	fmt.Println("-----------------------------")
	fmt.Println("Put the server into maintenance mode:")
	maintenanceMode := "true"
	err = ovc.SetMaintenanceMode(ServerList.Members[0].UUID.String(), maintenanceMode)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("MaintenanceMode value changed from",
			ServerList.Members[0].MaintenanceMode, "to", maintenanceMode)
	}

	// Change Power state using PATCH
	fmt.Println("-----------------------------")
	fmt.Println("Power off a server that is on, using the momentary press control:")
	power := map[string]interface{}{"powerState": "Off", "powerControl": "PressAndHold"}
	if power["powerState"] != ServerList.Members[0].PowerState {
		err = ovc.SetPowerState(ServerList.Members[0].UUID.String(), power)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("PowerState changed from", ServerList.Members[0].PowerState, "to", power["powerState"])
		}
	} else {
		fmt.Println("PowerState is already", power["powerState"])
	}

	fmt.Println("-----------------------------")
	fmt.Println("Power on a server that is off, using press and hold control:")
	power = map[string]interface{}{"powerState": "On", "powerControl": "MomentaryPress"}
	if power["powerState"] != ServerList.Members[0].PowerState {
		err = ovc.SetPowerState(ServerList.Members[0].UUID.String(), power)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("PowerState changed from", ServerList.Members[0].PowerState, "to", power["powerState"])
		}
	} else {
		fmt.Println("PowerState is already", power["powerState"])
	}

	fmt.Println("----------------------")
	fmt.Println("Get power status of a server:")
	powerState, err := serverName.GetPowerState()
	if err == nil {
		fmt.Println("Power state of the machine is ", powerState)
	} else {
		fmt.Println("Failed to fetch powerstate of the server: ", err)
	}
	//Trying to touggle power state of the server using put
	if powerState.String() == "On" {
		fmt.Println("Server is in powered on state ")
		serverName.PowerOff()
		powerState, _ := serverName.GetPowerState()
		fmt.Println("Power state of the machine is ", powerState)

	} else {

		fmt.Println("Server is in powered off state ")
		serverName.PowerOn()
		powerState, _ := serverName.GetPowerState()
		fmt.Println("Power state of the machine is ", powerState)

	}

	fmt.Println("--------------------")
	fmt.Println("Get ilo ipaddress of a server ")
	iloIpaddress := serverName.GetIloIPAddress()
	fmt.Println("ilo ip address of an server is =", iloIpaddress)

	fmt.Println("--------------------")
	fmt.Println("Get ilo ipaddress of all servers")
	for _, v := range ServerList.Members {
		fmt.Println("Server: ", v.Name, "ILO IP: ", v.GetIloIPAddress())
	}

	hardwareType, err := ovc.GetServerHardwareTypeByUri(ServerList.Members[0].ServerHardwareTypeURI)
	if hardwareType.Platform == "RackServer" {
		fmt.Println("-----------------------------")
		fmt.Println("Delete Server hardware:")
		if ServerList.Members[0].State == "NoProfileApplied" {
			err = ovc.DeleteServerHardware(ServerList.Members[0].URI)
			if err != nil {
				panic(err)
			} else {
				fmt.Println("Server hardware deleted successfully")
			}
		} else {
			fmt.Println("Server hardware cannot be removed while being used by a server profile.")
		}

	}

}
