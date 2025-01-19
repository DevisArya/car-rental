package main

import (
	"log"
	"os"

	"github.com/DevisArya/car-rental/app"
	"github.com/DevisArya/car-rental/config"
	"github.com/DevisArya/car-rental/helper"
	"github.com/DevisArya/car-rental/routes"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	db := config.NewDB()
	config.InitialMigration(db)
	config.SeedBookingTypes(db)

	validate := validator.New()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))

	appContainer := app.NewAppContainer(db, validate)
	routes.NewRouter(e, appContainer)

	if err := e.Start(os.Getenv("PORT")); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
