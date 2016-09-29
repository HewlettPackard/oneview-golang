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

func TestCreatePlanScript(t *testing.T) {
	var (
		d        *I3STest
		c        *I3SClient
		testName string
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_plan_script")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		// find out if the test plan script already exist
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testPlanScript, err := c.GetPlanScriptByName(testName)
		assert.NoError(t, err, "CreatePlanScript get the PlanScript error -> %s", err)

		if testPlanScript.URI.IsNil() {
			testPlanScript = PlanScript{
				Name: testName,
				Type: d.Tc.GetTestData(d.Env, "Type").(string),
			}
			err := c.CreatePlanScript(testPlanScript)
			assert.NoError(t, err, "CreatePlanScript error -> %s", err)

			err = c.CreatePlanScript(testPlanScript)
			assert.Error(t, err, "CreatePlanScript should error because the PlanScript already exists, err-> %s", err)

		} else {
			log.Warnf("The planScript already exist, so skipping CreatePlanScript test for %s", testName)
		}

		// reload the test profile that we just created
		testPlanScript, err = c.GetPlanScriptByName(testName)
		assert.NoError(t, err, "GetPlanScript error -> %s", err)
	}
}

func TestGetPlanScriptByName(t *testing.T) {
	var (
		d        *I3STest
		c        *I3SClient
		testName string
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_plan_script")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		testPlanScript, err := c.GetPlanScriptByName(testName)
		assert.NoError(t, err, "GetPlanScriptByName thew an error -> %s", err)
		assert.Equal(t, testName, testPlanScript.Name)

		testPlanScript, err = c.GetPlanScriptByName("bad")
		assert.NoError(t, err, "GetPlanScriptByName with fake name -> %s", err)
		assert.Equal(t, "", testPlanScript.Name)

	} else {
		d, c = getTestDriverU("test_plan_script")
		testName = d.Tc.GetTestData(d.Env, "Name").(string)
		data, err := c.GetPlanScriptByName(testName)
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestGetPlanScripts(t *testing.T) {
	var (
		c *I3SClient
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_plan_script")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		planScripts, err := c.GetPlanScripts("", "")
		assert.NoError(t, err, "GetPlanScript threw error -> %s, %+v\n", err, planScripts)

		planScripts, err = c.GetPlanScripts("", "name:asc")
		assert.NoError(t, err, "GetPlanScripts name:asc error -> %s, %+v\n", err, planScripts)

	} else {
		_, c = getTestDriverU("test_plan_script")
		data, err := c.GetPlanScripts("", "")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
	}
}

func TestDeletePlanScriptNotFound(t *testing.T) {
	var (
		c              *I3SClient
		testName       = "fake"
		testPlanScript PlanScript
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_plan_script")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}

		err := c.DeletePlanScript(testName)
		assert.NoError(t, err, "DeletePlanScript err-> %s", err)

		testPlanScript, err = c.GetPlanScriptByName(testName)
		assert.NoError(t, err, "GetPlanScriptByName with deleted plan script -> %+v", err)
		assert.Equal(t, "", testPlanScript.Name, fmt.Sprintf("Problem getting plan script name, %+v", testPlanScript))
	} else {
		_, c = getTestDriverU("test_plan_script")
		err := c.DeletePlanScript(testName)
		assert.Error(t, err, fmt.Sprintf("All ok, no error, caught as expected: %s,%+v\n", err, testPlanScript))
	}
}

func TestDeletePlanScript(t *testing.T) {
	var (
		d              *I3STest
		c              *I3SClient
		testName       string
		testPlanScript PlanScript
	)
	if os.Getenv("I3S_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_plan_script")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		testName = d.Tc.GetTestData(d.Env, "Name").(string)

		err := c.DeletePlanScript(testName)
		assert.NoError(t, err, "DeletePlanScript err-> %s", err)

		testPlanScript, err = c.GetPlanScriptByName(testName)
		assert.NoError(t, err, "GetPlanScriptByName with deleted plan script-> %+v", err)
		assert.Equal(t, "", testPlanScript.Name, fmt.Sprintf("Problem getting plan script name, %+v", testPlanScript))
	} else {
		_, c = getTestDriverU("test_plan_script")
		err := c.DeletePlanScript("footest")
		assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, testPlanScript))
	}

}
