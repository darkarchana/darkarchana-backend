package view

import "github.com/darkarchana/darkarchana-backend/view/view"

// ChapterRequest : Request Structure for Chapter API
type ChapterRequest struct {
	Request string       `json:"req"`
	Title   string       `json:"title"`
	Chapter view.Chapter `json:"chapter"`
}
