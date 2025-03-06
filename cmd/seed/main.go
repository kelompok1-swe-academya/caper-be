package main

import (
	"github.com/kelompok1-swe-academya/caper-be/internal/infra/database"
)

const SeedersFilePath = "data/seeders/"
const SeedersDevPath = SeedersFilePath + "dev/"
const SeedersProdPath = SeedersFilePath + "prod/"

func main() {
	psqlDB := database.NewPgsqlConn()
	defer psqlDB.Close()

	// var path string
	// if env.AppEnv.AppEnv == "production" {
	// 	path = SeedersProdPath
	// } else {
	// 	path = SeedersDevPath
	// }

	// validator := validator.Validator
	// uuid := uuid.UUID
	// bcrypt := bcrypt.Bcrypt
}
