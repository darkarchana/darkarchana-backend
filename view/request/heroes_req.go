package view

import "github.com/darkarchana/darkarchana-backend/model"

// HeroesRequest : Request Structure for Heroes API
type HeroesRequest struct {
	Request    string         `json:"req"`
	Filter     []model.Heroes `json:"filter"`
	FilterType string         `json:"filterType"`
}
