package bm

// Equipment are additional elements for the boat or
// the crew.
type Equipment struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
