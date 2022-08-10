package bm

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
)

// MMK Booking-Manager global constants.
const (
	BaseURL            = "https://www.booking-manager.com/api/v2/"
	AuthHeader         = "Authorization"
	TokenType          = "Bearer"
	APITokenContainer  = "MMK_API_TOKEN"
	RequestContentType = "application/json"
)

// // MMK Booking-Manager global errors.
var (
	errEmptyAuthKey = errors.New("you must provide a non-empty authentication key")
	errBadBaseURL   = errors.New("malformed base url, it must contain a trailing slash")
)

// Client manages communication with MMK's API.
type Client struct {
	BaseURL        *url.URL
	authentication string
	userAgent      string
	client         *http.Client
	common         service // Reuse a single struct instead of allocating one for each service on the heap.
	// Services
	Availability *AvailabilityService
	Countries    *CountryService
	Offers       *OffersService
	Reservation  *ReservationService
	Price        *PriceService
}

// NewClient returns a new MMK HTTP API client.
// You can pass a previously build http client, if none is provided then
// http.DefaultClient will be used.
//
// NewClient will lookup the environment for values to assign to the
// API token (`MMK_API_TOKEN`).
func NewClient(baseClient *http.Client) (mmk *Client, err error) {
	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	u, _ := url.Parse(BaseURL)

	mmk = &Client{
		BaseURL: u,
		client:  baseClient,
	}

	mmk.common.client = mmk

	// golang base user agent binding
	mmk.userAgent = strings.Join([]string{
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Version(),
	}, ";")

	// services for resources
	mmk.Countries = (*CountryService)(&mmk.common)
	mmk.Availability = (*AvailabilityService)(&mmk.common)
	mmk.Offers = (*OffersService)(&mmk.common)
	mmk.Reservation = (*ReservationService)(&mmk.common)
	mmk.Price = (*PriceService)(&mmk.common)

	// Parse authorization from specified environment variable
	tkn, ok := os.LookupEnv(APITokenContainer)
	if !ok {
		return nil, errEmptyAuthKey
	}
	mmk.authentication = tkn
	return
}

type service struct {
	client *Client
}

// NewAPIRequest is a wrapper around the http.NewRequest function.
//
// It will setup the authentication headers/parameters according to the client config.
func (c *Client) NewAPIRequest(method string, uri string, body interface{}) (req *http.Request, err error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, errBadBaseURL
	}

	u, err := c.BaseURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add(AuthHeader, strings.Join([]string{TokenType, c.authentication}, " "))
	req.Header.Set("Content-Type", RequestContentType)
	req.Header.Set("Accept", RequestContentType)
	req.Header.Set("User-Agent", c.userAgent)

	return
}

// Do sends an API request and returns the API response or returned as an
// error if an API error has occurred.
func (c *Client) Do(req *http.Request) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, _ := newResponse(resp)
	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	return response, nil
}

func newResponse(r *http.Response) (*Response, error) {
	var res Response
	c, err := ioutil.ReadAll(r.Body)
	if err == nil {
		res.content = c
	}
	err = json.NewDecoder(r.Body).Decode(&res)
	r.Body = io.NopCloser(bytes.NewBuffer(c))
	res.Response = r
	return &res, err
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have either no response
// body, or a JSON response body.
func CheckResponse(r *http.Response) error {
	if r.StatusCode >= http.StatusMultipleChoices {
		return newError(r)
	}
	return nil
}

/*
Constructor for Error
*/
func newError(r *http.Response) *Error {
	var e Error
	e.Response = r
	e.Code = r.StatusCode
	e.Message = r.Status
	c, err := ioutil.ReadAll(r.Body)
	if err == nil {
		e.Content = string(c)
	}
	r.Body = io.NopCloser(bytes.NewBuffer(c))
	return &e
}
