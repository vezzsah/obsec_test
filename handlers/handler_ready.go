package handlers

import (
	"net/http"

	"github.com/vezzsah/obsec_test/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSONToResponse(w, http.StatusOK, map[string]string{"status": "ok"})
}
