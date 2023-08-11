package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Internal server error ", msg)
	}
	type errorRes struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorRes{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//makes payload into jason string
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON res %v with error %v", payload, err)
		w.WriteHeader((500))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat) //pass in json data
}
