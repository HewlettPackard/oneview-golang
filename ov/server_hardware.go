package ov

import (
	"encoding/json"
	"github.com/docker/machine/log"
	"github.com/docker/machine/drivers/oneview/rest"
)

// get server hardware from ov
type ServerHardware struct {
	Type                   string  `json:"type,omitempty"`                  // "type": "server-hardware-3",
	Name                   string  `json:"name,omitempty"`                  // "name": "se05, bay 16",
	State                  string  `json:"state,omitempty"`                 // "state": "ProfileApplied",
	StateReason            string  `json:"stateReason,omitempty"`           // "stateReason": "NotApplicable",
	AssetTag               string  `json:"assetTag,omitempty"`              // "assetTag": "[Unknown]",
	Category               string  `json:"category,omitempty"`              // "category": "server-hardware",
	Created                string  `json:"created,omitempty"`               // "created": "2015-08-14T21:02:01.537Z",
	Description            Nstring `json:"description,omitempty"`           // "description": null,
	ETAG                   string  `json:"eTag,omitempty"`                  // "eTag": "1441147370086",
	FormFactor             string  `json:"formFactor,omitempty"`            // "formFactor": "HalfHeight",
	LicensingIntent        string  `json:"licensingIntent,omitempty"`       // "licensingIntent": "OneView",
	LocationURI            string  `json:"locationUri,omitempty"`           // "locationUri": "/rest/enclosures/092SN51207RR",
	MemoryMb               int     `json:"memoryMb,omitempty"`              // "memoryMb": 262144,
	Model                  string  `json:"model,omitempty"`                 // "model": "ProLiant BL460c Gen9",
	Modified               string  `json:"modified,omitempty"`              // "modified": "2015-09-01T22:42:50.086Z",
	MpDnsName              string  `json:"mpDnsName,omitempty"`             // "mpDnsName": "ILO2M25090RMW",
	MpFirwareVersion       string  `json:"mpFirmwareVersion,omitempty"`     // "mpFirmwareVersion": "2.03 Nov 07 2014",
	MpIpAddress            string  `json:"mpIpAddress,omitempty"`           // "mpIpAddress": "172.28.3.136",
	MpModel                string  `json:"mpModel,omitempty"`               // "mpModel": "iLO4",
	PartNumber             string  `json:"partNumber,omitempty"`            // "partNumber": "727021-B21",
	Position               int     `json:"position,omitempty"`              // "position": 16,
	PowerLock              bool    `json:"powerLock,omitempty"`             // "powerLock": false,
	PowerState             string  `json:"powerState,omitempty"`            // "powerState": "Off",
	ProcessorCoreCount     int     `json:"processorCoreCount,omitempty"`    // "processorCoreCount": 14,
	ProcessorCount         int     `json:"processorCount,omitempty"`        // "processorCount": 2,
	ProcessorSpeedMhz      int     `json:"processorSpeedMhz,omitempty"`     // "processorSpeedMhz": 2300,
	ProcessorType          string  `json:"processorType,omitempty"`         // "processorType": "Intel(R) Xeon(R) CPU E5-2695 v3 @ 2.30GHz",
	RefreshState           string  `json:"refreshState,omitempty"`          // "refreshState": "NotRefreshing",
	RomVersion             string  `json:"romVersion,omitempty"`            // "romVersion": "I36 11/03/2014",
	SerialNumber           string  `json:"serialNumber,omitempty"`          // "serialNumber": "2M25090RMW",
	ServerGroupURI         string  `json:"serverGroupUri,omitempty"`        // "serverGroupUri": "/rest/enclosure-groups/56ad0069-8362-42fd-b4e3-f5c5a69af039",
	ServerHardwareTypeURI  string  `json:"serverHardwareTypeUri,omitempty"` // "serverHardwareTypeUri": "/rest/server-hardware-types/DB7726F7-F601-4EA8-B4A6-D1EE1B32C07C",
	ServerProfileURI       string  `json:"serverProfileUri,omitempty"`      // "serverProfileUri": "/rest/server-profiles/9979b3a4-646a-4c3e-bca6-80ca0b403a93",
	ShortModel             string  `json:"shortModel,omitempty"`            // "shortModel": "BL460c Gen9",
	Status                 string  `json:"status,omitempty"`                // "status": "Warning",
	URI                    string  `json:"uri,omitempty"`                   // "uri": "/rest/server-hardware/30373237-3132-4D32-3235-303930524D57",
	UUID                   string  `json:"uuid,omitempty"`                  // "uuid": "30373237-3132-4D32-3235-303930524D57",
	VirtualSerialNumber    string  `json:"VirtualSerialNumber,omitempty"`   // "virtualSerialNumber": "",
	VirtualUUID            string  `json:"virtualUuid,omitempty"`           // "virtualUuid": "00000000-0000-0000-0000-000000000000"
	Client                 *OVClient
}

// get a server profiles
func (c *OVClient) GetServerHardware(uri string)(ServerHardware, error) {

	var hardware ServerHardware
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions( c.GetAuthHeaderMap() )

	// rest call
	data, err := c.RestAPICall(rest.GET, uri , nil)
	if err != nil {
		return hardware, err
	}

	log.Debugf("GetServerHardware %s", data)
	if err := json.Unmarshal([]byte(data), &hardware); err != nil {
		return hardware, err
	}
	hardware.Client = c
	return hardware, nil
}
