package dbhelper

import (
	"database/sql"

	"github.com/gauravlad21/ecommerce-golang/order-management/dbhelper/sqlc/dbsqlc"
)

func GetSqlcQuery(q *dbsqlc.Queries, tx ...*sql.Tx) *dbsqlc.Queries {
	var db *dbsqlc.Queries = q
	if len(tx) > 0 {
		db = q.WithTx(tx[0])
	}
	return db
}
