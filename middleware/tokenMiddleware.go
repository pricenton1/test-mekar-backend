package middleware

import (
	"log"
	"net/http"
	"test-mekar-backend/utils"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenValue := r.Header.Get("token")
		if len(tokenValue) == 0 {
			utils.ResponseWithoutPayload(w, http.StatusUnauthorized)
		} else {
			_, err := utils.JwtDecoder(tokenValue)
			if err != nil {
				utils.ResponseWithoutPayload(w, http.StatusUnauthorized)
				log.Print("Token Not Valid")
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
