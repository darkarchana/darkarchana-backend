package service

import (
	"github.com/darkarchana/darkarchana-backend/model"
	"github.com/darkarchana/darkarchana-backend/view"
)

// HeroesService : Interface for Heroes Related Service
type HeroesService interface {
	FindOne(view.HeroesRequest) (model.Heroes, error)
	FindAll(view.HeroesRequest) ([]model.Heroes, error)
}
