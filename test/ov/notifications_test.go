package ov

import (
	"os"
	"testing"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/docker/machine/libmachine/log"
	"github.com/stretchr/testify/assert"
)

func TestSendTestEmail(t *testing.T) {
	var (
		d        *OVTest
		c        *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		d, c = getTestDriverA("test_email_notification")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
			email := ov.TestEmailRequest{
				ToAddress:                    d.Tc.GetTestData(d.Env, "ToAddress").(string),,
				HtmlMessageBody:              d.Tc.GetTestData(d.Env, "HtmlMessageBody").(string),
				TextMessageBody:              d.Tc.GetTestData(d.Env, "TextMessageBody").(string),
				Subject:                      d.Tc.GetTestData(d.Env, "Subject").(string),
			}

			err := c.SendTestEmail(email)
			assert.NoError(t, err, "SendTestEmail error -> %s", err)

			err = c.SendTestEmail(email)
			assert.Error(t, err, "SendTestEmail should error becaue the network already exists, err -> %s", err)
		} else {
			log.Warnf("The email already exists so skipping SendTestEmail test for %s")
		}
	}
}

func TestGetEmailNotifications(t *testing.T) {
	var (
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_email_notification")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		emailNotifications, err := c.GetEmailNotifications("", "", "", "")
		assert.NoError(t, err, "GetEmailNotifications threw an error -> %s. %+v\n", err, emailNotifications)

	} else {
		_, c = getTestDriverU("test_email_notification")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetEmailNotifications("", "", "", "")
		assert.Error(t, err, fmt.Sprintf("All OK, no error, caught as expected: %s,%+v\n", err, data))

	}

	_, c = getTestDriverU("test_email_notification")
	data, err := c.GetEmailNotifications("", "", "", "")
	assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
}

func TestGetEmailNotificationsConfiguration(t *testing.T) {
	var (
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_email_notification")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		emailNotificationConfigurations, err := c.GetEmailNotificationsConfiguration("", "", "", "")
		assert.NoError(t, err, "GetEmailNotificationsConfiguration threw an error -> %s. %+v\n", err, emailNotificationConfigurations)

	} else {
		_, c = getTestDriverU("test_email_notification")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.GetEmailNotificationsConfiguration("", "", "", "")
		assert.Error(t, err, fmt.Sprintf("All OK, no error, caught as expected: %s,%+v\n", err, data))

	}

	_, c = getTestDriverU("test_email_notification")
	data, err := c.GetEmailNotificationsConfiguration("", "", "", "")
	assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
}

func GetEmailNotificationsByFilter(t *testing.T) {
	var (
		c *ov.OVClient
	)
	if os.Getenv("ONEVIEW_TEST_ACCEPTANCE") == "true" {
		_, c = getTestDriverA("test_email_notification")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		emailNotificationsFilter, err := c.GetEmailNotificationsByFilter("", "", "", "")
		assert.NoError(t, err, "emailNotificationsFilter threw an error -> %s. %+v\n", err, emailNotificationsFilter)

	} else {
		_, c = getTestDriverU("test_email_notification")
		if c == nil {
			t.Fatalf("Failed to execute getTestDriver() ")
		}
		data, err := c.emailNotificationsFilter("scope:Name", "", "", "")
		assert.Error(t, err, fmt.Sprintf("All OK, no error, caught as expected: %s,%+v\n", err, data))

	}

	_, c = getTestDriverU("test_email_notification")
	data, err := c.emailNotificationsFilter("", "", "", "")
	assert.Error(t, err, fmt.Sprintf("ALL ok, no error, caught as expected: %s,%+v\n", err, data))
}
