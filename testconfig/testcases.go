package testconfig

import (
	"os"

	"github.com/docker/machine/libmachine/log"
)

// test case objects
// { "name" : "PROTEST",
//   "cases": [
//      {
//         "name": "TestName1",   // function name of the test case
//         "enabled": true,        // flag to determine if we do the test case or skip
//         "TestData" : map[string][interface{ }] // list of test data by name
//         "ExpectsData" : map[string][interface{}] // expects results
//      },
//      {
//        ....
//      },
// ]
//}
type TestCases struct {
	Name        string                 `json:"name,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	TestData    map[string]interface{} `json:"test_data,omitempty"`    // map[string]interface{}{"test_data": []interface{}}
	ExpectsData map[string]interface{} `json:"expects_data,omitempty"` // map[string]interface{}{"expects_data": []interface{}}
}

// compare test data to a int num
func (tc *TestConfig) EqualFaceI(f interface{}, i int) bool {
	// convert the int to a float
	return float64(f.(float64)) == float64(i)
}

// Compare interface to string
// doesn't seem to really be needed....
func (tc *TestConfig) EqualFaceS(is interface{}, s string) bool {
	// convert the int to a float
	return string(is.(string)) == string(s)
}

// Get a test case by name
func (tc *TestConfig) GetTestCases(tc_name string) (t TestCases) {
	for _, t := range tc.Cases {
		if t.Name == tc_name {
			return t
		}
	}
	return t
}

// is test case enabled? defaults are always true
func (tc *TestConfig) IsTestEnabled(tc_name string) bool {
	if t := tc.GetTestCases(tc_name); t.Name == tc_name {
		return t.Enabled
	}
	if d := tc.GetTestCases("default"); d.Name != "" {
		return d.Enabled
	}
	log.Infof("tc no name -> %+v", tc.GetTestCases("foo"))
	log.Warn("Test config is using default true for enablement.")
	return true
}

// get expects data, returns interface because the data type is unknown
// use tc Equal arguments for type comparison, conversions
// or assert the type
func (tc *TestConfig) GetExpectsData(tc_name string, k string) interface{} {
	if t := tc.GetTestCases(tc_name); t.ExpectsData[k] != nil {
		log.Debugf("GetExpectsData(%s, %s) found -> %s", tc_name, k, t.ExpectsData[k])
		return t.ExpectsData[k]
	}
	if d := tc.GetTestCases("default"); d.ExpectsData[k] != nil {
		log.Debugf("GetExpectsData(%s, %s) found -> %s = %s", tc_name, k, "default", d.ExpectsData[k])
		return d.ExpectsData[k]
	}
	log.Errorf("Test config expects data not found %s for test case %s\n", k, tc_name)
	os.Exit(1)
	return nil
}

// get test data
func (tc *TestConfig) GetTestData(tc_name string, k string) interface{} {
	if t := tc.GetTestCases(tc_name); t.TestData[k] != nil {
		log.Debugf("GetTestData(%s, %s) found -> %s", tc_name, k, t.TestData[k])
		return t.TestData[k]
	}
	if d := tc.GetTestCases("default"); d.TestData[k] != nil {
		log.Debugf("GetTestData(%s, %s) found -> %s = %s", tc_name, k, "default", d.TestData[k])
		return d.TestData[k]
	}
	log.Errorf("Test config test data not found %s for test case %s\n", k, tc_name)
	log.Debugf("GetTestData(%s, %s) failed", tc_name, k)
	os.Exit(1)
	return nil
}
