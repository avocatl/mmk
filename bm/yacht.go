package bm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Yacht describes an offered YachtType product.
type Yacht struct {
	ID                   int64      `json:"id,omitempty"`
	ModelID              int64      `json:"modelID,omitempty"`
	ShipyardID           int64      `json:"shipyardID,omitempty"`
	CompanyID            int64      `json:"companyID,omitempty"`
	WC                   int64      `json:"wc,omitempty"`
	Berths               int64      `json:"berths,omitempty"`
	Cabins               int64      `json:"cabins,omitempty"`
	DefaultCheckInDay    int64      `json:"defaultCheckInDay,omitempty"`
	EquipmentIDs         []int64    `json:"equipmentIDs,omitempty"`
	Draught              float64    `json:"draught,omitempty"`
	Beam                 float64    `json:"beam,omitempty"`
	Length               float64    `json:"length,omitempty"`
	WaterCapacity        float64    `json:"waterCapacity,omitempty"`
	FuelCapacity         float64    `json:"fuelCapacity,omitempty"`
	Deposit              float64    `json:"deposit,omitempty"`
	CommissionPercentage float64    `json:"commissionPercentage,omitempty"`
	MainsailArea         float64    `json:"mainsailArea,omitempty"`
	GenoaArea            float64    `json:"genoaArea,omitempty"`
	Name                 string     `json:"name,omitempty"`
	Model                string     `json:"model,omitempty"`
	Year                 string     `json:"year,omitempty"`
	Kind                 string     `json:"kind,omitempty"`
	HomeBase             string     `json:"homeBase,omitempty"`
	Company              string     `json:"company,omitempty"`
	Engine               string     `json:"engine,omitempty"`
	MainsailType         string     `json:"mainsailType,omitempty"`
	GenoaType            string     `json:"genoaType,omitempty"`
	Images               []*Image   `json:"images,omitempty"`
	Products             []*Product `json:"products,omitempty"`
	CabinLayout          []*Layout  `json:"cabinLayout,omitempty"`
	BerthsLayout         []*Layout  `json:"berthsLayout,omitempty"`
}

type YachtService service

// GetYacht
func (os *YachtService) GetYacht(id int) (or *Yacht, err error) {
	var target string
	{
		target = fmt.Sprintf("yacht/%d", id)
	}

	req, err := os.client.NewAPIRequest(http.MethodGet, target, nil)
	if err != nil {
		return
	}

	res, err := os.client.Do(req)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &or); err != nil {
		return
	}

	return
}
