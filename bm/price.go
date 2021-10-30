package bm

import (
	"time"
)

// Price describes the base price for an operation with MMK.
type Price struct {
	YatchID            int64      `json:"yatchId,omitempty"`
	DateFrom           *time.Time `json:"dateFrom,omitempty"`
	DateTo             *time.Time `json:"dateTo,omitempty"`
	Product            string     `json:"product,omitempty"`
	Currency           string     `json:"currency,omitempty"`
	Price              float64    `json:"price,omitempty"`
	StartPrice         float64    `json:"startPrice,omitempty"`
	DiscountPercentage float64    `json:"discountPercentage,omitempty"`
}
