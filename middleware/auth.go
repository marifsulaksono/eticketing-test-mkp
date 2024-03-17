package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/marifsulaksono/eticketing-test-mkp/utils"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tokenString, ok := utils.GetTokenFromHeader(w, r)
		if !ok {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		token, err := utils.ValidateJWT(tokenString)
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorExpired:
				http.Error(w, "expired token", http.StatusUnauthorized)
				return
			default:
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		payload, ok := token.Claims.(*utils.JWTClaim)
		if !ok {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, utils.MKP_LOGINEDUSERID, payload.Id)
		ctx = context.WithValue(ctx, utils.MKP_LOGINEDUSERUSERNAME, payload.Username)
		ctx = context.WithValue(ctx, utils.MKP_LOGINEDUSEREMAIL, payload.Email)
		ctx = context.WithValue(ctx, utils.MKP_LOGINEDUSERROLE, payload.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
