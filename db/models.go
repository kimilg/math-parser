// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Equation struct {
	ID        int64
	Value     string
	Category  string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}
