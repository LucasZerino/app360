package middleware

import (
	"context"
	"net/http"

	"app360/cmd/api/response"
	"app360/internal/auth"
	"app360/internal/errors"
)

func AuthMiddleware(jwt *auth.JWTManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("auth_token")
			if err != nil {
				response.RespondWithError(w, errors.NewAPIError(http.StatusUnauthorized, "Não autorizado"))
				return
			}

			claims, err := jwt.ValidateToken(cookie.Value)
			if err != nil {
				response.RespondWithError(w, errors.NewAPIError(http.StatusUnauthorized, "Token inválido"))
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, "user_id", claims.UserID)
			ctx = context.WithValue(ctx, "email", claims.Email)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
