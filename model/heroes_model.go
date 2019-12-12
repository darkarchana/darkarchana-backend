package model

//Heroes : Model for Heroes
type Heroes struct {
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Signed bool   `json:"signed"`
}
