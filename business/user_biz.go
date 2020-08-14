package biz

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hoangduc02011998/todo-app/server/models"
)

// jwt claims ...
type JwtClaims struct {
	Name string      `json:"name"`
	User models.User `json:"user"`
	jwt.StandardClaims
}

func CreateUser(input models.User) *models.ResponseModel {
	user, _ := userRepo.GetByUsername(input.UserName)
	if user != nil {
		return &models.ResponseModel{
			Status:  models.APIStatus.Existed,
			Message: "User is available",
		}
	}

	err := userRepo.Insert(input)
	if err != nil {
		return &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: err.Error(),
		}
	}

	return &models.ResponseModel{
		Status:  models.APIStatus.Ok,
		Message: models.APIStatus.Ok,
	}
}

func UserLogin(input models.UserLogin) *models.ResponseModel {
	user, _ := userRepo.Login(models.User{UserName: input.UserName, Password: input.Password})
	if user == nil {
		return &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: "UserName or Password was wrong",
		}
	}

	token, err := createToken(*user)
	if err != nil {
		return &models.ResponseModel{
			Status:  models.APIStatus.Error,
			Message: err.Error(),
		}
	}

	return &models.ResponseModel{
		Status:  models.APIStatus.Ok,
		Message: "Login Success",
		Data:    token,
	}
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
