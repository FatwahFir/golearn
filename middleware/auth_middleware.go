package middleware

import (
	"net/http"

	"github.com/FatwahFir/golearn/helpers"
	"github.com/FatwahFir/golearn/model/response"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if "this-api-key" == request.Header.Get("X-API-KEY") {
		//Ok
		middleware.Handler.ServeHTTP(writer, request)

	} else {
		//UNAUTHORIZED
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		res := response.WebResponse{
			StatusCode: http.StatusUnauthorized,
			Status:     "UNAUTHORIZED",
		}

		helpers.WriteResponse(writer, res)
	}

}
