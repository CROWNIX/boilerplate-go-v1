package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"

	util_entity "gitea.qwertysystem.net/BETS/ts-utils/entity"
	util_http "gitea.qwertysystem.net/BETS/ts-utils/http"
	util_parser "gitea.qwertysystem.net/BETS/ts-utils/parser"

	"github.com/CROWNIX/boilerplate-go-v1/internal/user/dto"
	"github.com/CROWNIX/boilerplate-go-v1/internal/user/usecase"
)

func (c *UserController) GetUserList(ctx *fiber.Ctx) error {
	rawQuery := ctx.Queries()

	return util_http.RunWithTimeout(ctx, c.Timeout, func(ctxWithTimeout context.Context) (*dto.PaginatedGetUserListResponse, *util_entity.HttpError) {
		query, err := util_parser.ParseQuery[dto.GetUserQuery](rawQuery)
		if err != nil {
			return nil, util_entity.ToHttpError(err)
		}

		param := usecase.GetUserListParam{
			Ctx:         ctxWithTimeout,
			Query:       *query,
		}

		res, err := c.GetUserListUseCase.Invoke(param)
		if err != nil {
			return nil, util_entity.ToHttpError(err)
		}

		return res, nil
	}, "Successfully retrieve user list", fiber.StatusOK)
}
