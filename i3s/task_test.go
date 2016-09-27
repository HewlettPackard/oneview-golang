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
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// test unmarshalling a json payload that has progress
func TestProgressTaskJsonUnmarshal(t *testing.T) {
	var (
		task           *Task
		test_json_data = `{"type":"TaskResourceV2","data":null,"category":"tasks","taskType":"Background","stateReason":"Starting","associatedResource":{"resourceri":null,"associationType":"MANAGED_BY","resourceName":null,"resourceCategory":"deployment-plans"},"hidden":false,"percentComplete":0,"associatedTaskUri":null,"completedSteps":0,"computedPercentComplete":0,"expectedDuration":0,"parentTaskUri":null,"progressUpdates":[],"taskErrors":[],"taskOutput":[],"taskState":"Starting","taskStatus":"Applying","totalSteps":0,"userInitiated":false,"name":"Create","owner":"wenlock","created":"2015-09-10T20:12:50.006Z","eTag":"2","modified":"2015-09-10T20:12:50.083Z","uri":"/rest/tasks/BA3C8C84-4122-4A40-A3C3-D86CE9EDB5ED"}`
	)
	err := json.Unmarshal([]byte(test_json_data), &task)
	assert.NoError(t, err, fmt.Sprintf("Failed to unmarshal task object: %s, %+v\n", err, task))
}

// test unmarshalling a json payload that has a failure
func TestFailTaskJsonUnmarshal(t *testing.T) {
	var (
		task           *Task
		test_json_data = `{"type":"TaskResourceV2","data":null,"category":"tasks","taskType":"Background","stateReason":"ValidationError","associatedResource":{"resourceUri":null,"associationType":"MANAGED_BY","resourceName":null,"resourceCategory":"deployment-plans"},"hidden":true,"percentComplete":100,"associatedTaskUri":null,"completedSteps":0,"computedPercentComplete":100,"expectedDuration":0,"parentTaskUri":null,"progressUpdates":[],"taskErrors":[{"data":{},"errorCode":"MacTypeDiffGlobalMacType","details":"","nestedErrors":[],"errorSource":null,"recommendedActions":["Verify parameters and try again."],"message":"When macType is not user defined, mac type should be same as the global Mac assignment Virtual."}],"taskOutput":[],"taskState":"Error","taskStatus":"Unable to create deployment plan: deployment_plan_test01","totalSteps":0,"userInitiated":false,"name":"Create","owner":"wenlock","created":"2015-09-10T20:12:50.006Z","eTag":"3","modified":"2015-09-10T20:12:51.371Z","uri":"/rest/tasks/BA3C8C84-4122-4A40-A3C3-D86CE9EDB5ED"}`
	)
	err := json.Unmarshal([]byte(test_json_data), &task)
	assert.NoError(t, err, fmt.Sprintf("Failed to unmarshal task object: %s, %+v\n", err, task))
}

// test unmarshalling progress updates
// "progressUpdates":[{
//     "timestamp":"2015-09-10T22:50:06.016Z",
//     "statusUpdate":"",
//     "id":12564
//     },{
//     "timestamp":"2015-09-10T22:50:14.250Z",
//     "statusUpdate":"",
//     "id":12566
// }]

func TestProgressUpdatesTaskJsonUnmarshal(t *testing.T) {
	var (
		task           *Task
		test_json_data = `{"type":"TaskResourceV2","data":null,"category":"tasks","taskType":"User","stateReason":"Running","associatedResource":{"resourceUri":"/rest/deployment-plans/9e2e3adf-790b-4272-b8a7-e693cf39ade4","associationType":"MANAGED_BY","resourceName":"docker_machine_test01","resourceCategory":"deployment-plans"},"hidden":false,"percentComplete":50,"associatedTaskUri":null,"completedSteps":0,"computedPercentComplete":50,"expectedDuration":25,"parentTaskUri":null,"progressUpdates":[{"timestamp":"2015-09-10T22:50:06.016Z","statusUpdate":"","id":12564},{"timestamp":"2015-09-10T22:50:14.250Z","statusUpdate":"","id":12566}],"taskErrors":[],"taskOutput":[],"taskState":"Running","taskStatus":"Create","totalSteps":0,"userInitiated":false,"name":"Create","owner":"wenlock","created":"2015-09-10T22:50:05.406Z","eTag":"10","modified":"2015-09-10T22:50:14.250Z","uri":"/rest/tasks/DCB672A9-BC74-49DB-9D8F-798F3E82EC4C"}`
	)
	err := json.Unmarshal([]byte(test_json_data), &task)
	assert.NoError(t, err, fmt.Sprintf("Failed to unmarshal task object: %s, %+v\n", err, task))
}
