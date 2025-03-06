package main

import (
	"github.com/kelompok1-swe-academya/caper-be/internal/infra/database"
	"github.com/kelompok1-swe-academya/caper-be/internal/infra/env"
	"github.com/kelompok1-swe-academya/caper-be/internal/infra/server"
)

func main() {
	server := server.NewHttpServer()
	psqlDB := database.NewPgsqlConn()
	defer psqlDB.Close()

	server.MountMiddlewares()
	server.MountRoutes(psqlDB)
	server.Start(env.AppEnv.AppPort)
}
