package icsp

import (
	"encoding/json"
	"testing"

	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

// Test getting a valueitems from customattribute
func TestGetValueItems(t *testing.T) {
	var (
		d *ICSPTest
		s Server
	)
	d, _ = getTestDriverU()
	jsonServerData := d.Tc.GetTestData(d.Env, "ServerJSONString").(string)
	log.Debugf("jsonServerData => %s", jsonServerData)
	err := json.Unmarshal([]byte(jsonServerData), &s)
	assert.NoError(t, err, "Unmarshal Server threw error -> %s, %+v\n", err, jsonServerData)

	testKey1 := d.Tc.GetTestData(d.Env, "KeyTest1").(string)
	_, vi := s.GetValueItems(testKey1)
	assert.Equal(t, 1, len(vi), "Should find 1 valueitem")
}

// TestGetValueItem gets a valueitem with scope
func TestGetValueItem(t *testing.T) {
	var (
		d *ICSPTest
		s Server
	)
	d, _ = getTestDriverU()
	jsonServerData := d.Tc.GetTestData(d.Env, "ServerJSONString").(string)
	log.Debugf("jsonServerData => %s", jsonServerData)
	err := json.Unmarshal([]byte(jsonServerData), &s)
	assert.NoError(t, err, "Unmarshal Server threw error -> %s, %+v\n", err, jsonServerData)

	testKey1 := d.Tc.GetTestData(d.Env, "KeyTest1").(string)
	testScope1 := d.Tc.GetTestData(d.Env, "ScopeTest1").(string)
	expectsValue1 := d.Tc.GetExpectsData(d.Env, "ValueTest1").(string)

	_, v := s.GetValueItem(testKey1, testScope1)
	assert.Equal(t, expectsValue1, v.Value, "Should find ValueItem")
}

// TestSetValueItems
func TestSetValueItems(t *testing.T) {
	var (
		d *ICSPTest
		s Server
	)
	d, _ = getTestDriverU()
	jsonServerData := d.Tc.GetTestData(d.Env, "ServerJSONString").(string)
	log.Debugf("jsonServerData => %s", jsonServerData)
	err := json.Unmarshal([]byte(jsonServerData), &s)
	assert.NoError(t, err, "Unmarshal Server threw error -> %s, %+v\n", err, jsonServerData)

	// Try setting a ValueItem that doesn't exist
	s.SetValueItems("foo", ValueItem{Scope: "server", Value: "bar"})
	_, v := s.GetValueItem("foo", "server")
	assert.Equal(t, "bar", v.Value, "Should find bar from key foo")

}

// TestSetCustomAttribute
func TestSetCustomAttribute(t *testing.T) {
	var (
		d *ICSPTest
		s Server
	)
	// unit test case
	d, _ = getTestDriverU()
	jsonServerData := d.Tc.GetTestData(d.Env, "ServerJSONString").(string)
	log.Debugf("jsonServerData => %s", jsonServerData)
	err := json.Unmarshal([]byte(jsonServerData), &s)
	assert.NoError(t, err, "Unmarshal Server threw error -> %s, %+v\n", err, jsonServerData)

	testKey1 := d.Tc.GetTestData(d.Env, "KeyTest1").(string)
	testScope1 := d.Tc.GetTestData(d.Env, "ScopeTest1").(string)
	expectsValue1 := d.Tc.GetExpectsData(d.Env, "ValueTest1").(string)

	// complete test case 1, simple read of custom attributes
	_, testValue1 := s.GetValueItem(testKey1, testScope1)
	assert.Equal(t, expectsValue1, testValue1.Value, "Should return testcase 1 simple read")
	// TODO: test case 2 , setting a value on existing attribute
	testKey2 := d.Tc.GetTestData(d.Env, "KeyTest2Existing").(string)
	testScope2 := d.Tc.GetTestData(d.Env, "ScopeTest2Existing").(string)
	testValueNew2 := d.Tc.GetTestData(d.Env, "ValueTest2ExistingNew").(string)
	expectsValueOld2 := d.Tc.GetExpectsData(d.Env, "ValueTestOld2").(string)
	expectsValueNew2 := d.Tc.GetExpectsData(d.Env, "ValueTestNew2").(string)

	_, testValue2 := s.GetValueItem(testKey2, testScope2)
	assert.Equal(t, expectsValueOld2, testValue2.Value, "Should return testcase 2 old setting existing attribute")

	s.SetCustomAttribute(testKey2, testScope2, testValueNew2)
	_, testValue2 = s.GetValueItem(testKey2, testScope2)
	assert.Equal(t, expectsValueNew2, testValue2.Value, "Should return testcase 2 new setting existing attribute")

	// TODO: test case 3, appending a new attribute
	testKey3 := d.Tc.GetTestData(d.Env, "KeyTest3New").(string)
	testScope3 := d.Tc.GetTestData(d.Env, "ScopeTest3New").(string)
	testValue3New := d.Tc.GetTestData(d.Env, "ValueTest3New").(string)
	expectsValue3 := d.Tc.GetExpectsData(d.Env, "ValueTest3").(string)

	i, _ := s.GetValueItem(testKey3, testScope3)
	assert.Equal(t, -1, i, "Should not find test case 3 value")

	s.SetCustomAttribute(testKey3, testScope3, testValue3New)
	_, testValue3 := s.GetValueItem(testKey3, testScope3)
	assert.Equal(t, expectsValue3, testValue3.Value, "Should return testcase 3 result")
}
