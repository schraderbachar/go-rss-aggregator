package main

import (
	"fmt"
	"net/http"

	"github.com/schraderbachar/rss-aggregator/internal/auth"
	"github.com/schraderbachar/rss-aggregator/internal/database"
)

// our own custome type to handle the grabbing api key and the assocaited user
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// takes the api config so it can use the database and then it returns handler func so we can use it with the chi router
func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	//anon function here so we can query the databse
	return func(w http.ResponseWriter, r *http.Request) {
		//authenticated enpodint (aka logged in) gonna make it into a package
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, "Could not authenticate")
		}

		//user db query
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %s", err))
		}

		handler(w, r, user)
	}
}
