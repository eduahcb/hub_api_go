package responses

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   string      `json:"error"`
	Message interface{} `json:"message"`
}

func toJSON(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)
}

func JSON(w http.ResponseWriter, statusCode int, response interface{}) {
	toJSON(w, statusCode, response)
}

func ErrorJSON(w http.ResponseWriter, statusCode int, err error, response interface{}) {
	errorResponse := ErrorResponse{
		Error:   err.Error(),
		Message: response,
	}

	toJSON(w, statusCode, errorResponse)
}

func OK(w http.ResponseWriter, response interface{}) {
	toJSON(w, http.StatusOK, response)
}

func BadRequest(w http.ResponseWriter, err error) {
	errorResponse := ErrorResponse{
		Error:   "Bad Request",
		Message: err.Error(),
	}

	toJSON(w, http.StatusBadRequest, errorResponse)
}

func NotFound(w http.ResponseWriter, err error) {
	errorResponse := ErrorResponse{
		Error:   "Not Found",
		Message: err.Error(),
	}

	toJSON(w, http.StatusNotFound, errorResponse)
}

func Unauthorized(w http.ResponseWriter, err error) {
	errorResponse := ErrorResponse{
		Error:   "Unauthorized",
		Message: err.Error(),
	}

	toJSON(w, http.StatusUnauthorized, errorResponse)
}

func InternalServerError(w http.ResponseWriter, err error) {
	errorResponse := ErrorResponse{
		Error:   "Internal Server Error",
		Message: err.Error(),
	}

	toJSON(w, http.StatusInternalServerError, errorResponse)
}
