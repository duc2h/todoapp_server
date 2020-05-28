package api

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hoangduc02011998/todo_server/models"
	repo "github.com/hoangduc02011998/todo_server/repository"
	"github.com/hoangduc02011998/todo_server/repository/repoIml"
	"github.com/labstack/echo"
)

// jwt claims ...
type JwtClaims struct {
	Name string      `json:"name"`
	User models.User `json:"user"`
	jwt.StandardClaims
}

var userRepo repo.UserRepo

func InitUserRepo() {
	userRepo = repoIml.NewUserRepo(models.UserDB.Collection)
}

func UserPost(c echo.Context) error {
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	var user models.User
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error(),
		})
	}

	err = userRepo.Insert(user)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Invalid,
			Message: err.Error(),
		})
	}

	return Respond(c, &ResponseModel{
		Status:  APIStatus.Ok,
		Message: APIStatus.Ok,
		Data:    user,
	})
}

func UserLogin(c echo.Context) error {
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error()})
	}

	user := models.User{}
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error()})
	}

	userLogin, err := userRepo.Login(user)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error()})
	}

	token, err := createToken(userLogin)
	if err != nil {
		return Respond(c, &ResponseModel{
			Status:  APIStatus.Error,
			Message: err.Error()})
	}

	return Respond(c, &ResponseModel{
		Status:  APIStatus.Ok,
		Message: "Login Success",
		Data:    token,
	})
}

func createToken(user models.User) (string, error) {
	claims := JwtClaims{
		"edar",
		user,
		jwt.StandardClaims{
			Id:        "browser_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
