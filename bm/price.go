package bm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
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

// GetPriceRequest describes the request needed for requesting prices.
type GetPriceRequest struct {
	DateFrom     *MMKDateTime `json:"dateFrom,omitempty"`
	DateTo       *MMKDateTime `json:"dateTo,omitempty"`
	CompanyID    []int        `json:"companyId,omitempty"`
	CountryID    []int        `json:"countryId,omitempty"`
	ProductName  string       `json:"productName,omitempty"`
	YachtID      []int64      `json:"yachtId,omitempty"`
	Currency     string       `json:"currency,omitempty"`
	TripDuration int          `json:"tripDuration,omitempty"`
}

// PriceService operates over price requests.
type PriceService service

// GetPrice gets the price that a boat is in for.
func (ps *PriceService) GetPrice(pr *GetPriceRequest) (p *Price, err error) {
	v, _ := query.Values(pr)
	target := fmt.Sprintf("/prices?%s", v.Encode())

	req, err := ps.client.NewAPIRequest(http.MethodGet, target, nil)
	if err != nil {
		return
	}

	res, err := ps.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &p); err != nil {
		return
	}

	return
}
