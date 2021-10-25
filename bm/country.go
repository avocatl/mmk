package bm

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Country describes a single country in Booking Manager.
// It is a standard country entry with standard ISO elements
// for name, shortCode, longCode and numerical country id.
// worldRegion field links to an MMK defined world region
// available via /worldRegions call.
type Country struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Short       string `json:"short,omitempty"`
	Long        string `json:"long,omitempty"`
	WorldRegion int64  `json:"worldRegion,omitempty"`
}

// CountryService handles operations with MMK countries registry.
type CountryService service

// List retrieves a list of all ISO-3166 countries
// and their short, long codes and world region code.
func (cs *CountryService) List() (countries []*Country, err error) {
	var target string
	{
		target = "countries"
	}

	res, err := cs.get(target)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &countries); err != nil {
		return
	}

	return
}

// Get retrieves only one country by id.
func (cs *CountryService) Get(id string) (c *Country, err error) {
	var target string
	{
		target = fmt.Sprintf("country/%s", scapeQueryParams(id))
	}

	res, err := cs.get(target)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &c); err != nil {
		return
	}

	return
}

func (cs *CountryService) get(path string) (res *Response, err error) {
	req, err := cs.client.NewAPIRequest(http.MethodGet, path, nil)
	log.Print(req.URL.String())
	if err != nil {
		return
	}

	return cs.client.Do(req)
}
