package dto

import (
	util_dto "gitea.qwertysystem.net/BETS/ts-utils/dto"
)

type GetUserQuery struct {
	Page      int       `json:"page" transform:"int"`
	Limit     int       `json:"limit" transform:"int"`
	SortOrder int       `json:"sort_order" transform:"string"`
	SortBy    string    `json:"sort_by" transform:"string"`
	Search    string    `json:"search" transform:"string"`
	Columns   *[]string `json:"columns" transform:"array"`
}

type UserDB struct {
	ID   string `json:"id"  column:"u.id::text"`
	Name string `json:"name" column:"u.name"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetUserResult struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PaginatedGetUserListResponse = util_dto.PaginationResult[GetUserResult]