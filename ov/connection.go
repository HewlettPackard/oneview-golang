package ov

// ServerProfile , server profile object for ov
type Connection struct {
  ID                int     `json:"id,omitempty"`                // "id": 1,
  Name              string  `json:"name,omitempty"`              // "name": "",
  FunctionType      string  `json:"functionType,omitempty"`      // "functionType": "Ethernet",
  DeploymentStatus  string  `json:"deploymentStatus,omitempty"`  // "deploymentStatus": "Deployed",
  NetworkURI        string  `json:"networkUri,omitempty"`        // "networkUri": "/rest/ethernet-networks/cc3088c1-605b-4e99-8084-457f3be3a368",
  PortID            string  `json:"portId,omitempty"`            // "portId": "Flb 1:1-a",
  InterconnectURI   string  `json:"interconnectUri,omitempty"`   // "interconnectUri": "/rest/interconnects/5bfba9b4-4591-44f1-8abf-8cc877522ae5",
  MacType           string  `json:"macType,omitempty"`           // "macType": "Physical",
  WWPNType          string  `json:"wwpnType,omitempty"`          // "wwpnType": "Physical",
  MAC               string  `json:"mac,omitempty"`               // "mac": "34:64:A9:BB:E6:98",
  WWNN              Nstring `json:"wwnn,omitempty"`              // "wwnn": null,
  WWPN              Nstring `json:"wwpn,omitempty"`              // "wwpn": null,
  RequestedMbps     string  `json:"requestedMbps,omitempty"`     // "requestedMbps": "1000",
  AllocatedMbps     int     `json:"allocatedMbps,omitempty"`     // "allocatedMbps": 1000,
  MaximumMbps       int     `json:"maximumMbps,omitempty"`       // "maximumMbps": 2000,
}
