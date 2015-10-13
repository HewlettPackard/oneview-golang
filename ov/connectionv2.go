package ov

// Connectionv200 server profile object for ov
type Connectionv200 struct {
	AllocatedVFs int    `json:"allocatedVFs,omitempty"` // allocatedVFs The number of virtual functions allocated to this connection. This value will be null. integer read only
	RequestedVFs string `json:"requestedVFs,omitempty"` // requestedVFs This value can be "Auto" or 0. string
}
