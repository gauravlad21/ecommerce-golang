// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package dbsqlc

import (
	"database/sql"
)

type Product struct {
	ID              int32
	Name            string
	Weight          int32
	Unit            string
	Quantity        int32
	PricePerProduct float64
	Version         int32
	IsDeleted       sql.NullBool
	Created         sql.NullTime
	Updated         sql.NullTime
}