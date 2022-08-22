package controllers

import (
	"net/http"
)


type TestController struct {

}

func NewTestController() *TestController {
	return &TestController{}
}

// BatteryLevelChecker godoc
// @Summary Print Hello
// @Description test
// @Tags test
// @Accept json
// @Produce json
// @Failure 400 
// @Router /test [post]
func (hdl *TestController) PrintHello(resp http.ResponseWriter, req *http.Request) {
	bytes := []byte("Hello World")
	resp.Write(bytes)

	resp.Header().Set("Content-Type", "text/html")
}