package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID string) (string, error) {
	apiSecret := os.Getenv("SECRET_KEY")
	TokenTypeAccess := os.Getenv("TOKEN_TYPE")
	claim := jwt.RegisteredClaims{
		Issuer:    string(TokenTypeAccess),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
		Subject:   userID,
	}

	tokenWithClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return tokenWithClaim.SignedString([]byte(apiSecret))
}

func ValidateJWT(tokenString string) (uuid.UUID, error) {
	apiSecret := os.Getenv("SECRET_KEY")
	TokenTypeAccess := os.Getenv("TOKEN_TYPE")
	claim := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (any, error) {
		return []byte(apiSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	id, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.Nil, err
	}
	if issuer != string(TokenTypeAccess) {
		return uuid.Nil, errors.New("invalid issuer")
	}

	idUuid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, err
	}

	return idUuid, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	authorization := headers.Get("Authorization")
	if authorization == "" {
		return "", errors.New("missing authorization header")
	}

	splitAuth := strings.Split(authorization, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "Bearer" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

func MakeRefreshToken() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return hex.EncodeToString(key), err
}
