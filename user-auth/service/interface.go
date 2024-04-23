package service

import (
	"context"

	"github.com/gauravlad21/ecommerce-golang/user-auth/common"
	"github.com/gauravlad21/ecommerce-golang/user-auth/dbhelper"
)

type ServiceIF interface {
	Hello(ctx context.Context) string
	Signup(ctx context.Context, body *common.UserAuthBody) *common.Response
	Login(ctx context.Context, body *common.UserAuthBody) *common.LoginResposne
	RequireAuth(ctx context.Context, token string) (*common.User, error)
}

type ServiceStruct struct {
	DbOps dbhelper.DbOperationsIF
}

func New(dbOps dbhelper.DbOperationsIF) ServiceIF {
	return &ServiceStruct{
		DbOps: dbOps,
	}
}
