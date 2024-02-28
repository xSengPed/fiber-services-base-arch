package main

import (
	"go-fiber-basic/config"
	usercontroller "go-fiber-basic/controller/user"
	"go-fiber-basic/repository"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

func main() {

	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		slog.Error(err.Error())
	}

	dbConfig := new(config.DbConfig)

	err = env.Parse(dbConfig)
	if err != nil {
		slog.Error(err.Error())
	}

	repository.ConnectDatabase(dbConfig)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	userGroup := v1.Group("/user")                           // localhost:3000/api/v1/user
	userGroup.Post("/register", usercontroller.RegisterUser) // localhost:3000/api/v1/user/register

	app.Listen(":3000")
}
