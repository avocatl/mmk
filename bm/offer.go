package bm

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

// Offer is a product, price and time frame proposal for a
// booking operation.
type Offer struct {
	YachtId            int64        `json:"yachtId,omitempty"`
	StartBaseID        int64        `json:"startBaseId,omitempty"`
	EndBaseID          int64        `json:"endBaseId,omitempty"`
	StartPrice         float64      `json:"startPrice,omitempty"`
	Price              float64      `json:"price,omitempty"`
	DiscountPercentage float64      `json:"discountPercentage,omitempty"`
	Yacht              string       `json:"yacht,omitempty"`
	StartBase          string       `json:"startBase,omitempty"`
	EndBase            string       `json:"endBase,omitempty"`
	Product            string       `json:"product,omitempty"`
	Currency           string       `json:"currency,omitempty"`
	DateFrom           *MMKDateTime `json:"dateFrom,omitempty"`
	DateTo             *MMKDateTime `json:"dateTo,omitempty"`
}

// OfferOptions are params that could be passed to the offer request
type OfferOptions struct {
	DateFrom string `url:"dateFrom"`
	DateTo   string `url:"dateTo"`
	YachtIds *[]int `url:"yachtId"`
}

type OffersService service

// GetOffers returns offers based on passed parameters
func (os *OffersService) GetOffers(opts *OfferOptions) (or []*Offer, err error) {
	var target string
	{
		target = "offers"

		if opts != nil {
			v, _ := query.Values(opts)
			target = fmt.Sprintf("%s?%s", target, v.Encode())
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
