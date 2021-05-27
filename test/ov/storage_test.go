package ov

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/stretchr/testify/assert"
)

// test unmarshalling a json payload that has progress
func TestSanStorageOptions(t *testing.T) {
	var (
		sanstorageoptions *ov.SanStorageOptions
		testjsondata      = ` {
    "complianceControl": "CheckedMinimum",
    "manageSanStorage": true,
    "hostOSType": "Windows 2012 / WS2012 R2",
    "sanSystemCredentials": [

    ],
    "volumeAttachments": [
      {
        "id": 1,
        "associatedTemplateAttachmentId": "{template_attachmentID}",
        "lun": null,
        "lunType": "Auto",
        "storagePaths": [
          {
            "connectionId": 4,
            "networkUri": "/rest/fc-networks/{networkID}",
            "isEnabled": true,
            "targetSelector": "Auto",
            "targets": [

            ]
          }
        ],
        "volumeUri": null,
        "volume": {
          "templateUri": "/rest/storage-volume-templates/{volume_templateID}",
          "properties": {
            "snapshotPool": "/rest/storage-pools/{snapshotpoolID}",
            "isDeduplicated": false,
            "storagePool": "/rest/storage-pools/{poolID}",
            "name": "testvol7",
            "description": "",
            "provisioningType": "Thin",
            "size": 268435456,
            "templateVersion": "1.1",
            "isShareable": false
          },
          "initialScopeUris": null,
          "isPermanent": true
        },
        "volumeStorageSystemUri": "/rest/storage-systems/{storage_systemID}",
        "bootVolumePriority": "NotBootable"
      }
    ]
  }`
	)
	err := json.Unmarshal([]byte(testjsondata), &sanstorageoptions)
	assert.NoError(t, err, fmt.Sprintf("Failed to unmarshal task object: %s, %+v\n", err, sanstorageoptions))
	assert.Equal(t, "Windows 2012 / WS2012 R2", sanstorageoptions.HostOSType)
	assert.Equal(t, "Auto", sanstorageoptions.VolumeAttachments[0].StoragePaths[0].TargetSelector)
	assert.Equal(t, "Thin", sanstorageoptions.VolumeAttachments[0].Volume.Properties.ProvisioningType)
}
