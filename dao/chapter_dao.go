package dao

import "github.com/darkarchana/darkarchana-backend/model"

// ChapterDao : Interface for Chapter DAO
type ChapterDao interface {
	FindPage(model.DbOperate) (model.Chapter, error)
	FindChapter(model.DbOperate) ([]model.Chapter, error)
}
