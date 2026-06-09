package route

import (
	"time"

	util_service "gitea.qwertysystem.net/BETS/ts-utils/service"
	"github.com/CROWNIX/boilerplate-go-v1/internal/user/controller"
	"github.com/CROWNIX/boilerplate-go-v1/internal/user/usecase"
	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoute(
	app *fiber.App,
	userController controller.UserController,
) {
	apiGroup := app.Group("/users")
	apiGroup.Get("/", userController.GetUserList)
}

func SetUpUserController(app *fiber.App, sqlService util_service.PostgreSqlService) {
	timeout := 5 * time.Minute

	

	getUserListUseCase := usecase.MakeGetUserListUseCase(sqlService)
	userController := controller.MakeUserController(
		timeout, 
		getUserListUseCase,
	)

	SetUpUserRoute(app, *userController)
}
