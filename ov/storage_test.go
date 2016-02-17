package ov

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test unmarshalling a json payload that has progress
func TestSanStorageOptions(t *testing.T) {
	var (
		sanstorageoptions *SanStorageOptions
		testjsondata      = `{
            "hostOSType": "Windows 2012 / WS2012 R2",
            "manageSanStorage": true,
            "volumeAttachments": [{
                "id": 1,
                "lunType": "Auto",
                "volumeUri": "/rest/storage-volumes/{volumeID}",
                "volumeStoragePoolUri": "/rest/storage-pools/{poolID}",
                "volumeStorageSystemUri": "/rest/storage-systems/{systemID}",
                "storagePaths": [{
                    "storageTargetType": "TargetPorts",
                    "storageTargets": ["{WWPN1}", "{WWPN2}"],
                    "connectionId": 1,
                    "isEnabled": true
                }]
            },
            {
                "id": 2,
                "lunType": "Manual",
                "lun": "1",
                "volumeUri": "/rest/storage-volumes/{volumeID}",
                "volumeStoragePoolUri": "/rest/storage-pools/{poolID}",
                "volumeStorageSystemUri": "/rest/storage-systems/{systemID}",
                "storagePaths": [{
                    "storageTargetType": "Auto",
                    "storageTargets": [],
                    "connectionId": 1,
                    "isEnabled": true
                }]
            }]
        }`
	)
	err := json.Unmarshal([]byte(testjsondata), &sanstorageoptions)
	assert.NoError(t, err, fmt.Sprintf("Failed to unmarshal task object: %s, %+v\n", err, sanstorageoptions))
	assert.Equal(t, "Windows 2012 / WS2012 R2", sanstorageoptions.HostOSType)
	assert.Equal(t, "TargetPorts", sanstorageoptions.VolumeAttachments[0].StoragePaths[0].StorageTargetType)
	assert.Equal(t, "Auto", sanstorageoptions.VolumeAttachments[1].StoragePaths[0].StorageTargetType)
}
