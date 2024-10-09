package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Couldn't parse json: %v", payload)
		log.Println(err)
		w.WriteHeader(500)
	}
	w.WriteHeader(statusCode)
	w.Write(data)

}
