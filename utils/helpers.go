package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseBody[T any](w http.ResponseWriter, r *http.Request, params *T) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(params)
}

func WriteJSONToResponse(w http.ResponseWriter, resCode int, jsonObject any) {
	dat, err := json.Marshal(jsonObject)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resCode)
	_, err = w.Write(dat)
	if err != nil {
		log.Printf("Error writting JSON to Body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
