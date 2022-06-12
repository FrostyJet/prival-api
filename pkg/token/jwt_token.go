package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTToken struct {
	Secret string
}

func NewJWTToken(secret string) *JWTToken {
	return &JWTToken{
		Secret: secret,
	}
}

func (j *JWTToken) Create(userID string, duration time.Duration) (string, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	base64Str, err := jwtToken.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return base64Str, nil
}

func (j *JWTToken) Verify(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid signature")
		}

		return []byte(j.Secret), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		return nil, err
	}

	return jwtToken.Claims.(*Payload), nil
}
