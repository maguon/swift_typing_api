package util

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type options struct {
	signingMethod jwt.SigningMethod
	signingKey    interface{}
	expired       int
	tokenType     string
}

const defaultKey = "gin-go"

var defaultOptions = options{
	tokenType:     "Bearer",
	expired:       7200,
	signingMethod: jwt.SigningMethodHS256,
	signingKey:    []byte(defaultKey),
}

func GenerateAccess(userID int) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(defaultOptions.signingMethod, jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Duration(defaultOptions.expired) * time.Second).Unix(),
		NotBefore: now.Unix(),
		Id:        strconv.Itoa(userID),
	})

	tokenString, err := token.SignedString(defaultOptions.signingKey)
	if err != nil {
		return "", errors.New("generate token fail")
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(defaultKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenInvalid
			} else {
				return nil, ErrTokenInvalid
			}
		}
	} else if !token.Valid {
		return nil, ErrTokenInvalid
	}

	return token.Claims.(*jwt.StandardClaims), nil
}
