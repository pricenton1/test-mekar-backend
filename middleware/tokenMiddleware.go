package middleware

import (
	"fmt"
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
				fmt.Println(err)
				utils.ResponseWithoutPayload(w, http.StatusUnauthorized)
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
