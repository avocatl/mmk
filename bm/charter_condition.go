package bm

import "time"

// CharterCondition response description
type CharterCondition struct {
	Period          *Period   `json:"period,omitempty"`
	CheckInDays     []*int    `json:"checkInDays,omitempty"`
	CheckOutDays    []*int    `json:"checkOutDays,omitempty"`
	MinimumDuration int       `json:"minimumDuration,omitempty"`
	Overnight       bool      `json:"overnight,omitempty"`
	CheckInTime     time.Time `json:"checkInTime,omitempty"`
	CheckOutTime    time.Time `json:"checkOutTime,omitempty"`
	LicenseRequired bool      `json:"licenseRequired,omitempty"`
	SkipperRequired bool      `json:"skipperRequired,omitempty"`
}
