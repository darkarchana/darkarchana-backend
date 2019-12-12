package controller

import (
	"encoding/json"
	"log"
	"net/http"

	service "github.com/darkarchana/darkarchana-backend/service/serviceimpl"
	"github.com/darkarchana/darkarchana-backend/view"
)

// Heroes : Heroes API
func Heroes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientReq view.HeroesRequest
		var status view.Status

		if r.Method == http.MethodGet {
			err := json.NewDecoder(r.Body).Decode(&clientReq)
			if err != nil {
				log.Print(err)
				status.Code = http.StatusBadRequest
				status.Response = err
			}

			switch clientReq.Request {
			case "findOne":
				data, err := service.HeroesServiceImpl().FindOne(clientReq)
				if err != nil {
					log.Print(err)
					status.Code = http.StatusBadRequest
					status.Response = err
				} else {
					status.Code = http.StatusOK
					status.Response = data
				}
			case "findAll":
				data, err := service.HeroesServiceImpl().FindAll(clientReq)
				if err != nil {
					log.Print(err)
					status.Code = http.StatusBadRequest
					status.Response = err
				} else {
					status.Code = http.StatusOK
					status.Response = data
				}
			default:
				status.Code = http.StatusBadRequest
				status.Response = "Request is not Valid"
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status.Code)
			if err != nil {
				log.Print(err)
				status.Response = err
			} else {
				if status.Code != http.StatusBadRequest {
					json.NewEncoder(w).Encode(status.Response)
				} else {
					json.NewEncoder(w).Encode(status)
				}
			}
		}
	}
}
