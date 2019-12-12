package service

import (
	dao "github.com/darkarchana/darkarchana-backend/dao/daoimpl"
	"github.com/darkarchana/darkarchana-backend/model"
	"github.com/darkarchana/darkarchana-backend/service"
	"github.com/darkarchana/darkarchana-backend/view"
	"go.mongodb.org/mongo-driver/bson"
)

type impl struct{}

// FindOne : Override Method FindOne on HeroesService Interface
func (implementation *impl) FindOne(heroesReq view.HeroesRequest) (model.Heroes, error) {
	filter := filterQuery(heroesReq)
	dbOperate := model.DbOperate{
		Collection: "test",
		Option:     model.DbOption{},
		Filter:     filter,
	}
	data, err := dao.HeroesDaoImpl().FindOne(dbOperate)
	return data, err
}

// FindAll : Override Method FindAll on HeroesService Interface
func (implementation *impl) FindAll(heroesReq view.HeroesRequest) ([]model.Heroes, error) {
	filter := filterQuery(heroesReq)
	dbOperate := model.DbOperate{
		Collection: "test",
		Option:     model.DbOption{},
		Filter:     filter,
	}
	data, err := dao.HeroesDaoImpl().FindAll(dbOperate)
	return data, err
}

func filterQuery(heroesReq view.HeroesRequest) bson.M {
	filter := bson.M{}
	switch heroesReq.FilterType {
	case "name":
		filter["$or"] = []bson.M{}
		for _, s := range heroesReq.Filter {
			filter["$or"] = append(filter["$or"].([]bson.M), bson.M{"name": s.Name})
		}
	}
	return filter
}

// HeroesServiceImpl : Implementation of Interface HeroesService
func HeroesServiceImpl() service.HeroesService {
	var service service.HeroesService = &impl{}
	return service
}
