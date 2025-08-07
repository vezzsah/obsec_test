package handlers

import (
	"net/http"

	"github.com/vezzsah/obsec_test/internal/database"
)

type ApiConfig struct {
	DbQueries  *database.Queries
	HttpClient *http.Client
	Env        string
}
