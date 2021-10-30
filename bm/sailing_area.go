package bm

// SailingArea describes a single sailing area in Booking Manager.
type SailingArea struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
