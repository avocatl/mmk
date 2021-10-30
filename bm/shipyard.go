package bm

// Shipyard is a place where ships are built and repaired.
type Shipyard struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ShortName string `json:"shortName,omitempty"`
}
