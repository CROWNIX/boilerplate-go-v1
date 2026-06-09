package route

import (
	util_provider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"github.com/CROWNIX/boilerplate-go-v1/internal/user/controller"
	"github.com/CROWNIX/boilerplate-go-v1/internal/user/usecase"

	"time"

	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoute(
	app *fiber.App,
	userController controller.UserController,
) {
	apiGroup := app.Group("/users")
	apiGroup.Get("/", userController.GetUserList)
}

func SetUpUserController(app *fiber.App, serviceProvider util_provider.ServiceProvider) {
	timeout := 5 * time.Minute
	getUserListUseCase := usecase.MakeGetUserListUseCase(serviceProvider)
	userController := controller.MakeUserController(
		timeout, 
		getUserListUseCase,
	)

	SetUpUserRoute(app, *userController)
}
