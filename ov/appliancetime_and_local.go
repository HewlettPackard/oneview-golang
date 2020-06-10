package ov
import (
    "encoding/json"
    "fmt"
    "github.com/HewlettPackard/oneview-golang/rest"
    "github.com/HewlettPackard/oneview-golang/utils"
    "github.com/docker/machine/libmachine/log"
)
type ApplianceTimeandLocal struct {
    Type              string          `json:"type,omitempty"`
    Locale            string          `json:"locale,omitempty"`
    LocaleDisplayName string        `json:"localeDisplayName,omitempty"`
    DateTime          utils.Nstring   `json:"dateTime,omitempty"`
    Timezone          utils.Nstring   `json:"timezone,omitempty"`
    NtpServers        []utils.Nstring  `json:"ntpServers,omitempty"` // "ntpServers":[]
    PollingInterval   int             `json:"pollingInterval"`
    Category          string          `json:"category,omitempty"`
    URI               utils.Nstring   `json:"uri,omitempty"`
    ETAG              string          `json:"eTag,omitempty"`
    Modified          string          `json:"modified,omitempty"`
    Created           string          `json:"created,omitempty"`
}
func (c *OVClient) CreateApplianceTimeandLocal(timelocal ApplianceTimeandLocal) error {
    log.Infof("Initializing creation of time and local for %s.", timelocal.Locale)
    var (
        uri = "/rest/appliance/configuration/time-locale"
        t   = (&Task{}).NewProfileTask(c)
    )
    // refresh login
    c.RefreshLogin()
    c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
    t.ResetTask()
    log.Debugf("REST : %s \n %+v\n", uri, timelocal)
    log.Debugf("task -> %+v", t)
    data, err := c.RestAPICall(rest.POST, uri, timelocal)
    if err != nil {
        t.TaskIsDone = true
        log.Errorf("Error submitting new time and local request: %s", err)
        return err
    }
    log.Debugf("Response New timelocalwork %s", data)
    if err := json.Unmarshal(data, &t); err != nil {
        t.TaskIsDone = true
        log.Errorf("Error with task un-marshal: %s", err)
        return err
    }
    err = t.Wait()
    if err != nil {
        return err
    }
    return nil
}
func (c *OVClient) GetApplianceTimeandLocals(filter string, sort string, start string, count string)) (ApplianceTimeandLocal, error) {
    var (
        uri                    = "/rest/appliance/configuration/time-locale"
        q                      = make(map[string]interface{})
        applianceTimeandLocals ApplianceTimeandLocal
    )
    if len(filter) > 0 {
        q["filter"] = filter
    }
    if sort != "" {
        q["sort"] = sort
    }
    if start != "" {
        q["start"] = start
    }
    if count != "" {
        q["count"] = count
    }
    // refresh login
    c.RefreshLogin()
    c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
    // Setup query
    if len(q) > 0 {
        c.SetQueryString(q)
    }
    data, err := c.RestAPICall(rest.GET, uri, nil)
    if err != nil {
        return applianceTimeandLocals, err
    }
    log.Debugf("GetapplianceTimeandLocals %s", data)
    if err := json.Unmarshal(data, &applianceTimeandLocals); err != nil {
        return applianceTimeandLocals, err
    }
    return applianceTimeandLocals, nil
}
