package bm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Offer is a product, price and time frame proposal for a
// booking operation.
type Offer struct {
	YachtId            int64      `json:"yachtId,omitempty"`
	StartBaseID        int64      `json:"startBaseId,omitempty"`
	EndBaseID          int64      `json:"endBaseId,omitempty"`
	StartPrice         float64    `json:"startPrice,omitempty"`
	Price              float64    `json:"price,omitempty"`
	DiscountPercentage float64    `json:"discountPercentage,omitempty"`
	Yacht              string     `json:"yacht,omitempty"`
	StartBase          string     `json:"startBase,omitempty"`
	EndBase            string     `json:"endBase,omitempty"`
	Product            string     `json:"product,omitempty"`
	Currency           string     `json:"currency,omitempty"`
	DateFrom           *time.Time `json:"dateFrom,omitempty"`
	DateTo             *time.Time `json:"dateTo,omitempty"`
}

// OfferParams are params that could be passed to the offer request
type OfferParams struct {
	DateFrom *time.Time
	DateTo   *time.Time
	YachtId  *int64
}

type OffersService service

// GetOffers returns offers based on passed parameters
func (os *OffersService) GetOffers(op *OfferParams) (or []*Offer, err error) {
	var target string
	{
		target = fmt.Sprintf(
			"offers?dateFrom=%s&dateTo=%s",
			op.DateFrom.Format("20060102T000000"),
			op.DateTo.Format("20060102T000000"),
		)
		if op.YachtId != nil {
			target = fmt.Sprintf("%s&yachtId=%d", target, op.YachtId)
		}
	}

	req, err := os.client.NewAPIRequest(http.MethodGet, target, nil)
	if err != nil {
		return
	}

	res, err := os.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &or); err != nil {
		return
	}

	return
}
