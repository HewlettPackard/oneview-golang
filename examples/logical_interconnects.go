package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func newTrue() *bool {
	b := true
	return &b
}
func newFalse() *bool {
	b := false
	return &b
}

func main() {
	var (
		clientOV        *ov.OVClient
		id              = "d4468f89-4442-4324-9c01-624c7382db2d"
		macAddress      = "94:57:A5:67:2C:BE"
		internalVlan    = "504"
		interconnectURI = "/rest/interconnects/aca6687f-1370-46cd-b832-7e3192dbddfd"
		externalVlan    = "504"
		tcId            = "1"
	)
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		800,
		"*")

	fmt.Println("....  Logical Interconnects Collection .....")
	logicalInterconnectList, _ := ovc.GetLogicalInterconnects("", "0", "10")
	fmt.Println(logicalInterconnectList)

	fmt.Println("....  Logical Interconnect by Id.....")
	lig, _ := ovc.GetLogicalInterconnectById(id)
	fmt.Println(lig)

	fmt.Println("....  Logical Interconnect PortMonitor.....")
	portMonitor, _ := ovc.GetLogicalInterconnectPortMonitor(id)
	fmt.Println(portMonitor)

	fmt.Println("....  Logical Interconnect EthernetSettings.....")
	ethernetSettings, _ := ovc.GetLogicalInterconnectEthernetSettings(id)
	fmt.Println(ethernetSettings)

	fmt.Println("....  Logical Interconnect Firmware.....")
	firmware, _ := ovc.GetLogicalInterconnectFirmware(id)
	fmt.Println(firmware)

	fmt.Println("....  Logical Interconnect SNMPConfiguration.....")
	snmpconfig, _ := ovc.GetLogicalInterconnectSNMPConfiguration(id)
	fmt.Println(snmpconfig)

	fmt.Println("....  Logical Interconnect Forwarding Information.....")
	var filter []string
	fi, _ := ovc.GetLogicalInterconnectForwardingInformation(filter, id)
	fmt.Println(fi)

	fmt.Println("....  Logical Interconnect Forwarding Information By Mac Address.....")
	fi_mac, _ := ovc.GetLogicalInterconnectForwardingInformationByMacAddress(macAddress, id)
	fmt.Println(fi_mac)

	fmt.Println("....  Logical Interconnect Forwarding Information By Internal Vlan.....")
	fi_intern_vlan, _ := ovc.GetLogicalInterconnectForwardingInformationByInternalVlan(internalVlan, id)
	fmt.Println(fi_intern_vlan)

	fmt.Println("....  Logical Interconnect Forwarding Information By Interconnect URI and ExternalVlan.....")
	fi_interconnect_external, _ := ovc.GetLogicalInterconnectForwardingInformationByInterconnectAndExternalVlan(interconnectURI, externalVlan, id)
	fmt.Println(fi_interconnect_external)

	fmt.Println("....  Logical Interconnect Internal VLAN IDs for the provisioned networks.....")
	fi_internal_vlan, _ := ovc.GetLogicalInternalVlans(id)
	fmt.Println(fi_internal_vlan)

	fmt.Println("....  Logical Interconnect QOS Configuration.....")
	fi_qos_config, _ := ovc.GetLogicalQosAggregatedConfiguration(id)
	fmt.Println(fi_qos_config)

	fmt.Println("....  Logical Interconnect Unassigned Ports for Port Monitor.....")
	port_monitor_ports := ovc.GetUnassignedPortsForPortMonitor(id)
	fmt.Println(port_monitor_ports)

	fmt.Println("....  Logical Interconnect Unassigned Uplink Ports for Port Monitor.....")
	uplink_port_monitor_ports, _ := ovc.GetUnassignedUplinkPortsForPortMonitor(id)
	fmt.Println(uplink_port_monitor_ports)

	fmt.Println("....  Logical Interconnect Telemetry Configuration.....")
	telemetry_config, _ := ovc.GetTelemetryConfigurations(id, "1")
	fmt.Println(telemetry_config)

	fmt.Println("....  Updating Logical Interconnect Consistent State.....")
	var liUris []utils.Nstring
	liUris = append(liUris, utils.NewNstring("/rest/logical-interconnects/d4468f89-4442-4324-9c01-624c7382db2d"))
	liCompliance := ov.LogicalInterconnectCompliance{Type: "li-compliance", LogicalInterconnectUris: liUris, Description: ""}
	err_compliance := ovc.UpdateLogicalInterconnectConsistentState(liCompliance)
	if err_compliance != nil {
		fmt.Println("Could not update ConsistentState of Logical Interconnect", err_compliance)
	}

	fmt.Println("....  Updating Logical Interconnect EthernetSetting.....")
	liEthernetSettings := ov.EthernetSettings{Type: "EthernetInterconnectSettingsV4", InterconnectType: "Ethernet", URI: utils.NewNstring("/rest/logical-interconnects/d4468f89-4442-4324-9c01-624c7382db2d/ethernetSettings"), ID: "d4468f89-4442-4324-9c01-624c7382db2d"}
	err_ethernet := ovc.UpdateLogicalInterconnectEthernetSettings(liEthernetSettings, id)
	if err_ethernet != nil {
		fmt.Println("Could not update Ethernet Settings of Logical Interconnect", err_ethernet)
	}

	fmt.Println("....  Updating Logical Interconnect Firmware.....")
	liFirmware := ov.Firmware{Command: "Update", EthernetActivationDelay: 5, EthernetActivationType: "Parallel", FcActivationDelay: 5, FcActivationType: "Parallel", Force: false, SppUri: utils.NewNstring("/rest/firmware-drivers/SPP_2018_06_20180709_for_HPE_Synergy_Z7550-96524")}
	err_firmware := ovc.UpdateLogicalInterconnectFirmware(liFirmware, id)
	if err_firmware != nil {
		fmt.Println("Could not update Firmware of Logical Interconnect", err_firmware)
	}

	fmt.Println("....  Updating Logical Interconnect InternalNetworks.....")
	var internalNetworks []utils.Nstring
	internalNetworks = append(internalNetworks, utils.NewNstring("/rest/ethernet-networks/a71b9c9e-b044-48ee-8e4e-26ced1a9a9ef"))
	err_networks := ovc.UpdateLogicalInterconnectInternalNetworks(internalNetworks, id)
	if err_networks != nil {
		fmt.Println("Could not update Internal Networks of Logical Interconnect", err_networks)
	}

	fmt.Println("....  Updating Logical Interconnect QOS Configuration.....")
	liActiveQosConfig := ov.ActiveQosConfig{Type: "QosConfiguration", Category: "qos-aggregated-configuration", ConfigType: "Passthrough"}
	liQosConfig := ov.QosConfiguration{Type: "qos-aggregated-configuration", Category: "qos-aggregated-configuration", ActiveQosConfig: liActiveQosConfig}

	err_qos := ovc.UpdateLogicalInterconnectQosConfigurations(liQosConfig, id)
	if err_qos != nil {
		fmt.Println("Could not update QOS Configuration of Logical Interconnect", err_qos)
	}

	fmt.Println("....  Updating Logical Interconnect SNMP Configuration.....")
	liSNMPConfig := ov.SnmpConfiguration{Type: "snmp-configuration", Category: "snmp-configuration", V3Enabled: newTrue()}

	err_snmp := ovc.UpdateLogicalInterconnectSNMPConfigurations(liSNMPConfig, id)
	if err_snmp != nil {
		fmt.Println("Could not update SNMP Configuration of Logical Interconnect", err_snmp)
	}

	fmt.Println("....  Updating Logical Interconnect Configuration.....")
	err_conf := ovc.UpdateLogicalInterconnectConfigurations(id)
	if err_conf != nil {
		fmt.Println("Could not update Configuration of Logical Interconnect", err_conf)
	}

	fmt.Println("....  Updating Logical Interconnect Port Monitor Configuration.....")
	liPMConfig := ov.PortMonitor{Type: "port-monitor", Category: "port-monitor", ETAG: "8a302a85-ec4d-4214-a3e0-10ef71d28769", Name: "name2095641007-1533682087640"}

	err_pm := ovc.UpdateLogicalInterconnectPortMonitor(liPMConfig, id)
	if err_pm != nil {
		fmt.Println("Could not update PortMonitor Configuration of Logical Interconnect", err_pm)
	}

	fmt.Println("....  Updating Logical Interconnect Telemetry  Configuration.....")
	liTMConfig := ov.TelemetryConfiguration{Type: "telemetry-configuration", EnableTelemetry: newTrue(), SampleInterval: 300, SampleCount: 12, Name: "name771327580-1533682118441"}

	err_tm := ovc.UpdateLogicalInterconnectTelemetryConfigurations(liTMConfig, id, tcId)
	if err_tm != nil {
		fmt.Println("Could not update PortMonitor Configuration of Logical Interconnect", err_tm)
	}

}
