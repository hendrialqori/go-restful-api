package exception

import (
	"go-restful-api/helper"
	"go-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writter http.ResponseWriter, request *http.Request, err any) {
	if notFoundError(writter, err) {
		return
	}

	if validatorErrors(writter, err) {
		return
	}

	internalServerError(writter, err)
}

func notFoundError(writter http.ResponseWriter, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToJson(writter, webResponse)
		return true
	} else {
		return false
	}
}

func validatorErrors(writter http.ResponseWriter, err any) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToJson(writter, webResponse)
		return true
	} else {
		return false
	}

}

func internalServerError(writter http.ResponseWriter, err any) {
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   500,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToJson(writter, webResponse)
}
