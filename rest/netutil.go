package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/docker/machine/log"
)

// Options for REST call
type Options struct {
  Headers map[string]string
	Query map[string]string
}

// Client - generic REST api client
type Client struct {
	Method
	User       string
	Password   string
	Domain     string
	APIKey     string
	APIVersion int
	SSLVerify  bool
	Endpoint   string
	Option     Options
}

// NewClient - get a new network client
func (c *Client) NewClient(user, key, endpoint string) *Client {
	var options Options
	return &Client{User: user, APIKey: key, Endpoint: endpoint, Option: options}
}

// isOkStatus - check the return status of the response
func (c *Client) isOkStatus(code int) bool {
	codes := map[int]bool{
		200: true,
		201: true,
		204: true,
		400: false,
		404: false,
		500: false,
		409: false,
		406: false,
	}

	return codes[code]
}

// SetQueryString - set the query strings to use
func (c *Client) SetQueryString(query map[string]string) {
	// TODO: uuencode the query String
	c.Option.Query = query
}

// GetQueryString - get a query string for url
func (c *Client) GetQueryString(u *url.URL) {
	if len(c.Option.Query) == 0 {
		return
	}
	parameters := url.Values{}
	for k, v := range  c.Option.Query {
		parameters.Add(k, v)
		u.RawQuery = parameters.Encode()
	}
	return
}

// SetAuthHeaderOptins - set the Headers Options
func (c *Client) SetAuthHeaderOptions(headers map[string]string) {
	c.Option.Headers = headers
}

// RestAPICall - general rest method caller
func (c *Client) RestAPICall(method Method, path string, options interface{}) ([]byte, error) {
	log.Debugf("RestAPICall %s - %s%s", method, c.Sanatize(c.Endpoint), path)

	var (
		Url *url.URL
		err error
		req *http.Request
	)

	Url, err = url.Parse(c.Sanatize(c.Endpoint))
	if err != nil {
		return nil, err
	}
	Url.Path += path

	// Manage the query string
	c.GetQueryString(Url)

	// TODO: this should have a real cert
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	if options != nil {
		OptionsJSON, err := json.Marshal(options)
		if err != nil {
			return nil, err
		}
		// TODO: remove comment when done with dev
		// fmt.Printf("*******  %+v\n",bytes.NewBuffer(OptionsJSON))
		req, err = http.NewRequest(string(method), Url.String(), bytes.NewBuffer(OptionsJSON))
	} else {
		req, err = http.NewRequest(string(method), Url.String(), nil)
	}

	if err != nil {
		return nil, fmt.Errorf("Error with request: %v - %q", Url, err)
	}

	// build the auth headerU
	for k, v := range  c.Option.Headers {
		log.Debugf("Headers -> %s -> %+v\n",k,v)
		req.Header.Add(k, v)
  }

	// req.SetBasicAuth(c.User, c.APIKey)
	req.Method = fmt.Sprintf("%s", method)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: CLeanup Later
	// DEBUGGING WHILE WE WORK
	// DEBUGGING WHILE WE WORK
	// fmt.Printf("METHOD --> %+v\n",method)
	// fmt.Printf("REQ    --> %+v\n",req)
	// fmt.Printf("RESP   --> %+v\n",resp)
	// fmt.Printf("ERROR  --> %+v\n",err)
	// DEBUGGING WHILE WE WORK

	data, err := ioutil.ReadAll(resp.Body)

	if !c.isOkStatus(resp.StatusCode) {
		type apiErr struct {
			Err string `json:"details"`
		}
		var outErr apiErr
		json.Unmarshal(data, &outErr)
		return nil, fmt.Errorf("Error in response: %s\n Response Status: %s", outErr.Err, resp.Status)
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}
