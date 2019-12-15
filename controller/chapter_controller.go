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

		if r.Method == http.MethodGet {
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
			w.Header().Set("Access-Control-Allow-Origin", "*")
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
