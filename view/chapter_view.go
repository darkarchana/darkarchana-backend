package view

//Chapter : View for Chapter
type Chapter struct {
	Chapter int16 `json:"chapter"`
	Page    int16 `json:"page"`
	Link    bool  `json:"link"`
}
