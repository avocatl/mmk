package bm

// Company represents a booking manager organizational user.
type Company struct {
	ID                 int64  `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	Address            string `json:"address,omitempty"`
	City               string `json:"city,omitempty"`
	ZipCode            string `json:"zip,omitempty"`
	Country            string `json:"country,omitempty"`
	Telephone          string `json:"telephone,omitempty"`
	Telephone2         string `json:"telephone2,omitempty"`
	Mobile             string `json:"mobile,omitempty"`
	VATCode            string `json:"vatCode,omitempty"`
	Email              string `json:"email,omitempty"`
	Web                string `json:"web,omitempty"`
	BankAccountNumber  string `json:"bankAccountNumber,omitempty"`
	TermsAndConditions string `json:"termsAndConditions,omitempty"`
}
