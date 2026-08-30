package ov

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testFC struct {
	mock.Mock
	OVClient
}

func (f *testFC) GetFCNetworkByName(name string) (FCNetwork, error) {
	fcNets, err := f.GetFCNetworks(fmt.Sprintf("name matches '%s'", name), "name:asc", "", "")
	if fcNets.Total > 0 {
		return fcNets.Members[0], err
	}

	return FCNetwork{}, err

}

func (f *testFC) GetFCNetworks(filter string, sort string, start string, count string) (FCNetworkList, error) {
	fmt.Println("holla")
	_f_ret := f.Called("", "name:asc", "", "")
	if _f_ret.Get(1) != nil {
		return _f_ret.Get(0).(FCNetworkList), _f_ret.Error(1)
	}
	return _f_ret.Get(0).(FCNetworkList), _f_ret.Error(1)

}

func TestFcNetworkGetByName(t *testing.T) {
	var fc_result = FCNetwork{
		AutoLoginRedistribution: false,
		Description:             "Test FC Network",
		LinkStabilityTime:       30,
		FabricType:              "FabricAttach",
		Name:                    "testName",
		Type:                    "fc-networkV4", //The Type value is for API>500.
	}

	var fc_result_array = []FCNetwork{fc_result}

	var fc_result_list = FCNetworkList{
		Total:       1,
		Count:       1,
		Start:       0,
		PrevPageURI: "null",
		NextPageURI: "null",
		URI:         "/rest/server-profiles?",
		Members:     fc_result_array,
	}

	fcObj := &testFC{}

	fcObj.On("GetFCNetworks", "", "name:asc", "", "").Return(fc_result_list, nil)

	fc_found, _ := fcObj.GetFCNetworkByName("testName")
	assert.Equal(t, fc_found.Name, "testName", "both should be equal")

	fcObj.AssertCalled(t, "GetFCNetworks", "", "name:asc", "", "")
	fcObj.AssertExpectations(t)

}
