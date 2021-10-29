package bm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
)

// AvailabilityResponse descriptor.
type AvailabilityResponse struct {
	ID                   int64      `json:"id,omitempty"`
	YatchID              int64      `json:"yatchId,omitempty"`
	Status               int64      `json:"status,omitempty"`
	BaseFromID           int64      `json:"baseFromId,omitempty"`
	BaseToID             int64      `json:"baseToId,omitempty"`
	DateFrom             *time.Time `json:"dateFrom,omitempty"`
	DateTo               *time.Time `json:"dateTo,omitempty"`
	OptionExpirationDate *time.Time `json:"optionExpirationDate,omitempty"`
}

// AvailabilityService operates over availability requests.
type AvailabilityService service

// AvailabilityOptions describes valid query string parameters
// for availability requests.
type AvailabilityOptions struct {
	CompanyID int `url:"companyId,omitempty"`
}

// GetAvailability returns availability for specific year.
func (as *AvailabilityService) GetAvailability(year int, opts *AvailabilityOptions) (ar []*AvailabilityResponse, err error) {
	var target string
	{
		target = fmt.Sprintf("/availability/%s", strconv.Itoa(year))

		if opts != nil {
			v, _ := query.Values(opts)
			target = fmt.Sprintf("%s?%s", target, v.Encode())
		}
	}

	req, err := as.client.NewAPIRequest(http.MethodGet, target, nil)
	if err != nil {
		return
	}

	res, err := as.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &ar); err != nil {
		return
	}

	return
}

// ShortAvailabilityFormat supported by MMK.
type ShortAvailabilityFormat int

// Supported short availability formats.
const (
	Binary ShortAvailabilityFormat = iota + 1
	Hex
	Status
)

// ShortAvailabilityOptions available on MMK's API.
type ShortAvailabilityOptions struct {
	CompanyID int64                   `url:"companyId,omitempty"`
	YatchID   int64                   `url:"yatchId,omitempty"`
	Format    ShortAvailabilityFormat `url:"format,omitempty"`
}

// ShortAvailabilityResponse descriptor.
type ShortAvailabilityResponse struct {
	YatchID int64  `json:"yatchId,omitempty"`
	BS      string `json:"bs,omitempty"`
}

// GetShortAvailability returns availability for specific year.
// Different status format will retrieve results in different format.
//
// Formats and definition.
// 1 Binary - availability information for the resource in binary format.
// Each availabilityInfo is 365 characters long (or 366 in case of the leap year)
// and it represents the whole year (single character is one day of the year).
// First character is January 1st, second character is January 2nd, and so on.
// If character is equal to “0” it means that yacht is available on that day,
// otherwise character is equalto “1”.
// 2 Hex - availability information for the resource in hexadecimal format.
// Each availabilityInfo is 92 characters long and it represents the whole year.
// First character is January 1, January 2, January 3 and January 4, second character is January 5,
// January 6, January 7 and January 8, and so on.
// Last characters should be discarded depending on how long the requested year is.
// Example, if availabilityInfo is “fe03f0..” in binary it is “0000 0001 1111 1100 0000 1111..”
// and it means that yacht is available from January 1 until January 8, but it is not available from
// January 8 until January 15, and so on.
// 3 Status - Response is the same as the Binary response with one difference –
// booked days are not represented with character “1” for all types of reservations but instead
// reservation status id is used.
// Example, if availabilityInfo is “2222 2220 0000...” it means that boat is under status Option
// from January 1 until January 8, and it is free from January 8 etc.
func (as *AvailabilityService) GetShortAvailability(year int, opts *ShortAvailabilityOptions) (sar *ShortAvailabilityResponse, err error) {
	var target string
	{
		target = fmt.Sprintf("/shortAvailability/%s", strconv.Itoa(year))

		if opts != nil {
			v, _ := query.Values(opts)
			target = fmt.Sprintf("%s?%s", target, v.Encode())
		}
	}

	req, err := as.client.NewAPIRequest(http.MethodGet, target, nil)
	if err != nil {
		return
	}

	res, err := as.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &sar); err != nil {
		return
	}

	return
}
