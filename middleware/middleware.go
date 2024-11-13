package middleware

import (
	"cinema/service"
	"cinema/utils"
	"context"
	"net/http"
)

func AuthMiddleware(authService service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token := r.Header.Get("Authorization")
			if len(token) > 7 && token[:7] == "Bearer " {
				token = token[7:]
			}

			if token == "" {
				utils.SendJSONResponse(w, http.StatusUnauthorized, "Token required", nil)
				return
			}

			userID, err := authService.VerifyToken(token)
			if err != nil {
				utils.SendJSONResponse(w, http.StatusUnauthorized, "Invalid or expired token", nil)
				return
			}

			ctx := context.WithValue(r.Context(), "userID", userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
