package controllers

import (
	"net/http"
)

type NotFoundController struct {
}

func (hdl *NotFoundController) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	respondError(resp, 404, "not valid in default Controller")
}
