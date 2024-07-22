package model

type RouteItem struct {
	Duration     string `json:"duration"`
	Price        string `json:"price"`
	Location     string `json:"location"`
	WhatToExpect string `json:"what_to_expect"`
	Transport    string `json:"transport"`
}
