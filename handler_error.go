package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request) {
	//write the request
	respondWithError(w, 400, "Something went wrong")
}
