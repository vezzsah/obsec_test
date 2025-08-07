package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vezzsah/obsec_test/internal/database"
	"github.com/vezzsah/obsec_test/nistlayer"
	"github.com/vezzsah/obsec_test/utils"
)

func (cfg *ApiConfig) GetProjectCPEs(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {
	req_name := r.URL.Query().Get("project_name")
	if req_name != "" {
		proj, err := cfg.DbQueries.GetProjectByNameAndCreator(r.Context(), database.GetProjectByNameAndCreatorParams{
			ProjectName: req_name,
			Creator:     user_id,
		})
		if err != nil {
			utils.RespondWithError(w, err, http.StatusNotFound, "not found")
			return
		}

		cves, err := cfg.DbQueries.GetAllCPEByProject(r.Context(), proj.ID)
		if err != nil {
			utils.RespondWithError(w, err, http.StatusInternalServerError, "internal error")
		}
		utils.WriteJSONToResponse(w, http.StatusOK, cves)
	}

	utils.RespondWithError(w, nil, http.StatusNotFound, "not found")
}

// cpe example: cpe:2.3:o:microsoft:windows_10:1511
// cpe format= cpe:2.3: {part} : {vendor} : {product} : {version} : {update} : {edition} : {language}
// {part} can be h = hardware, o = operating system, a = application
func (cfg *ApiConfig) RegisterCPE(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {
	params := RegisterCPEParams{}
	err := utils.ParseBody(w, r, &params)
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError, "could not decode parameters")
		return
	}

	found, err := cfg.DbQueries.CheckIfProjectExistByUserIdAndName(r.Context(), database.CheckIfProjectExistByUserIdAndNameParams{
		Creator:     user_id,
		ProjectName: params.ProjectName,
	})
	if err != nil {
		utils.RespondWithError(w, err, http.StatusInternalServerError, "problem while accessing DB")
		return
	}
	if !found {
		utils.RespondWithError(w, err, http.StatusUnauthorized, "unauthorized")
		return
	}

	if len(params.CPEData) == 0 {
		utils.RespondWithError(w, nil, http.StatusBadRequest, "at least one CPE is required")
		return
	}

	project, err := cfg.DbQueries.GetProjectByNameAndCreator(r.Context(), database.GetProjectByNameAndCreatorParams{
		ProjectName: params.ProjectName,
		Creator:     user_id,
	})
	if err != nil {
		utils.RespondWithError(w, err, http.StatusNotFound, "project not found")
		return
	}

	//First check if all CPEs are valid before writing anything to DB.
	for _, cpe := range params.CPEData {
		if cpe.Part == "" || cpe.Vendor == "" || cpe.Product == "" || cpe.Version == "" {
			utils.RespondWithError(w, nil, http.StatusBadRequest, "part, vendor, product and version of each cpe are required")
			return
		}
		cpeString := GenerateCPEString(cpe)
		if !nistlayer.ValidateIfCPEExists(cpeString) {
			utils.RespondWithError(w, nil, http.StatusBadRequest, "inexistent cpe")
			return
		}
		found, err = cfg.DbQueries.CheckIfCPEExistByProjectName(r.Context(), database.CheckIfCPEExistByProjectNameParams{
			Cpe:       cpeString,
			ProjectID: project.ID,
		})
		if err != nil {
			utils.RespondWithError(w, err, http.StatusInternalServerError, "error while cheking validity project")
			return
		}
		if found {
			utils.RespondWithError(w, err, http.StatusConflict, "project name already in use")
			return
		}
	}

	//Store CPEs and CVE in DB
	for _, cpe := range params.CPEData {
		cpestring := GenerateCPEString(cpe)
		storedCpe, err := cfg.DbQueries.StoreCPE(r.Context(), database.StoreCPEParams{
			Cpe:       cpestring,
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
			UpdatedAt: time.Now().UTC().Format(time.RFC3339),
			ProjectID: project.ID,
		})
		if err != nil {
			utils.RespondWithError(w, err, http.StatusInternalServerError, "error while storing cpe")
		}

		cves, err := nistlayer.GETCVEs(*cfg.HttpClient, cpestring)
		if err != nil {
			utils.RespondWithError(w, err, http.StatusInternalServerError, "error while fetching cve")
		}

		if len(cves.Vulnerabilities) > 0 {
			for _, cve := range cves.Vulnerabilities {
				_, err = cfg.DbQueries.StoreCVE(r.Context(), database.StoreCVEParams{
					ID:        fmt.Sprintf("%s_%s_%s", cve.Cve.ID, storedCpe.Cpe, project.ProjectName),
					Cve:       cve.Cve.ID,
					Descrip:   cve.Cve.Descriptions[0].Value,
					CreatedAt: time.Now().UTC().Format(time.RFC3339),
					UpdatedAt: time.Now().UTC().Format(time.RFC3339),
					Cpe:       storedCpe.ID,
					Project:   project.ID,
				})
				if err != nil {
					utils.RespondWithError(w, err, http.StatusInternalServerError, "error while storing CVEs")
				}
			}
		}

	}

	utils.WriteJSONToResponse(w, http.StatusCreated, nil)

}

func ValidateCPEs(cpes []CPE) error {
	for _, cpe := range cpes {
		if cpe.Part == "" || cpe.Vendor == "" || cpe.Product == "" || cpe.Version == "" {
			return errors.New("incomplete cpe")
		}
		cpeString := GenerateCPEString(cpe)
		if !nistlayer.ValidateIfCPEExists(cpeString) {
			return errors.New("none existant cpe")
		}
	}

	return nil
}

func GenerateCPEString(cpe CPE) string {
	// cpe format= cpe:2.3: {part} : {vendor} : {product} : {version} : {update} : {edition} : {language}
	part := cpe.Part
	vendor := cpe.Vendor
	product := cpe.Product
	version := cpe.Version
	update := "*"
	edition := "*"
	language := "*"

	if cpe.Update != "" {
		update = cpe.Update
	}
	if cpe.Edition != "" {
		edition = cpe.Edition
	}
	if cpe.Language != "" {
		language = cpe.Language
	}
	return fmt.Sprintf("cpe:2.3:%s:%s:%s:%s:%s:%s:%s", part, vendor, product, version, update, edition, language)
}
