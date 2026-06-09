package usecase

import (
	"context"

	util_functions "gitea.qwertysystem.net/BETS/ts-utils/functions"
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
	SqlService util_service.PostgreSqlService
}

type GetUserListServices struct {
	SqlService util_service.PostgreSqlService
}

func MakeGetUserListUseCase(sqlService util_service.PostgreSqlService) *GetUserListUseCase {
    return &GetUserListUseCase{SqlService: sqlService}
}

func (u *GetUserListUseCase) Invoke(param GetUserListParam) (*dto.PaginatedGetUserListResponse, error) {
	queryString, args, err := query_builder.GetUserListBuilder(param.Query)
	
	var result []dto.PaginatedGetUserListResponse
	err = u.SqlService.SelectMany(
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
