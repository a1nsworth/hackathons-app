package auth

import (
	"log/slog"

	"hackathons-app/internal/config"
	"hackathons-app/internal/jwt"
	"hackathons-app/internal/services"
)

type Auth struct {
	log    *slog.Logger
	config config.JWT
	user   services.UserService
}

func NewAuth(log *slog.Logger, config config.JWT, user services.UserService) *Auth {
	return &Auth{
		log:    log.WithGroup("auth"),
		config: config,
		user:   user,
	}
}

func (a Auth) JWTAuthentication(token string) (*jwt.Payload, error) {
	return jwt.DecodeJwt(
		token,
		a.config.AccessSecret,
	)
}
