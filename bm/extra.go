package bm

// Extra describes additional or complementary items
// for items offered as products on MMK.
type Extra struct {
	Name                  string           `json:"name,omitempty"`
	Obligatory            bool             `json:"obligatory,omitempty"`
	Price                 float64          `json:"price,omitempty"`
	Unit                  string           `json:"unit,omitempty"`
	PayableInBase         bool             `json:"payableInBase,omitempty"`
	IncludedDepositWaiver bool             `json:"includedDepositWaiver,omitempty"`
	ValidDaysFrom         int64            `json:"validDaysFrom,omitempty"`
	ValidDaysTo           int64            `json:"validDaysTo,omitempty"`
	MinNumberOfPersons    int64            `json:"minNumberOfPersons,omitempty"`
	ManNumberOfPersons    int64            `json:"manNumberOfPersons,omitempty"`
	ValidForBases         []*ValidForBases `json:"validForBases,omitempty"`
}
