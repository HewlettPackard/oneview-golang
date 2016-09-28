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

func TestCreateOSBuildPlan(t *testing.T) {
	var (
		d        *I3STest
		c        *I3SClient
		testName string
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_os_build_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test os build plan already exist
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testOSBuildPlan, err := c.GetOSBuildPlanByName(testName)
		assert.NoError(t, err, "CreateOSBuildPlan get the OS BuildPlan error -> %s", err)

		if testOSBuildPlan.URI.IsNil() {
			testOSBuildPlan = OSBuildPlan{
				Name: testName,
				Type: d.Tc.GetTestData(d.Env, "Type").(string),
			}
			err := c.CreateOSBuildPlan(testOSBuildPlan)
			assert.NoError(t, err, "CreateOSBuildPlan error -> %s", err)

			err = c.CreateOSBuildPlan(testOSBuildPlan)
			assert.Error(t, err, "CreateOSBuildPlan should error because the OSBuildPlan already exists, err-> %s", err)

		} else {
			log.Warnf("The osBuildPlan already exist, so skipping CreateOSBuildPlan test for %s", testName)
		}

		// reload the os build plan that we just created
		testOSBuildPlan, err = c.GetOSBuildPlanByName(testName)
		assert.NoError(t, err, "GetOSBuildPlan error -> %s", err)
	}
}

func TestGetOSBuildPlanByName(t *testing.T) {
	var (
		d        *I3STest
		c        *I3SClient
		testName string
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_os_build_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testOSBuildPlan, err := c.GetOSBuildPlanByName(testName)
		assert.NoError(t, err, "GetOSBuildPlanByName thew an error -> %s", err)
		assert.Equal(t, testName, testOSBuildPlan.Name)

		testOSBuildPlan, err = c.GetOSBuildPlanByName("bad")
		assert.NoError(t, err, "GetOSBuildPlanByName with fake name -> %s", err)
		assert.Equal(t, "", testOSBuildPlan.Name)

	} else {
		d, c = getTestDriverU("test_os_build_plan")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetOSBuildPlanByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetOSBuildPlans(t *testing.T) {
	var (
		c *I3SClient
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_os_build_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		osBuildPlans, err := c.GetOSBuildPlans("", "")
		assert.NoError(t, err, "GetOSBuildPlans threw error -> %s, %+v\n", err, osBuildPlans)

		osBuildPlans, err = c.GetOSBuildPlans("", "name:asc")
		assert.NoError(t, err, "GetOSBuildPlans name:asc error -> %s, %+v\n", err, osBuildPlans)

	} else {
		_, c = getTestDriverU("test_os_build_plan")
		data, err := c.GetOSBuildPlans("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeleteOSBuildPlanNotFound(t *testing.T) {
	var (
		c               *I3SClient
		testName        = "fake"
		testOSBuildPlan OSBuildPlan
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_os_build_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteOSBuildPlan(testName)
		assert.NoError(t, err, "DeleteOSBuildPlan err-> %s", err)

		testOSBuildPlan, err = c.GetOSBuildPlanByName(testName)
		assert.NoError(t, err, "GetOSBuildPlanByName with deleted os build plan -> %+v", err)
		assert.Equal(t, "", testOSBuildPlan.Name, fmt.Sprintf("Problem getting os build plan by name, %+v", testOSBuildPlan))
	} else {
		_, c = getTestDriverU("test_os_build_plan")
		err := c.DeleteOSBuildPlan(testName)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testOSBuildPlan))
	}
}

func TestDeleteOSBuildPlan(t *testing.T) {
	var (
		d               *I3STest
		c               *I3SClient
		testName        string
		testOSBuildPlan OSBuildPlan
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_os_build_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeleteOSBuildPlan(testName)
		assert.NoError(t, err, "DeleteOSBuildPlan err-> %s", err)

		testOSBuildPlan, err = c.GetOSBuildPlanByName(testName)
		assert.NoError(t, err, "GetOSBuildPlanByName with deleted os build plan-> %+v", err)
		assert.Equal(t, "", testOSBuildPlan.Name, fmt.Sprintf("Problem getting os build plan name, %+v", testOSBuildPlan))
	} else {
		_, c = getTestDriverU("test_os_build_plan")
		err := c.DeleteOSBuildPlan("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testOSBuildPlan))
	}

}
