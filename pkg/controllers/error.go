package controllers

import (
	"encoding/json"
	"net/http"
)

func respondJSON(resp http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)
	resp.Write([]byte(response))
}

func respondError(resp http.ResponseWriter, code int, message string) {
	respondJSON(resp, code, map[string]string{"error": message})
}