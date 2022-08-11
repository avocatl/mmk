package bm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Price describes the base price for an operation with MMK.
type Price struct {
	YatchID            int64        `json:"yatchId,omitempty"`
	DateFrom           *MMKDateTime `json:"dateFrom,omitempty"`
	DateTo             *MMKDateTime `json:"dateTo,omitempty"`
	Product            string       `json:"product,omitempty"`
	Currency           string       `json:"currency,omitempty"`
	Price              float64      `json:"price,omitempty"`
	StartPrice         float64      `json:"startPrice,omitempty"`
	DiscountPercentage float64      `json:"discountPercentage,omitempty"`
}

// GetPriceRequest describes the request needed for requesting prices.
type GetPriceRequest struct {
	DateFrom     *MMKDateTime `url:"dateFrom,omitempty"`
	DateTo       *MMKDateTime `url:"dateTo,omitempty"`
	CompanyID    []int        `url:"companyId,omitempty"`
	CountryID    []int        `url:"countryId,omitempty"`
	ProductName  string       `url:"productName,omitempty"`
	YachtID      []int64      `url:"yachtId,omitempty"`
	Currency     string       `url:"currency,omitempty"`
	TripDuration int          `url:"tripDuration,omitempty"`
}

// PriceService operates over price requests.
type PriceService service

// GetPrice gets the price that a boat is in for.
func (ps *PriceService) GetPrice(pr *GetPriceRequest) (p *Price, err error) {
	target := fmt.Sprintf("/prices?yachtId=%d&dateFrom=%s&dateTo=%s", pr.YachtID[0], pr.DateFrom.Format("2006-01-02T15:04:05"), pr.DateTo.Format("2006-01-02T15:04:05"))

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
