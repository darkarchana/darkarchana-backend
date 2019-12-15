package controller

import (
	"encoding/json"
	"log"
	"net/http"

	service "github.com/darkarchana/darkarchana-backend/service/serviceimpl"
	"github.com/darkarchana/darkarchana-backend/view"
)

// Chapter : Chapter API
func Chapter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientReq view.ChapterRequest
		var status view.Status

		if r.Method == http.MethodOptions {
			addCorsHeader(w)
			w.WriteHeader(status.Code)
		} else if r.Method == http.MethodGet {
			err := json.NewDecoder(r.Body).Decode(&clientReq)
			if err != nil {
				log.Print(err)
				status.Code = http.StatusBadRequest
				status.Response = err
			}

			switch clientReq.Request {
			case "findPage":
				data, err := service.ChapterServiceImpl().FindPage(clientReq)
				if err != nil {
					log.Print(err)
					status.Code = http.StatusBadRequest
					status.Response = err
				} else {
					status.Code = http.StatusOK
					status.Response = data
				}
			case "findChapter":
				data, err := service.ChapterServiceImpl().FindChapter(clientReq)
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
			addCorsHeader(w)
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

func addCorsHeader(w http.ResponseWriter) {
	w.Header().Add("Connection", "keep-alive")
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	w.Header().Add("Access-Control-Max-Age", "86400")
}
