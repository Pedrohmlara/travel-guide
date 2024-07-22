package model

type ItineraryRequest struct {
	Location              string `json:"location"`
	BudgetInLocalCurrency string `json:"budget_in_local_currency"`
	Days                  string `json:"days"`
	Avoid                 string `json:"avoid"`
}
