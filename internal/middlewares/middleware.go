package middlewares

import "github.com/ahargunyllib/hackathon-fiber-starter/pkg/jwt"

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
