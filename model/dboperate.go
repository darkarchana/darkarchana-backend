package model

//DbOption : Model for Database Operation Options
type DbOption struct {
	FindLimit int64
}

//DbOperate : Model for Database Operation
type DbOperate struct {
	Collection string
	Option     DbOption
	Filter     interface{}
	Data       interface{}
}
