package controller

import (
	"time"

	"github.com/CROWNIX/boilerplate-go-v1/internal/user/dto"
	"github.com/CROWNIX/boilerplate-go-v1/internal/user/usecase"
	"github.com/CROWNIX/boilerplate-go-v1/pkg/entity"
)

type UserController struct {
	Timeout time.Duration
	GetUserListUseCase     entity.UseCase[usecase.GetUserListParam, *dto.PaginatedGetUserListResponse]
}

func MakeUserController(
	timeout time.Duration,
	getUserListUseCase entity.UseCase[usecase.GetUserListParam, *dto.PaginatedGetUserListResponse],
) *UserController {
	return &UserController{
		Timeout:                       timeout,
		GetUserListUseCase:     getUserListUseCase,
	}
}
