package bm

import "time"

// Offer is a product, price and time frame proposal for a
// booking operation.
type Offer struct {
	YatchID            int64      `json:"yatchId,omitempty"`
	StartBaseID        int64      `json:"startBaseId,omitempty"`
	EndBaseID          int64      `json:"endBaseId,omitempty"`
	StartPrice         float64    `json:"startPrice,omitempty"`
	Price              float64    `json:"price,omitempty"`
	DiscountPercentage float64    `json:"discountPercentage,omitempty"`
	Yatch              string     `json:"yatch,omitempty"`
	StartBase          string     `json:"startBase,omitempty"`
	EndBase            string     `json:"endBase,omitempty"`
	Product            string     `json:"product,omitempty"`
	Currency           string     `json:"currency,omitempty"`
	DateFrom           *time.Time `json:"dateFrom,omitempty"`
	DateTo             *time.Time `json:"dateTo,omitempty"`
}
