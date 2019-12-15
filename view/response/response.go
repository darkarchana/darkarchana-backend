package view

// Status : Response to Clients
type Status struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}
