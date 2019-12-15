package service

import "github.com/darkarchana/darkarchana-backend/view"

// ChapterService : Interface for Chapter Related Service
type ChapterService interface {
	FindPage(view.ChapterRequest) (view.Chapter, error)
	FindChapter(view.ChapterRequest) ([]view.Chapter, error)
}
