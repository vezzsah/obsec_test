package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	if apiCfg.DbQueries != nil {
		v1Router.Delete("/users", apiCfg.ResetUsers)
		v1Router.Post("/users", apiCfg.CreateNewUser)
		v1Router.Post("/login", apiCfg.LogInUser)

		v1Router.Post("/projects", apiCfg.MiddlewareAuth(apiCfg.CreateProject))
		v1Router.Get("/projects", apiCfg.MiddlewareAuth(apiCfg.ViewProject))

		v1Router.Post("/projects/cpes", apiCfg.MiddlewareAuth(apiCfg.RegisterCPE))
		v1Router.Get("/projects/cpes", apiCfg.MiddlewareAuth(apiCfg.GetProjectCPEs))

		v1Router.Post("/projects/cves", apiCfg.MiddlewareAuth(apiCfg.ResolveProjectCVE))
		v1Router.Get("/projects/cves", apiCfg.MiddlewareAuth(apiCfg.GetProjectCVEs))
	}

	v1Router.Get("/healthz", handlers.HandlerReadiness)

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: time.Minute,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
