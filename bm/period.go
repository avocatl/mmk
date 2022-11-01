package bm

// Period for a charter condition
type Period struct {
	DateFrom *MMKDateTime `json:"dateFrom,omitempty"`
	DateTo   *MMKDateTime `json:"dateTo,omitempty"`
}
