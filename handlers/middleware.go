package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/vezzsah/obsec_test/internal/auth"
	"github.com/vezzsah/obsec_test/utils"
)

type authedHandler func(http.ResponseWriter, *http.Request, uuid.UUID)

func (cfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			utils.RespondWithError(w, err, http.StatusForbidden, "forbidden")
			return
		}

		userId, err := auth.ValidateJWT(token)
		if err != nil {
			utils.RespondWithError(w, err, http.StatusForbidden, "forbidden")
			return
		}

		handler(w, r, userId)
	}
}
