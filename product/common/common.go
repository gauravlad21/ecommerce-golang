package common

import (
	"context"

	"github.com/gin-gonic/gin"
)

func GetContext(ctx *gin.Context) context.Context {
	if user, ok := ctx.Get("user"); ok {
		usr := user.(*User)
		return context.WithValue(context.Background(), "user", usr)
	}
	return context.Background()
}
