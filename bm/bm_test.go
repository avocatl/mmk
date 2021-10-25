package bm

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	tMux    *http.ServeMux
	tServer *httptest.Server
	tClient *Client
)

func TestNewClient(t *testing.T) {
	func() {
		setEnv()
		setup()
	}()
	defer func() {
		unsetEnv()
		teardown()
	}()
	var c = http.DefaultClient
	{
		c.Timeout = 25 * time.Second
	}

	tests := []struct {
		name   string
		client *http.Client
	}{
		{
			"nil returns a valid client",
			nil,
		},
		{
			"a passed client is decorated",
			c,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.client)
			if err != nil {
				t.Errorf("not nil error received: %v", err)
			}
		})
	}
}

func TestNewClientWithEnvVars(t *testing.T) {
	setEnv()
	defer unsetEnv()

	var c = http.DefaultClient
	{
		c.Timeout = 25 * time.Second
	}

	tests := []struct {
		name   string
		client *http.Client
	}{
		{
			"nil returns a valid client",
			nil,
		},
		{
			"a passed client is decorated",
			c,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.client)
			if err != nil {
				t.Errorf("not nil error received: %v", err)
			}

			if got.authentication == "" {
				t.Errorf("got empty api key %v, value %s expected", got.authentication, "token_X12b31ggg23")
			}
		})
	}
}

func TestClient_NewAPIRequest(t *testing.T) {
	func() {
		setEnv()
		setup()
	}()
	defer func() {
		unsetEnv()
		teardown()
	}()
	b := []string{"hello", "bye"}
	inURL, outURL := "test", tServer.URL+"/test"
	inBody, outBody := b, `["hello","bye"]`+"\n"
	req, _ := tClient.NewAPIRequest("GET", inURL, inBody)

	testHeader(t, req, "Accept", RequestContentType)
	testHeader(t, req, AuthHeader, "Bearer test_token")
	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}
}

func TestClient_NewAPIRequest_ErrTrailingSlash(t *testing.T) {
	uri, _ := url.Parse("http://localhost")
	tClient = &Client{
		BaseURL: uri,
	}
	_, err := tClient.NewAPIRequest("GET", "test", nil)

	if err == nil {
		t.Errorf("expected error %v not occurred, got %v", errBadBaseURL, err)
	}
}

func TestClient_NewAPIRequest_HTTPReqNativeError(t *testing.T) {
	func() {
		setEnv()
		setup()
	}()
	defer func() {
		unsetEnv()
		teardown()
	}()
	_, err := tClient.NewAPIRequest("\\\\\\", "test", nil)

	if err == nil {
		t.Fatal("nil error produced")
	}

	if !strings.Contains(err.Error(), "invalid method") {
		t.Errorf("unexpected err received %v", err)
	}
}

func TestClient_NewAPIRequest_OrgTokenOverApiKey(t *testing.T) {
	func() {
		setEnv()
		setup()
	}()
	defer func() {
		unsetEnv()
		teardown()
	}()
	req, _ := tClient.NewAPIRequest("GET", "test", nil)

	testHeader(t, req, AuthHeader, "Bearer test_token")
}

func TestClient_NewAPIRequest_ErrorBodySerialization(t *testing.T) {
	func() {
		setEnv()
		setup()
	}()
	defer func() {
		unsetEnv()
		teardown()
	}()
	b := make(chan int)
	_, err := tClient.NewAPIRequest("GET", "test", b)

	if err == nil {
		t.Fatal("nil error produced")
	}

	if !strings.Contains(err.Error(), "unsupported type") {
		t.Errorf("unexpected err received %v", err)
	}
}

func TestClient_NewAPIRequest_NativeURLParseError(t *testing.T) {
	func() {
		setEnv()
		setup()
	}()
	defer func() {
		unsetEnv()
		teardown()
	}()
	_, err := tClient.NewAPIRequest("GET", ":", nil)

	if err == nil {
		t.Fatal("nil error produced")
	}

	if !strings.Contains(err.Error(), "parse") {
		t.Errorf("unexpected err received %v", err)
	}
}

func TestClient_Do(t *testing.T) {
	setEnv()
	setup()
	defer func() {
		unsetEnv()
		teardown()
	}()

	tMux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, AuthHeader, "Bearer test_token")
		w.WriteHeader(http.StatusOK)
	})

	req, _ := tClient.NewAPIRequest(http.MethodGet, "/test", nil)
	res, err := tClient.Do(req)

	if err != nil {
		t.Errorf("unexpected error received: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("request failed: %+v", res)
	}
}

func TestClient_DoErrInvalidJSON(t *testing.T) {
	func() {
		setEnv()
		setup()
	}()
	defer func() {
		unsetEnv()
		teardown()
	}()
	tMux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, AuthHeader, "Bearer test_token")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{"))
	})
	req, _ := tClient.NewAPIRequest("GET", "test", nil)
	req.URL = nil
	_, err := tClient.Do(req)

	if err == nil {
		t.Error(err)
	}
	if !strings.Contains(err.Error(), "nil Request.URL") {
		t.Errorf("unexpected response, got %v", err)
	}
}

func TestClient_DoErr(t *testing.T) {
	func() {
		setEnv()
		setup()
	}()
	defer func() {
		unsetEnv()
		teardown()
	}()
	req, _ := tClient.NewAPIRequest("GET", "test", nil)
	req.URL = nil
	_, err := tClient.Do(req)

	if err == nil {
		t.Error(err)
	}

	if !strings.Contains(err.Error(), "nil Request.URL") {
		t.Errorf("unexpected response, got %v", err)
	}
}

func TestCheckResponse(t *testing.T) {
	res1 := &http.Response{
		StatusCode: http.StatusNotFound,
		Status:     http.StatusText(http.StatusNotFound),
		Body:       ioutil.NopCloser(strings.NewReader("not found ok")),
	}

	res3 := &http.Response{
		StatusCode: http.StatusNotFound,
		Status:     http.StatusText(http.StatusNotFound),
		Body:       ioutil.NopCloser(strings.NewReader("")),
	}

	res2 := &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(strings.NewReader("success ok")),
	}

	tests := []struct {
		name string
		code string
		arg  *http.Response
	}{
		{
			"successful response",
			"",
			res2,
		},
		{
			"not found response",
			"Not Found",
			res1,
		},
		{
			"success with empty body",
			"",
			res3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckResponse(tt.arg); err != nil {
				if !strings.Contains(err.Error(), tt.code) {
					t.Error(err)
				}
			}
		})
	}
}

// <----- Testing helpers ----->

// the parameter indicates if you want to prepare your tests against the US sandbox
// just to be used when doing integration testing.
func setup() {
	tm := http.NewServeMux()
	ts := httptest.NewServer(tMux)
	tc, _ := NewClient(nil)
	u, _ := url.Parse(ts.URL + "/")
	tc.BaseURL = u

	tMux = tm
	tServer = ts
	tClient = tc
}

func teardown() {
	tServer.Close()
}

func setEnv() {
	_ = os.Setenv(APITokenContainer, "test_token")
}

func unsetEnv() {
	_ = os.Unsetenv(APITokenContainer)
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}
