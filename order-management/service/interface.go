package service

import (
	"context"
	"database/sql"

	"github.com/gauravlad21/ecommerce-golang/common"
	omCommon "github.com/gauravlad21/ecommerce-golang/order-management/common"
	"github.com/gauravlad21/ecommerce-golang/order-management/dbhelper"
)

type ServiceIF interface {
	Hello(ctx context.Context) string

	InsertOrder(ctx context.Context, req *omCommon.Order, tx ...*sql.Tx) *common.Response
	InsertSubOrder(ctx context.Context, req *omCommon.SubOrder, tx ...*sql.Tx) *common.Response
	GetOrder(ctx context.Context, req *omCommon.GetOrderRequest, tx ...*sql.Tx) (res *omCommon.GetOrderResponse, err error)
}

type ServiceStruct struct {
	DbOps dbhelper.DbOperationsIF
}

func New(dbOps dbhelper.DbOperationsIF) ServiceIF {
	return &ServiceStruct{
		DbOps: dbOps,
	}
}
