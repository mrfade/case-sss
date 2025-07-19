package postgre

import (
	"context"

	"github.com/mrfade/case-sss/pkg/request"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func WithRequest(ctx context.Context, request *request.Request, count *int64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if request == nil {
			return db
		}

		if request.PageNumber > 0 && request.PageSize > 0 {
			offset := (request.PageNumber - 1) * request.PageSize
			db.Offset(offset).Limit(request.PageSize)
		}

		if len(request.Sorts) > 0 {
			columns := make([]clause.OrderByColumn, 0, len(request.Sorts))
			for field, direction := range request.Sorts {
				columns = append(columns, clause.OrderByColumn{
					Column: clause.Column{Name: field},
					Desc:   direction == "desc",
				})
			}

			db.Clauses(clause.OrderBy{
				Columns: columns,
			})
		}

		if len(request.Filters) > 0 {
			for field, value := range request.Filters {
				db.Where(clause.Eq{Column: field, Value: value})
			}
		}

		db.Count(count)

		return db
	}
}
