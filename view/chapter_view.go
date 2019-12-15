package view

//Chapter : View for Chapter
type Chapter struct {
	Chapter int32  `json:"chapter"`
	Page    int32  `json:"page"`
	Link    string `json:"link"`
}
