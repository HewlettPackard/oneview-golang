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
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateGoldenImage(t *testing.T) {
	var (
		d        *I3STest
		c        *i3s.I3SClient
		testName string
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_golden_image")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test golden image already exist
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testGoldenImage, err := c.GetGoldenImageByName(testName)
		assert.NoError(t, err, "CreateGoldenImage get the GoldenImage error -> %s", err)

		if testGoldenImage.URI.IsNil() {
			testGoldenImage = i3s.GoldenImage{
				Name: testName,
				Type: d.Tc.GetTestData(d.Env, "Type").(string),
			}
			err := c.CreateGoldenImage(testGoldenImage)
			assert.NoError(t, err, "CreateGoldenImage error -> %s", err)

			err = c.CreateGoldenImage(testGoldenImage)
			assert.Error(t, err, "CreateGoldenImage should error because the GoldenImage already exists, err-> %s", err)

		} else {
			log.Warnf("The goldenImage already exist, so skipping CreateGoldenImage test for %s", testName)
		}

		// reload the test profile that we just created
		testGoldenImage, err = c.GetGoldenImageByName(testName)
		assert.NoError(t, err, "GetGoldenImage error -> %s", err)
	}
}

func TestGetGoldenImageByName(t *testing.T) {
	var (
		d        *I3STest
		c        *i3s.I3SClient
		testName string
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_golden_image")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testGoldenImage, err := c.GetGoldenImageByName(testName)
		assert.NoError(t, err, "GetGoldenImageByName thew an error -> %s", err)
		assert.Equal(t, testName, testGoldenImage.Name)

		testGoldenImage, err = c.GetGoldenImageByName("bad")
		assert.NoError(t, err, "GetGoldenImageByName with fake name -> %s", err)
		assert.Equal(t, "", testGoldenImage.Name)

	} else {
		d, c = getTestDriverU("test_golden_image")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetGoldenImageByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetGoldenImages(t *testing.T) {
	var (
		c *i3s.I3SClient
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_golden_image")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		goldenImages, err := c.GetGoldenImages("", "")
		assert.NoError(t, err, "GetGoldenImages threw error -> %s, %+v\n", err, goldenImages)

		goldenImages, err = c.GetGoldenImages("", "name:asc")
		assert.NoError(t, err, "GetGoldenImages name:asc error -> %s, %+v\n", err, goldenImages)

	} else {
		_, c = getTestDriverU("test_golden_image")
		data, err := c.GetGoldenImages("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeleteGoldenImageNotFound(t *testing.T) {
	var (
		c               *i3s.I3SClient
		testName        = "fake"
		testGoldenImage i3s.GoldenImage
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_golden_image")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteGoldenImage(testName)
		assert.NoError(t, err, "DeleteGoldenImage err-> %s", err)

		testGoldenImage, err = c.GetGoldenImageByName(testName)
		assert.NoError(t, err, "GetGoldenImageByName with deleted golden image -> %+v", err)
		assert.Equal(t, "", testGoldenImage.Name, fmt.Sprintf("Problem getting golden image name, %+v", testGoldenImage))
	} else {
		_, c = getTestDriverU("test_golden_image")
		err := c.DeleteGoldenImage(testName)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testGoldenImage))
	}
}

func TestDeleteGoldenImage(t *testing.T) {
	var (
		d               *I3STest
		c               *i3s.I3SClient
		testName        string
		testGoldenImage i3s.GoldenImage
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_golden_image")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeleteGoldenImage(testName)
		assert.NoError(t, err, "DeleteGoldenImage err-> %s", err)

		testGoldenImage, err = c.GetGoldenImageByName(testName)
		assert.NoError(t, err, "GetGoldenImageByName with deleted golden image-> %+v", err)
		assert.Equal(t, "", testGoldenImage.Name, fmt.Sprintf("Problem getting golden image name, %+v", testGoldenImage))
	} else {
		_, c = getTestDriverU("test_golden_image")
		err := c.DeleteGoldenImage("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testGoldenImage))
	}

}
