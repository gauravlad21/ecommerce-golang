package dbhelper

import (
	"context"
	"database/sql"
	"fmt"

	omCommon "github.com/gauravlad21/ecommerce-golang/order-management/common"
	"github.com/gauravlad21/ecommerce-golang/order-management/dbhelper/sqlc/dbsqlc"
	_ "github.com/lib/pq"
)

type DbOperationsIF interface {
	Exec(ctx context.Context, query string) error
	CloseDb(ctx context.Context) error

	InsertOrder(ctx context.Context, req *omCommon.Order, tx ...*sql.Tx) (err error)
	InsertSubOrder(ctx context.Context, req *omCommon.SubOrder, tx ...*sql.Tx) (err error)
	GetOrder(ctx context.Context, req *omCommon.GetOrderRequest, tx ...*sql.Tx) (res *omCommon.GetOrderResponse, err error)
	// UpdateUserOrderStatus(ctx context.Context, req *omCommon.UpdateSubOrderStatus, tx ...*sql.Tx) (id int32, err error)
	// UpdateSubOrderStatus(ctx context.Context, req *productCommon.Product, tx ...*sql.Tx) (id int32, err error)
}

type DbOps struct {
	DB     *sql.DB
	DbSqlc *dbsqlc.Queries
}

func New(db *sql.DB) DbOperationsIF {
	return &DbOps{DbSqlc: dbsqlc.New(db), DB: db}
}

func (dbOps *DbOps) Exec(ctx context.Context, query string) error {
	_, err := dbOps.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (dbOps *DbOps) CloseDb(ctx context.Context) error {
	return dbOps.DB.Close()
}

// start from below
func (dbOps *DbOps) InsertOrder(ctx context.Context, req *omCommon.Order, tx ...*sql.Tx) (err error) {
	params := dbsqlc.InsertOrderParams{UserID: req.UserID, TotalPrice: req.TotalPrice, UniqueOrderID: req.UniqueOrderID}
	_, err = GetSqlcQuery(dbOps.DbSqlc, tx...).InsertOrder(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (dbOps *DbOps) InsertSubOrder(ctx context.Context, req *omCommon.SubOrder, tx ...*sql.Tx) (err error) {
	params := dbsqlc.InsertSubOrderParams{UniqueOrderID: req.UniqueOrderID, ProductID: req.ProductID, Quantity: req.Quantity, Status: req.Status}
	_, err = GetSqlcQuery(dbOps.DbSqlc, tx...).InsertSubOrder(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (dbOps *DbOps) GetOrder(ctx context.Context, req *omCommon.GetOrderRequest, tx ...*sql.Tx) (res *omCommon.GetOrderResponse, err error) {
	params := dbsqlc.GetOrderParams{UniqueOrderID: req.UniqueOrderID, UserID: req.UserID}
	rows, err := GetSqlcQuery(dbOps.DbSqlc, tx...).GetOrder(ctx, params)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	row := rows[0]
	res = &omCommon.GetOrderResponse{UniqueOrderID: row.UniqueOrderID, UserID: row.UserID, TotalPrice: row.TotalPrice, Orderstatus: row.Orderstatus}
	subOrder := []*omCommon.SubOrderResponse{}
	for _, ro := range rows {
		subOrder = append(subOrder, &omCommon.SubOrderResponse{
			ProductID:      ro.ProductID,
			Quantity:       ro.Quantity,
			Suborderstatus: ro.Suborderstatus,
		})
	}
	res.SubOrderResponse = subOrder
	return res, nil

}
