package bm

// Layout of an offered product.
type Layout struct {
	Name   string `json:"name,omitempty"`
	Amount int64  `json:"amount,omitempty"`
}
