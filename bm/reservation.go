package bm

import "time"

// Reservation represents a booking and its links
// to a yatch and client.
type Reservation struct {
	DateFrom    *time.Time `json:"dateFrom,omitempty"`
	DateTo      *time.Time `json:"dateTo,omitempty"`
	YatchID     int64      `json:"yatchId,omitempty"`
	Status      int64      `json:"status,omitempty"`
	ProductName string     `json:"productName,omitempty"`
	BaseFromID  int64      `json:"baseFromId,omitempty"`
	BaseToID    int64      `json:"baseToId,omitempty"`
	ClientName  string     `json:"clientName,omitempty"`
	ClientID    int64      `json:"clientId,omitempty"`
	Currency    string     `json:"currency,omitempty"`
}
