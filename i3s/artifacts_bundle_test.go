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
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateArtifactsBundle(t *testing.T) {
	var (
		d        *I3STest
		c        *I3SClient
		testName string
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_artifacts_bundle")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test artifactsBundle already exist
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testArtifactsBundle, err := c.GetArtifactsBundleByName(testName)
		assert.NoError(t, err, "CreateArtifactsBundle get the ArtifactsBundle error -> %s", err)

		if testArtifactsBundle.URI.IsNil() {
			testArtifactsBundle = InputArtifactsBundle{
				Name: testName,
				Type: d.Tc.GetTestData(d.Env, "Type").(string),
			}
			err := c.CreateArtifactsBundle(testArtifactsBundle)
			assert.NoError(t, err, "CreateArtifactsBundle error -> %s", err)

			err = c.CreateArtifactsBundle(testArtifactsBundle)
			assert.Error(t, err, "CreateArtifactsBundle should error because the ArtifactsBundle already exists, err-> %s", err)

		} else {
			log.Warnf("The ArtifactsBundle already exist, so skipping CreateArtifactsBundle test for %s", testName)
		}

		// reload the test artifact bundle that we just created
		testArtifactsBundle, err = c.GetArtifactsBundleByName(testName)
		assert.NoError(t, err, "GetArtifactsBundle error -> %s", err)
	}
}

func TestGetArtifactsBundleByName(t *testing.T) {
	var (
		d        *I3STest
		c        *I3SClient
		testName string
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_artifacts_bundle")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testArtifactsBundle, err := c.GetArtifactsBundleByName(testName)
		assert.NoError(t, err, "GetArtifactsBundleByName thew an error -> %s", err)
		assert.Equal(t, testName, testArtifactsBundle.Name)

		testArtifactsBundle, err = c.GetArtifactsBundleByName("bad")
		assert.NoError(t, err, "GetArtifactsBundleByName with fake name -> %s", err)
		assert.Equal(t, "", testArtifactsBundle.Name)

	} else {
		d, c = getTestDriverU("test_artifacts_bundle")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetArtifactsBundleByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetArtifactsBundles(t *testing.T) {
	var (
		c *I3SClient
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_artifacts_bundle")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		artifactsBundles, err := c.GetArtifactsBundles("", "")
		assert.NoError(t, err, "GetArtifactsBundles threw error -> %s, %+v\n", err, artifactsBundles)

		artifactsBundles, err = c.GetArtifactsBundles("", "name:asc")
		assert.NoError(t, err, "GetArtifactsBundles name:asc error -> %s, %+v\n", err, artifactsBundles)

	} else {
		_, c = getTestDriverU("test_artifacts_bundle")
		data, err := c.GetArtifactsBundles("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeleteArtifactsBundleNotFound(t *testing.T) {
	var (
		c                   *I3SClient
		testName            = "fake"
		testArtifactsBundle ArtifactsBundle
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_artifacts_bundle")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteArtifactsBundle(testName)
		assert.NoError(t, err, "DeleteArtifactsBundle err-> %s", err)

		testArtifactsBundle, err = c.GetArtifactsBundleByName(testName)
		assert.NoError(t, err, "GetArtifactsBundleByName with deleted artifacts bundle -> %+v", err)
		assert.Equal(t, "", testArtifactsBundle.Name, fmt.Sprintf("Problem getting artifacts bundle name, %+v", testArtifactsBundle))
	} else {
		_, c = getTestDriverU("test_artifacts_bundle")
		err := c.DeleteArtifactsBundle(testName)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testArtifactsBundle))
	}
}

func TestDeleteArtifactsBundle(t *testing.T) {
	var (
		d                   *I3STest
		c                   *I3SClient
		testName            string
		testArtifactsBundle ArtifactsBundle
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_deployment_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeleteArtifactsBundle(testName)
		assert.NoError(t, err, "DeleteArtifactsBundle err-> %s", err)

		testArtifactsBundle, err = c.GetArtifactsBundleByName(testName)
		assert.NoError(t, err, "GetArtifactsBundleByName with deleted artifactsBundle-> %+v", err)
		assert.Equal(t, "", testArtifactsBundle.Name, fmt.Sprintf("Problem getting artifactsBundle name, %+v", testArtifactsBundle))
	} else {
		_, c = getTestDriverU("test_artifacts_bundle")
		err := c.DeleteArtifactsBundle("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testArtifactsBundle))
	}

}
