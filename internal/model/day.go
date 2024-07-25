package model

type Day struct {
	Summary   string      `json:"summary"`
	TotalCost string      `json:"total_cost"`
	Route     []RouteItem `json:"route"`
}
