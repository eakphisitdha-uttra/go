package responses

import (
	"microservice/logs"
	"net/http"
)

type NoData struct{}

type NoDetail struct{}

// 200
func SuccessOK(data interface{}) (int, Success) {
	return http.StatusOK, Success{
		Status:  http.StatusOK,
		Message: "Get successfully",
		Data:    data,
	}
}

// 201
func SuccessCreated(data interface{}) (int, Success) {
	return http.StatusCreated, Success{
		Status:  http.StatusCreated,
		Message: "created successfully",
		Data:    data,
	}
}

// 202
func SuccessAccepted(data interface{}) (int, Success) {
	return http.StatusAccepted, Success{
		Status:  http.StatusAccepted,
		Message: "Accepted successfully",
		Data:    data,
	}
}

// 400
func ErrorBadRequested(err error, details interface{}) (int, Error) {
	return http.StatusBadRequest, Error{
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
		Message: err.Error(),
		Details: details,
	}
}

// 404
func ErrorNotFound(err error, details interface{}) (int, Error) {
	return http.StatusNotFound, Error{
		Status:  http.StatusNotFound,
		Error:   "Not Found",
		Message: err.Error(),
		Details: details,
	}
}

// 422
func ErrorValidated(details interface{}) (int, Error) {
	return http.StatusUnprocessableEntity, Error{
		Status:  http.StatusUnprocessableEntity,
		Error:   "Unprocessable Entity",
		Message: "Validation failed",
		Details: details,
	}
}

// 500
func ErrorInternalServer(err error, details interface{}) (int, Error) {
	logs.Error(err.Error())
	return http.StatusInternalServerError, Error{
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
		Message: err.Error(),
		Details: details,
	}
}
