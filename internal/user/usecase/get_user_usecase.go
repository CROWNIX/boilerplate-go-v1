package usecase

import (
	"context"
	"os"

	util_db "gitea.qwertysystem.net/BETS/ts-utils/db"
	util_functions "gitea.qwertysystem.net/BETS/ts-utils/functions"
	util_provider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	util_service "gitea.qwertysystem.net/BETS/ts-utils/service"

	"github.com/CROWNIX/boilerplate-go-v1/internal/user/dto"
	query_builder "github.com/CROWNIX/boilerplate-go-v1/internal/user/query_builder"
)

type GetUserListParam struct {
	Ctx         context.Context
	Query       dto.GetUserQuery
	CompanyCode string
}

type GetUserListUseCase struct {
	ServiceProvider util_provider.ServiceProvider
}

type GetUserListServices struct {
	SqlService util_service.PostgreSqlService
}

func MakeGetUserListUseCase(
	serviceProvider util_provider.ServiceProvider,
) *GetUserListUseCase {
	return &GetUserListUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *GetUserListUseCase) MakeService() GetUserListServices {
	return GetUserListServices{
		SqlService: u.ServiceProvider.MakePostgreSqlService(util_db.DBName(os.Getenv("DATABASE_NAME"))),
	}
}

func (u *GetUserListUseCase) Invoke(param GetUserListParam) (*dto.PaginatedGetUserListResponse, error) {
	services := u.MakeService()

	queryString, args, err := query_builder.GetUserListBuilder(param.Query)
	
	var result []dto.PaginatedGetUserListResponse
	err = services.SqlService.SelectMany(
		&result,
		param.Ctx,
		queryString,
		args...,
	)
	if err != nil {
		return nil, err
	}

	return util_functions.MakePointer(util_functions.FormatPaginationResult(result)), nil
}
