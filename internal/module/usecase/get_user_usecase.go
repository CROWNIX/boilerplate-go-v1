package usecase

import (
	"context"

	util_db "gitea.qwertysystem.net/BETS/ts-utils/db"
	util_functions "gitea.qwertysystem.net/BETS/ts-utils/functions"
	util_provider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	util_service "gitea.qwertysystem.net/BETS/ts-utils/service"

	"github.com/CROWNIX/boilerplate-go-v1/internal/module/dto"
	query_builder "github.com/CROWNIX/boilerplate-go-v1/internal/module/query_builder"
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
	FixedAssetService util_service.PostgreSqlService
}

func MakeGetUserListUseCase(
	serviceProvider util_provider.ServiceProvider,
) *GetUserListUseCase {
	return &GetUserListUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *GetUserListUseCase) InitService(companyCode string) {}

func (u *GetUserListUseCase) MakeService(companyCode string) GetUserListServices {
	fixedAssetDbName := util_db.FixedAssetDBName.CompanyDBName(companyCode)

	return GetUserListServices{
		FixedAssetService: u.ServiceProvider.MakePostgreSqlService(fixedAssetDbName),
	}
}

func (u *GetUserListUseCase) Invoke(param GetUserListParam) (*dto.PaginatedGetUserListResponse, error) {
	services := u.MakeService(param.CompanyCode)

	var err error
	var result []dto.PaginatedGetUserListResponse

	queryString, args, err := query_builder.GetUserListBuilder(param.Query)

	err = services.FixedAssetService.SelectMany(
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
