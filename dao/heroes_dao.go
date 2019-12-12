package dao

import "github.com/darkarchana/darkarchana-backend/model"

// HeroesDao : Interface for Heroes DAO
type HeroesDao interface {
	FindOne(model.DbOperate) (model.Heroes, error)
	FindAll(model.DbOperate) ([]model.Heroes, error)
}
