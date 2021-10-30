package bm

// Service is an addition to any reservation that is not tangible.
type Service struct {
	Nme   string  `json:"nme,omitempty"`
	Total float64 `json:"total,omitempty"`
	ID    int64   `json:"id,omitempty"`
	Rate  float64 `json:"rate,omitempty"`
	Code  string  `json:"code,omitempty"`
}
