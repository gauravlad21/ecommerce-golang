package urlsmappings

import (
	"github.com/gauravlad21/ecommerce-golang/user-auth/controller"

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

		{Url: "/signup", Method: POST, Handler: []gin.HandlerFunc{controller.Signup}},
		{Url: "/login", Method: POST, Handler: []gin.HandlerFunc{controller.Login}},
		{Url: "/validate", Method: GET, Handler: []gin.HandlerFunc{controller.RequireAuth, controller.Validate}},
		{Url: "/authorize", Method: POST, Handler: []gin.HandlerFunc{controller.Authorized}},
	}
}
