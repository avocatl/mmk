package bm

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Reservation represents a booking and its links
// to a yatch and client.
type Reservation struct {
	ID             *int64       `json:"ID,omitempty"`
	DateFrom       *MMKDateTime `json:"dateFrom,omitempty"`
	DateTo         *MMKDateTime `json:"dateTo,omitempty"`
	ExpirationDate *MMKDateTime `json:"ExpirationDate,omitempty"`
	YatchID        int64        `json:"yatchId,omitempty"`
	Status         int64        `json:"status,omitempty"`
	ProductName    string       `json:"productName,omitempty"`
	BaseFromID     int64        `json:"baseFromId,omitempty"`
	BaseToID       int64        `json:"baseToId,omitempty"`
	ClientName     string       `json:"clientName,omitempty"`
	ClientID       int64        `json:"clientId,omitempty"`
	Currency       string       `json:"currency,omitempty"`
}

// ReservationRequest describes the request to query or update reservations
type ReservationRequest struct {
	Year *int   `json:"year,omitempty"`
	ID   *int64 `json:"ID,omitempty"`
}

// ReservationList describes the struct for a list of Reservation
type ReservationList struct {
	Reservations []*Reservation `json:"Reservations,omitempty"`
}

// ReservationService operates over reservation requests.
type ReservationService service

// GetReservation gets a reservation using the reservation id.
func (rsrv *ReservationService) GetReservation(rr *ReservationRequest) (rl *ReservationList, err error) {
	var target string
	if rr.Year != nil {
		target = fmt.Sprintf("reservations/%d", rr.Year)
	} else if rr.ID != nil {
		target = fmt.Sprintf("reservation/%d", rr.ID)
	} else {
		err = errors.New("error processing request, either year or id needs to be set")
		return
	}

	req, err := rsrv.client.NewAPIRequest(http.MethodGet, target, nil)
	if err != nil {
		return
	}

	res, err := rsrv.client.Do(req)
	if err != nil {
		return
	}

	if rr.Year != nil {
		if err = json.Unmarshal(res.content, &rl); err != nil {
			return
		}
	} else {
		var r *Reservation
		if err = json.Unmarshal(res.content, &r); err != nil {
			return
		}
		rl.Reservations = append(rl.Reservations, r)
	}

	return
}

// CreateInfo sends a request to create an info reservation.
func (rsrv *ReservationService) CreateInfo() error {
	return errors.New("creating info is not supported by mmk")
}

// CreateOption sends a request to create an option reservation.
func (rsrv *ReservationService) CreateOption(rr *Reservation) (r *Reservation, err error) {
	target := fmt.Sprintf("reservation")

	req, err := rsrv.client.NewAPIRequest(http.MethodPost, target, rr)
	if err != nil {
		return
	}

	res, err := rsrv.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &r); err != nil {
		return
	}

	return
}

// CreateBooking sends a post request to create a booking reservation.
func (rsrv *ReservationService) CreateBooking(rr *ReservationRequest) (r *Reservation, err error) {
	var target string
	if rr.ID != nil {
		target = fmt.Sprintf("reservation/%d", rr.ID)
	} else {
		err = errors.New("id needs to be set when creating a reservation")
		return
	}

	req, err := rsrv.client.NewAPIRequest(http.MethodPut, target, nil)
	if err != nil {
		return
	}

	res, err := rsrv.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &r); err != nil {
		return
	}

	return
}
