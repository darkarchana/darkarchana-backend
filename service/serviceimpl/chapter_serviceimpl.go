package service

import (
	"errors"
	"log"

	dao "github.com/darkarchana/darkarchana-backend/dao/daoimpl"
	"github.com/darkarchana/darkarchana-backend/model"
	"github.com/darkarchana/darkarchana-backend/service"
	"github.com/darkarchana/darkarchana-backend/view"
	"go.mongodb.org/mongo-driver/bson"
)

type chapterImpl struct{}

// FindPage : Override Method FindPage on ChapterService Interface
func (implementation *chapterImpl) FindPage(chapterReq view.ChapterRequest) (view.Chapter, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	filter, err := filterChapterQuery(chapterReq)
	if err != nil {
		return view.Chapter{}, err
	}
	dbOperate := model.DbOperate{
		Collection: chapterReq.Title,
		Option:     model.DbOption{},
		Filter:     filter,
	}
	data, err := dao.ChapterDaoImpl().FindPage(dbOperate)
	view := modelToView(data)
	return view, err

}

// FindChapter: Override Method FindChapter on ChapterService Interface
func (implementation *chapterImpl) FindChapter(chapterReq view.ChapterRequest) ([]view.Chapter, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	filter, err := filterChapterQuery(chapterReq)
	if err != nil {
		return []view.Chapter{}, err
	}
	dbOperate := model.DbOperate{
		Collection: chapterReq.Title,
		Option:     model.DbOption{},
		Filter:     filter,
	}
	data, err := dao.ChapterDaoImpl().FindChapter(dbOperate)
	views := modelsToViews(data)
	return views, err

}

func filterChapterQuery(chapterReq view.ChapterRequest) (bson.M, error) {
	filter := bson.M{}
	if chapterReq.Chapter.Chapter > 0 {
		filter["chapter"] = chapterReq.Chapter.Chapter
	} else {
		return nil, errors.New("No Chapter Selected")
	}

	if chapterReq.Chapter.Page > 0 {
		filter["page"] = chapterReq.Chapter.Page
	}
	return filter, nil
}

func modelsToViews(models []model.Chapter) []view.Chapter {
	views := []view.Chapter{}
	for _, v := range models {
		views = append(views, modelToView(v))
	}
	return views
}

func modelToView(model model.Chapter) view.Chapter {
	view := view.Chapter{
		Chapter: model.Chapter,
		Page:    model.Page,
		Link:    model.Link,
	}
	return view
}

// ChapterServiceImpl : Implementation of Interface ChapterService
func ChapterServiceImpl() service.ChapterService {
	var service service.ChapterService = &chapterImpl{}
	return service
}
