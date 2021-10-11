package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuanyu90221/ethereum_blockchain_services/configs"
)

func SetupRouter() *gin.Engine {
	initRouter := gin.Default()
	initRouter.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	initRouter.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s response successfully", configs.GetEnvConfig().APP_NAME),
			"status":  "health",
		})
	})

	SetBlockRouter(initRouter)
	SetTransactionRouter(initRouter)
	initRouter.Use(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("%s is not Support", ctx.Request.RequestURI),
		})
	})
	return initRouter
}
