package models

type JSONresult struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

/** just using for define swag doc */

type JSONsuccessResult struct {
	Code int `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
	Data interface{} `json:"data"`
}

type JSONfailResult struct {
	Code int `json:"code" example:"400"`
	Message string `json:"message" example:"error"`
	Data interface{} `json:"data"`
}