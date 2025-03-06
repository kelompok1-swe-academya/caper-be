package controller

import (
	"github.com/ahargunyllib/hackathon-fiber-starter/domain/contracts"
	"github.com/ahargunyllib/hackathon-fiber-starter/domain/dto"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/helpers/http/response"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authService contracts.AuthService
}

func InitAuthController(router fiber.Router, authService contracts.AuthService) {
	controller := authController{
		authService: authService,
	}

	authGroup := router.Group("/auth")
	authGroup.Post("/register", controller.registerUser)
	authGroup.Post("/login", controller.loginUser)
}

func (ac *authController) registerUser(ctx *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := ac.authService.RegisterUser(ctx.Context(), req)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}

func (ac *authController) loginUser(ctx *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := ac.authService.LoginUser(ctx.Context(), req)
	if err != nil {
		return err
	}

	return response.SendResponse(ctx, fiber.StatusOK, res)
}
