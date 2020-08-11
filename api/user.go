package api

import (
	"encoding/json"
	"io/ioutil"

	biz "github.com/hoangduc02011998/todo_server/business"
	"github.com/hoangduc02011998/todo_server/models"
	"github.com/labstack/echo"
)

func UserPost(c echo.Context) error {
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return models.Respond(c, &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: err.Error(),
		})
	}

	var user models.User
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return models.Respond(c, &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: err.Error(),
		})
	}

	if user.UserName == "" || user.Password == "" || user.PasswordConfirm == "" {
		return models.Respond(c, &models.ResponseModel{
			Status:  models.APIStatus.Invalid,
			Message: "UserName, Password, PasswordConfirm is required",
		})
	}

	return models.Respond(c, biz.CreateUser(user))
}

func UserLogin(c echo.Context) error {
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return models.Respond(c, &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: err.Error()})
	}

	userLogin := models.UserLogin{}
	err = json.Unmarshal(bodyBytes, &userLogin)
	if err != nil {
		return models.Respond(c, &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: err.Error()})
	}

	if userLogin.UserName == "" {
		return models.Respond(c, &models.ResponseModel{
			Status:  models.APIStatus.Invalid,
			Message: "UserName is required",
		})
	}

	return models.Respond(c, biz.UserLogin(userLogin))
}
