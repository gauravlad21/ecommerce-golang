package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gauravlad21/ecommerce-golang/common"
	omCommon "github.com/gauravlad21/ecommerce-golang/order-management/common"
	"github.com/google/uuid"
)

func (s *ServiceStruct) Hello(ctx context.Context) string {
	return "hello from order-management service"
}

func (s *ServiceStruct) InsertOrder(ctx context.Context, req *omCommon.Order, tx ...*sql.Tx) *common.Response {
	if req == nil {
		return common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "body is nil")
	}
	req.UniqueOrderID = uuid.New().String() // generate unique order id

	err := s.DbOps.InsertOrder(ctx, req)
	if err != nil {
		return common.GetErrResponse(common.StatusCode_INTERNAL_ERROR, err)
	}
	return common.GetDefaultResponse()
}

func (s *ServiceStruct) InsertSubOrder(ctx context.Context, req *omCommon.SubOrder, tx ...*sql.Tx) *common.Response {
	if req == nil || req.UniqueOrderID == "" || req.Status == "" {
		return common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "body is nil")
	}

	err := s.DbOps.InsertSubOrder(ctx, req)
	if err != nil {
		return common.GetErrResponse(common.StatusCode_INTERNAL_ERROR, err)
	}
	return common.GetDefaultResponse()
}

// todo: need validation
func (s *ServiceStruct) GetOrder(ctx context.Context, req *omCommon.GetOrderRequest, tx ...*sql.Tx) (res *omCommon.GetOrderResponse, err error) {
	if req == nil || req.UniqueOrderID == "" {
		return nil, fmt.Errorf("body is nil")
	}

	res, err = s.DbOps.GetOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
