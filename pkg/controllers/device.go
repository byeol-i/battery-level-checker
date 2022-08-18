package controllers

import (
	"net/http"
)

type DeviceController struct {

}

func NewDeviceController() *DeviceController {
	return &DeviceController{}
}

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

