package controllers

import (
	"encoding/json"
	"net/http"

	dbSvc "github.com/byeol-i/battery-level-checker/pkg/svc/db"
	"github.com/byeol-i/battery-level-checker/pkg/user"
)

type UserControllers struct {

}

func NewUserController() *UserControllers {
	return &UserControllers{}
}

// AddUser godoc
// @Summary Register New user
// @Description Register New user
// @Tags user
// @Accept json
// @Produce json
// @Param userInfo body models.Spec true "add user"
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{}
// @Router /user [post]
func (hdl *UserControllers) AddNewUser(resp http.ResponseWriter, req *http.Request) {
	
	var userSpec user.UserImpl
	err := json.NewDecoder(req.Body).Decode(&userSpec)
	if err != nil {
		respondError(resp, http.StatusBadRequest, "invalid format")
	}

	err = user.UserValidator(&userSpec)
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())	
	}

	newUser := user.NewUser()
	newUser.SetId(userSpec.Id)
	newUser.SetName(userSpec.Name)

	err = dbSvc.CallAddNewUser(&newUser.UserImpl)

	respondJSON(resp, http.StatusOK, "success", nil)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user ID"
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{}
// @Router /user/{userID} [delete]
func (hdl *UserControllers) DeleteUser(resp http.ResponseWriter, req *http.Request) {
	respondJSON(resp, http.StatusOK, "success", nil)
}