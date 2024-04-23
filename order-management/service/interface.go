package service

import (
	"context"
	"database/sql"

	"github.com/gauravlad21/ecommerce-golang/order-management/common"
	"github.com/gauravlad21/ecommerce-golang/order-management/dbhelper"
)

type ServiceIF interface {
	Hello(ctx context.Context) string

	InsertOrder(ctx context.Context, req *common.Order, tx ...*sql.Tx) *common.Response
	InsertSubOrder(ctx context.Context, req *common.SubOrder, tx ...*sql.Tx) *common.Response
	GetOrder(ctx context.Context, req *common.GetOrderRequest, tx ...*sql.Tx) (res *common.GetOrderResponse, err error)
}

type ServiceStruct struct {
	DbOps dbhelper.DbOperationsIF
}

func New(dbOps dbhelper.DbOperationsIF) ServiceIF {
	return &ServiceStruct{
		DbOps: dbOps,
	}
}
