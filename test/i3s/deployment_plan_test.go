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

func TestCreateDeploymentPlan(t *testing.T) {
	var (
		d        *I3STest
		c        *i3s.I3SClient
		testName string
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_deployment_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test deployment plan already exist
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testDeploymentPlan, err := c.GetDeploymentPlanByName(testName)
		assert.NoError(t, err, "CreateDeploymentPlan get the DeploymentPlan error -> %s", err)

		if testDeploymentPlan.URI.IsNil() {
			testDeploymentPlan = i3s.DeploymentPlan{
				Name: testName,
				Type: d.Tc.GetTestData(d.Env, "Type").(string),
			}
			err := c.CreateDeploymentPlan(testDeploymentPlan)
			assert.NoError(t, err, "CreateDeploymentPlan error -> %s", err)

			err = c.CreateDeploymentPlan(testDeploymentPlan)
			assert.Error(t, err, "CreateDeploymentPlan should error because the DeploymentPlan already exists, err-> %s", err)

		} else {
			log.Warnf("The deploymentPlan already exist, so skipping CreateDeploymentPlan test for %s", testName)
		}

		// reload the test profile that we just created
		testDeploymentPlan, err = c.GetDeploymentPlanByName(testName)
		assert.NoError(t, err, "GetDeplymentPlan error -> %s", err)
	}
}

func TestGetDeploymentPlanByName(t *testing.T) {
	var (
		d        *I3STest
		c        *i3s.I3SClient
		testName string
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_deployment_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testDeploymentPlan, err := c.GetDeploymentPlanByName(testName)
		assert.NoError(t, err, "GetDeploymentPlanByName thew an error -> %s", err)
		assert.Equal(t, testName, testDeploymentPlan.Name)

		testDeploymentPlan, err = c.GetDeploymentPlanByName("bad")
		assert.NoError(t, err, "GetDeploymentPlanByName with fake name -> %s", err)
		assert.Equal(t, "", testDeploymentPlan.Name)

	} else {
		d, c = getTestDriverU("test_deployment_plan")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetDeploymentPlanByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetDeploymentPlans(t *testing.T) {
	var (
		c *i3s.I3SClient
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_deployment_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		deploymentPlans, err := c.GetDeploymentPlans("", "", "", "", "")
		assert.NoError(t, err, "GetDeploymentPlans threw error -> %s, %+v\n", err, deploymentPlans)

		deploymentPlans, err = c.GetDeploymentPlans("", "", "", "name:asc", "")
		assert.NoError(t, err, "GetDeploymentPlans name:asc error -> %s, %+v\n", err, deploymentPlans)

	} else {
		_, c = getTestDriverU("test_deployment_plan")
		data, err := c.GetDeploymentPlans("", "", "", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeleteDeploymentPlanNotFound(t *testing.T) {
	var (
		c                  *i3s.I3SClient
		testName           = "fake"
		testDeploymentPlan i3s.DeploymentPlan
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_deployment_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeleteDeploymentPlan(testName)
		assert.NoError(t, err, "DeleteDeploymentPlan err-> %s", err)

		testDeploymentPlan, err = c.GetDeploymentPlanByName(testName)
		assert.NoError(t, err, "GetDeploymentPlanByName with deleted deployment plan -> %+v", err)
		assert.Equal(t, "", testDeploymentPlan.Name, fmt.Sprintf("Problem getting deployment plan name, %+v", testDeploymentPlan))
	} else {
		_, c = getTestDriverU("test_deployment_plan")
		err := c.DeleteDeploymentPlan(testName)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testDeploymentPlan))
	}
}

func TestDeleteDeploymentPlan(t *testing.T) {
	var (
		d                  *I3STest
		c                  *i3s.I3SClient
		testName           string
		testDeploymentPlan i3s.DeploymentPlan
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_deployment_plan")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeleteDeploymentPlan(testName)
		assert.NoError(t, err, "DeleteDeploymentPlan err-> %s", err)

		testDeploymentPlan, err = c.GetDeploymentPlanByName(testName)
		assert.NoError(t, err, "GetDeploymentPlanByName with deleted deployment plan-> %+v", err)
		assert.Equal(t, "", testDeploymentPlan.Name, fmt.Sprintf("Problem getting deployment plan name, %+v", testDeploymentPlan))
	} else {
		_, c = getTestDriverU("test_deployment_plan")
		err := c.DeleteDeploymentPlan("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testDeploymentPlan))
	}

}
