package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	//write the request
	respondWithJSON(w, 200, struct{}{}) //all we care about in this case is the res code so return empty
}
