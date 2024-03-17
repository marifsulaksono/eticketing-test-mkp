package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/marifsulaksono/eticketing-test-mkp/entities"
)

type JWTClaim struct {
	Id       int
	Username string
	Email    string
	Role     string
	jwt.RegisteredClaims
}

const (
	MKP_LOGINEDUSERID       string = "logined-user-id"
	MKP_LOGINEDUSERUSERNAME string = "logined-user-username"
	MKP_LOGINEDUSEREMAIL    string = "logined-user-email"
	MKP_LOGINEDUSERROLE     string = "logined-user-role"
)

func GetTokenFromHeader(w http.ResponseWriter, r *http.Request) (string, bool) {
	authHeader := r.Header.Get("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		return "", false
	}

	return strings.Replace(authHeader, "Bearer ", "", -1), true
}

func ValidateJWT(ts string) (*jwt.Token, error) {
	claims := &JWTClaim{}

	token, err := jwt.ParseWithClaims(ts, claims, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	return token, err
}

func GenerateToken(user entities.User) (string, error) {
	// Create token claim
	jwtExpTime := time.Now().Add(time.Hour * 168) // 1 week token duration
	claims := &JWTClaim{
		Id:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "eticketing-mkp",
			ExpiresAt: jwt.NewNumericDate(jwtExpTime),
		},
	}

	// Generate JWT Token
	tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenAlgorithm.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}
