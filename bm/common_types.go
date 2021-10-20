package bm

import (
	"fmt"
	"net/http"
)

// ProductEnum is a type alias for supported mmk products.
type ProductEnum string

const (
	BareBoat ProductEnum = "bareboat"
	Crewed   ProductEnum = "crewed"
	Cabin    ProductEnum = "cabin"
	Flotilla ProductEnum = "flotilla"
	Power    ProductEnum = "power"
	Berth    ProductEnum = "breth"
	Regatta  ProductEnum = "regatta"
)

// ItemTypeEnum is a type alias for supported items on mmk.
type ItemTypeEnum string

const (
	ModelItem       ItemTypeEnum = "model"
	ReservationItem ItemTypeEnum = "reservation"
	UserItem        ItemTypeEnum = "item"
	YatchItem       ItemTypeEnum = "yatch"
)

// VAT calculation values.
type VAT struct {
	Base     float64
	Rate     float64
	Total    float64
	BaseAlt  float64
	TotalAlt float64
}

// Bi is something I yet have to understand.
type Bi struct {
	Name  string  `json:"name,omitempty"`
	Value float64 `json:"value,omitempty"`
}

/*
Error reports details on a failed API request.
*/
type Error struct {
	Code     int            `json:"code"`
	Message  string         `json:"message"`
	Content  string         `json:"content,omitempty"`
	Response *http.Response `json:"response"` // the full response that produced the error
}

// Error function complies with the error interface
func (e *Error) Error() string {
	return fmt.Sprintf("%v:\n%v", e.Message, e.Content)
}

// Response represents an API response.
//
// This wraps the standard http.Response returned from MMK
// and provides convenient access to things the decoded body.
type Response struct {
	*http.Response
	content []byte
}
