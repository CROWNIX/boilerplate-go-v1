package app

import (
	"os"

	util_provider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"github.com/CROWNIX/boilerplate-go-v1/internal/user/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type App struct {
	app *fiber.App
}

func MakeApp() *App {
	return &App{app: fiber.New()}
}

func (a *App) Run(serviceProvider util_provider.ServiceProvider) {
	a.app.Use(cors.New())
	a.app.Use(logger.New())

	setupRoute(a.app, serviceProvider)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	a.app.Listen(":" + port)
}

func setupRoute(app *fiber.App, serviceProvider util_provider.ServiceProvider) {
	route.SetUpUserController(app, serviceProvider)
}
