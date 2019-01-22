package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
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
		clientOV     *ov.OVClient
		lig_name     = "DemoLogicalInterConnectGroup"
		lig_type     = "logical-interconnect-groupV5"
		new_lig_name = "RenamedLogicalInterConnectGroup"
	)
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		800)

	fmt.Println("#..........Getting Logical Interconnect Collection.....")
	sort := "name:desc"
	logicalInterconnectGroupList, _ := ovc.GetLogicalInterconnectGroups("", sort)
	fmt.Println(logicalInterconnectGroupList)

	fmt.Println("#..........Creating Logical Interconnect .....#")
	locationEntry_first := ov.LocationEntry{Type: "Bay", RelativeValue: 3}
	locationEntry_second := ov.LocationEntry{Type: "Enclosure", RelativeValue: 1}
	locationEntries1 := new([]ov.LocationEntry)
	*locationEntries1 = append(*locationEntries1, locationEntry_first)
	*locationEntries1 = append(*locationEntries1, locationEntry_second)
	locationEntry_third := ov.LocationEntry{Type: "Bay", RelativeValue: 6}
	locationEntry_four := ov.LocationEntry{Type: "Enclosure", RelativeValue: 1}
	locationEntries2 := new([]ov.LocationEntry)
	*locationEntries2 = append(*locationEntries2, locationEntry_third)
	*locationEntries2 = append(*locationEntries2, locationEntry_four)

	logicalLocation1 := ov.LogicalLocation{LocationEntries: *locationEntries1}
	logicalLocation2 := ov.LogicalLocation{LocationEntries: *locationEntries2}
	interconnectMapEntryTemplate1 := ov.InterconnectMapEntryTemplate{LogicalLocation: logicalLocation1, PermittedInterconnectTypeUri: "/rest/interconnect-types/59080afb-85b5-43ae-8c69-27c08cb91f3a", EnclosureIndex: 1}
	interconnectMapEntryTemplate2 := ov.InterconnectMapEntryTemplate{LogicalLocation: logicalLocation2, PermittedInterconnectTypeUri: "/rest/interconnect-types/59080afb-85b5-43ae-8c69-27c08cb91f3a", EnclosureIndex: 1}
	interconnectMapEntryTemplates := new([]ov.InterconnectMapEntryTemplate)
	*interconnectMapEntryTemplates = append(*interconnectMapEntryTemplates, interconnectMapEntryTemplate1)

	*interconnectMapEntryTemplates = append(*interconnectMapEntryTemplates, interconnectMapEntryTemplate2)

	interconnectMapTemplate := ov.InterconnectMapTemplate{InterconnectMapEntryTemplates: *interconnectMapEntryTemplates}
	fmt.Println(&interconnectMapTemplate)

	enclosureIndexes := []int{1}

	ethernetSettings := ov.EthernetSettings{Type: "EthernetInterconnectSettingsV4", URI: "/ethernetSettings", Name: "defaultEthernetSwitchSettings", ID: "e2f5a9ac-8b0e-435c-8a3d-db00407fc7c7",
		InterconnectType: "Ethernet", EnableIgmpSnooping: newFalse(), IgmpIdleTimeoutInterval: 260, EnableFastMacCacheFailover: newTrue(),
		MacRefreshInterval: 5, EnableNetworkLoopProtection: newTrue(), EnablePauseFloodProtection: newTrue(), EnableRichTLV: newFalse()}
	telemetryConfig := ov.TelemetryConfiguration{Type: "telemetry-configuration", EnableTelemetry: true, SampleCount: 12, SampleInterval: 300}
	snmpConfig := ov.SnmpConfiguration{Type: "snmp-configuration", Enabled: newFalse(), Category: "snmp-configuration", V3Enabled: true}
	qosActiveConfig := ov.ActiveQosConfig{Type: "QosConfiguration", Category: "qos-aggregated-configuration", ConfigType: "Passthrough"}
	qosConfig := ov.QosConfiguration{ActiveQosConfig: qosActiveConfig, Type: "qos-aggregated-configuration", Category: "qos-aggregated-configuration"}

	logicalInterconnectGroup := ov.LogicalInterconnectGroup{Type: lig_type, EthernetSettings: &ethernetSettings, Name: lig_name, TelemetryConfiguration: &telemetryConfig,
		InterconnectMapTemplate: &interconnectMapTemplate, EnclosureType: "SY12000", EnclosureIndexes: enclosureIndexes,
		InterconnectBaySet: 3, RedundancyType: "Redundant", SnmpConfiguration: &snmpConfig, QosConfiguration: &qosConfig}
	er := ovc.CreateLogicalInterconnectGroup(logicalInterconnectGroup)
	if er != nil {
		fmt.Println("........Logical Interconnect Group Creation failed:", er)
	}
	if er == nil {
		fmt.Println(".....Logical Interconnect Group Creation Success....")
	}

	fmt.Println("....  Logical Interconnect Group by Name.....")
	lig, _ := ovc.GetLogicalInterconnectGroupByName(lig_name)
	fmt.Println(lig)

	fmt.Println("... Logical Interconnect Group by URI ....")
	uri := lig.URI
	lig_uri, _ := ovc.GetLogicalInterconnectGroupByUri(uri)
	fmt.Println(lig_uri)
	fmt.Println("... Updating LogicalInterconnectGroup ...")
	fmt.Println("")
	lig_uri.Name = new_lig_name
	err := ovc.UpdateLogicalInterconnectGroup(lig_uri)
	if err != nil {
		panic(err)
	}
	if err == nil {
		fmt.Println(".....Updated Logical Interconnect Group Successfully....")
	}
	fmt.Println("... Deleting LogicalInterconnectGroup ...")
	del_err := ovc.DeleteLogicalInterconnectGroup(lig_uri.Name)
	if del_err != nil {
		panic(del_err)
	}
	if del_err == nil {
		fmt.Println(".....Deleted Logical Interconnect Group Successfully....")
	}
}
