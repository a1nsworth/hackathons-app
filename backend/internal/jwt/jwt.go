package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	jwt.RegisteredClaims
}

func NewPayload(subject string) *Payload {
	return &Payload{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: subject,
		},
	}
}

func (p *Payload) CreateJwt(secret string, exp time.Duration) (string, error) {
	now := time.Now()

	p.NotBefore = nil
	p.IssuedAt = jwt.NewNumericDate(now)
	p.ExpiresAt = jwt.NewNumericDate(now.Add(exp))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, p)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeJwt(tokenString string, secret string) (*Payload, error) {
	parsedToken, err := jwt.ParseWithClaims(
		tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*Payload); ok && parsedToken.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("can`t parse token")
	}
}
