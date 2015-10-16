package ov

import "github.com/docker/machine/drivers/oneview/utils"

// firmware additional properties introduced in 200
// "FirmwareOnly" - Updates the firmware without powering down the server hardware using using HP Smart Update Tools.
// "FirmwareAndOSDrivers" - Updates the firmware and OS drivers without powering down the server hardware using HP Smart Update Tools.
// "FirmwareOnlyOfflineMode" - Manages the firmware through HP OneView. Selecting this option requires the server hardware to be powered down.
type FirmwareOptionv200 struct {
	FirmwareInstallType string `json:"firmwareInstallType,omitempty"` // Specifies the way a Service Pack for ProLiant (SPP) is installed. This field is used if the 'manageFirmware' field is true. Possible values are
}

// ServerProfilev200 - v200 changes to ServerProfile
type ServerProfilev200 struct {
	TemplateCompliance       string        `json:"templateCompliance,omitempty"`       // v2 Compliant, NonCompliant, Unknown
	ServerProfileTemplateURI utils.Nstring `json:"serverProfileTemplateUri,omitempty"` // undocmented option
}
