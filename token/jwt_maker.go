package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

// JWT maker is JSON Web Token MAker
type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("Key size must be atleast %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (j *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)

	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(j.secretKey))
}

func (j *JWTMaker) VerifyToken(token string) (*Payload, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) //type checking syntax

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(j.secretKey), nil

	})

	if err != nil {
		return nil, err
	}

	payload, ok := parsedToken.Claims.(*Payload)

	if !ok {
		return nil, jwt.ErrInvalidType
	}

	return payload, nil

}
