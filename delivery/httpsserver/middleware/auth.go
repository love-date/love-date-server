package middleware

import (
	"context"
	"love-date/delivery/httpsserver/response"
	"love-date/pkg/errhandling/httpmapper"
	"love-date/pkg/jwttoken"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			response.Fail("Malformed token", http.StatusUnauthorized).ToJSON(w)

			return
		} else {
			tokenString := authHeader[1]
			isValid, claims, jErr := jwttoken.ValidateJWT(tokenString)
			if !isValid {
				msg, code := httpmapper.Error(jErr)
				response.Fail(msg, code).ToJSON(w)

				return
			}

			ctx := context.WithValue(r.Context(), "user_id", claims.UserID)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

	})
}
