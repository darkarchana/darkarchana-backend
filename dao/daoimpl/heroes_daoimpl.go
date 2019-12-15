package dao

import (
	"context"
	"log"

	"github.com/darkarchana/darkarchana-backend/dao"
	"github.com/darkarchana/darkarchana-backend/database"
	"github.com/darkarchana/darkarchana-backend/model"
	"go.mongodb.org/mongo-driver/bson"
)

type heroesImpl struct{}

// FindOne : Override Method FindOne on HeroesDaoInterface
func (implementation *heroesImpl) FindOne(dbOperate model.DbOperate) (model.Heroes, error) {
	// Connection of Database
	database.MongoDbConnect()
	defer database.MongoDbDisconnect()

	var result model.Heroes
	err := database.MongoDbFindOne(dbOperate).Decode(&result)
	if err != nil {
		log.Print(err)
		return result, err
	}
	return result, nil
}

// FindAll : Override Method FindAll on HeroesDaoInterface
func (implementation *heroesImpl) FindAll(dbOperate model.DbOperate) ([]model.Heroes, error) {
	// Connection of Database
	database.MongoDbConnect()
	defer database.MongoDbDisconnect()

	var results []model.Heroes
	cur, err := database.MongoDbFind(dbOperate)
	if err != nil {
		log.Print(err)
	}

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		data := &bson.D{}
		err := cur.Decode(&data)
		if err != nil {
			log.Print(err)
		}
		dataMapping := data.Map()
		results = append(results, model.Heroes{
			Name:   dataMapping["name"].(string),
			Alias:  dataMapping["alias"].(string),
			Signed: dataMapping["signed"].(bool),
		})
	}

	if err := cur.Err(); err != nil {
		log.Print(err)
	}

	return results, err
}

// HeroesDaoImpl : Implementation of Interface HeroesDao
func HeroesDaoImpl() dao.HeroesDao {
	var dao dao.HeroesDao = &heroesImpl{}
	return dao
}
