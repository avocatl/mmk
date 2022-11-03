package bm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CharterCondition response description
type CharterCondition struct {
	Period          *Period   `json:"period,omitempty"`
	CheckInDays     []*int    `json:"checkInDays,omitempty"`
	CheckOutDays    []*int    `json:"checkOutDays,omitempty"`
	MinimumDuration int       `json:"minimumDuration,omitempty"`
	Overnight       bool      `json:"overnight,omitempty"`
	CheckInTime     time.Time `json:"defaultCheckInTime,omitempty"`
	CheckOutTime    time.Time `json:"defaultCheckOutTime,omitempty"`
	LicenseRequired bool      `json:"licenseRequired,omitempty"`
	SkipperRequired bool      `json:"requiredSkipperLicense,omitempty"`
}

type CharterConditionService service

// GetCharterCondition
func (os *CharterConditionService) GetCharterCondition(id int) (or *Yacht, err error) {
	var target string
	{
		target = fmt.Sprintf("yacht/%d", id)
	}

	req, err := os.client.NewAPIRequest(http.MethodGet, target, nil)
	if err != nil {
		return
	}

	res, err := os.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &or); err != nil {
		return
	}

	return
}
