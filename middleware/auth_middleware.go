package middleware

import (
	"go-restful-api/helper"
	"go-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writter http.ResponseWriter, request *http.Request) {
	var (
		apiKey = "OMKEGAS!"
	)

	if apiKey == request.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(writter, request)
	} else {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToJson(writter, webResponse)
	}
}
