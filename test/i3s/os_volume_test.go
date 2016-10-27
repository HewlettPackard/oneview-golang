/*
(c) Copyright [2015] Hewlett Packard Enterprise Development LP

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package i3s

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/i3s"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetOSVolumeByName(t *testing.T) {
	var (
		d        *I3STest
		c        *i3s.I3SClient
		testName string
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_os_volume")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testOSVolume, err := c.GetOSVolumeByName(testName)
		assert.NoError(t, err, "GetOSVolumeByName thew an error -> %s", err)
		assert.Equal(t, testName, testOSVolume.Name)

		testOSVolume, err = c.GetOSVolumeByName("bad")
		assert.NoError(t, err, "GetOSVolumeByName with fake name -> %s", err)
		assert.Equal(t, "", testOSVolume.Name)

	} else {
		d, c = getTestDriverU("test_os_volume")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetOSVolumeByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetOSVolumes(t *testing.T) {
	var (
		c *i3s.I3SClient
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_os_volume")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		osVolumes, err := c.GetOSVolumes("", "")
		assert.NoError(t, err, "GetOSVolumes threw error -> %s, %+v\n", err, osVolumes)

		osVolumes, err = c.GetOSVolumes("", "name:asc")
		assert.NoError(t, err, "GetOSVolumes name:asc error -> %s, %+v\n", err, osVolumes)

	} else {
		_, c = getTestDriverU("test_os_volume")
		data, err := c.GetOSVolumes("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}
