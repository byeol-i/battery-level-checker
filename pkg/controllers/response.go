package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/byeol-i/battery-level-checker/pkg/models"
)

func respondJSON(resp http.ResponseWriter, status int, message string, payload interface{}) {
	response, err := json.Marshal(models.JSONresult{
		Code: status,
		Message: message,
		Data: payload,
	})
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
		return
	} else {
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(status)
		resp.Write([]byte(response))
		return
	}
}

func respondError(resp http.ResponseWriter, code int, message string) {
	errMsg := strings.Replace(message, "rpc error: code = Unknown desc = ", "", 1)
		
	respondJSON(resp, code, "error",  map[string]string{"error": errMsg})
}