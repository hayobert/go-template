package svc

import (
	"go-template/etc"
	"go-template/middleware"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func InitService() (*gin.Engine, error) {
	g := gin.New()

	//普罗米修斯的作用
	p := ginprometheus.NewPrometheus("gin")
	p.Use(g)

	//设置gin debug/release版本
	gin.SetMode(etc.GetConfig().Mode)

	g.Use(gin.Recovery())
	g.Use(middleware.Logger())

	//初始化路由
	InitRoute(g)

	return g, nil
}
