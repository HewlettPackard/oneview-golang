package ov

// logical drive options
type LogicalDrive struct {
  Bootable   bool `json:"bootable,omitempty"`    // "bootable": true,
  RaidLevel  string `json:"raidLevel,omitempty"` // "raidLevel": "RAID1"
}

type LocalStorageOptions struct {                                            // "localStorage": {
  ManageLocalStorage   bool            `json:"manageLocalStorage,omitempty"` // "manageLocalStorage": true,
  LogicalDrives        []LogicalDrive  `json:"logicalDrives,omitempty"`      // "logicalDrives": [],
  Initialize           bool            `json:"initialize,omitempty"`         // 				"initialize": true
}                                                                            // 		},


// 		"sanStorage": {
// 				"volumeAttachments": [],
// 				"manageSanStorage": false
// 		},
