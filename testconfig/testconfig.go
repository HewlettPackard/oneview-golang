package testconfig

// A test configuration management package for loading
// and executing test data and expects data that can be
// configured for various test target environments.
import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/docker/machine/libmachine/log"
)

//
// test case objects
//  see testcases.go for test case methods
// { "name" : "PROTEST",
//   "cases": [
//      {
//        ...
//      },
// ]
//}
type TestConfig struct {
	Cases []TestCases `json:"cases,omitempty"` // "cases":[]
	Name  string      `json:"name,omitempty"`  // "name": "PROTEST",
}

// setup Testing
func (tc *TestConfig) NewTestConfig() *TestConfig {
	return &TestConfig{Name: "",
		Cases: []TestCases{}}
}

// UnMarshall json to data
func (tc *TestConfig) UnMarshallTestingConfig(json_data []byte) {
	tc.Cases = []TestCases{}
	if err := json.Unmarshal(json_data, &tc); err != nil {
		log.Errorf("Error with un-marshalling test config data: %s", err)
		os.Exit(1)
	}
}

// get config for testing
// Examples
// cv := tc.GetExpectsData("TestGetAPIVersion", "CurrentVersion")
// log.Infof("tc test_data -> %s\n", tc.EqualFaceI(cv, 120))
// log.Infof("2 tc compare s -> %s\n", tc.GetExpectsData("TestGetAPIVersion", "FakeData") == "foo")
// log.Infof("tc compare s -> %s\n", tc.EqualFaceS(tc.GetExpectsData("TestGetAPIVersion", "FakeData"), "foo"))
// log.Infof("get no test data -> %+v \n", tc.GetTestData("TestGetAPIVersion", "Surprise"))
// log.Infof("is test enabled -> %+v \n", tc.IsTestEnabled("TestGetAPIVersion"))
//
func (tc *TestConfig) GetTestingConfiguration(config_name string) {
	var (
		package_root  string
		Pkg           PackageInfo
		test_data_dir string
	)
	if os.Getenv("ONEVIEW_DEBUG") == "true" {
		log.SetDebug(true)
	}
	package_root = os.Getenv("TESTCONFIG_PACKAGE_ROOT_PATH")
	if found, package_full_dir := Pkg.GetPackageRootDir(package_root); found == true {
		test_data_dir = Pkg.JoinPath([]string{package_full_dir,
			os.Getenv("TESTCONFIG_JSON_DATA_DIR")})
		if found, _ := Pkg.DirExists(test_data_dir); found == true {
			files, _ := ioutil.ReadDir(test_data_dir)
			for _, f := range files {
				log.Debugf("working on f -> %+v", f)
				data, err := ioutil.ReadFile(Pkg.JoinPath([]string{test_data_dir, f.Name()}))
				if err != nil {
					log.Errorf("File error : %v\n", err)
					os.Exit(1)
				}
				tc.UnMarshallTestingConfig(data)
				if tc.Name == config_name {
					log.Debugf("Found test data, %s \n", tc.Name)
					log.Debugf("tc -> %+v \n", tc)
					return
				}
			}
		} else {
			log.Errorf("No configuration found in %s, %s \n", test_data_dir, config_name)
		}
	}
	log.Errorf("No configuration found in %s, %s \n", package_root, config_name)
	os.Exit(1)
	return
}
