package ov

import (
  "strings"
  "encoding/json"
  "time"

  "github.com/docker/machine/log"
  "github.com/docker/machine/drivers/oneview/rest"
)
// associated resource
type AssociatedResource struct {
  ResourceName     string `json:"resourceName,omitempty"`     // "resourceName": "se05, bay 16",
  AssociationType  string `json:"associationType,omitempty"`  // "associationType": "MANAGED_BY",
  ResourceCateogry string `json:"resourceCategory,omitempty"` // "resourceCategory": "server-hardware",
  ResourceURI      string `json:"resourceUri,omitempty"`      // "resourceUri": "/rest/server-hardware/30373237-3132-4D32-3235-303930524D57"
}

// task state
type TaskState int

const(
	T_COMPLETED   TaskState = 1 + iota
	T_ERROR
	T_INERRUPTED
  T_KILLED
  T_NEW
  T_PENDING
  T_RUNNING
  T_STARTING
  T_STOPPING
  T_SUSPENDED
  T_TERMINATED
  T_UNKNOWN
  T_WARNING
)

var taskstate = [...]string {
  "Completed",   // Completed Task has been completed.
  "Error",       // Error Task has terminated with an error.
  "Interrupted", // Interrupted Task has been interrupted.
  "Killed",      // Killed Task has been killed.
  "New",         // New Task is new.
  "Pending",     // Pending Task is in pending state.
  "Running",     // Running Task is running.
  "Starting",    // Starting Task is starting.
  "Stopping",    // Stopping Task is stopping.
  "Suspended",   // Suspended Task is suspended.
  "Terminated",  // Terminated Task has been terminated.
  "Unknown",     // Unknown State of task is unknown.
  "Warning",     // Warning Task has terminated with a warning.
}

func (ts TaskState) String() string { return taskstate[ts-1] }
func (ts TaskState) Equal(s string) (bool) {return (strings.ToUpper(s) == strings.ToUpper(ts.String()))}

// task type
type TaskType int

const(
  T_APPLIANCE   TaskType = 1 + iota
  T_BACKGROUND
  T_USER
)

var tasktype = [...]string {
  "Applicance",  // Appliance Task is appliance initiated and shows in notification panel.
  "Background",  // Background Task is appliance initiated and does not show in notification panel.
  "User",        // User Task is user initiated and shows in notification panel.
}

func (tt TaskType) String() string { return tasktype[tt-1] }
func (tt TaskType) Equal(s string) (bool) {return (strings.ToUpper(s) == strings.ToUpper(tt.String()))}

// Task structure

type Task struct {
  Type                    string             `json:"type,omitempty"`               // "type": "TaskResourceV2",
  Data                    Nstring            `json:"data,omitempty"`               // "data": null,
  Category                string             `json:"category,omitempty"`           // "category": "tasks",
  Hidden                  bool               `json:"hidden,omitempty"`             // "hidden": false,
  StateReason             string             `json:"stateReason,omitempty"`        // "stateReason": null,
  User                    string             `json:"User,omitempty"`               // "taskType": "User",
  AssociatedRes           AssociatedResource `json:"associatedResource,omitempty"` // "associatedResource": { },
  PercentComplete         int                `json:"percentComplete,omitempty"`    // "percentComplete": 0,
  AssociatedTaskUri       Nstring            `json:"associatedTaskUri,omitempty"`  // "associatedTaskUri": null,
  CompletedSteps          int                `json:"completedSteps,omitempty"`     // "completedSteps": 0,
  ComputedPercentComplete int                `json:"computedPercentComplete,omitempty"` //     "computedPercentComplete": 0,
  ExpectedDuration        int                `json:"expectedDuration,omitempty"`   // "expectedDuration": 300,
  ParentTaskUri           Nstring            `json:"parentTaskUri,omitempty"`      // "parentTaskUri": null,
  ProgressUpdates         []string           `json:"progressUpdates,omitempty"`    // "progressUpdates": [],
  TaskErrors              []string           `json:"taskErrors,omitempty"`         // "taskErrors": [],
  TaskOutput              []string           `json:"taskOutput,omitempty"`         // "taskOutput": [],
  TaskState               string             `json:"taskState,omitempty"`          // "taskState": "New",
  TaskStatus              string             `json:"taskStatus,omitempty"`         // "taskStatus": "Power off Server: se05, bay 16",
  TotalSteps              int                `json:"totalSteps,omitempty"`         // "totalSteps": 0,
  UserInitiated           bool               `json:"userInitiated,omitempty"`      // "userInitiated": true,
  Name                    string             `json:"name,omitempty"`               // "name": "Power off",
  Owner                   string             `json:"owner,omitempty"`              // "owner": "wenlock",
  ETAG                    string             `json:"eTag,omitempty"`               // "eTag": "0",
  Created                 string             `json:"created,omitempty"`            // "created": "2015-09-07T03:25:54.844Z",
  Modified                string             `json:"modified,omitempty"`           // "modified": "2015-09-07T03:25:54.844Z",
  URI                     Nstring             `json:"uri,omitempty"`                // "uri": "/rest/tasks/145F808A-A8DD-4E1B-8C86-C2379C97B3B2"
  TaskIsDone  bool             // when true, task are done
	Timeout     int              // time before timeout on Executor
	WaitTime    time.Duration    // time between task checks
  Client      *OVClient
}

// Create New Task
func ( t *Task ) NewProfileTask(c *OVClient)(*Task) {
	return &Task{ TaskIsDone:  false,
                Client:      c,
                URI:         "",
                Name:        "",
                Owner:       "",
  							Timeout:     36, // default 6min
  							WaitTime:    10} // default 10sec, impacts Timeout
}

// reset the power task back to off
func ( t *Task) ResetTask() {
	t.TaskIsDone  = false
	t.URI         = ""
  t.Name        = ""
  t.Owner       = ""
}

// Get the current status
func ( t *Task ) GetCurrentTaskStatus()(error) {
  log.Debugf("Working on getting current task status")
	var (
		uri  = t.URI
	)
	if uri != "" {
		data, err := t.Client.RestAPICall(rest.GET, uri.String(), nil)
		if err != nil {
			return err
		}
		log.Debugf("data: %s",data)
		if err := json.Unmarshal([]byte(data), &t); err != nil {
			return err
		}
	} else {
		log.Debugf("Unable to get current task, no URI found")
	}
	return nil
}
