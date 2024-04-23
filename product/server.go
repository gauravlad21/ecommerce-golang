package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/gauravlad21/ecommerce-golang/product/common"
	"github.com/gauravlad21/ecommerce-golang/product/controller"
	urlmap "github.com/gauravlad21/ecommerce-golang/product/urls_mappings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func startServer(port string) {
	server := gin.New()
	server.Use(gin.Recovery())

	m := urlmap.GetUrlMaps()
	for _, urlMap := range m {
		url := fmt.Sprintf(viper.GetString("url-prefix") + urlMap.Url)
		switch urlMap.Method {
		case urlmap.GET:
			server.GET(url, urlMap.Handler...)
		case urlmap.POST:
			server.POST(url, urlMap.Handler...)
		case urlmap.DELETE:
			server.DELETE(url, urlMap.Handler...)
		case urlmap.PUT:
			server.PUT(url, urlMap.Handler...)
		case urlmap.PATCH:
			server.PATCH(url, urlMap.Handler...)
		}
	}

	server.Run(":" + port) // ":5002"
}

func initAndStartServer() {
	ctx := context.Background()

	controller.InitializeHandlers()
	controller.StartupHook(ctx)

	port := viper.GetString("port")
	startServer(port)
}

func main() {

	defaultPath := "default-path"
	var configPath string
	flag.StringVar(&configPath, "config", defaultPath, "local config path")

	flag.Parse()

	common.ReadConfigFile(configPath)
	initAndStartServer()
}
