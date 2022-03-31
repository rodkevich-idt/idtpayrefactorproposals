package respond

import (
	"encoding/json"
	"log"
	"net/http"

	"gateway/internal/utility/message"
)

type Standard struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta,omitempty"`
}

type Meta struct {
	Size  int   `json:"size"`
	Total int64 `json:"total"`
}

func Json(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.WriteHeader(statusCode)

	if payload == nil {
		return
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		Error(w, http.StatusInternalServerError, message.ErrInternalError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
		Error(w, http.StatusInternalServerError, message.ErrInternalError)
		return
	}
}

func Status(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}
