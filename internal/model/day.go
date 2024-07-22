package model

type Day struct {
	Date      string      `json:"date"`
	Day       string      `json:"day"`
	Summary   string      `json:"summary"`
	TotalCost string      `json:"total_cost"`
	Route     []RouteItem `json:"route"`
}
