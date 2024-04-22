package controller

import (
	"context"

	"github.com/gauravlad21/ecommerce-golang/product/dbhelper"
	"github.com/gauravlad21/ecommerce-golang/product/service"
)

var dbOpsIf dbhelper.DbOperationsIF
var serviceRepo service.ServiceIF

func InitializeHandlers() {

	if dbOpsIf == nil {
		dbOpsIf = dbhelper.GetDbOps()
	}

	if serviceRepo == nil {
		serviceRepo = service.New(dbOpsIf)
	}
}

func StartupHook(ctx context.Context) {
	if serviceRepo == nil {
		InitializeHandlers()
	}
}
