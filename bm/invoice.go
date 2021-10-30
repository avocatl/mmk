package bm

import "time"

type Invoice struct {
	InvoiceType             int64      `json:"type,omitempty"`
	InvoiceCode             string     `json:"invoiceCode,omitempty"`
	ReservationNumber       string     `json:"reservationNumber,omitempty"`
	InvoiceDate             *time.Time `json:"invoiceDate,omitempty"`
	Client                  string     `json:"client,omitempty"`
	ClientCode              string     `json:"clientCode,omitempty"`
	ClientVATCode           string     `json:"clientVATCode,omitempty"`
	ClientID                int64      `json:"clientId,omitempty"`
	GuestName               string     `json:"guestName,omitempty"`
	Currency                string     `json:"currency,omitempty"`
	ExchangeRate            float64    `json:"exchangeRate,omitempty"`
	AltCurrency             string     `json:"altCurrency,omitempty"`
	AltExchangeRate         float64    `json:"altExchangeRate,omitempty"`
	Resource                string     `json:"resource,omitempty"`
	ResourceType            string     `json:"resourceType,omitempty"`
	ResourceCode            int64      `json:"resourceCode,omitempty"`
	TotalPrice              float64    `json:"totalPrice,omitempty"`
	TotalPriceWithoutTax    float64    `json:"totalPriceWithoutTax,omitempty"`
	Rate                    float64    `json:"rate,omitempty"`
	TotalAltPrice           float64    `json:"totalAltPrice,omitempty"`
	TotalAltPriceWithoutTax float64    `json:"totalAltPriceWithoutTax,omitempty"`
	BaseFrom                string     `json:"baseFrom,omitempty"`
	BaseTo                  string     `json:"baseTo,omitempty"`
	AlreadyTransferred      bool       `json:"alreadyTransferred,omitempty"`
	ServiceDateFrom         *time.Time `json:"serviceDateFrom,omitempty"`
	ServiceDateTo           *time.Time `json:"serviceDateTo,omitempty"`
	PaymentMethodName       string     `json:"paymentMethodName,omitempty"`
	PaymentMethodType       string     `json:"paymentMethodType,omitempty"`
	AgencyID                int64      `json:"agencyId,omitempty"`
	AgencyName              string     `json:"agencyName,omitempty"`
	AgencyCode              string     `json:"agencyCode,omitempty"`
	AgencyVATCode           string     `json:"agencyVatCode,omitempty"`
	RelatedReservationID    int64      `json:"relatedReservationID,omitempty"`
	ReservationID           int64      `json:"reservationId,omitempty"`
	RelatedInvoiceNumber    string     `json:"relatedInvoiceNumber,omitempty"`
	VAT                     VAT        `json:"vat,omitempty"`
	Services                []*Service `json:"services,omitempty"`
	BI                      []*Bi      `json:"bi,omitempty"`
}
