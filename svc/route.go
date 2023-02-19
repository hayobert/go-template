package svc

import (
	"net/http"

	"go-template/api/route"
	_ "go-template/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//初始化路由
func InitRoute(g *gin.Engine) {
	g.GET("/healthchech", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "true"})
	})

	g.GET("swagger/index", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.InitIndex(g)
}
