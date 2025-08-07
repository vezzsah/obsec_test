package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/vezzsah/obsec_test/handlers"
	"github.com/vezzsah/obsec_test/internal/database"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("warning: assuming default configuration. .env unreadable: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	apiCfg := handlers.ApiConfig{
		HttpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}

	apiCfg.Env = os.Getenv("ENVIRONMENT")
	if apiCfg.Env == "" {
		log.Fatal("No Environment Set")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("DATABASE_URL environment variable is not set")
		log.Println("Running without CRUD endpoints")
	} else {
		db, err := sql.Open("libsql", dbURL)
		if err != nil {
			log.Fatal(err)
		}
		apiCfg.DbQueries = database.New(db)
		log.Println("Connected to database")
	}

	mux := http.NewServeMux()

	if apiCfg.DbQueries != nil {
		/*Users Management*/
		mux.HandleFunc("DELETE /api/users", apiCfg.ResetUsers)
		mux.HandleFunc("POST /api/users", apiCfg.CreateNewUser)
		mux.HandleFunc("POST /api/login", apiCfg.LogInUser)

		/*Project Management*/
		mux.HandleFunc("POST /api/projects", apiCfg.MiddlewareAuth(apiCfg.CreateProject))
		mux.HandleFunc("GET /api/projects", apiCfg.MiddlewareAuth(apiCfg.ViewProject))

		/*CPE per Project, Management*/
		mux.HandleFunc("POST /api/projects/cpes", apiCfg.MiddlewareAuth(apiCfg.RegisterCPE))
		mux.HandleFunc("GET /api/projects/cpes", apiCfg.MiddlewareAuth(apiCfg.GetProjectCPEs))

		/*TODO: CPE Search*/
		//mux.HandleFunc("GET /api/cpes", apiCfg.MiddlewareAuth(apiCfg.SearchCPE))

		/*CVE per Project Management*/
		mux.HandleFunc("GET /api/projects/cve", apiCfg.MiddlewareAuth(apiCfg.GetProjectCVEs))
		mux.HandleFunc("POST /api/projects/cve", apiCfg.MiddlewareAuth(apiCfg.ResolveProjectCVE))
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: time.Minute,
	}

	log.Println("Ready")
	err = server.ListenAndServe()
	if err == http.ErrServerClosed {
		return
	}
}
