package urlsmappings

import (
	"github.com/gauravlad21/ecommerce-golang/product/controller"

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

		{Url: "/product", Method: POST, Handler: []gin.HandlerFunc{controller.AddProduct}},
		{Url: "/product", Method: PUT, Handler: []gin.HandlerFunc{controller.UpdateProduct}},
		{Url: "/product", Method: DELETE, Handler: []gin.HandlerFunc{controller.DeleteProduct}},
		{Url: "/product", Method: GET, Handler: []gin.HandlerFunc{controller.GetProduct}},
	}
}