package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Role int

const (
	Admin Role = iota
	User
)

type Payload struct {
	jwt.RegisteredClaims
	Role Role `json:"role"`
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

func (p *Payload) DecodeJwt(tokenString string, secret string) (*Payload, error) {
	parsedToken, err := jwt.ParseWithClaims(
		tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},
	)
	if err != nil {
		return nil, nil
	}

	if claims, ok := parsedToken.Claims.(*Payload); ok && parsedToken.Valid {
		return claims, nil
	} else {
		return nil, nil
	}
}
