package urlsmappings

import (
	"github.com/gauravlad21/ecommerce-golang/order-management/controller"

	"github.com/gin-gonic/gin"
)

const (
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type urlMap struct {
	Url     string
	Method  string
	Handler []gin.HandlerFunc
}

var urlsMappings []*urlMap

func GetUrlMaps() []*urlMap {
	return urlsMappings
}

func init() {
	urlsMappings = []*urlMap{
		// testing endpoints
		{Url: "/hello", Method: GET, Handler: []gin.HandlerFunc{controller.Hello}},

		{Url: "/userorder", Method: POST, Handler: []gin.HandlerFunc{controller.InsertOrder}},
		{Url: "/suborder", Method: POST, Handler: []gin.HandlerFunc{controller.InsertSubOrder}},
		{Url: "/order", Method: GET, Handler: []gin.HandlerFunc{controller.GetOrder}},
	}
}
