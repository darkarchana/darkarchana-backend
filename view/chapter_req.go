package view

// ChapterRequest : Request Structure for Chapter API
type ChapterRequest struct {
	Request string  `json:"req"`
	Title   string  `json:"title"`
	Chapter Chapter `json:"chapter"`
}
