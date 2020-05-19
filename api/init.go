package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// StatusEnum ...
type StatusEnum struct {
	Ok           string
	Error        string
	Invalid      string
	NotFound     string
	Forbidden    string
	Existed      string
	Unauthorized string
}

// APIStatus Published enum
var APIStatus = &StatusEnum{
	Ok:           "OK",
	Error:        "ERROR",
	Invalid:      "INVALID",
	NotFound:     "NOT_FOUND",
	Forbidden:    "FORBIDDEN",
	Existed:      "EXISTED",
	Unauthorized: "UNAUTHORIZED",
}

type ResponseModel struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type APIRespose struct {
	context echo.Context
}

func Respond(context echo.Context, response *ResponseModel) error {

	switch response.Status {
	case APIStatus.Ok:
		return context.JSON(http.StatusOK, response)
	case APIStatus.Error:
		return context.JSON(http.StatusInternalServerError, response)
	case APIStatus.Forbidden:
		return context.JSON(http.StatusForbidden, response)
	case APIStatus.Invalid:
		return context.JSON(http.StatusBadRequest, response)
	case APIStatus.NotFound:
		return context.JSON(http.StatusNotFound, response)
	case APIStatus.Unauthorized:
		return context.JSON(http.StatusUnauthorized, response)
	case APIStatus.Existed:
		return context.JSON(http.StatusConflict, response)
	}

	return context.JSON(http.StatusBadRequest, response)
}
