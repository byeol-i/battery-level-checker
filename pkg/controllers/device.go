package controllers

import (
	"net/http"
)

type DeviceController struct {

}

func NewDeviceController() *DeviceController {
	return &DeviceController{}
}


// GetThing godoc
// @Summary This is the summary for getting a thing by its UUID
// @Description This is the description for getting a thing by its UUID. Which can be longer,
// @Description and can continue over multiple lines
// @ID get-thing
// @Tags Thing
// @Param uuid path string true "The UUID of a thing"
// @Success 200 {object} ThingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /thing/{uuid} [get]
func (hdl *DeviceController) GetBattery(resp http.ResponseWriter, req *http.Request) {
	bytes := []byte(`
	{
		level: 20%,
		time: 
		status: charging,
	}`)
	resp.Write(bytes)

	resp.Header().Set("Content-Type", "text/html")
}

func (hdl *DeviceController) UpdateBattery(resp http.ResponseWriter, req *http.Request) {

} 

func (hdl *DeviceController) PrintHello(resp http.ResponseWriter, req *http.Request) {
	bytes := []byte("Hello World")
	resp.Write(bytes)

	resp.Header().Set("Content-Type", "text/html")
}

