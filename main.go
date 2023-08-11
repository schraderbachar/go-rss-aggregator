package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/schraderbachar/rss-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port not in env")
	}

	dbURL := os.Getenv("DB_URL")
	if portString == "" {
		log.Fatal("DB_URL not in env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to databse")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	//* connecting to server
	router := chi.NewRouter()

	//people can make requests from a browser to our server

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness) //path to see if server is live and running
	v1Router.Get("/err", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser) //route to create user

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe() //code stops here and handles http reqs. if it goes wrong, then the log logs the error

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PORT: ", portString)
}
