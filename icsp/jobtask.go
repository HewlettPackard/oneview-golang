package icsp

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/docker/machine/drivers/oneview/rest"
	"github.com/docker/machine/drivers/oneview/utils"
	"github.com/docker/machine/log"
)

// ODSUri  returned from create server for job uri task
type ODSUri struct {
	URI utils.Nstring `json:"uri,omitempty"` // uri of job
}

// JobTask holds a Job ODSUri and task status
type JobTask struct {
	Job                    // copy of the original job
	JobURI   ODSUri        // link to the job
	IsDone   bool          // when true, task are done
	Timeout  int           // time before timeout on Executor
	WaitTime time.Duration // time between task checks
	Client   *ICSPClient   // reference to a client
}

// NewJobTask create a new job task
func (jt *JobTask) NewJobTask(c *ICSPClient) *JobTask {
	return &JobTask{
		IsDone:   false,
		Client:   c,
		Timeout:  144, // default 24min
		WaitTime: 10}  // default 10sec, impacts Timeout
}

// Reset - reset job task
func (jt *JobTask) Reset() {
	jt.IsDone = false
}

// GetCurrentStatus - Get the current status
func (jt *JobTask) GetCurrentStatus() error {
	log.Debugf("Working on getting current job status")
	if jt.JobURI.URI != "" {
		log.Debugf(jt.JobURI.URI.String())
		data, err := jt.Client.RestAPICall(rest.GET, jt.JobURI.URI.String(), nil)
		if err != nil {
			return err
		}
		log.Debugf("data: %s", data)
		if err := json.Unmarshal([]byte(data), &jt); err != nil {
			return err
		}
	} else {
		log.Debugf("Unable to get current job, no URI found")
	}
	if JOB_STATUS_ERROR.Equal(jt.Status) {
		var errmsg string
		errmsg = ""
		for _, je := range jt.JobResult {
			if je.JobResultErrorDetails != "" {
				errmsg += je.JobMessage + " \n" + je.JobResultErrorDetails + "\n"
			}
		}
		return errors.New(errmsg)
	}
	return nil
}

// GetLastStatusUpdate get the last status from JobProgress
func (jt *JobTask) GetLastStatusUpdate() string {
	lastjobstep := len(jt.JobProgress)
	if lastjobstep > 0 {
		return jt.JobProgress[lastjobstep-1].CurrentStepName
	}
	return ""
}

// GetComplettedStatus  get the message from JobResult
func (jt *JobTask) GetComplettedStatus() string {
	lastjobstep := len(jt.JobResult)
	if lastjobstep > 0 {
		return jt.JobResult[lastjobstep-1].JobMessage
	}
	return ""

}

// Wait - wait on job task to complete
func (jt *JobTask) Wait() error {
	var (
		currenttime int
	)
	log.Debugf("task : %+v", jt)

	for !jt.IsDone && (currenttime < jt.Timeout) {
		if err := jt.GetCurrentStatus(); err != nil {
			jt.IsDone = true
			return err
		}
		if jt.JobURI.URI != "" && JOB_RUNNING_YES.Equal(jt.Running) {
			jt.IsDone = true
		}
		if jt.JobURI.URI != "" {
			log.Debugf("Waiting for job to complete, %s ", jt.Description)
			lastjobstep := len(jt.JobProgress)
			if lastjobstep > 0 {
				stepscompleted := jt.JobProgress[lastjobstep-1].JobCompletedSteps
				totalcompleted := jt.JobProgress[lastjobstep-1].JobTotalSteps
				progress := 0
				if totalcompleted > 0 {
					progress = stepscompleted / totalcompleted
				}
				log.Infof("Waiting on, %s, %d%%, %s", jt.Description, progress, jt.GetLastStatusUpdate())
			}
		} else {
			log.Info("Waiting on job creation.")
		}

		// wait time before next check
		time.Sleep(time.Millisecond * (1000 * jt.WaitTime)) // wait 10sec before checking the status again
		currenttime++
	}
	if !(currenttime < jt.Timeout) {
		log.Warn("Task timed out.")
	}

	if JOB_RUNNING_NO.Equal(jt.Running) {
		log.Infof("Job, %s, completed", jt.GetComplettedStatus())
	}
	return nil
}
