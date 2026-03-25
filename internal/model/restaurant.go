package model

type Restaurant struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Cuisine  string  `json:"cuisine"`
	Rating   float64 `json:"rating"`
	IsOpen   bool    `json:"is_open"`
	ETAInMin int     `json:"eta_in_min"`
}