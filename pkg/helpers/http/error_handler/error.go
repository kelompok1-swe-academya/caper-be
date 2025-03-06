package errorhandler

import (
	"errors"

	"github.com/kelompok1-swe-academya/caper-be/domain"
	"github.com/kelompok1-swe-academya/caper-be/pkg/helpers/http/response"
	"github.com/kelompok1-swe-academya/caper-be/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var valErr validator.ValidationErrors
	if errors.As(err, &valErr) {
		return response.SendResponse(c, fiber.StatusUnprocessableEntity, valErr)
	}

	var reqErr *domain.RequestError
	if errors.As(err, &reqErr) {
		return response.SendResponse(c, reqErr.StatusCode, reqErr)
	}

	return response.SendResponse(c, fiber.StatusInternalServerError, err)
}
