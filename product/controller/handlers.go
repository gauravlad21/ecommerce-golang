package controller

import (
	"net/http"
	"strconv"

	"github.com/gauravlad21/ecommerce-golang/common"
	productCommon "github.com/gauravlad21/ecommerce-golang/product/common"
	"github.com/gin-gonic/gin"
)

func Hello(oldctx *gin.Context) {
	ctx := common.GetContext(oldctx)
	msg := serviceRepo.Hello(ctx)
	oldctx.JSON(200, msg)
}

func AddProduct(c *gin.Context) {
	body := &productCommon.Product{}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read body"))
		return
	}
	res := serviceRepo.AddProduct(common.GetContext(c), body)
	c.JSON(http.StatusOK, res)
}

func UpdateProduct(c *gin.Context) {
	body := &productCommon.UpdateProductQuantity{}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read body"))
		return
	}
	res := serviceRepo.UpdateProductQuantity(common.GetContext(c), body)
	c.JSON(http.StatusOK, res)
}

func DeleteProduct(c *gin.Context) {
	idStr := c.Request.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read id"))
		return
	}
	res := serviceRepo.DeleteProduct(common.GetContext(c), int32(id))
	c.JSON(http.StatusOK, res)
}

func GetProduct(c *gin.Context) {
	idStr := c.Request.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read id"))
		return
	}
	res := serviceRepo.GetProduct(common.GetContext(c), int32(id))
	c.JSON(http.StatusOK, res)
}
