package middleware

import (
	"log"
	"net/http"

	"github.com/4ndr3ye/DemoServer/handler"
	"github.com/4ndr3ye/DemoServer/security"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		token, err := req.Cookie("DemoCookie")
		if err != nil {
			handler.ErrorHandler(err)
			http.Error(w, "Forbidden", http.StatusForbidden)
			log.Println("No Cookie header")
			return
		}
		isValid, _ := security.ValidateToken(token.Value)
		if !isValid && token.Value != "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			log.Println("Not valid JWT")
			return
		}
		next.ServeHTTP(w, req)
	})
}
