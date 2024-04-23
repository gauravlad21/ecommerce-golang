package service

import (
	"context"

	"github.com/gauravlad21/ecommerce-golang/product/common"
	"github.com/gauravlad21/ecommerce-golang/product/dbhelper"
)

type ServiceIF interface {
	Hello(ctx context.Context) string
	AddProduct(ctx context.Context, product *common.Product) *common.Response
	GetProduct(ctx context.Context, id int32) *common.Product
	UpdateProductQuantity(ctx context.Context, req *common.UpdateProductQuantity) *common.Response
	DeleteProduct(ctx context.Context, id int32) *common.Response
}

type ServiceStruct struct {
	DbOps dbhelper.DbOperationsIF
}

func New(dbOps dbhelper.DbOperationsIF) ServiceIF {
	return &ServiceStruct{
		DbOps: dbOps,
	}
}
