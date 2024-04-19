package common

import (
	"context"

	"github.com/gin-gonic/gin"
)

func GetContext(ctx *gin.Context) context.Context {
	// return context.WithValue(context.Background(), common.GetLoggerKey(), ctx.MustGet(common.GetLoggerKey().String()).(*logging.Logger))
	return context.Background()
}
