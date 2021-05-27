package ov

type FirmwareDrivers struct {
	BaselineShortName string `json:"baselineShortName,omitempty"`
	BundleSize        int    `json:"bundleSize,omitempty"`
	BundleType        string `json:"bundleType,omitempty"`
	Category          string `json:"category,omitempty"`
	Created           string `json:"created,omitempty"`
	Description       string `json:"description,omitempty"`
	ETAG              string `json:"eTag,omitempty"`
	EsxiOsDriverMetaData
}
