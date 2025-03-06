package middlewares

import (
	"strings"

	"github.com/kelompok1-swe-academya/caper-be/domain"
	"github.com/kelompok1-swe-academya/caper-be/internal/infra/env"
	"github.com/gofiber/fiber/v2"
)

func ApiKey() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		apiKey := ctx.Get("x-api-key")
		if apiKey == "" {
			return domain.ErrNoAPIKey
		}

		keySlice := strings.Split(apiKey, " ")
		if len(keySlice) != 2 {
			return domain.ErrInvalidAPIKey
		}

		key := keySlice[1]
		if key != env.AppEnv.ApiKey {
			return domain.ErrInvalidAPIKey
		}

		return ctx.Next()
	}
}
