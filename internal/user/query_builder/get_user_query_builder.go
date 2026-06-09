package query_builder

import (
	util_db "gitea.qwertysystem.net/BETS/ts-utils/db"
	util_sql_query "gitea.qwertysystem.net/BETS/ts-utils/sql_query"

	"github.com/CROWNIX/boilerplate-go-v1/internal/user/dto"
)

func GetUserListBuilder(query dto.GetUserQuery) (string, []any, error) {
	paginationQuery := util_sql_query.Pagination{
		Page:      query.Page,
		Limit:     query.Limit,
		SortOrder: query.SortOrder,
		SortBy:    query.SortBy,
	}

	searchQuery := []string{}

	columnToSearchMap := map[string]string{
		"name": "u.name",
	}

	if query.Columns != nil {
		for _, column := range *query.Columns {
			if mapped, ok := columnToSearchMap[column]; ok {
				searchQuery = append(searchQuery, mapped)
			}
		}
	}

	return util_sql_query.NewSQLSelectBuilder[dto.UserDB](util_db.UserTableName, "u").
		Where(map[string]util_sql_query.SQLCondition{
			"u.deleted_at": {
				Operator: util_sql_query.SQLOperatorIsNull,
			},
		}).
		Search(query.Search, searchQuery).
		Paginate(paginationQuery).
		Build()
}