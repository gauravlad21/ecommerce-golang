package controller

import (
	"net/http"
	"strconv"

	"github.com/gauravlad21/ecommerce-golang/order-management/common"
	"github.com/gin-gonic/gin"
)

func Hello(oldctx *gin.Context) {
	ctx := common.GetContext(oldctx)
	msg := serviceRepo.Hello(ctx)
	oldctx.JSON(200, msg)
}

func InsertOrder(c *gin.Context) {
	body := &common.Order{}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read body"))
		return
	}
	res := serviceRepo.InsertOrder(common.GetContext(c), body)
	c.JSON(http.StatusOK, res)
}

func InsertSubOrder(c *gin.Context) {
	body := &common.SubOrder{}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read body"))
		return
	}
	res := serviceRepo.InsertSubOrder(common.GetContext(c), body)
	c.JSON(http.StatusOK, res)
}

func GetOrder(c *gin.Context) {
	unique_order_id := c.Request.URL.Query().Get("unique_order_id")
	user_id_str := c.Request.URL.Query().Get("user_id")
	user_id, err := strconv.Atoi(user_id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read user_id"))
		return
	}

	req := &common.GetOrderRequest{UserID: int32(user_id), UniqueOrderID: unique_order_id}
	res, err := serviceRepo.GetOrder(common.GetContext(c), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, res)
}
