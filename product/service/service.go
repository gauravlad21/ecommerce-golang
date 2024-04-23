package service

import (
	"context"
	"fmt"

	"github.com/gauravlad21/ecommerce-golang/common"
	productCommon "github.com/gauravlad21/ecommerce-golang/product/common"
)

func (s *ServiceStruct) Hello(ctx context.Context) string {
	return "hello from user-auth service"
}

func (s *ServiceStruct) AddProduct(ctx context.Context, product *productCommon.Product) *common.Response {
	// validation
	if product == nil || product.Name == "" || product.PricePerItem == 0 || product.Weight == 0 || product.Quantity == 0 || productCommon.StringToUnit(product.Unit) != "" {
		return common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "body is nil or input invalid")
	}
	id, err := s.DbOps.InsertProduct(ctx, product)
	if err != nil {
		return &common.Response{StatusCode: common.StatusCode_INTERNAL_ERROR, ErrorMsg: []string{err.Error()}}
	}
	fmt.Printf("product inserted with Id %v\n", id)
	return common.GetDefaultResponse()
}

func (s *ServiceStruct) GetProduct(ctx context.Context, id int32) *productCommon.Product {
	product, err := s.DbOps.GetProduct(ctx, id)
	if err != nil || product == nil {
		return &productCommon.Product{}
	}
	return product
}

func (s *ServiceStruct) UpdateProductQuantity(ctx context.Context, req *productCommon.UpdateProductQuantity) *common.Response {
	if req == nil || req.DescreaseQuantityCount == 0 {
		return common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "input is nil")
	}
	err := s.DbOps.UpdateProductQuantity(ctx, req)
	if err != nil {
		return common.GetErrResponse(common.StatusCode_INTERNAL_ERROR, err)
	}
	return common.GetDefaultResponse()
}

func (s *ServiceStruct) DeleteProduct(ctx context.Context, id int32) *common.Response {
	err := s.DbOps.DeleteProduct(ctx, id)
	if err != nil {
		return common.GetErrResponse(common.StatusCode_INTERNAL_ERROR, err)
	}
	return common.GetDefaultResponse()
}
