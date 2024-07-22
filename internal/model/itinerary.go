package model

type Itinerary struct {
	Summary   string `json:"summary"`
	TotalCost string `json:"total_cost"`
	Days      []Day  `json:"days"`
}
