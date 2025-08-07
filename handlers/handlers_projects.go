package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vezzsah/obsec_test/internal/database"
	"github.com/vezzsah/obsec_test/utils"
)

func (cfg *ApiConfig) CreateProject(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {
	params := CreateProjectParams{}
	err := utils.ParseBody(w, r, &params)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError, "could not decode parameters")
		return
	}

	if params.Name == "" {
		utils.RespondWithError(w, nil, http.StatusBadRequest, "name required")
		return
	}

	found, err := cfg.DbQueries.CheckIfProjectExistByName(r.Context(), params.Name)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError, "error while storing project")
		return
	}
	if found {
		utils.RespondWithError(w, err, http.StatusConflict, "project name already in use")
		return
	}

	proj, err := cfg.DbQueries.CreateProject(r.Context(), database.CreateProjectParams{
		ProjectName: params.Name,
		CreatedAt:   time.Now().UTC().Format(time.RFC3339),
		UpdatedAt:   time.Now().UTC().Format(time.RFC3339),
		Creator:     user_id,
	})
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError, "error while storing project")
	}

	utils.WriteJSONToResponse(w, http.StatusCreated, proj)
}

func (cfg *ApiConfig) ViewProject(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {
	req_name := r.URL.Query().Get("project_name")
	if req_name != "" {
		res, err := cfg.DbQueries.GetProjectByNameAndCreator(r.Context(), database.GetProjectByNameAndCreatorParams{
			ProjectName: req_name,
			Creator:     user_id,
		})
		if err != nil {
			utils.RespondWithError(w, err, http.StatusNotFound, "not found")
			return
		}
		utils.WriteJSONToResponse(w, http.StatusOK, res)
	} else {
		res, err := cfg.DbQueries.GetProjectsByUser(r.Context(), user_id)
		if err != nil {
			utils.RespondWithError(w, err, http.StatusInternalServerError, "internal server error")
			return
		}
		utils.WriteJSONToResponse(w, http.StatusOK, res)
	}
}
