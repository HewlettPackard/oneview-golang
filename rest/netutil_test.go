package rest

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var empty = &Client{}

func Test200Request(t *testing.T) {
	response := "This is the response"

	ts, endpoint, path := getServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})
	defer ts.Close()

	c := empty.NewClient("", "", endpoint)
	res, err := c.RestAPICall(GET, path, nil)

	if err != nil {
		t.Logf("Error received when performing request to %q: %s", ts.URL, err.Error())
		t.Fail()
	}

	if str := string(res); response != str {
		t.Logf("Expected response %q, received %q", response, str)
		t.Fail()
	}
}

func Test404Request(t *testing.T) {
	ts, endpoint, path := getServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	defer ts.Close()

	c := empty.NewClient("", "", endpoint)
	res, err := c.RestAPICall(POST, path, nil)

	if len(res) != 0 {
		t.Logf("Response included content, it should be empty: %q", string(res))
		t.Fail()
	}

	if err == nil {
		t.Logf("Error is nil, actual error expected.")
		t.Fail()
	}
}

func TestHeaderAuthenticaiton(t *testing.T) {
	var (
		token = "abcdef123"
		msg   = "Token invalid"
		out   = "Ok!"
	)

	ts, endpoint, path := getServer(func(w http.ResponseWriter, r *http.Request) {
		if local := r.Header.Get("X-Auth-Token"); local != token {
			http.Error(w, fmt.Sprintf(`{ "details": %q }`, msg), http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(out))
	})
	defer ts.Close()

	c := empty.NewClient("", "", endpoint)

	res, err := c.RestAPICall(GET, path, nil)

	if len(res) != 0 {
		t.Logf("Response included content, it should be empty: %q", string(res))
		t.Fail()
	}

	if err != nil && !strings.Contains(err.Error(), msg) {
		t.Logf("Expected error to contain %q. Original message: %q", msg, err.Error())
		t.Fail()
	}

	c.SetAuthHeaderOptions(map[string]string{"X-Auth-Token": token})
	res, err = c.RestAPICall(GET, path, nil)

	if err != nil {
		t.Logf("Unexpected error: %s", err.Error())
		t.Fail()
	}

	if s := string(res); s != out {
		t.Logf("Expected response body to be %q, got %q instead", out, s)
		t.Fail()
	}
}

func TestNewHTTPClientTLSConfig(t *testing.T) {
	t.Run("SSLVerify true enables cert verification", func(t *testing.T) {
		c := &Client{SSLVerify: true}
		hc := c.newHTTPClient()
		transport, ok := hc.Transport.(*http.Transport)
		if !ok {
			t.Fatalf("expected *http.Transport, got %T", hc.Transport)
		}

		if transport.TLSClientConfig == nil {
			t.Fatal("expected TLSClientConfig to be initialized")
		}

		if transport.TLSClientConfig.MinVersion != tls.VersionTLS13 {
			t.Fatalf("expected MinVersion TLS1.3, got %v", transport.TLSClientConfig.MinVersion)
		}

		if transport.TLSClientConfig.InsecureSkipVerify {
			t.Fatal("expected InsecureSkipVerify=false when SSLVerify=true")
		}
	})

	t.Run("SSLVerify false disables cert verification", func(t *testing.T) {
		c := &Client{SSLVerify: false}
		hc := c.newHTTPClient()
		transport, ok := hc.Transport.(*http.Transport)
		if !ok {
			t.Fatalf("expected *http.Transport, got %T", hc.Transport)
		}

		if transport.TLSClientConfig == nil {
			t.Fatal("expected TLSClientConfig to be initialized")
		}

		if transport.TLSClientConfig.MinVersion != tls.VersionTLS13 {
			t.Fatalf("expected MinVersion TLS1.3, got %v", transport.TLSClientConfig.MinVersion)
		}

		if !transport.TLSClientConfig.InsecureSkipVerify {
			t.Fatal("expected InsecureSkipVerify=true when SSLVerify=false")
		}
	})
}

func getServer(h func(http.ResponseWriter, *http.Request)) (*httptest.Server, string, string) {
	ts := httptest.NewServer(http.HandlerFunc(h))
	endpoint, path := getEndpointAndPath(ts)
	return ts, endpoint, path
}

func getEndpointAndPath(ts *httptest.Server) (string, string) {
	if ts == nil {
		return "", ""
	}

	// We don't need to check for errors here, since "ts" was created by the
	// Go HTTP test library
	u, _ := url.Parse(ts.URL)

	return fmt.Sprintf("%s://%s", u.Scheme, u.Host),
		fmt.Sprintf("%s?%s", u.Path, u.RawQuery)
}
