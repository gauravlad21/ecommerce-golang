// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package dbsqlc

import (
	"database/sql"
)

type User struct {
	ID       int32
	Email    string
	Password string
	UserType string
	Created  sql.NullTime
	Updated  sql.NullTime
}
