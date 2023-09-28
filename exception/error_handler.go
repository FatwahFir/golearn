package exception

import (
	"net/http"

	"github.com/FatwahFir/golearn/helpers"
	"github.com/FatwahFir/golearn/model/response"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	res := response.WebResponse{
		StatusCode: http.StatusInternalServerError,
		Status:     "INTERNAL SERVER ERROR",
	}

	helpers.WriteResponse(writer, res)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFounError)

	if ok {

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		res := response.WebResponse{
			StatusCode: http.StatusNotFound,
			Status:     "NOT FOUND",
			Data:       exception.Error,
		}

		helpers.WriteResponse(writer, res)

		return true
	} else {
		return false
	}
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		res := response.WebResponse{
			StatusCode: http.StatusBadRequest,
			Status:     "BAD REQUEST",
			Data:       exception.Error(),
		}

		helpers.WriteResponse(writer, res)

		return true
	} else {
		return false
	}
}
