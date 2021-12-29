package bm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Base is a location for the boat.
type Base struct {
	ID           int64   `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	City         string  `json:"city,omitempty"`
	Country      string  `json:"country,omitempty"`
	Latitude     string  `json:"latitude,omitempty"`
	Longitude    string  `json:"longitude,omitempty"`
	CountryID    int64   `json:"countryId,omitempty"`
	SailingAreas []int64 `json:"sailingAreas,omitempty"`
}

// ValidForBases contains a range of bases
// where any given product or extra is valid as
// departure and destination.
type ValidForBases struct {
	From int64 `json:"from,omitempty"`
	To   int64 `json:"to,omitempty"`
}

type BaseService service

// List retrieves a list of bases that are assigned for each boat.
// Base is the exact location where boat is located.
// One base can belong to more sailing areas. In that case,
// multiple dedicated sailing areas are separated by comma.
//
// See: https://app.swaggerhub.com/apis-docs/mmksystems/bm-api/2.0.2#/Booking/getBases
func (bs *BaseService) List() (bases []*Base, err error) {
	var target string
	{
		target = "bases"
	}

	res, err := bs.get(target)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &bases); err != nil {
		return
	}

	return
}

// Get returns only one base with the specified id.
//
// See: https://app.swaggerhub.com/apis-docs/mmksystems/bm-api/2.0.2#/Booking/getBaseById
func (bs *BaseService) Get(id string) (b *Base, err error) {
	var target string
	{
		target = fmt.Sprintf("base/%s", escapePath(id))
	}

	res, err := bs.get(target)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &b); err != nil {
		return
	}

	return
}

func (cs *BaseService) get(path string) (res *Response, err error) {
	req, err := cs.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return
	}

	return cs.client.Do(req)
}
