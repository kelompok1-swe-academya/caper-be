package middlewares

import "github.com/kelompok1-swe-academya/caper-be/pkg/jwt"

type Middleware struct {
	jwt jwt.JwtInterface
}

func NewMiddleware(
	jwt jwt.JwtInterface,
) *Middleware {
	return &Middleware{
		jwt: jwt,
	}
}
