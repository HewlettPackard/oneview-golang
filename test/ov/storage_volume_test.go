package ov

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateStorageVolume(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_storage_volume")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test ethernet network already exist
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testSVol, err := c.GetStorageVolumeByName(testName)
		assert.NoError(t, err, "CreateStorageVolume get the Storage Volume error -> %s", err)

		pMap := d.Tc.GetTestData(d.Env, "ProvisioningParameters").(map[string]interface{})

		//provParams := ov.ProvisioningParameters{StoragePoolUri: utils.NewNstring(pMap["storagePoolUri"].(string)), RequestedCapacity: pMap["requestedCapacity"].(string), ProvisionType: pMap["provisionType"].(string), Shareable: pMap["shareable"].(bool)}

		if testSVol.URI.IsNil() {
			testSVol = ov.StorageVolumeV3{
				Name:             testName,
				StorageSystemUri: utils.NewNstring(d.Tc.GetTestData(d.Env, "StorageSystemUri").(string)),
				Type:             d.Tc.GetTestData(d.Env, "Type").(string),
				//				ProvisioningParameters: provParams,
			}

			// not changed after this TODO: update to storage volume tests
			err := c.CreateStorageVolume(testSVol)
			assert.NoError(t, err, "CreateStorageVolume error -> %s", err)

			err = c.CreateStorageVolume(testSVol)
			assert.Error(t, err, "CreateStorageVolume should error because the Storage volume already exists, err-> %s", err)

		} else {
			log.Warnf("The storage volume already exist, so skipping CreateStorageVolume test for %s", testName)
		}

		// reload the test profile that we just created
		testSVol, err = c.GetStorageVolumeByName(testName)
		assert.NoError(t, err, "GetStorageVolumeByName error -> %s", err)
	}

}

func TestGetStorageVolumeByName(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testName string
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_storage_volume")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testSVol, err := c.GetStorageVolumeByName(testName)
		assert.NoError(t, err, "GetStorageVolumeByName thew an error -> %s", err)
		assert.Equal(t, testName, testSVol.Name)

		testSVol, err = c.GetStorageVolumeByName("bad")
		assert.NoError(t, err, "GetStorageVolumeByName with fake name -> %s", err)
		assert.Equal(t, "", testSVol.Name)

	} else {
		d, c = getTestDriverU("test_storage_volume")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetStorageVolumeByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetStorageVolumes(t *testing.T) {
	var (
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_storage_volume")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		sVols, err := c.GetStorageVolumes("", "")
		assert.NoError(t, err, "GetStorageVolumes threw error -> %s, %+v\n", err, sVols)

		sVols, err = c.GetStorageVolumes("", "name:asc")
		assert.NoError(t, err, "GetStorageVolumes name:asc error -> %s, %+v\n", err, sVols)

	} else {
		_, c = getTestDriverU("test_storage_volume")
		data, err := c.GetStorageVolumes("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeleteStorageVolumeNotFound(t *testing.T) {
	var (
		c        *ov.OVClient
		testName = "fake"
		testSVol ov.StorageVolumeV3
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_storage_volume")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteStorageVolume(testName)
		assert.NoError(t, err, "DeleteStorageVolume err-> %s", err)

		testSVol, err = c.GetStorageVolumeByName(testName)
		assert.NoError(t, err, "GetStorageVolumeByName with deleted storage volume -> %+v", err)
		assert.Equal(t, "", testSVol.Name, fmt.Sprintf("Problem getting storage volume name, %+v", testSVol))
	} else {
		_, c = getTestDriverU("test_storage_volume")
		err := c.DeleteStorageVolume(testName)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testSVol))
	}
}

func TestDeleteStorageVolume(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
		testName string
		testSVol ov.StorageVolumeV3
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_storage_volume")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeleteStorageVolume(testName)
		assert.NoError(t, err, "DeleteStorageVolume err-> %s", err)

		testSVol, err = c.GetStorageVolumeByName(testName)
		assert.NoError(t, err, "GetStorageVolumeByName with deleted storage volume -> %+v", err)
		assert.Equal(t, "", testSVol.Name, fmt.Sprintf("Problem getting storage volume name, %+v", testSVol))
	} else {
		_, c = getTestDriverU("test_storage_volume")
		err := c.DeleteStorageVolume("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testSVol))
	}
}
