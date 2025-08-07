package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/vezzsah/obsec_test/internal/database"
	"github.com/vezzsah/obsec_test/utils"
)

func (cfg *ApiConfig) GetProjectCVEs(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {
	req_name := r.URL.Query().Get("project_name")
	if req_name == "" {
		utils.RespondWithError(w, nil, http.StatusBadRequest, "requires project_name")
	}

	proj, err := cfg.DbQueries.GetProjectByNameAndCreator(r.Context(), database.GetProjectByNameAndCreatorParams{
		ProjectName: req_name,
		Creator:     user_id,
	})
	if err != nil {
		utils.RespondWithError(w, err, http.StatusNotFound, "not found")
		return
	}

	res, err := cfg.DbQueries.GetAllCVEByProject(r.Context(), proj.ID)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError, "error while fetching CVEs")
	}

	utils.WriteJSONToResponse(w, http.StatusOK, res)
}

func (cfg *ApiConfig) ResolveProjectCVE(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {
	params := ResolveProjectCVEParams{}
	err := utils.ParseBody(w, r, &params)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError, "could not decode parameters")
		return
	}

	dbRes, err := cfg.DbQueries.GetProjectAndCPE(r.Context(), database.GetProjectAndCPEParams{
		ProjectName: params.ProjectName,
		Creator:     user_id,
		Cpe:         params.CPE,
		Cve:         params.CVE,
	})
	if err != nil {
		utils.RespondWithError(w, err, http.StatusNotFound, "not found")
		return
	}

	err = cfg.DbQueries.UpdateCVE(r.Context(), database.UpdateCVEParams{
		Solved:  true,
		Cve:     dbRes.Cvestring.String,
		Project: dbRes.Projectid,
		Cpe:     dbRes.Cpeid.String,
	})
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError, "error while updating CVE")
		return
	}

	utils.WriteJSONToResponse(w, http.StatusOK, nil)
}
