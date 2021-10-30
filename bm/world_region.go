package bm

// WorldRegion describes a single world region in Booking Manager.
type WorldRegion struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
