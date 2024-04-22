package service

import (
	"context"

	"github.com/gauravlad21/ecommerce-golang/common"
	productCommon "github.com/gauravlad21/ecommerce-golang/product/common"
	"github.com/gauravlad21/ecommerce-golang/product/dbhelper"
)

type ServiceIF interface {
	Hello(ctx context.Context) string
	AddProduct(ctx context.Context, product *productCommon.Product) *common.Response
	GetProduct(ctx context.Context, id int32) *productCommon.Product
	UpdateProductQuantity(ctx context.Context, req *productCommon.UpdateProductQuantity) *common.Response
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
