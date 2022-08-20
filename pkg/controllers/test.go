package controllers

import (
	"net/http"
)

// GetMap godoc
// @Summary Get Map Example
// @Description get map
// @ID get-map
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /test [get]
type TestController struct {

}

func NewTestController() *TestController {
	return &TestController{}
}

func (hdl *TestController) PrintHello(resp http.ResponseWriter, req *http.Request) {
	bytes := []byte("Hello World")
	resp.Write(bytes)

	resp.Header().Set("Content-Type", "text/html")
}