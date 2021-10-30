package bm

// Product represents an offered item on MMK booking system.
type Product struct {
	Name   string   `json:"name,omitempty"`
	Extras []*Extra `json:"extras,omitempty"`
}
