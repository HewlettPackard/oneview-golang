package ov

import (
	"encoding/json"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type SasInterconnectList struct {
	Type        string             `json:"type,omitempty"`        // "Type": "SasInterconnectListV3"
	URI         utils.Nstring      `json:"uri,omitempty"`         // "uri": "/rest/sas-interconnects?sort=name:asc"
	Category    string             `json:"category,omitempty"`    // "category": sas-interconnects
	Start       int                `json:"start,omitempty"`       // "start": 0,
	Count       int                `json:"count,omitempty"`       // "count": 1,
	Total       int                `json:"total,omitempty"`       // "total": 1,
	PrevPageURI utils.Nstring      `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring      `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	Members     []SasInterconnects `json:"members,omitempty"`     // "members": []
}

type SasInterconnects struct {
	Type                      string                   `json:"type,omitempty"`                      // "type": "SasInterconnectListV3"
	URI                       utils.Nstring            `json:"uri,omitempty"`                       // "uri": "/rest/sas-interconnects/TWT732W0CY"
	Category                  string                   `json:"category"`                            // "category": "sas-interconnects"
	ETag                      string                   `json:"eTag,omitempty"`                      // "eTag": "463bd328-ffc8-40ae-9603-6136fa9e6e58",
	Created                   string                   `json:"created,omitempty"`                   // "created": "2018-08-02T15:49:59.963Z",
	Modified                  string                   `json:"modified,omitempty"`                  // "modified": "2018-12-03T18:26:43.335Z",
	RefreshState              string                   `json:"refreshState,omitempty"`              // "refreshState": "NotRefreshing"
	StateReason               string                   `json:"stateReason,omitempty"`               //	"stateReason": null
	InterconnectLocation      SASInterconnectLocation  `json:"interconnectLocation"`                // "interconnectLocation": {}
	PowerState                string                   `json:"powerState,omitempty"`                // "powerState": "On",
	HardResetState            string                   `json:"hardResetState,omitempty"`            // "hardResetState": "Normal"
	SoftResetState            string                   `json:"softResetState,omitempty"`            // "softResetState": "Normal"
	InterconnectIP            string                   `json:"interconnectIP,omitempty"`            // "interconnectIP": "fe80::9eb6:54ff:fe91:2170"
	FirmwareVersion           string                   `json:"firmwareVersion,omitempty"`           // "firmwareVersion": "1.5.11.0",
	ProductName               string                   `json:"productName,omitempty"`               // "productName": "Synergy 12Gb SAS Connection Module"
	Model                     string                   `json:"model,omitempty"`                     // "model": "Synergy 12Gb SAS Connection Module"
	SasWWN                    string                   `json:"sasWWN,omitempty"`                    // "sasWWN": "50014380421BE900"
	EnclosureURI              utils.Nstring            `json:"enclosureUri,omitempty"`              // "enclosureUri": "/rest/enclosures/013645CN759000AC"
	SerialNumber              string                   `json:"serialNumber,omitempty"`              // "serialNumber": "TWT732W0CY",
	PartNumber                string                   `json:"partNumber,omitempty"`                // "partNumber": "755985-B21",
	SasLogicalInterconnectURI utils.Nstring            `json:"sasLogicalInterconnectUri,omitempty"` // "sasLogicalInterconnectUri": "/rest/sas-logical-interconnects/c6e17ed8-de41-4d53-aa50-2da58a0d63b8"
	InterconnectTypeURI       utils.Nstring            `json:"interconnectTypeUri,omitempty"`       //"interconnectTypeURI": "/rest/sas-interconnect-types/Synergy12GbSASConnectionModule"
	PortCount                 int                      `json:"portCount,omitempty"`                 // "portCount": 12,
	EnclosureName             string                   `json:"enclosureName,omitempty"`             // "enclosureName": "CEC"
	UIDState                  string                   `json:"uidState,omitempty"`                  // "uidState": "Off"
	SasPorts                  []SasPorts               `json:"sasPorts,omitempty"`                  // "sasPorts": []
	RemoteSupportSettings     SASRemoteSupportSettings `json:"remoteSupportSettings,omitempty"`     // "remoteSupportSettings": {}
	SparePartNumber           string                   `json:"sparePartNumber,omitempty"`           // "sparePartNumber": "758686-001"
	ScopesURI                 string                   `json:"scopesUri,omitempty"`                 // "scopesUri": "/rest/scopes/resources/rest/sas-interconnects/TWT732W0CY",
	LogicalSasInterconnectURI string                   `json:"logicalSasInterconnectUri,omitempty"` //"logicalSasInterconnectUri": "/rest/sas-logical-interconnects/c6e17ed8-de41-4d53-aa50-2da58a0d63b8"
	Description               utils.Nstring            `json:"description,omitempty"`               // Description": null
	State                     string                   `json:"state,omnitempty"`                    // "state": "Configured"
	Status                    string                   `json:"status,omitempty"`                    // "status": "OK"
	Name                      string                   `json:"name,omitempty"`                      // "name": "CEC, interconnect 4"
}

type SASInterconnectLocation struct {
	LocationEntries []SASLocationEntries `json:"locationEntries"` // "locationEntries": []
}

type SASLocationEntries struct {
	Type  string `json:"type"`  // "type": "bay"
	Value string `json:"value"` // "value": "4"
}

type SasPorts struct {
	Type               string        `json:"type,omitempty"`               // "type": "sas-port",
	URI                string        `json:"uri,omitempty"`                //"uri": "/rest/sas-interconnects/TWT732W0CY/sas-ports/e0ff4925-b690-45ab-a186-9f5a9d7362dd"
	Category           string        `json:"category,omitempty"`           // "category": "sas-ports"
	ETag               string        `json:"eTag,omitempty"`               // "eTag": "2019-04-10T16:15:11.080Z",
	Created            string        `json:"created,omitempty"`            // "created": "2019-04-10T13:40:56.059Z",
	Modified           string        `json:"modified,omitempty"`           // "modified": "2019-04-10T16:15:11.080Z",
	PortIdentifier     string        `json:"portIdentifier,omitempty"`     // "portIdentifier": "11",
	PhyCount           int           `json:"phyCount,omitempty"`           // "phyCount": 4,
	PortName           string        `json:"portName,omitempty"`           // "portName": "11",
	PortType           string        `json:"portType,omitempty"`           // "portType": "Downlink",
	PortLocation       string        `json:"portLocation,omitempty"`       // "portLocation": "11",
	Enabled            bool          `json:"enabled,omitempty"`            // "enabled": true,
	PortStatusReason   string        `json:"portStatusReason,omitempty"`   // "portStatusReason": "None",
	LinkedPortURI      utils.Nstring `json:"linkedPortUri,omitempty"`      // "linkedPortUri": null,
	ContainerDeviceURI string        `json:"containerDeviceUri,omitempty"` // "containerDeviceUri": "/rest/sas-interconnects/TWT732W0CY",
	Description        string        `json:"description,omitempty"`        // "description": null,
	State              string        `json:"state,omitempty"`              // "state": "Unlinked",
	Status             string        `json:"status,omitempty"`             // "status": "DISABLED",
	Name               string        `json:"name,omitempty"`               // "name": "11"
}

type SASRemoteSupportSettings struct {
	RemoteSupportURI           string `json:"remoteSupportUri,omitempty"`           // "remoteSupportUri": "/rest/support/sas-interconnects/TWT732W0CY"
	SupportDataCollectionsURI  string `json:"supportDataCollectionsUri,omitempty"`  // "supportDataCollectionsUri": "/rest/support/data-collections?deviceID=TWT732W0CY&category=sas-interconnects"
	Destination                string `json:"destination,omitempty"`                // "destination": ""
	RemoteSupportCurrentState  string `json:"remoteSupportCurrentState,omitempty"`  // "remoteSupportCurrentState": "Unregistered"
	SupportDataCollectionState string `json:"supportDataCollectionState,omitempty"` // "supportDataCollectionState": null
	SupportDataCollectionType  string `json:"supportDataCollectionType,omitempty"`  // "supportDataCollectionType": null
	SupportState               string `json:"supportState,omitempty"`               // "supportState": "Disabled"
}

func (c *OVClient) GetSasInterconnects(count string, filter []string, query string, sort string, start string) (SasInterconnectList, error) {
	var (
		uri              = "/rest/sas-interconnects"
		q                map[string]interface{}
		sasInterconnects SasInterconnectList
	)
	q = make(map[string]interface{})
	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
	}

	if start != "" {
		q["start"] = start
	}

	if count != "" {
		q["count"] = count
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return sasInterconnects, err
	}

	log.Debugf("GetSasInterconnects %s", data)
	if err := json.Unmarshal([]byte(data), &sasInterconnects); err != nil {
		return sasInterconnects, err
	}
	return sasInterconnects, nil
}
