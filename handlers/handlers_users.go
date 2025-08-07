package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/vezzsah/obsec_test/internal/auth"
	"github.com/vezzsah/obsec_test/internal/database"
	"github.com/vezzsah/obsec_test/utils"
)

func (cfg *ApiConfig) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	param := CreateUserParams{}
	err := decoder.Decode(&param)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest, "incorrect input")
		return
	}

	found, err := cfg.DbQueries.CheckIfUserExistByEmail(r.Context(), param.Email)
	if err != nil || found == 1 {
		utils.RespondWithError(w, err, http.StatusConflict, "email already in use")
		return
	}

	param.Password, err = auth.HashPassword(param.Password)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError, "error while creating user")
		return
	}

	dbparams := database.CreateUserParams{
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
		Email:     param.Email,
		Hashedp:   param.Password,
	}
	log.Printf("dbParams: %v", dbparams)
	createdUser, err := cfg.DbQueries.CreateUser(r.Context(), dbparams)
	if err != nil {
		log.Printf("user: %v", createdUser)
		log.Printf("time: %s", time.Now().UTC().Format(time.RFC3339))
		utils.RespondWithError(w, err, http.StatusInternalServerError, "something went wrong while storing user")
		return
	}

	utils.WriteJSONToResponse(w, http.StatusCreated, createdUser)
}

func (cfg *ApiConfig) LogInUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	param := LogInUserParams{}
	err := decoder.Decode(&param)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusBadRequest, "incorrect input")
		return
	}

	user, err := cfg.DbQueries.GetUserByEmail(r.Context(), param.Email)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusUnauthorized, "unauthorized")
		return
	}

	err = auth.CheckPasswordHash(param.Password, user.Hashedp)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusUnauthorized, "unauthorized")
		return
	}

	jwtToken, err := auth.MakeJWT(*user.ID)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError, "something went wrong")
		return
	}

	userWithoutPass := struct {
		ID        string `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Email     string `json:"email"`
		Token     string `json:"token"`
	}{
		ID:        *user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Token:     jwtToken,
	}
	utils.WriteJSONToResponse(w, http.StatusOK, userWithoutPass)
}

func (cfg *ApiConfig) ResetUsers(w http.ResponseWriter, r *http.Request) {
	if cfg.Env != "dev" {
		utils.RespondWithError(w, nil, http.StatusForbidden, "forbidden")
		return
	}

	err := cfg.DbQueries.DeleteAllUsers(r.Context())
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError, "something went wrong while deleting users")
		return
	}
}
