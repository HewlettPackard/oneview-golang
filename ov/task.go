package ov

import (
  "strings"
  "encoding/json"
  "time"
  "errors"

  "github.com/docker/machine/log"
  "github.com/docker/machine/drivers/oneview/rest"
  "github.com/docker/machine/drivers/oneview/utils"
)
// associated resource
type AssociatedResource struct {
  ResourceName     utils.Nstring `json:"resourceName,omitempty"`     // "resourceName": "se05, bay 16",
  AssociationType  string `json:"associationType,omitempty"`  // "associationType": "MANAGED_BY",
  ResourceCateogry string `json:"resourceCategory,omitempty"` // "resourceCategory": "server-hardware",
  ResourceURI      utils.Nstring `json:"resourceUri,omitempty"`      // "resourceUri": "/rest/server-hardware/30373237-3132-4D32-3235-303930524D57"
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

// Task Error
type TaskError struct {
  Data                map[string]interface{} `json:"data,omitempty"`               // "data":{},
  ErrorCode           string                 `json:"errorCode,omitempty"`          // "errorCode":"MacTypeDiffGlobalMacType",
  Details             string                 `json:"details,omitempty"`            // "details":"",
  NestedErrors        []string               `json:"nestedErrors,omitempty"`       // "nestedErrors":[],
  Message             string                 `json:"message,omitempty"`            // "message":"When macType is not user defined, mac type should be same as the global Mac assignment Virtual."
  ErrorSource         utils.Nstring                `json:"errorSource,omitempty"`        // "errorSource":null,
  RecommendedActions  []string               `json:"recommendedActions,omitempty"` // "recommendedActions":["Verify parameters and try again."],
}

// Task Progress Updates
type ProgressUpdate struct {
  TimeStamp     string   `json:"timestamp,omitempty"`    // "timestamp":"2015-09-10T22:50:14.250Z",
  StatusUpdate  string   `json:"statusUpdate,omitempty"` // "statusUpdate":"Apply server settings.",
  ID            int      `json:"id,omitempty"`           // "id":12566
}

// Task structure
type Task struct {
  Type                    string             `json:"type,omitempty"`               // "type": "TaskResourceV2",
  Data                    utils.Nstring            `json:"data,omitempty"`               // "data": null,
  Category                string             `json:"category,omitempty"`           // "category": "tasks",
  Hidden                  bool               `json:"hidden,omitempty"`             // "hidden": false,
  StateReason             string             `json:"stateReason,omitempty"`        // "stateReason": null,
  User                    string             `json:"User,omitempty"`               // "taskType": "User",
  AssociatedRes           AssociatedResource `json:"associatedResource,omitempty"` // "associatedResource": { },
  PercentComplete         int                `json:"percentComplete,omitempty"`    // "percentComplete": 0,
  AssociatedTaskUri       utils.Nstring            `json:"associatedTaskUri,omitempty"`  // "associatedTaskUri": null,
  CompletedSteps          int                `json:"completedSteps,omitempty"`     // "completedSteps": 0,
  ComputedPercentComplete int                `json:"computedPercentComplete,omitempty"` //     "computedPercentComplete": 0,
  ExpectedDuration        int                `json:"expectedDuration,omitempty"`   // "expectedDuration": 300,
  ParentTaskUri           utils.Nstring            `json:"parentTaskUri,omitempty"`      // "parentTaskUri": null,
  ProgressUpdates         []ProgressUpdate   `json:"progressUpdates,omitempty"`    // "progressUpdates": [],
  TaskErrors              []TaskError        `json:"taskErrors,omitempty"`         // "taskErrors": [],
  TaskOutput              []string           `json:"taskOutput,omitempty"`         // "taskOutput": [],
  TaskState               string             `json:"taskState,omitempty"`          // "taskState": "New",
  TaskStatus              string             `json:"taskStatus,omitempty"`         // "taskStatus": "Power off Server: se05, bay 16",
  TaskType                string             `json:"taskType,omitempty"`
  TotalSteps              int                `json:"totalSteps,omitempty"`         // "totalSteps": 0,
  UserInitiated           bool               `json:"userInitiated,omitempty"`      // "userInitiated": true,
  Name                    string             `json:"name,omitempty"`               // "name": "Power off",
  Owner                   string             `json:"owner,omitempty"`              // "owner": "wenlock",
  ETAG                    string             `json:"eTag,omitempty"`               // "eTag": "0",
  Created                 string             `json:"created,omitempty"`            // "created": "2015-09-07T03:25:54.844Z",
  Modified                string             `json:"modified,omitempty"`           // "modified": "2015-09-07T03:25:54.844Z",
  URI                     utils.Nstring             `json:"uri,omitempty"`                // "uri": "/rest/tasks/145F808A-A8DD-4E1B-8C86-C2379C97B3B2"
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
  							Timeout:     144, // default 24min
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
    log.Debugf(uri.String())
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
  if (len(t.TaskErrors) > 0) {
    var errmsg string
    errmsg = ""
    for _, te := range t.TaskErrors {
      errmsg += te.Message + " \n" + strings.Join(te.RecommendedActions, " ")
    }
    return errors.New(errmsg)
  }
	return nil
}

// wait on task to complete
func ( t *Task ) Wait()(error) {
  var (
		currenttime int = 0
	)
	log.Debugf("task : %+v", t)

  for !t.TaskIsDone && (currenttime < t.Timeout) {
    if err := t.GetCurrentTaskStatus(); err != nil {
      t.TaskIsDone = true
      return err
    }
    if t.URI != "" && T_COMPLETED.Equal(t.TaskState) {
      t.TaskIsDone = true
    }
    if t.URI != "" {
      log.Debugf("Waiting for task to complete, for %s ", t.Name)
      log.Infof("Waiting on, %s, %d%%, %s.", t.Name, t.ComputedPercentComplete, t.TaskStatus)
    } else {
      log.Info("Waiting on task creation.")
    }

    // wait time before next check
    time.Sleep(time.Millisecond * (1000 * t.WaitTime)) // wait 10sec before checking the status again
    currenttime++
  }
  if !(currenttime < t.Timeout) {
    log.Warn("Task timed out.")
  }

  if (t.Name != "") {
    log.Infof("Task, %s, completed", t.Name)
  }
  return nil
}
