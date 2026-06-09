package app

import (
	"os"

	util_provider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	util_service "gitea.qwertysystem.net/BETS/ts-utils/service"
	util_db "gitea.qwertysystem.net/BETS/ts-utils/db"
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

	sqlService := serviceProvider.MakePostgreSqlService(
        util_db.DBName(os.Getenv("DATABASE_NAME")),
    )

	setupRoute(a.app, sqlService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	a.app.Listen(":" + port)
}

func setupRoute(app *fiber.App, sqlService util_service.PostgreSqlService) {
	route.SetUpUserController(app, sqlService)
}
