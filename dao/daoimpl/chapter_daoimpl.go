package dao

import (
	"context"
	"log"

	"github.com/darkarchana/darkarchana-backend/dao"
	"github.com/darkarchana/darkarchana-backend/database"
	"github.com/darkarchana/darkarchana-backend/model"
	"go.mongodb.org/mongo-driver/bson"
)

type chapterImpl struct{}

// FindPage : Override Method FindPage on ChapterDaoInterface
func (implementation *chapterImpl) FindPage(dbOperate model.DbOperate) (model.Chapter, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Connection of Database
	database.MongoDbConnect()
	defer database.MongoDbDisconnect()

	var result model.Chapter
	err := database.MongoDbFindOne(dbOperate).Decode(&result)
	if err != nil {
		log.Print(err)
		return result, err
	}
	return result, nil
}

// FindChapter : Override Method FindChapter on ChapterDaoInterface
func (implementation *chapterImpl) FindChapter(dbOperate model.DbOperate) ([]model.Chapter, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Connection of Database
	database.MongoDbConnect()
	defer database.MongoDbDisconnect()

	var results []model.Chapter
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
		results = append(results, model.Chapter{
			Chapter: dataMapping["chapter"].(int32),
			Page:    dataMapping["page"].(int32),
			Link:    dataMapping["link"].(string),
		})
	}

	if err := cur.Err(); err != nil {
		log.Print(err)
	}

	return results, err
}

// ChapterDaoImpl : Implementation of Interface ChapterDao
func ChapterDaoImpl() dao.ChapterDao {
	var dao dao.ChapterDao = &chapterImpl{}
	return dao
}
