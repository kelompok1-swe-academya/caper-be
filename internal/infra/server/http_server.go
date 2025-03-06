package server

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	userCtr "github.com/kelompok1-swe-academya/caper-be/internal/app/user/controller"
	userRepo "github.com/kelompok1-swe-academya/caper-be/internal/app/user/repository"
	userSvc "github.com/kelompok1-swe-academya/caper-be/internal/app/user/service"
	"github.com/kelompok1-swe-academya/caper-be/internal/infra/env"
	"github.com/kelompok1-swe-academya/caper-be/internal/middlewares"
	"github.com/kelompok1-swe-academya/caper-be/pkg/bcrypt"
	errorhandler "github.com/kelompok1-swe-academya/caper-be/pkg/helpers/http/error_handler"
	"github.com/kelompok1-swe-academya/caper-be/pkg/helpers/http/response"
	"github.com/kelompok1-swe-academya/caper-be/pkg/jwt"
	"github.com/kelompok1-swe-academya/caper-be/pkg/log"
	timePkg "github.com/kelompok1-swe-academya/caper-be/pkg/time"
	"github.com/kelompok1-swe-academya/caper-be/pkg/uuid"
	"github.com/kelompok1-swe-academya/caper-be/pkg/validator"
)

type HttpServer interface {
	Start(part string)
	MountMiddlewares()
	MountRoutes(db *sqlx.DB)
	GetApp() *fiber.App
}

type httpServer struct {
	app *fiber.App
}

func NewHttpServer() HttpServer {
	config := fiber.Config{
		CaseSensitive: true,
		AppName:       "Hackathon Fiber Starter",
		ServerHeader:  "Hackathon Fiber Starter",
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
		ErrorHandler:  errorhandler.ErrorHandler,
	}

	app := fiber.New(config)

	return &httpServer{
		app: app,
	}
}

func (s *httpServer) GetApp() *fiber.App {
	return s.app
}

func (s *httpServer) Start(port string) {
	if port[0] != ':' {
		port = ":" + port
	}

	err := s.app.Listen(port)

	if err != nil {
		log.Fatal(log.LogInfo{
			"error": err.Error(),
		}, "[SERVER][Start] failed to start server")
	}
}

func (s *httpServer) MountMiddlewares() {
	s.app.Use(middlewares.LoggerConfig())
	s.app.Use(middlewares.Helmet())
	s.app.Use(middlewares.Compress())
	s.app.Use(middlewares.Cors())
	if env.AppEnv.AppEnv != "development" {
		s.app.Use(middlewares.ApiKey())
	}
	s.app.Use(middlewares.RecoverConfig())
}

func (s *httpServer) MountRoutes(db *sqlx.DB) {
	bcrypt := bcrypt.Bcrypt
	_ = timePkg.Time
	uuid := uuid.UUID
	validator := validator.Validator
	jwt := jwt.Jwt

	_ = middlewares.NewMiddleware(jwt)

	s.app.Get("/", func(c *fiber.Ctx) error {
		return response.SendResponse(c, fiber.StatusOK, "hai manies😘")
	})

	api := s.app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return response.SendResponse(c, fiber.StatusOK, "hai manies😘")
	})

	userRepository := userRepo.NewUserRepository(db)

	userService := userSvc.NewUserService(userRepository, validator, uuid, bcrypt)

	userCtr.InitNewController(v1, userService)

	s.app.Use(func(c *fiber.Ctx) error {
		return c.SendFile("./web/not-found.html")
	})
}
